package parser

import (
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

	err error
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
		p.parseBreak,
	}

	return p
}

func (p *parser) AST(str string, tokens []lexer.Token) (*Block, error) {

	p.end = len(tokens)
	p.tokens = tokens
	p.str = str

	rootBlock := newBlock("")
	addBuiltInFunc(rootBlock)

	for {
		if p.done() {
			break
		}

		n := p.parse(rootBlock)
		if n == nil {
			return nil, p.generateErr()
		}
		p.err = nil
		// p.printNode(n)
		rootBlock.AddNode(n)
	}

	return rootBlock, nil
}

func addBuiltInFunc(b *Block) {
	b.setNode(lexer.Token{
		Type: lexer.T_FUNC_SYMBOL,
		Str:  "print",
	})
}

func (p *parser) generateErr() error {

	if p.err == nil {
		return fmt.Errorf("unknown error around %s at line %d", p.current().Str, p.current().Line)

	}
	return p.err
}

func (pr *parser) parse(b *Block) *Node {
	for _, p := range pr.parsers {
		start := pr.pos
		if t := p(b); t != nil {
			return t
		}
		pr.pos = start
	}
	return nil
}

func (p *parser) undefinedSymbol(t lexer.Token) {
	p.err = fmt.Errorf("undefined symbol %s, line %d", t.Str, t.Line)
}

func (p *parser) multidefinition(last, first lexer.Token) {
	p.err = fmt.Errorf("symbol %s redefined line %d, defined at line %d already", last.Str, last.Line, first.Line)
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
