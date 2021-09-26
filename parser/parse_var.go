package parser

import (
	"tee/lexer"
)

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
func (p *parser) parseExpression(b *block) (root *node) {

	pos := p.pos
	defer func() {
		if root == nil { // reset pos if tokens cannot be processed
			p.pos = pos
		}
	}()

	if p.current().Type != lexer.T_VAR && p.current().Type != lexer.T_DOUBLE && p.current().Type != lexer.T_INT && p.current().Type != lexer.T_STRING {
		return nil
	}

	if !p.canPeek() || (p.peek().Type != lexer.T_MATH_OPS && p.peek().Type != lexer.T_CMP_OPS) {
		defer p.next()

		if p.current().Type == lexer.T_VAR {
			return b.lookup(p.current()) // EXP -> VAR
		} else {
			return &node{token: p.current()} // EXP -> NUM
		}
	}

	//EXP -> EXP OP EXP
	leftToken := p.current()

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

	if opsNode.PrecendenceCmp(expressionNode) {

		opsNode.AddChild(expressionNode.PopChild())
		expressionNode.LeftChild(opsNode)

		return expressionNode
	}

	opsNode.AddChild(expressionNode)

	return opsNode
}
