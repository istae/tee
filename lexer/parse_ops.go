package lexer

// =
func (l *lexer) parseOps() *Token {

	switch l.current() {
	case '+', '-', '*', '/':
		{
			t := &Token{Type: T_OPS, Start: l.pos, End: l.pos + 1}
			l.next()
			return t
		}
	default:
		return nil
	}
}
