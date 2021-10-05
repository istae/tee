package eval

import (
	"fmt"
	"strconv"
	"tee/lexer"
	"tee/parser"
)

type valueType string

const (
	num     valueType = "num"
	boolean valueType = "bool"
)

type value struct {
	token   lexer.Token
	valType valueType
	val     interface{}
}

func (v value) Print() {
	fmt.Printf("%v %t - %v\n", v.token, v.val, v.val)
}

type eval struct {
	evals map[*parser.Node]*value
}

func NewEval() *eval {

	return &eval{
		evals: make(map[*parser.Node]*value),
	}
}

func (e *eval) lookup(n *parser.Node) *value {

	v, ok := e.evals[n]
	if ok {
		return v
	}

	v = &value{token: n.Token, val: n.Token.Str}
	if n.Token.Type == lexer.T_NUM {
		v.val = toFloat(n.Token.Str)
		v.valType = "num"
	}

	e.evals[n] = v

	return v
}

func (e *eval) Eval(b *parser.Block) (values []*value) {

	for _, n := range b.Nodes {
		v := e.result(n)
		// v.Print()
		values = append(values, v)
		// fmt.Println("--------------------------------")
	}

	return
}

func (e *eval) result(n *parser.Node) *value {

	// fmt.Println(n.Token)

	if n.Token.Type == lexer.T_EQUAL {
		left := e.lookup(n.Children[0])
		res := e.result(n.Children[1])
		left.val = res.val
		left.valType = res.valType
		fmt.Printf("equal: %v\n", left)
		return left
	}

	if n.Token.Type == lexer.T_MATH_OPS {

		left := e.result(n.Children[0])
		right := e.result(n.Children[1])

		fmt.Printf("ops: %s %v %v\n", n.Token.Str, left, right)

		e.lookup(n).val = mathOpsFloat(n.Token.Str, getFloat(left), getFloat(right))
		e.lookup(n).valType = num
	}

	if n.Token.Type == lexer.T_CMP_OPS {

		left := e.result(n.Children[0])
		right := e.result(n.Children[1])

		fmt.Printf("cmp: %s %v %v\n", n.Token.Str, left, right)

		e.lookup(n).val = cmpOpsFloat(n.Token.Str, getFloat(left), getFloat(right))
		e.lookup(n).valType = boolean
	}

	if n.Token.Type == lexer.T_IF {

		exp := e.result(n.Children[0])
		body := n.Children[1:]

		fmt.Printf("exp: %v\n", exp.val)

		if exp.val.(bool) {
			for _, b := range body {
				e.result(b)
			}
		}
	}

	if n.Token.Type == lexer.T_FOR {
		body := n.Children[1:]

		for e.result(n.Children[0]).val.(bool) {
			for _, b := range body {
				e.result(b)
			}
		}
	}

	if n.Token.Type == lexer.T_FUNC_CALL {

		var argValues []*value

		for _, arg := range n.Children[1:] {
			argValues = append(argValues, e.result(arg))
		}

		fNode := n.Children[0]
		fArgs := fNode.Children[0]
		fBody := fNode.Children[1:]

		for i, arg := range fArgs.Children {
			v := e.lookup(arg)
			v.val = argValues[i].val
			v.valType = argValues[i].valType
		}

		for _, b := range fBody {
			e.result(b)
		}
	}

	return e.lookup(n)
}

func getFloat(v *value) float64 {

	// fmt.Println("getFloat", v.token)

	if v.token.Type == lexer.T_NUM {
		return v.val.(float64)
	}

	if v.valType == num {
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
