package parser

import (
	"fmt"
	"tee/lexer"
)

func (p *parser) parseFor(b *block) (root *node) {

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

	forBlock := newBlock()
	b.AddChild(forBlock)

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

	n := &node{
		token: forToken,
	}
	n.AddChild(exp)
	n.AddChild(forBlock.nodes...)

	return n

}
