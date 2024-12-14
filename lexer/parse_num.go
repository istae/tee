package lexer

// [0-9]+(.[0-9])?
func (l *lexer) parseNum() TokenType {

	if !isDigit(l.current()) {
		return T_UNKNOWN
	}

	for {
		if l.next() {
			break
		}

		if !isDigit(l.current()) {
			break
		}
	}

	if !l.done() && l.current() == '.' && isDigit(l.peek()) {

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
		return T_NUM

	}

	// if l.done() || !(isSpace(l.current()) || l.current() == '\n') {
	// 	l.pos = pos
	// 	return nil
	// }

	return T_NUM
}
