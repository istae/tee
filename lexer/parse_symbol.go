package lexer

// [a-b, A-b]+
func (l *lexer) parseSymbol() TokenType {

	if !isAlpha(l.current()) {
		return T_UNKNOWN
	}

	for {
		if l.next() {
			break
		}

		if !isAlpha(l.current()) {
			break
		}
	}

	return T_SYMBOL
}
