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
			i.evalVarDeclaration(n)
		case *parser.Assignment:
			i.evalAssignment(n)
		case *parser.PrintStatement:
			i.evalPrintStatement(n)
		default:
			fmt.Printf("Unknown node type: %T\n", n)
	}
}

func (i *Interpreter) evalVarDeclaration(node *parser.VarDeclaration) {
	// Store the variable in the enviornment
	i.env[node.Name] = node.Value
}

func (i *Interpreter) evalAssignment(node *parser.Assignment) {
	leftVal, ok1 := i.env[node.Left]
	rightVal, ok2 := i.env[node.Right]

	if !ok1 || !ok2 {
		fmt.Printf("Variable not found: %s or %s\n", node.Left, node.Right)
		return
	}

	i.env[node.VarName] = leftVal + rightVal
}

func (i *Interpreter) evalPrintStatement(node *parser.PrintStatement) {
	val, ok := i.env[node.VarName]
	if !ok {
		fmt.Printf("Variable not found: %s\n", node.VarName)
		return
	}
	fmt.Println(val)
}