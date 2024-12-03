package types

import "fmt"

type Expression interface {
	Concrete(env) any
}

/* Parser Abstract Syntax Tree Nodes */

type NumberLiteral struct {
	Value int
}
type StringLiteral struct {
	Value string
}
type BoolLiteral struct {
	Value bool
}
type Variable struct {
	Value string
}
type BinaryExpression struct {
	Lhs      Expression
	Operator rune
	Rhs      Expression
}

/* Interpreter Evaluations */

func (number NumberLiteral) Concrete(env env) any {
	return number.Value
}
func (str StringLiteral) Concrete(env env) any {
	return str.Value
}
func (boolean BoolLiteral) Concrete(env env) any {
	return boolean.Value
}
func (variable Variable) Concrete(env env) any {
	val, ok := env[variable.Value]
	if !ok {
		fmt.Printf("Variable not found in environment: %s\n", variable.Value)
		return nil
	} else {
		return val
	}
}
func (binary BinaryExpression) Concrete(env env) any {
	return binary.perform(env)
}
