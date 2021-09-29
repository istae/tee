package parser

import (
	"fmt"
	"tee/lexer"
)

func (p *parser) parseIf(b *block) (root *node) {

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

	fmt.Println("! if !")

	if p.next() {
		return nil
	}

	fmt.Println(p.current().Str)

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
	b.AddChild(ifBlock)

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

	n := &node{
		token: ifToken,
	}
	n.AddChild(exp)
	n.AddChild(ifBlock.nodes...)

	return n

}
