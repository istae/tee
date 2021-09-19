package parser

import (
	"errors"
	"fmt"
	"tee/lexer"
)

type parser struct {
	tokens []lexer.Token

	pos int
	end int
	str string
}

func NewParser() *parser {

	return &parser{}
}

func (p *parser) AST(str string, tokens []lexer.Token) error {

	p.end = len(tokens)
	p.tokens = tokens
	p.str = str

	rootBlock := newBlock()

	for {
		if p.next() {
			break
		}

		n := p.parseVar(rootBlock)
		p.printNode(n)
		if n == nil {
			return errors.New("asd")
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

func (p *parser) canPeek() bool {
	return p.pos+1 < p.end
}

func (p *parser) peek() lexer.Token {
	return p.tokens[p.pos+1]
}
