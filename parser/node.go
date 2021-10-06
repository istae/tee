package parser

import (
	"fmt"
	"tee/lexer"
)

var (
	precendence = map[string]int{
		"/": 5, "*": 5,
		"+": 4, "-": 4,
		">": 3, "<": 3,
	}
)

type NodeType string

const (
	T_NUM      = "T_NUM"
	T_SYMBOL   = "T_SYMBOL"
	T_STRING   = "T_STRING"
	T_ASSIGN   = "T_ASSIGN"
	T_MATH_OPS = "T_MATH_OPS"
	T_CMP_OPS  = "T_CMP_OPS"
	T_FUNC_DEf = "T_FUNC_DEf"

	T_FOR     = "T_FOR"
	T_IF      = "T_IF"
	T_ELSE    = "T_ELSE"
	T_NEWLINE = "T_NEWLINE"
	T_COMMENT = "T_COMMENT"
)

type Node struct {
	Children []*Node
	parent   *Node
	Token    lexer.Token
	// Type  NodeType
	block *Block
}

// if n has higher precedence
func (n *Node) PrecendenceCmp(x *Node) bool {

	xPre := precendence[x.Token.Str]
	if xPre == 0 {
		return false
	}

	nPre := precendence[n.Token.Str]
	if nPre == 0 {
		return false
	}

	fmt.Println(x.Token, n.Token)

	return nPre > xPre
}

func (n *Node) AddChildAndParent(children ...*Node) {
	for _, c := range children {
		c.parent = n
		n.Children = append(n.Children, c)
	}
}

func (n *Node) AddChild(children ...*Node) {
	n.Children = append(n.Children, children...)
}

func (n *Node) LeftChild(c *Node) {
	c.parent = n
	n.Children = append([]*Node{c}, n.Children...)
}

func (n *Node) PopChild() *Node {
	var c *Node
	c, n.Children = n.Children[0], n.Children[1:]
	return c
}
