package lexer

import (
	"plusLang/tokens"
	"unicode"
)

type Lexer struct {
	input        []rune
	readPosition int
}

func New(input string) *Lexer {
	return &Lexer{input: []rune(input)}
}

func (l *Lexer) NextToken() tokens.Token {
	l.JumpWhiteSpace()

	if l.readPosition >= len(l.input) {
		return l.newToken(tokens.EOF, "", tokens.EOF)
	}

	var tok tokens.Token

	switch true {
	case l.tokenIs(tokens.TO_STRING):
		tok = l.newToken(tokens.TO_STRING, tokens.TO_STRING, tokens.FUNCTION)
	case l.tokenIs(tokens.TO_NUM):
		tok = l.newToken(tokens.TO_NUM, tokens.TO_NUM, tokens.FUNCTION)
	case l.tokenIs(tokens.MAP):
		tok = l.newToken(tokens.MAP, tokens.MAP, tokens.FUNCTION)
	case l.tokenIs(tokens.PLUS):
		tok = l.newToken(tokens.PLUS, tokens.PLUS, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.MINUS):
		tok = l.newToken(tokens.MINUS, tokens.MINUS, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.MULT):
		tok = l.newToken(tokens.MULT, tokens.MULT, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.DIV):
		tok = l.newToken(tokens.DIV, tokens.DIV, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.EQUAL_OR_GREATER_THEN):
		tok = l.newToken(tokens.EQUAL_OR_GREATER_THEN, tokens.EQUAL_OR_GREATER_THEN, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.EQUAL_OR_LESS_THEN):
		tok = l.newToken(tokens.EQUAL_OR_LESS_THEN, tokens.EQUAL_OR_LESS_THEN, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.LESS_THEN):
		tok = l.newToken(tokens.LESS_THEN, tokens.LESS_THEN, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.GREATER_THEN):
		tok = l.newToken(tokens.GREATER_THEN, tokens.GREATER_THEN, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.NOT):
		tok = l.newToken(tokens.NOT, tokens.NOT, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.EQUAL):
		tok = l.newToken(tokens.EQUAL, tokens.EQUAL, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.APPEND):
		tok = l.newToken(tokens.APPEND, tokens.APPEND, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.REDUCE):
		tok = l.newToken(tokens.REDUCE, tokens.REDUCE, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.SWAP):
		tok = l.newToken(tokens.SWAP, tokens.SWAP, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.COMPOSE):
		tok = l.newToken(tokens.COMPOSE, tokens.COMPOSE, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.DELETE):
		tok = l.newToken(tokens.DELETE, tokens.DELETE, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.TERNARY):
		tok = l.newToken(tokens.TERNARY, tokens.TERNARY, tokens.FUNCTION)
		break
	case l.tokenIs(tokens.POP):
		tok = l.newToken(tokens.POP, tokens.POP, tokens.ARGUMENT)
		break
	case l.tokenIs(tokens.POP_ALL):
		tok = l.newToken(tokens.POP_ALL, tokens.POP_ALL, tokens.ARGUMENT)
		break
	case l.tokenIs("\""):
		tok = l.readString()
	case l.tokenIsInt():
		tok = l.readNumber()
		break
	case l.tokenIs(tokens.FUNCTION_DEF_START):
		tok = l.readFunction()
	default:
		tok = l.newToken(tokens.ILLEGAL, string(l.getCurChar()), tokens.ILLEGAL)
	}

	return tok
}

func (l *Lexer) newToken(tokenType tokens.TokenType, ch string, tokenRole tokens.TokenRole) tokens.Token {
	l.readPosition += len(ch)
	return tokens.New(tokenType, ch, tokenRole)
}

func (l *Lexer) tokenIs(token string) bool {
	for i := 0; i < len(token); i++ {
		if i+l.readPosition >= len(l.input) {
			return false
		}
		if []rune(token)[i] != l.input[i+l.readPosition] {
			return false
		}
	}

	return true
}

func (l *Lexer) getCurChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) tokenIsInt() bool {
	return unicode.IsDigit(l.getCurChar())
}

func (l *Lexer) readString() tokens.Token {
	l.readPosition++
	stringLiteral := ""

	for {
		curChar := l.getCurChar()
		l.readPosition++

		if curChar == 0 || curChar == '"' {
			break
		}

		stringLiteral += string(curChar)
	}

	return tokens.New(tokens.STRING, stringLiteral, tokens.ARGUMENT)
}

func (l *Lexer) readNumber() tokens.Token {
	numberLiteral := ""
	numberType := tokens.INT

	for {
		if l.getCurChar() == ',' {
			l.readPosition++
			if numberType == tokens.FLOAT {
				continue
			}

			numberLiteral += "."
			numberType = tokens.FLOAT
			continue
		}

		if !unicode.IsDigit(l.getCurChar()) {
			break
		}

		numberLiteral += string(l.getCurChar())
		l.readPosition++
	}

	return tokens.New(tokens.TokenType(numberType), numberLiteral, tokens.ARGUMENT)
}

func (l *Lexer) readFunction() tokens.Token {
	l.readPosition++
	position := l.readPosition
	for l.getCurChar() != []rune(tokens.FUNCTION_DEF_END)[0] {
		l.readPosition++
	}

	tok := tokens.New(tokens.FUNCTION, string(l.input[position:l.readPosition]), tokens.ARGUMENT)
	l.readPosition++
	return tok
}

func (l *Lexer) JumpWhiteSpace() {
	for unicode.IsSpace(l.getCurChar()) {
		l.readPosition++
	}
}
