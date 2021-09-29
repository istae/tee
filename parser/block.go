package parser

import "tee/lexer"

type block struct {
	children []*block
	symbols  map[string]*node
	parent   *block
	nodes    []*node
}

func (b *block) lookup(t lexer.Token) *node {

	var n *node
	look := b

	for {

		if look == nil {
			break
		}

		n = look.symbols[t.Str]
		if n == nil {
			look = b.parent
		} else {
			break
		}
	}

	if n == nil {
		n = &node{token: t}
		b.symbols[t.Str] = n
	}

	return n
}

func (b *block) AddNode(n *node) {
	b.nodes = append(b.nodes, n)
	n.block = b
}

func (b *block) AddChild(child *block) {
	child.parent = b
	b.children = append(b.children, child)
}

func newBlock() *block {

	return &block{
		symbols: map[string]*node{},
	}
}