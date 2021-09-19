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

var (
	precendence = map[string]int{"/": 5, "*": 5, "+": 4, "-": 4}
)

func NewParser() *parser {

	return &parser{}
}

type node struct {
	children []*node
	parent   *node
	token    lexer.Token
}

// if n has higher precedence
func (n *node) PrecendenceCmp(x *node) bool {

	xPre := precendence[x.token.Str]
	// fmt.Printf("%s %d\n", x.token.Str, xPre)
	if xPre == 0 {
		return false
	}

	nPre := precendence[n.token.Str]
	// fmt.Printf("%s %d\n", n.token.Str, nPre)
	if nPre == 0 {
		return false
	}

	return nPre > xPre
}

func (n *node) AddChild(c *node) {
	c.parent = n
	n.children = append(n.children, c)
}

func (n *node) LeftChild(c *node) {
	c.parent = n
	n.children = append([]*node{c}, n.children...)
}

func (n *node) PopChild() *node {
	var c *node
	c, n.children = n.children[0], n.children[1:]
	return c
}

type block struct {
	// children []*block
	symbols map[string]*node
	// parent   *block
	root *node
}

func (b *block) lookup(t lexer.Token) *node {
	n := b.symbols[t.Str]
	if n == nil {
		n = &node{token: t}
		b.symbols[t.Str] = n
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

	fmt.Printf("%s id %d", n.token.Str, n.token.Start)
	if n.parent != nil {
		fmt.Printf(" <- %d", n.parent.token.Start)
	}
	fmt.Println()

	for _, c := range n.children {
		p.printNode(c)
	}
}

/*
GRAMMER:
VAR = EXP
*/
func (p *parser) parseVar(b *block) *node {

	if p.current().Type != lexer.T_VAR {
		return nil
	}

	pos := p.pos

	var varToken = p.current()

	if p.next() {
		p.pos = pos
		return nil
	}

	if p.current().Type != lexer.T_EQUAL {
		return nil
	}

	equalToken := p.current()

	if p.next() {
		p.pos = pos
		return nil
	}

	expNode := p.parseExpression(b)
	if expNode == nil {
		return nil
	}

	equalNode := &node{token: equalToken}

	varNode := b.lookup(varToken)

	equalNode.AddChild(varNode)
	equalNode.AddChild(expNode)

	return equalNode
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
			return b.lookup(p.current()) // EXP -> VAR
		} else {
			return &node{token: p.current()} // EXP -> NUM
		}
	}

	//EXP -> EXP OP EXP
	leftToken := p.current()
	fmt.Println(leftToken.Str)

	if p.next() {
		return nil
	}

	opsToken := p.current()

	if p.next() {
		return nil
	}

	expressionNode := p.parseExpression(b)
	if expressionNode == nil {
		return nil
	}

	opsNode := &node{token: opsToken}

	var leftNode *node
	if leftToken.Type == lexer.T_VAR {
		leftNode = b.lookup(leftToken)
	} else {
		leftNode = &node{token: leftToken}
	}

	opsNode.AddChild(leftNode)

	// if opsNode has higher presedence than rightNode,
	/* ex: 3 * 4 + 2
		*							+
	3		+					*		2
		4		2  becomes	3		4
	*/

	fmt.Printf("%s %s\n", opsNode.token.Str, expressionNode.token.Str)
	if opsNode.PrecendenceCmp(expressionNode) {

		opsNode.AddChild(expressionNode.PopChild())
		expressionNode.LeftChild(opsNode)

		return expressionNode
	}

	opsNode.AddChild(expressionNode)

	return opsNode
}

func deepestLowerPrecedenceNode(n *node, r *node) {

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
