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

type node struct {
	children []*node
	parent   *node
	token    lexer.Token
}

func (n *node) AddChild(c *node) {
	n.children = append(n.children, c)
}

type block struct {
	// children []*block
	symbols map[string]*node
	// parent   *block
	root *node
}

func (b *block) lookup(s string, t lexer.Token) *node {
	n := b.symbols[s]
	if n == nil {
		n = &node{token: t}
		b.symbols[s] = n
	}
	return n
}

func newBlock() *block {

	return &block{
		symbols: map[string]*node{},
		root:    &node{},
	}

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

		n := p.parseExpression(rootBlock)
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

	fmt.Printf("%s\n", n.token.Str(p.str))

	// if n.parent != nil {
	// 	fmt.Printf("\n->%s\n", n.parent.token.Str(p.str))
	// }

	for i, c := range n.children {
		fmt.Printf("(%d)\n", i)
		p.printNode(c)
	}
}

/*
EXP -> VAR | NUM
EXP -> EXP OP EXP
*/
func (p *parser) parseExpression(b *block) *node {

	if p.current().Type != lexer.T_VAR && p.current().Type != lexer.T_DOUBLE && p.current().Type != lexer.T_INT {
		return nil
	}

	if !p.canPeek() || p.peek().Type != lexer.T_OPS {
		defer p.next()

		if p.current().Type == lexer.T_VAR {
			return b.lookup(p.current().Str(p.str), p.current()) // EXP -> VAR
		} else {
			return &node{token: p.current()} // EXP -> NUM
		}
	}

	//EXP -> EXP OP EXP
	leftToken := p.current()

	if p.next() {
		return nil
	}

	opsToken := p.current()

	if p.next() {
		return nil
	}

	expNode := p.parseExpression(b)
	if expNode == nil {
		return nil
	}

	opsNode := &node{token: opsToken}

	leftNode := b.lookup(leftToken.Str(p.str), leftToken)

	leftNode.parent = opsNode
	expNode.parent = opsNode

	opsNode.AddChild(leftNode)
	opsNode.AddChild(expNode)

	return opsNode
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

/*
GRAMMER:
VAR = EXP
*/
// func (p *parser) parseVar(b *block) *node {

// 	if p.current().Type != lexer.T_VAR {
// 		return nil
// 	}

// 	pos := p.pos

// 	var (
// 		varToken    = p.current()
// 		varTokenStr = p.current().Str(p.str)
// 	)

// 	if p.next() {
// 		p.pos = pos
// 		return nil
// 	}

// 	if p.current().Type != lexer.T_EQUAL {
// 		return nil
// 	}

// 	equalToken := p.current()

// 	if p.next() {
// 		p.pos = pos
// 		return nil
// 	}

// 	expNode := p.parseExpression(b)
// 	if expNode == nil {
// 		return nil
// 	}

// 	equalNode := &node{token: equalToken}

// 	varNode := b.lookup(varTokenStr, varToken)

// 	equalNode.AddChild(varNode)
// 	equalNode.AddChild(expNode)

// 	varNode.parent = equalNode
// 	expNode.parent = equalNode

// 	return equalNode
// }
