package parser

import (
	"fmt"
	"tee/lexer"
)

func (p *parser) parseIf(b *Block) (root *Node) {

	ifToken := p.current()

	if p.current().Type != lexer.T_IF {
		return nil
	}

	if p.next() {
		return nil
	}

	exp := p.parseExpression(b)
	if exp == nil {
		return nil
	}

	fmt.Println("~~~ if exp", exp.Token)

	if p.current().Type != lexer.T_OPEN_BRACKET {
		return nil
	}

	if p.next() {
		return nil
	}

	ifBlock := newBlock(ifBlock)
	ifBlock.Parent(b)

	for {
		if p.done() {
			break
		}

		n := p.parse(ifBlock)
		if n == nil {
			break
		}
		ifBlock.AddNode(n)
	}

	if p.done() {
		return nil
	}

	if p.current().Type != lexer.T_CLOSE_BRACKET {
		return nil
	}

	p.next()

	n := &Node{
		Token: ifToken,
	}
	n.AddChildAndParent(exp)
	n.AddChildAndParent(ifBlock.Nodes...)

	return n

}
