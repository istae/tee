package lexer

// [0-9]+(.[0-9])?
func (l *lexer) parseNum() *Token {

	pos := l.pos

	if !isDigit(l.current()) {
		return nil
	}

	for {
		if l.next() {
			break
		}

		if !isDigit(l.current()) {
			break
		}
	}

	if l.current() == '.' && l.canPeek() && isDigit(l.peek()) {

		for {
			if l.next() {
				break
			}
			if !isDigit(l.current()) {
				break
			}
		}

		// if l.done() || !(isSpace(l.current()) || l.current() == '\n') {
		// 	l.pos = pos
		// 	return nil
		// }
		return &Token{Type: T_DOUBLE, Start: pos, End: l.pos, Str: l.str[pos:l.pos]}

	}

	// if l.done() || !(isSpace(l.current()) || l.current() == '\n') {
	// 	l.pos = pos
	// 	return nil
	// }

	return &Token{Type: T_INT, Start: pos, End: l.pos, Str: l.str[pos:l.pos]}
}
