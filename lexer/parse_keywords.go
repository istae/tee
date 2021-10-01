package lexer

import "strings"

func (l *lexer) parseKeyword() *Token {

	pos := l.pos

	// newline
	if l.current() == '\n' {
		l.next()
		return &Token{Type: T_NEWLINE, Str: "\n"}
	}

	// comment
	if l.current() == '/' && l.canPeek() && l.peek() == '/' {
		for {
			if l.next() {
				break
			}

			if l.current() == '\n' {
				l.next()
				return &Token{Type: T_COMMENT, Str: "//"}
			}
		}

		l.pos = pos
		return nil
	}

	if l.current() == '{' {
		l.next()
		return &Token{Type: T_OPEN_BRACKET, Str: "{"}
	}

	if l.current() == '}' {
		l.next()
		return &Token{Type: T_CLOSE_BRACKET, Str: "}"}
	}

	if strings.HasPrefix(l.str[l.pos:], "if ") {
		l.nextN(3)
		return &Token{Type: T_IF, Str: "if "}
	}

	if strings.HasPrefix(l.str[l.pos:], "for ") {
		l.nextN(4)
		return &Token{Type: T_FOR, Str: "for "}
	}

	if strings.HasPrefix(l.str[l.pos:], "func ") {
		l.nextN(5)
		return &Token{Type: T_FUNC, Str: "func "}
	}

	if l.current() == '(' {
		l.next()
		return &Token{Type: T_OPEN_PARS, Str: "("}
	}

	if l.current() == ')' {
		l.next()
		return &Token{Type: T_CLOSE_PARS, Str: ")"}
	}

	if l.current() == ',' {
		l.next()
		return &Token{Type: T_COMMA, Str: ","}
	}

	l.pos = pos
	return nil
}
