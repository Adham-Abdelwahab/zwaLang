package bin

import "zwa/utils"

type Tokens = utils.Queue[token]

type Lexer struct {
	utils.Queue[rune]
	tokens Tokens
}

func Lex(content string) Tokens {
	lexer := Lexer{utils.NewQueue[rune](), utils.NewQueue[token]()}
	for _, c := range content {
		lexer.Push(c)
	}
	lexer.Tokenize()
	return lexer.tokens
}

// skipWhitespace skips over whitespace characters and advances the lexer's position if necessary
func (lexer *Lexer) skipWhitespace() (rune, bool) {
	if c, ok := lexer.Pop(); ok {
		switch c {
		case ' ', '\t', '\n', '\r':
			return lexer.skipWhitespace()
		default:
			return c, true
		}
	} else {
		return c, false
	}
}

// isLetter checks if the given character is a letter
func isLetter(c rune) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

// readIdentifier reads an identifier from the input and returns it as a string
func (lexer *Lexer) readIdentifier(c rune) string {
	ident := string(c)
	peek, ok := lexer.Peek()
	if !ok || !isLetter(peek) {
		return ident
	}
	next, _ := lexer.Pop()
	return ident + lexer.readIdentifier(next)
}

// isDigit checks if the given character is a digit
func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

// readNumber reads a number from the input and returns it as a string
func (lexer *Lexer) readNumber(c rune) string {
	num := string(c)
	peek, ok := lexer.Peek()
	if !ok || !isDigit(peek) {
		return num
	}
	next, _ := lexer.Pop()
	return num + lexer.readNumber(next)
}

func (lexer *Lexer) readString() string {
	if next, ok := lexer.Pop(); ok {
		if next == '"' {
			return ""
		} else {
			return string(next) + lexer.readString()
		}
	} else {
		print("ERROR: String failed to terminate\n")
		return "ERROR"
	}
}

func (lexer *Lexer) Tokenize() Tokens {
	c, ok := lexer.skipWhitespace()
	if !ok {
		return lexer.tokens
	}

	var tok = token{Literal: string(c)}
	var Type TokenType
	switch c {
	case '=':
		Type = EQUALS
	case '+':
		Type = PLUS
	case '-':
		Type = MINUS
	case '*':
		Type = ASTERISK
	case '/':
		Type = SLASH
	case ':':
		Type = COLON
	case '(':
		Type = LPAREN
	case ')':
		Type = RPAREN
	case '%':
		Type = MODULO
	case '&':
		Type = AND
	case '|':
		Type = OR
	case '"':
		Type = STRING
		tok.Literal = lexer.readString()
	default:
		switch {
		case isLetter(c):
			identifier := lexer.readIdentifier(c)
			switch identifier {
			case "show":
				Type = SHOW
			case "true", "false":
				Type = BOOL
			default:
				Type = IDENT
			}
			tok.Literal = identifier
		case isDigit(c):
			Type = NUMBER
			tok.Literal = lexer.readNumber(c)
		default:
			Type = ILLEGAL
		}
	}

	tok.Type = Type
	lexer.tokens.Push(tok)
	return lexer.Tokenize()
}
