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
	result := i.evalExpression(node.Expression)
	// technically we should check if the variable is already in the enviornment
	i.env[node.VarName] = result
}

func (i *Interpreter) evalPrintStatement(node *parser.PrintStatement) {
	fmt.Println(i.evalExpression(node.Expression))
}

func (i *Interpreter) evalExpression(expr parser.Expression) int {
	switch e := expr.(type) {
		case *parser.NumberLiteral:
			return e.Value
		case *parser.Variable:
			val, ok := i.env[e.Name]
			if !ok {
				fmt.Printf("Variable not found: %s\n", e.Name)
			}
			return val
		case *parser.BinaryExpression:
			leftVal := i.evalExpression(e.Left)
			rightVal := i.evalExpression(e.Right)
			switch e.Operator {
				case "+":
					return leftVal + rightVal
				case "-":
					return leftVal - rightVal
				case "*":
					return leftVal * rightVal
				case "/":
					return leftVal / rightVal
				case "%":
					return leftVal % rightVal
				default:
					fmt.Printf("Unknown operator: %s\n", e.Operator)
					return 0
			}
		default:
			fmt.Printf("Unknown expression type: %T\n", e)
			return 0
	}	
}