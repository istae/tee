package lexer

// [a-b, A-b]+
func (l *lexer) parseVar() *Token {

	pos := l.pos

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

	return &Token{Type: T_VAR, Start: pos, End: l.pos, Str: l.str[pos:l.pos]}
}
