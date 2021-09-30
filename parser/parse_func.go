package parser

import (
	"fmt"
	"tee/lexer"
)

/*

func name() {}
func name(a) {}
func name(a,b) {}
*/

func (p *parser) parseFunc(b *block) (root *node) {

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

	funcSym := b.lookup(p.current())

	if p.next() {
		return nil
	}

	if p.current().Type != lexer.T_OPEN_PARS {
		return nil
	}

	if p.next() {
		return nil
	}

	args := &node{}
	funcBlock := newBlock()
	funcBlock.AddNode(args) // ?

	// no arg func
	if p.current().Type == lexer.T_CLOSE_PARS {
		goto body
	}

	if p.current().Type != lexer.T_SYMBOL {
		return nil
	}

	args.AddChild(b.lookup(p.current()))

	if p.next() {
		return nil
	}

	for {
		if p.done() {
			break
		}

		if p.current().Type == lexer.T_COMMA && p.canPeek() && p.peek().Type == lexer.T_SYMBOL {
			p.next()
			args.AddChild(b.lookup(p.current()))
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

	b.AddChild(funcBlock)

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

	funcSym.token.Type = lexer.T_FUNC_SYMBOL
	fmt.Println(funcSym.token)

	funcSym.AddChild(args)
	funcSym.AddChild(funcBlock.nodes...)

	return funcSym

}

/*
	asd()
	asd(a)
	asd(a,b)
*/
func (p *parser) parseCall(b *block) (root *node) {

	pos := p.pos
	defer func() {
		if root == nil {
			p.pos = pos
		}
	}()

	callSymbol := b.lookup(p.current())

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

	// no args
	if p.current().Type == lexer.T_CLOSE_PARS {
		goto done
	}

	if p.current().Type != lexer.T_SYMBOL {
		return nil
	}

	callSymbol.AddChild(b.lookup(p.current()))

	if p.next() {
		return nil
	}

	for {
		if p.done() {
			break
		}

		if p.current().Type == lexer.T_COMMA && p.canPeek() && p.peek().Type == lexer.T_SYMBOL {
			p.next()
			callSymbol.AddChild(b.lookup(p.current()))
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

	callSymbol.token.Type = lexer.T_FUNC_CALL
	fmt.Println(callSymbol.token)
	p.next()

	return callSymbol
}
