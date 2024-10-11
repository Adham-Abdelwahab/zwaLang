package interpreter

import (
	"fmt"
	"zwaLang/src/parser"
)

// Interpreter structure with enviornment for variables
type Interpreter struct {
	env map[string]int
}

// New Interpreter creates an interpreter with an empty enviornment
func NewInterpreter() *Interpreter {
	return &Interpreter{env: make(map[string]int)}
}

// -- Reciever Functions --

// Eval walks the AST and executes each node
func (i *Interpreter) Eval(nodes []parser.Node) {
	for _, node := range nodes {
		i.evalNode(node)
	}
}

// evalNode handles each node type: declarations, assignments, and print statements
func (i *Interpreter) evalNode(node parser.Node) {
	switch n := node.(type) {
		case *parser.VarDeclaration:
			//i.evalVarDeclaration(n)
		case *parser.Assignment:
			//i.evalAssignment(n)
		case *parser.PrintStatement:
			//i.evalPrintStatement(n)
		default:
			fmt.Printf("Unknown node type: %T\n", n)
	}
}