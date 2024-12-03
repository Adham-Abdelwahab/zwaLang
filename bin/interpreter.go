package bin

type Interpreter struct {
	AST
	env map[string]any
}

func Interpret(ast AST) {
	interpreter := Interpreter{ast, make(map[string]any)}
	interpreter.interpret()
}

func (interpreter *Interpreter) interpret() {
	statement, ok := interpreter.Pop()
	if !ok {
		return
	}
	statement.Evaluate(interpreter.env)
	interpreter.interpret()
}
