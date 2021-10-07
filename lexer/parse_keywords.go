package lexer

import "strings"

func (l *lexer) parseKeyword() *Token {

	pos := l.pos

	// newline
	if l.current() == '\n' {
		l.next()
		return &Token{Type: T_NEWLINE}
	}

	// comment
	if l.current() == '/' && l.canPeek() && l.peek() == '/' {
		for {
			if l.next() {
				break
			}

			if l.current() == '\n' {
				l.next()
				return &Token{Type: T_COMMENT}
			}
		}

		l.pos = pos
		return nil
	}

	if l.current() == '{' {
		l.next()
		return &Token{Type: T_OPEN_BRACKET}
	}

	if l.current() == '}' {
		l.next()
		return &Token{Type: T_CLOSE_BRACKET}
	}

	if strings.HasPrefix(l.str[l.pos:], "if") {
		l.nextN(3)
		return &Token{Type: T_IF}
	}

	if strings.HasPrefix(l.str[l.pos:], "for") {
		l.nextN(3)
		if isWhiteSpace(l.current(), false) {
			return &Token{Type: T_FOR}
		}
		l.pos = pos
	}

	if strings.HasPrefix(l.str[l.pos:], "func") {
		l.nextN(4)
		if isWhiteSpace(l.current(), false) {
			return &Token{Type: T_FUNC}
		}
		l.pos = pos
	}

	if strings.HasPrefix(l.str[l.pos:], "break") {
		l.nextN(5)
		if isWhiteSpace(l.current(), true) {
			return &Token{Type: T_BREAK}
		}
		l.pos = pos
	}

	if l.current() == '(' {
		l.next()
		return &Token{Type: T_OPEN_PARS}
	}

	if l.current() == ')' {
		l.next()
		return &Token{Type: T_CLOSE_PARS}
	}

	if l.current() == ',' {
		l.next()
		return &Token{Type: T_COMMA}
	}

	l.pos = pos
	return nil
}
