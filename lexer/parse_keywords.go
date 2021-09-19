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

	// if strings.HasPrefix(l.str[l.pos:], "for") {
	// 	l.nextN(3)
	// 	if l.canPeek() && isSpace(l.peek()) {
	// 		return &token{Type: T_FOR, str: "for"}
	// 	}
	// }

	l.pos = pos
	return nil
}
