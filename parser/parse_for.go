package parser

import (
	"fmt"
	"tee/lexer"
)

func (p *parser) parseBreak(b *Block) (root *Node) {

	if p.current().Type == lexer.T_BREAK {
		look := b
		for look != nil {
			if look.blockType == forBlock {
				fmt.Println(p.current())
				defer p.next()
				return &Node{
					Token: p.current(),
				}
			}

			if look.blockType != ifBlock {
				break
			}
			look = look.parent
		}

		p.undefinedSymbol(p.current())
	}

	return nil
}

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

	forBlock := newBlock(forBlock)
	forBlock.Parent(b)

	for {
		if p.done() {
			break
		}

		n := p.parse(forBlock)
		if n == nil {
			if p.current().Type == lexer.T_BREAK {
				forBlock.AddNode(&Node{Token: p.current()})
			}
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
	n.AddChildAndParent(exp)
	n.AddChildAndParent(forBlock.Nodes...)

	return n

}
