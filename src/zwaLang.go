package main

import (
	"fmt"
	"zwaLang/src/lexer"
	//"zwaLang/src/token"
	"zwaLang/src/parser"
	"zwaLang/src/interpreter"
	"os"
	"path/filepath"
)

// -- Main --
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide a file name")
		os.Exit(1)
	}
	if filepath.Ext(os.Args[1]) != ".zwa" {
		fmt.Println("File must have a .zwa extension")
		os.Exit(1)
	}
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}
	contentString := string(content)

	l := lexer.NewLexer(contentString)
	//for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
	// 	fmt.Printf("%+v\n", tok)
	//}

	p := parser.NewParser(l)
	//fmt.Printf("%+v\n", p)
	ast := p.ParseProgram()
	//fmt.Printf("%+v\n", ast)

	// for _, node := range ast {
	// 	fmt.Printf("%+v\n", node)
	// }
	i := interpreter.NewInterpreter()
	i.Eval(ast)
}
