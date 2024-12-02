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

	AND TokenType = "AND" // &
	OR  TokenType = "OR"  // |

	// Expressions
	NUMBER TokenType = "NUMBER" // Number
	STRING TokenType = "STRING" // String literal
	BOOL   TokenType = "BOOL"
	LPAREN TokenType = "LPAREN" // (
	RPAREN TokenType = "RPAREN" // )

	// Misc
	EOF     TokenType = "EOF"     // End of File
	ILLEGAL TokenType = "ILLEGAL" // Illegal token
)

// Token structure
type token struct {
	Type    TokenType
	Literal string
}
