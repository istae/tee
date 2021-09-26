package parser

import "tee/lexer"

func (p *parser) parseNewline(b *block) *node {

	if p.current().Type == lexer.T_NEWLINE {
		defer p.next()
		return &node{
			token: p.current(),
		}
	}

	return nil
}
