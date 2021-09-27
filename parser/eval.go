package parser

import (
	"fmt"
	"strconv"
	"tee/lexer"
)

type value struct {
	tokenType lexer.TokenType
	val       interface{}
}

func (v value) Print() {
	fmt.Println(v.val)
}

type eval struct {
	evals map[*node]*value
}

func NewEval() *eval {

	return &eval{
		evals: make(map[*node]*value),
	}
}

func (e *eval) lookup(n *node) *value {

	v, ok := e.evals[n]
	if ok {
		return v
	}

	v = &value{tokenType: n.token.Type, val: n.token.Str}
	e.evals[n] = v

	return v
}

func (e *eval) Eval(b *block) (values []*value) {

	for _, n := range b.nodes {
		values = append(values, e.result(n))
	}

	return
}

func (e *eval) result(n *node) *value {

	if n.token.Type == lexer.T_EQUAL {
		e.lookup(n.children[0]).val = e.result(n.children[1])
		return e.lookup(n.children[0])
	}

	if n.token.Type == lexer.T_MATH_OPS {

		left := e.result(n.children[0])
		right := e.result(n.children[1])

		if left.tokenType == lexer.T_INT && right.tokenType == lexer.T_INT {
			leftInt, _ := strconv.Atoi(left.val.(string))
			rightInt, _ := strconv.Atoi(right.val.(string))
			e.lookup(n).val = leftInt + rightInt
		}

	}

	return e.lookup(n)
}
