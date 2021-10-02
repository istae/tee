package parser

import (
	"tee/lexer"
)

/*

func name() {}
func name(a) {}
func name(a,b) {}
*/

func (p *parser) parseFunc(b *Block) (root *Node) {

	pos := p.pos
	defer func() {
		if root == nil {
			p.pos = pos
		}
	}()

	if p.current().Type != lexer.T_FUNC {
		return nil
	}

	if p.next() {
		return nil
	}

	if p.current().Type != lexer.T_SYMBOL {
		return nil
	}

	funcToken := p.current()
	funcSym := b.getNode(funcToken)

	if funcSym != nil {
		p.multidefinition(funcToken, funcSym.Token)
		return nil
	}

	if p.next() {
		return nil
	}

	if p.current().Type != lexer.T_OPEN_PARS {
		return nil
	}

	if p.next() {
		return nil
	}

	args := &Node{}
	funcBlock := newBlock()
	funcBlock.AddNode(args) // ?

	// no arg func
	if p.current().Type == lexer.T_CLOSE_PARS {
		goto body
	}

	if p.current().Type != lexer.T_SYMBOL {
		return nil
	}

	args.AddChild(funcBlock.setNode(p.current()))

	if p.next() {
		return nil
	}

	for {
		if p.done() {
			break
		}

		if p.current().Type == lexer.T_COMMA && p.canPeek() && p.peek().Type == lexer.T_SYMBOL {
			p.next()
			args.AddChild(funcBlock.setNode(p.current()))
			p.next()
		} else {
			break
		}
	}

	if p.done() {
		return nil
	}

	if p.current().Type != lexer.T_CLOSE_PARS {
		return nil
	}

body:
	if p.next() {
		return nil
	}

	if p.current().Type != lexer.T_OPEN_BRACKET {
		return nil
	}

	if p.next() {
		return nil
	}

	funcBlock.Parent(b)

	for {
		if p.done() {
			break
		}

		n := p.parse(funcBlock)
		if n == nil {
			break
		}
		funcBlock.AddNode(n)
	}

	if p.done() {
		return nil
	}

	if p.current().Type != lexer.T_CLOSE_BRACKET {
		return nil
	}

	p.next()

	funcSym = b.setNode(funcToken)
	funcSym.Token.Type = lexer.T_FUNC_SYMBOL

	funcSym.AddChild(args)
	funcSym.AddChild(funcBlock.Nodes...)

	return funcSym

}

/*
	asd()
	asd(a)
	asd(a,b)
*/
func (p *parser) parseCall(b *Block) (root *Node) {

	pos := p.pos
	defer func() {
		if root == nil {
			p.pos = pos
		}
	}()

	callToken := p.current()

	if p.current().Type != lexer.T_SYMBOL {
		return nil
	}

	if p.next() {
		return nil
	}

	if p.current().Type != lexer.T_OPEN_PARS {
		return nil
	}

	if p.next() {
		return nil
	}

	callToken.Type = lexer.T_FUNC_CALL
	n := &Node{
		Token: callToken,
	}

	var arg *Node

	// no args
	if p.current().Type == lexer.T_CLOSE_PARS {
		goto done
	}

	if p.current().Type != lexer.T_SYMBOL {
		return nil
	}

	arg = b.getNode(p.current())
	if arg == nil {
		p.undefinedSymbol(p.current())
		return nil
	}

	n.AddChild(arg)

	if p.next() {
		return nil
	}

	for {
		if p.done() {
			break
		}

		if p.current().Type == lexer.T_COMMA && p.canPeek() && p.peek().Type == lexer.T_SYMBOL {
			p.next()
			arg = b.getNode(p.current())
			if arg == nil {
				p.undefinedSymbol(p.current())
				return nil
			}
			p.next()
		} else {
			break
		}
	}

	if p.done() {
		return nil
	}

	if p.current().Type != lexer.T_CLOSE_PARS {
		return nil
	}

done:
	p.next()

	return n
}
