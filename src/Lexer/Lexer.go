package Lexer

import (
	"fmt"
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
	//l.readChar()
	return l;
}

func main() {
	testInputString := `x: number = 10
						y: number = 20
						z = x + y
						show z`
	
	fmt.Println(testInputString);
}