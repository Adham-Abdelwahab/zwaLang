package lexer

import (
	"unicode"
	"zwaLang/src/token"
)

type Lexer struct {
	input 			string
	position 		int // current position in input (points to current char)
	readPosition 	int // reading position (after current char)
	ch 				byte // current char under examination
}

// NewLexer initializes a new lexer with the given input string
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l;
}

// -- Receiver Functions --

// Read the next character in the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0; // Indicates EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken generates the next token from the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
		case '=':
			tok = newToken(token.ASSIGN, l.ch)
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case ':':
			tok = newToken(token.COLON, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			if isLetter(l.ch) {
				identifier := l.readIdentifier()
				tok.Type = lookupIdentifier(identifier)
				tok.Literal = identifier
				return tok
			} else if isDigit(l.ch) {
				tok.Type = token.NUMBER
				tok.Literal = l.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}

	l.readChar()
	return tok
}

// skipWhitespace skips over whitespace characters and advances the lexer's position if necessary
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar() // Read the next character
	}
}

// readNumber reads a number from the input and returns it as a string
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// readIdentifier reads an identifier from the input and returns it as a string
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// -- Helper Functions --

// newToken creates a new token with the given type and literal
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter checks if the given character is a letter
func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

// isDigit checks if the given character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// lookupIdentifier checks if the given identifier is a keyword (i.e. show) and returns the corresponding token type
func lookupIdentifier(identifier string) token.TokenType {
	switch identifier {
		case "show":
			return token.SHOW
		case "number":
			return token.NATURAL_NUMBER_TYPE
		default:
			return token.IDENT
	}
}