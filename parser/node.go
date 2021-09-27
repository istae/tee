package parser

import "tee/lexer"

var (
	precendence = map[string]int{"/": 5, "*": 5, "+": 4, "-": 4}
)

type node struct {
	children []*node
	parent   *node
	token    lexer.Token
	block    *block
}

// if n has higher precedence
func (n *node) PrecendenceCmp(x *node) bool {

	xPre := precendence[x.token.Str]
	if xPre == 0 {
		return false
	}

	nPre := precendence[n.token.Str]
	if nPre == 0 {
		return false
	}

	return nPre > xPre
}

func (n *node) AddChild(children ...*node) {
	for _, c := range children {
		c.parent = n
		n.children = append(n.children, c)
	}
}

func (n *node) LeftChild(c *node) {
	c.parent = n
	n.children = append([]*node{c}, n.children...)
}

func (n *node) PopChild() *node {
	var c *node
	c, n.children = n.children[0], n.children[1:]
	return c
}
