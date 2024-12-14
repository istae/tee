package lexer

func (l *lexer) parseString() TokenType {

	// "*"
	if l.current() == '"' {

		for {
			if l.next() {
				break
			}

			if l.current() == '"' {
				l.next()
				return T_STRING
			}
		}
	}

	return T_UNKNOWN
}
