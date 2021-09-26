package parser

import (
	"errors"
	"fmt"
	"tee/lexer"
)

type parserFunc func(*block) *node

type parser struct {
	tokens []lexer.Token

	parsers []parserFunc

	pos int
	end int
	str string
}

func NewParser() *parser {

	p := &parser{}

	p.parsers = []parserFunc{
		p.parseVar,
		p.parseNewline,
		p.parseIf,
	}

	return p
}

func (p *parser) AST(str string, tokens []lexer.Token) error {

	p.end = len(tokens)
	p.tokens = tokens
	p.str = str

	rootBlock := newBlock()

	for {
		if p.done() {
			break
		}

		n := p.parse(rootBlock)
		p.printNode(n)
		if n == nil {
			return errors.New("~~ERROR~~")
		}
		rootBlock.AddNode(n)
	}

	return nil

}

func (p *parser) parse(b *block) *node {
	for _, p := range p.parsers {
		if t := p(b); t != nil {
			return t
		}
	}
	return nil
}

func (p *parser) printNode(n *node) {

	if n == nil {
		return
	}

	fmt.Printf("%s id %d", n.token.Str, n.token.Start)
	if n.parent != nil {
		fmt.Printf(" <- %d", n.parent.token.Start)
	}
	fmt.Println()

	for _, c := range n.children {
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
