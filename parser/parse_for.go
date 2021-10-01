package parser

import (
	"tee/lexer"
)

func (p *parser) parseFor(b *Block) (root *Node) {

	pos := p.pos
	defer func() {
		if root == nil {
			p.pos = pos
		}
	}()

	forToken := p.current()

	if p.current().Type != lexer.T_FOR {
		return nil
	}

	if p.next() {
		return nil
	}

	exp := p.parseExpression(b)
	if exp == nil {
		return nil
	}

	if p.current().Type != lexer.T_OPEN_BRACKET {
		return nil
	}

	if p.next() {
		return nil
	}

	forBlock := newBlock()
	forBlock.Parent(b)

	for {
		if p.done() {
			break
		}

		n := p.parse(forBlock)
		if n == nil {
			break
		}
		forBlock.AddNode(n)
	}

	if p.done() {
		return nil
	}

	if p.current().Type != lexer.T_CLOSE_BRACKET {
		return nil
	}

	p.next()

	n := &Node{
		Token: forToken,
	}
	n.AddChild(exp)
	n.AddChild(forBlock.Nodes...)

	return n

}
