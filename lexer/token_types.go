package lexer

type TokenType int

const (
	T_UNKNOWN TokenType = iota
	T_NUM
	T_SYMBOL
	T_STRING
	T_EQUAL
	T_MATH_OPS
	T_CMP_OPS
	T_FUNC
	T_BREAK
	T_FOR
	T_IF
	T_ELSE
	T_OPEN_BRACKET
	T_CLOSE_BRACKET
	T_OPEN_PARS
	T_CLOSE_PARS
	T_NEWLINE
	T_COMMENT
	T_COMMA

	// below are not produced by the lexer.
	// The parser reassigns tokens with the types below accordingly.
	// Normally, the parser nodes should have node types, but we don't
	T_FUNC_CALL
	T_FUNC_SYMBOL
)

func (t TokenType) String() string {
	switch t {
	case T_UNKNOWN:
		return "T_UNKNOWN"
	case T_NUM:
		return "T_NUM"
	case T_SYMBOL:
		return "T_SYMBOL"
	case T_STRING:
		return "T_STRING"
	case T_EQUAL:
		return "T_EQUAL"
	case T_MATH_OPS:
		return "T_MATH_OPS"
	case T_CMP_OPS:
		return "T_CMP_OPS"
	case T_FUNC:
		return "T_FUNC"
	case T_BREAK:
		return "T_BREAK"
	case T_FOR:
		return "T_FOR"
	case T_IF:
		return "T_IF"
	case T_ELSE:
		return "T_ELSE"
	case T_OPEN_BRACKET:
		return "T_OPEN_BRACKET"
	case T_CLOSE_BRACKET:
		return "T_CLOSE_BRACKET"
	case T_OPEN_PARS:
		return "T_OPEN_PARS"
	case T_CLOSE_PARS:
		return "T_CLOSE_PARS"
	case T_NEWLINE:
		return "T_NEWLINE"
	case T_COMMENT:
		return "T_COMMENT"
	case T_COMMA:
		return "T_COMMA"
	case T_FUNC_CALL:
		return "T_FUNC_CALL"
	case T_FUNC_SYMBOL:
		return "T_FUNC_SYMBOL"
	default:
		return ""
	}
}
