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
GRAMMER:
VAR = EXP
*/
func (p *parser) parseVar(b *block) *node {

	pos := p.pos

	if p.current().Type == lexer.T_VAR {

		var (
			varT    = p.current()
			varTStr = p.current().Str(p.str)
		)

		if p.next() {
			p.pos = pos
			return nil
		}

		if p.current().Type != lexer.T_EQUAL {
			return nil
		}

		eqT := p.current()

		if p.next() {
			p.pos = pos
			return nil
		}

		eqN := &node{token: eqT}
		varN := b.lookup(varTStr, varT)
		varN.parent = eqN
		eqN.children = append(eqN.children, varN)

		expN := p.parseExpression(b)
		if expN == nil {
			return nil
		}

		expN.parent = eqN
		eqN.children = append(eqN.children, expN)

		return eqN
	}

	return nil
}

/*
EXP -> VAR | NUM
EXP -> EXP OP EXP
*/
func (p *parser) parseExpression(b *block) *node {

	if p.current().Type != lexer.T_VAR && p.current().Type != lexer.T_DOUBLE && p.current().Type != lexer.T_INT {
		return nil
	}

	// EXP -> VAR
	if p.current().Type == lexer.T_VAR {
		if !p.canPeek() || p.peek().Type != lexer.T_OPS {
			defer p.next()
			return b.lookup(p.current().Str(p.str), p.current())
		}
	}

	// EXP -> NUM
	if p.current().Type == lexer.T_DOUBLE || p.current().Type == lexer.T_INT {
		if !p.canPeek() || p.peek().Type != lexer.T_OPS {
			defer p.next()
			return &node{token: p.current()}
		}
	}

	left := p.current()

	if p.next() {
		return nil
	}

	// EXP -> EXP OP EXP
	if p.current().Type == lexer.T_OPS {

		ops := p.current()

		if p.next() {
			return nil
		}

		right := p.parseExpression(b)
		if right == nil {
			return nil
		}

		ret := &node{token: ops}

		right.parent = ret

		leftN := b.lookup(left.Str(p.str), left)
		leftN.parent = ret

		ret.children = append(ret.children, leftN)
		ret.children = append(ret.children, right)

		return ret

	}

	return nil

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
