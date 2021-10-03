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

type TokenType string

const (
	T_NUM           = "T_NUM"
	T_SYMBOL        = "T_SYMBOL"
	T_STRING        = "T_STRING"
	T_EQUAL         = "T_EQUAL"
	T_MATH_OPS      = "T_MATH_OPS"
	T_CMP_OPS       = "T_CMP_OPS"
	T_FUNC          = "T_FUNC"
	T_FOR           = "T_FOR"
	T_IF            = "T_IF"
	T_ELSE          = "T_ELSE"
	T_OPEN_BRACKET  = "T_OPEN_BRACKET"
	T_CLOSE_BRACKET = "T_CLOSE_BRACKET"
	T_OPEN_PARS     = "T_OPEN_PARS"
	T_CLOSE_PARS    = "T_CLOSE_PARS"
	T_NEWLINE       = "T_NEWLINE"
	T_COMMENT       = "T_COMMENT"
	T_COMMA         = "T_COMMA"

	// below are not produced by the lexer.
	// The parser reassigns tokens with the types below accordingly.
	// Normally, the parser nodes should have node types, but we don't
	T_FUNC_CALL   = "T_FUNC_CALL"
	T_FUNC_SYMBOL = "T_FUNC_SYMBOL"
)

type Token struct {
	Type  TokenType
	Start int
	End   int
	Line  int
	Str   string
}

func NewLexer() *lexer {

	l := &lexer{line: 1}

	l.parsers = []parserFunc{
		l.parseKeyword,
		l.parseNum,
		l.parseString,
		l.parseSymbol,
		l.parseOps,
	}

	return l

}

func (l *lexer) Read(s string) (tokens []Token, err error) {
	l.end = len(s)
	l.str = s

	for {

		l.skipWhiteSpace()

		if l.done() {
			break
		}

		pos := l.pos
		var t *Token

		for _, p := range l.parsers {
			if t = p(); t != nil {
				break
			}
		}

		if t == nil {
			err = fmt.Errorf("at line %d: %w", l.line, errSyntax)
			return
		} else {
			t.Start = pos
			t.End = l.pos
			t.Line = l.line
			t.Str = l.str[pos:l.pos]
			fmt.Printf("token: %v\n", t)
			tokens = append(tokens, *t)
			if t.Type == T_NEWLINE {
				l.line++
			}
		}
	}

	return
}

func (l *lexer) done() bool {
	return l.pos >= l.end
}

func (l *lexer) doneN(n int) bool {
	return (l.pos + n) >= l.end
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

func isWhiteSpace(b byte) bool {
	switch b {
	case '\t', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}

func (l *lexer) skipWhiteSpace() {
	for ; l.pos < l.end; l.pos++ {
		if !isWhiteSpace(l.str[l.pos]) {
			break
		}
	}
}
