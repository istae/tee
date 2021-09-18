package lexer

import (
	"errors"
	"fmt"
)

var errSyntax = errors.New("syntax Error")

type parserFunc func() *Token

type lexer struct {
	line int
	pos  int
	end  int
	// tokens []token
	str string
	// c      byte
	parsers []parserFunc
}

type TokenType int

const (
	T_INT = iota
	T_DOUBLE
	T_EQUAL
	T_VAR
	T_OPS
	T_FOR
	T_IF
	T_ELSE
	T_OPEN_BRACKET
	T_CLOSE_BRACKET
	T_NEWLINE
	T_COMMENT
)

type Token struct {
	Type  TokenType
	Start int
	End   int
}

func (t Token) Str(s string) string {
	return s[t.Start:t.End]
}

func NewLexer() *lexer {

	l := &lexer{}

	l.parsers = []parserFunc{l.parseKeyword, l.parseNum, l.parseVar, l.parseOps}

	return l

}

func (l *lexer) Read(s string) (tokens []Token, err error) {
	l.end = len(s)
	l.str = s

	for {

		l.skipSpace()

		if l.done() {
			break
		}

		t := func() *Token {
			for _, p := range l.parsers {
				if t := p(); t != nil {
					return t
				}
			}
			return nil
		}()

		if t == nil {
			err = fmt.Errorf("syntax error at line %d\n%s:%w", l.line, l.str[l.pos:], errSyntax)
			return
		} else {
			tokens = append(tokens, *t)
			fmt.Printf("token: %s\n", l.str[t.Start:t.End])
		}
	}

	return
}

func (l *lexer) done() bool {
	return l.pos >= l.end
}

func (l *lexer) next() bool {
	l.pos++
	return l.done()
}

func (l *lexer) nextN(n int) bool {
	l.pos += n
	return l.done()
}

func (l *lexer) canPeek() bool {
	return (l.pos + 1) < l.end
}

func (l *lexer) peek() byte {
	return l.str[l.pos+1]
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

func isSpace(b byte) bool {
	switch b {
	case '\t', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}

func (l *lexer) skipSpace() {
	for ; l.pos < l.end; l.pos++ {
		if !isSpace(l.str[l.pos]) {
			break
		}
	}
}
