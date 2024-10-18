package main

import (
	//"fmt"
	"zwaLang/src/lexer"
	//"zwaLang/src/token"
	"zwaLang/src/parser"
	"zwaLang/src/interpreter"
)

// -- Main --

func main() {
	testInputString := `x: number = 10
						y: number = 20
						t: number = 20
						z = x + y + t
						show z
						t = 100
						show t
						x = x * y
						show x
						x = x - y
						show x
						t = t / y
						show t
						t = (3 + 2) * 5
						show t`
	//fmt.Println(testInputString)

	l := lexer.NewLexer(testInputString)
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
