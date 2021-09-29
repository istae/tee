package lexer

func (l *lexer) parseOps() *Token {

	switch l.current() {
	case '+', '-', '*', '/':
		{
			defer l.next()
			return &Token{Type: T_MATH_OPS, Start: l.pos, End: l.pos + 1, Str: l.str[l.pos : l.pos+1]}
		}
	case '=':
		{
			if l.canPeek() && l.peek() == '=' {
				return &Token{Type: T_CMP_OPS, Start: l.pos, End: l.pos + 2, Str: "=="}
			}
			defer l.next()
			return &Token{Type: T_EQUAL, Start: l.pos, End: l.pos + 1, Str: "="}
		}
	case '<', '>':
		{
			defer l.next()
			return &Token{Type: T_CMP_OPS, Start: l.pos, End: l.pos + 1, Str: l.str[l.pos : l.pos+1]}
		}
	}

	return nil
}
