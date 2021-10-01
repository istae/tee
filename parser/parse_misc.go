package parser

import "tee/lexer"

func (p *parser) parseNewline(b *Block) *Node {

	if p.current().Type == lexer.T_NEWLINE {
		defer p.next()
		return &Node{
			Token: p.current(),
		}
	}

	if p.current().Type == lexer.T_COMMENT {
		defer p.next()
		return &Node{
			Token: p.current(),
		}
	}

	return nil
}
