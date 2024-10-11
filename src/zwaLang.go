package main

import (
	"fmt"
	"zwaLang/src/lexer"
	"zwaLang/src/token"
)

// -- Main --

func main() {
	testInputString := `x: number = 10
						y: number = 20
						z = x + y
						show z`
	fmt.Println(testInputString)

	l := lexer.NewLexer(testInputString)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
