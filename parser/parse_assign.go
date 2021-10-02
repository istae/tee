package parser

import (
	"tee/lexer"
)

/*
GRAMMER:
VAR = EXP
*/
func (p *parser) parseAssign(b *Block) (root *Node) {

	pos := p.pos
	defer func() {
		if root == nil { // reset pos if tokens cannot be processed
			p.pos = pos
		}
	}()

	if p.current().Type != lexer.T_SYMBOL {
		return nil
	}

	var varToken = p.current()

	if p.next() {
		return nil
	}

	if p.current().Type != lexer.T_EQUAL {
		return nil
	}

	equalToken := p.current()

	if p.next() {
		return nil
	}

	expNode := p.parseExpression(b)
	if expNode == nil {
		return nil
	}

	if !p.done() && !(p.current().Type == lexer.T_NEWLINE || p.current().Type == lexer.T_COMMENT) {
		return nil
	}

	p.next()

	equalNode := &Node{Token: equalToken}

	if n := b.getNode(varToken); n != nil && n.Token.Type != lexer.T_SYMBOL {
		p.multidefinition(varToken, n.Token)
		return nil
	}

	equalNode.AddChild(b.setNode(varToken))
	equalNode.AddChild(expNode)

	return equalNode
}

/*
EXP -> VAR | NUM | FUNC_CALL
EXP -> EXP OP EXP
*/
func (p *parser) parseExpression(b *Block) (root *Node) {

	pos := p.pos
	defer func() {
		if root == nil {
			p.pos = pos
		}
	}()

	var leftNode *Node

	t := p.current().Type
	if t != lexer.T_SYMBOL && t != lexer.T_NUM && t != lexer.T_STRING {
		return nil
	}

	leftNode = p.parseCall(b)

	if leftNode != nil {
		f := b.getNode(leftNode.Token)
		if f == nil || f.Token.Type != lexer.T_FUNC_SYMBOL {
			p.undefinedSymbol(leftNode.Token)
			return nil
		}
	}

	if leftNode == nil {
		if p.current().Type == lexer.T_NUM || p.current().Type == lexer.T_STRING {
			leftNode = &Node{Token: p.current()}
		} else {
			leftNode = b.getNode(p.current())
			if leftNode == nil {
				p.undefinedSymbol(p.current())
				return nil
			}
		}
		if p.next() {
			return nil
		}
	}

	// EXP -> SYM, // EXP -> NUM
	if p.done() || (p.current().Type != lexer.T_MATH_OPS && p.current().Type != lexer.T_CMP_OPS) {
		return leftNode
	}

	// EXP -> EXP OP EXP

	opsToken := p.current()

	if p.next() {
		return nil
	}

	expressionNode := p.parseExpression(b)
	if expressionNode == nil {
		return nil
	}

	opsNode := &Node{Token: opsToken}

	opsNode.AddChild(leftNode)

	// if opsNode has higher presedence than rightNode,
	/* ex: 3 * 4 + 2
		*							+
	3		+					*		2
		4		2  becomes	3		4
	*/
	if opsNode.PrecendenceCmp(expressionNode) {

		opsNode.AddChild(expressionNode.PopChild())
		expressionNode.LeftChild(opsNode)

		return expressionNode
	}

	opsNode.AddChild(expressionNode)

	return opsNode
}
