package tokens

type TokenType string
type TokenRole string

type Token struct {
	Type    TokenType
	Literal string
	Role    TokenRole
}

func New(tokenType TokenType, literal string, role TokenRole) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
		Role:    role,
	}
}

//Token roles
const (
	FUNCTION = "FUNCTION"
	ARGUMENT = "ARGUMENT"
)

//Token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	NONE    = "NONE"

	INT     = "INT"
	POP     = "."
	POP_ALL = "A"

	PLUSS   = "+"
	MINUS   = "-"
	APPEND  = "]"
	REDUCE  = "/"
	SWAP    = "S"
	COMPOSE = "o"
	DELETE  = "D"
	COPY    = "C"
	MAP     = "-<"
	TERNARY = "?"

	NOT                   = "!"
	EQUAL                 = "="
	NOT_EQUAL             = "!="
	GREATER_THEN          = ">"
	EQUAL_OR_GREATER_THEN = ">="
	LESS_THEN             = "<"
	EQUAL_OR_LESS_THEN    = "<="

	FUNCTION_DEF_START = "("
	FUNCTION_DEF_END   = ")"
)
