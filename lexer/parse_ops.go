package lexer

func (l *lexer) parseOps() TokenType {

	switch l.current() {
	case '+', '-', '*', '/':
		{
			l.next()
			return T_MATH_OPS
		}
	case '=':
		{
			if l.peek() == '=' {
				l.nextN(2)
				return T_CMP_OPS
			}
			l.next()
			return T_EQUAL
		}
	case '<', '>':
		{
			l.next()
			return T_CMP_OPS
		}
	}

	return T_UNKNOWN
}
