package lexer

import (
	"errors"
	"fmt"
)

var errSyntax = errors.New("syntax Error")

type tokenizerFunc func() TokenType

type lexer struct {
	line int
	pos  int
	end  int
	str  string
}

type Token struct {
	Type  TokenType
	Start int
	End   int
	Line  int
	Str   string
}

func NewLexer() *lexer {
	return &lexer{line: 1}

}

func (l *lexer) Read(s string) (tokens []Token, err error) {

	l.end = len(s)
	l.str = s

	tokenizers := []tokenizerFunc{
		l.parseKeyword,
		l.parseNum,
		l.parseString,
		l.parseSymbol,
		l.parseOps,
	}

	for {

		l.skipWhiteSpace()

		if l.done() {
			break
		}

		start := l.pos
		var tokenType TokenType

		for _, tr := range tokenizers {

			if tokenType = tr(); tokenType == T_UNKNOWN {
				l.pos = start
				continue
			}

			tokens = append(tokens, Token{
				Type:  tokenType,
				Start: start,
				End:   l.pos,
				Line:  l.line,
				Str:   l.str[start:l.pos],
			})

			if tokenType == T_NEWLINE {
				l.line++
			}

			fmt.Printf("token: %v\n", tokenType)

			break
		}

		if tokenType == T_UNKNOWN {
			err = fmt.Errorf("at line %d: %w", l.line, errSyntax)
			return
		}
	}

	return
}

func (l *lexer) done() bool {
	return l.pos >= l.end
}

func (l *lexer) nextN(n int) bool {
	l.pos += n
	return l.done()
}

func (l *lexer) next() bool {
	return l.nextN(1)
}

func (l *lexer) peekN(n int) byte {

	if (l.pos + n) < l.end {
		return l.str[l.pos+n]
	}

	return 0
}

func (l *lexer) peek() byte {
	return l.peekN(1)
}

func (l *lexer) current() byte {
	b := l.str[l.pos]
	return b
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isWhiteSpace(b byte, nl bool) bool {
	switch b {
	case '\t', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	case '\n':
		if nl {
			return true
		}
	}

	return false
}

func (l *lexer) skipWhiteSpace() {
	for ; l.pos < l.end; l.pos++ {
		if !isWhiteSpace(l.str[l.pos], false) {
			break
		}
	}
}
