package lexer

func (l *lexer) parseString() *Token {

	pos := l.pos

	// "*"
	if l.current() == '"' {

		for {
			if l.next() {
				break
			}

			if l.current() == '"' {
				l.next()
				return &Token{Type: T_STRING, Start: pos, End: l.pos, Str: l.str[pos:l.pos]}
			}
		}
	}

	l.pos = pos
	return nil
}
