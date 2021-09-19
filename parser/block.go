package parser

import "tee/lexer"

type block struct {
	// children []*block
	symbols map[string]*node
	// parent   *block
	root *node
}

func (b *block) lookup(t lexer.Token) *node {
	n := b.symbols[t.Str]
	if n == nil {
		n = &node{token: t}
		b.symbols[t.Str] = n
	}
	return n
}

func newBlock() *block {

	return &block{
		symbols: map[string]*node{},
		root:    &node{},
	}
}
