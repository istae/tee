package lexer

// [a-b, A-b]+
func (l *lexer) parseSymbol() *Token {

	if !isAlpha(l.current()) {
		return nil
	}

	for {
		if l.next() {
			break
		}

		if !isAlpha(l.current()) {
			break
		}
	}

	return &Token{Type: T_SYMBOL}
}
