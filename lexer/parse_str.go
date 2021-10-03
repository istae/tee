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
				return &Token{Type: T_STRING}
			}
		}
	}

	l.pos = pos
	return nil
}
