package bin

// Token types
type TokenType string

const (
	IDENT TokenType = "IDENT" // Identifiers

	// Statements
	EQUALS TokenType = "EQUALS" // =
	COLON  TokenType = "COLON"  // :
	SHOW   TokenType = "SHOW"   // show

	// Operators
	PLUS     TokenType = "PLUS"     // +
	MINUS    TokenType = "MINUS"    // -
	ASTERISK TokenType = "ASTERISK" // *
	SLASH    TokenType = "SLASH"    // / {division}
	MODULO   TokenType = "MODULO"   // %

	// Expressions
	NUMBER              TokenType = "NUMBER"              // Number
	STRING              TokenType = "STRING"              // String literal
	NATURAL_NUMBER_TYPE TokenType = "NATURAL_NUMBER_TYPE" // Natural Number Type
	LPAREN              TokenType = "LPAREN"              // (
	RPAREN              TokenType = "RPAREN"              // )

	// Misc
	EOF     TokenType = "EOF"     // End of File
	ILLEGAL TokenType = "ILLEGAL" // Illegal token
)

// Token structure
type token struct {
	Type    TokenType
	Literal string
}
