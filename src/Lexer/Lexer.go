package main

import (
	"fmt"
	"unicode"
)

// Token types
type TokenType string

const (
	NUMBER 			TokenType = "NUMBER"					// Number	
	NATURAL_NUMBER_TYPE 	TokenType = "NATURAL_NUMBER_TYPE" 		// Natural Number Type
	IDENT 			TokenType = "IDENT"						// Identifiers
	ASSIGN 			TokenType = "ASSIGN"					// =
	PLUS 			TokenType = "PLUS"						// +
	COLON 			TokenType = "COLON"						// :
	SHOW 			TokenType = "SHOW"						// show
	EOF 			TokenType = "EOF"						// End of File
	ILLEGAL 		TokenType = "ILLEGAL"					// Illegal token
)

// Token structure
type Token struct {
	Type 	TokenType
	Literal string
}

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
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
		case '=':
			tok = newToken(ASSIGN, l.ch)
		case '+':
			tok = newToken(PLUS, l.ch)
		case ':':
			tok = newToken(COLON, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = EOF
		default:
			if isLetter(l.ch) {
				identifier := l.readIdentifier()
				tok.Type = lookupIdentifier(identifier)
				tok.Literal = identifier
				return tok
			} else if isDigit(l.ch) {
				tok.Type = NUMBER
				tok.Literal = l.readNumber()
				return tok
			} else {
				tok = newToken(ILLEGAL, l.ch)
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
func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
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
func lookupIdentifier(identifier string) TokenType {
	switch identifier {
		case "show":
			return SHOW
		case "number":
			return NATURAL_NUMBER_TYPE
		default:
			return IDENT
	}
}

// -- Main --

func main() {
	testInputString := `x: number = 10
						y: number = 20
						z = x + y
						show z`
	fmt.Println(testInputString)

	l := NewLexer(testInputString)

	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}