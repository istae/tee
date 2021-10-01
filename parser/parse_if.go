package parser

import (
	"tee/lexer"
)

func (p *parser) parseIf(b *Block) (root *Node) {

	pos := p.pos
	defer func() {
		if root == nil {
			p.pos = pos
		}
	}()

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

	if p.current().Type != lexer.T_OPEN_BRACKET {
		return nil
	}

	if p.next() {
		return nil
	}

	ifBlock := newBlock()
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
	n.AddChild(exp)
	n.AddChild(ifBlock.Nodes...)

	return n

}
