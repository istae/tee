package lexer

import "strings"

func (l *lexer) parseKeyword() TokenType {

	// newline
	if l.current() == '\n' {
		l.next()
		return T_NEWLINE
	}

	// comment
	if l.current() == '/' && l.peek() == '/' {
		for {
			if l.next() {
				break
			}

			if l.current() == '\n' {
				l.next()
				return T_COMMENT
			}
		}

		return T_UNKNOWN
	}

	if l.current() == '{' {
		l.next()
		return T_OPEN_BRACKET
	}

	if l.current() == '}' {
		l.next()
		return T_CLOSE_BRACKET
	}

	if strings.HasPrefix(l.str[l.pos:], "if") {
		l.nextN(3)
		return T_IF
	}

	if strings.HasPrefix(l.str[l.pos:], "for") {
		if isWhiteSpace(l.peekN(3), false) {
			l.nextN(3)
			return T_FOR
		}
	}

	if strings.HasPrefix(l.str[l.pos:], "func") {
		if isWhiteSpace(l.peekN(4), false) {
			l.nextN(4)
			return T_FUNC
		}
	}

	if strings.HasPrefix(l.str[l.pos:], "break") {
		if isWhiteSpace(l.peekN(5), true) {
			l.nextN(5)
			return T_BREAK
		}
	}

	if l.current() == '(' {
		l.next()
		return T_OPEN_PARS
	}

	if l.current() == ')' {
		l.next()
		return T_CLOSE_PARS
	}

	if l.current() == ',' {
		l.next()
		return T_COMMA
	}

	return T_UNKNOWN

}
