package lexer

func (l *lexer) parseOps() *Token {

	switch l.current() {
	case '+', '-', '*', '/':
		{
			defer l.next()
			return &Token{Type: T_MATH_OPS}
		}
	case '=':
		{
			if l.canPeek() && l.peek() == '=' {
				defer l.nextN(2)
				return &Token{Type: T_CMP_OPS}
			}
			defer l.next()
			return &Token{Type: T_EQUAL}
		}
	case '<', '>':
		{
			defer l.next()
			return &Token{Type: T_CMP_OPS}
		}
	}

	return nil
}
