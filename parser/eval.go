package parser

import (
	"fmt"
	"strconv"
	"tee/lexer"
)

type value struct {
	tokenType lexer.TokenType
	varType   string
	val       interface{}
}

func (v value) Print() {
	fmt.Printf("%t - %v\n", v.val, v.val)
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
		v := e.result(n)
		v.Print()
		values = append(values, v)
		fmt.Println("--------------------------------")
	}

	return
}

func (e *eval) result(n *node) *value {

	// fmt.Println(n.token.Str, n.token.Type)

	if n.token.Type == lexer.T_EQUAL {
		left := e.lookup(n.children[0])
		res := e.result(n.children[1])
		left.val = res.val
		left.varType = res.varType
		return left
	}

	if n.token.Type == lexer.T_MATH_OPS {

		left := e.result(n.children[0])
		right := e.result(n.children[1])

		fmt.Printf("left %v  right %v\n", left, right)

		e.lookup(n).val = mathOpsFloat(n.token.Str, getFloat(left), getFloat(right))
		e.lookup(n).varType = "num"
	}

	if n.token.Type == lexer.T_CMP_OPS {

		left := e.result(n.children[0])
		right := e.result(n.children[1])

		fmt.Printf("left %v right %v\n", left, right)

		e.lookup(n).val = cmpOpsFloat(n.token.Str, getFloat(left), getFloat(right))
	}

	if n.token.Type == lexer.T_IF {

		exp := e.result(n.children[0])
		body := n.children[1:]

		if exp.val.(bool) {
			for _, b := range body {
				e.result(b)
			}
		}

		e.lookup(n).val = exp.val
	}

	if n.token.Type == lexer.T_FOR {
		body := n.children[1:]

		for e.result(n.children[0]).val.(bool) {
			for _, b := range body {
				e.result(b)
			}
		}
	}

	return e.lookup(n)
}

func getFloat(v *value) float64 {

	if v.tokenType == lexer.T_NUM {
		return toFloat(v.val.(string))
	}

	if v.varType == "num" {
		return v.val.(float64)
	}

	return 0
}

func mathOpsFloat(ops string, left, right float64) float64 {
	switch ops {
	case "+":
		return left + right
	case "-":
		return left - right
	case "/":
		return left / right
	case "*":
		return left * right
	}
	return 0
}

func cmpOpsFloat(ops string, left, right float64) bool {
	switch ops {
	case "<":
		return left < right
	case ">":
		return left > right
	case "==":
		return left == right
	}
	return false
}

func toFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 32)
	return f
}
