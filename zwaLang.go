package main

import (
	"fmt"
	"os"
	"path/filepath"
	compiler "zwa/bin"
)

// -- Main --
func main() {
	if len(os.Args) < 2 {
		help(os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "exec":
		if len(os.Args) < 3 {
			fmt.Println("Missing parameter, provide at least one file name")
			os.Exit(1)
		} else {
			for _, filename := range os.Args[2:] {
				if filepath.Ext(filename) != ".zwa" {
					fmt.Println("Failed to compile", filename, ": file must have a .zwa extension")
				} else {
					fmt.Println("Output of", filename)
					compile(filename)
				}
			}
		}

	default:
		help(os.Args[0])
	}
}

func help(name string) {
	fmt.Println("usage:", name, "[exec | help] <file name>")
}

func compile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filename, err)
		os.Exit(1)
	}

	l := compiler.NewLexer(string(content))
	//for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
	// 	fmt.Printf("%+v\n", tok)
	//}

	p := compiler.NewParser(l)
	//fmt.Printf("%+v\n", p)
	ast := p.ParseProgram()
	//fmt.Printf("%+v\n", ast)

	// for _, node := range ast {
	// 	fmt.Printf("%+v\n", node)
	// }
	i := compiler.NewInterpreter()
	i.Eval(ast)
}
