package token

// Token types
type TokenType string

const (
	NUMBER 					TokenType = "NUMBER"					// Number	
	NATURAL_NUMBER_TYPE 	TokenType = "NATURAL_NUMBER_TYPE" 		// Natural Number Type
	IDENT 					TokenType = "IDENT"						// Identifiers
	ASSIGN 					TokenType = "ASSIGN"					// =
	PLUS 					TokenType = "PLUS"						// +
	COLON 					TokenType = "COLON"						// :
	SHOW 					TokenType = "SHOW"						// show
	EOF 					TokenType = "EOF"						// End of File
	ILLEGAL 				TokenType = "ILLEGAL"					// Illegal token
)

// Token structure
type Token struct {
	Type 	TokenType
	Literal string
}