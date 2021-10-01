package parser

import (
	"errors"
	"fmt"
	"tee/lexer"
)

type parserFunc func(*Block) *Node

type parser struct {
	tokens []lexer.Token

	parsers []parserFunc

	pos int
	end int
	str string

	errStr string
}

func NewParser() *parser {

	p := &parser{}

	p.parsers = []parserFunc{
		p.parseAssign,
		p.parseNewline,
		p.parseIf,
		p.parseFor,
		p.parseFunc,
		p.parseCall,
	}

	return p
}

func (p *parser) AST(str string, tokens []lexer.Token) (*Block, error) {

	p.end = len(tokens)
	p.tokens = tokens
	p.str = str

	rootBlock := newBlock()

	for {
		if p.done() {
			break
		}

		n := p.parse(rootBlock)
		if n == nil {
			return nil, errors.New(p.errStr)
		}
		p.printNode(n)
		rootBlock.AddNode(n)
	}

	return rootBlock, nil
}

func (p *parser) parse(b *Block) *Node {
	for _, p := range p.parsers {
		if t := p(b); t != nil {
			return t
		}
	}
	return nil
}

func (p *parser) undefinedSymbol(t lexer.Token) {
	p.errStr = fmt.Sprintf("\nundefined symbol %s, line %d\n", t.Str, t.Line)
}

func (p *parser) multidefinition(t lexer.Token) {
	p.errStr = fmt.Sprintf("\n symbol %s redefined line %d\n", t.Str, t.Line)
}

func (p *parser) printNode(n *Node) {

	if n == nil {
		return
	}

	fmt.Printf("%s id %d", n.Token.Str, n.Token.Start)
	if n.parent != nil {
		fmt.Printf(" <- %d", n.parent.Token.Start)
	}

	fmt.Println()

	for _, c := range n.Children {
		p.printNode(c)
	}
}

func (p *parser) current() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) next() bool {
	p.pos++
	return p.pos >= p.end
}

func (p *parser) done() bool {
	return p.pos >= p.end
}

func (p *parser) canPeek() bool {
	return p.pos+1 < p.end
}

func (p *parser) peek() lexer.Token {
	return p.tokens[p.pos+1]
}
