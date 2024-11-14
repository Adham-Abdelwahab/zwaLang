package bin

import (
	"fmt"
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
func (i *Interpreter) Eval(nodes []Node) {
	for _, node := range nodes {
		i.evalNode(node)
	}
}

// evalNode handles each node type: declarations, assignments, and print statements
func (i *Interpreter) evalNode(node Node) {
	switch n := node.(type) {
	case *VarDeclaration:
		i.evalVarDeclaration(n)
	case *Assignment:
		i.evalAssignment(n)
	case *PrintStatement:
		i.evalPrintStatement(n)
	default:
		fmt.Printf("Unknown node type: %T\n", n)
	}
}

func (i *Interpreter) evalVarDeclaration(node *VarDeclaration) {
	// Store the variable in the enviornment
	i.env[node.Name] = node.Value
}

func (i *Interpreter) evalAssignment(node *Assignment) {
	result := i.evalExpression(node.Expression)
	// technically we should check if the variable is already in the enviornment
	i.env[node.VarName] = result
}

func (i *Interpreter) evalPrintStatement(node *PrintStatement) {
	fmt.Println(i.evalExpression(node.Expression))
}

func (i *Interpreter) evalExpression(expr Expression) int {
	switch e := expr.(type) {
	case *NumberLiteral:
		return e.Value
	case *Variable:
		val, ok := i.env[e.Name]
		if !ok {
			fmt.Printf("Variable not found: %s\n", e.Name)
		}
		return val
	case *BinaryExpression:
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
