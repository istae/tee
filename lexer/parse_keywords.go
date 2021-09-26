package lexer

func (l *lexer) parseKeyword() *Token {

	pos := l.pos

	// newline
	if l.current() == '\n' {
		l.next()
		return &Token{Type: T_NEWLINE, Start: pos, End: l.pos, Str: "\n"}
	}

	// comment
	if l.current() == '/' && l.canPeek() && l.peek() == '/' {
		for {
			if l.next() {
				break
			}

			if l.current() == '\n' {
				l.next()
				return &Token{Type: T_COMMENT, Start: pos, End: l.pos, Str: "//"}
			}
		}

		l.pos = pos
		return nil
	}

	if l.current() == '{' {
		l.next()
		return &Token{Type: T_OPEN_BRACKET, Start: pos, End: l.pos, Str: "{"}
	}

	if l.current() == '}' {
		l.next()
		return &Token{Type: T_CLOSE_BRACKET, Start: pos, End: l.pos, Str: "}"}
	}

	if l.current() == 'i' && l.canPeek() && l.peek() == 'f' {
		l.nextN(2)

		if !l.done() && isWhiteSpace(l.current()) {
			l.next()
			return &Token{Type: T_IF, Start: pos, End: l.pos, Str: "if"}
		}
	}

	// if strings.HasPrefix(l.str[l.pos:], "for") {
	// 	l.nextN(3)
	// 	if l.canPeek() && isSpace(l.peek()) {
	// 		return &token{Type: T_FOR, str: "for"}
	// 	}
	// }

	l.pos = pos
	return nil
}
