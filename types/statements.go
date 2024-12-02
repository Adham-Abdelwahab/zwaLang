package types

import "fmt"

type env = map[string]any

type Statement interface {
	Evaluate(env)
}

/* Parser Abstract Syntax Tree Nodes */

type Assignment struct {
	Name  string
	Value Expression
}

type Print struct {
	Value Expression
}

/* Interpreter Evaluations */

func (assign Assignment) Evaluate(env env) {
	expr := assign.Value.Concrete(env)
	env[assign.Name] = expr
}

func (print Print) Evaluate(env env) {
	expr := print.Value.Concrete(env)
	fmt.Println(expr)
}
