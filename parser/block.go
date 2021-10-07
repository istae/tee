package parser

import (
	"tee/lexer"
)

type blockType string

const (
	forBlock  blockType = "for"
	ifBlock   blockType = "if"
	funcBlock blockType = "func"
)

type Block struct {
	symbols   map[string]*Node
	parent    *Block
	Nodes     []*Node
	blockType blockType
}

func (b *Block) getOrSetNode(t lexer.Token) *Node {

	var n *Node
	look := b

	for {

		if look == nil {
			break
		}

		n = look.symbols[t.Str]
		if n == nil {
			look = look.parent
		} else {
			break
		}
	}

	if n == nil {
		n = &Node{Token: t}
		b.symbols[t.Str] = n
	}

	return n
}

func (b *Block) getNode(t lexer.Token) *Node {

	var n *Node
	look := b

	for {

		if look == nil {
			break
		}

		n = look.symbols[t.Str]
		if n == nil {
			look = look.parent
		} else {
			break
		}
	}

	return n
}

func (b *Block) setNode(t lexer.Token) *Node {
	n := &Node{Token: t}
	b.symbols[t.Str] = n
	return n
}

func (b *Block) AddNode(n *Node) {
	b.Nodes = append(b.Nodes, n)
	n.block = b
}

func (b *Block) Parent(p *Block) {
	b.parent = p
}

func newBlock(t blockType) *Block {
	return &Block{
		symbols:   map[string]*Node{},
		blockType: t,
	}
}
