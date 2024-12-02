package bin

import (
	"fmt"
	. "zwa/types"
	"zwa/utils"
)

type AST = utils.Queue[Statement]

type Parser struct {
	Tokens
	ast AST
}

func Parse(tokens Tokens) AST {
	parser := Parser{tokens, utils.NewQueue[Statement]()}
	return parser.parseStatements()
}

func (parser *Parser) parseStatements() AST {
	stmt, ok := parser.Pop()
	if !ok {
		return parser.ast
	}

	var statement Statement
	switch stmt.Type {
	case IDENT:
		parser.Pop() // Remove Equals Sign
		statement = Assignment{Name: stmt.Literal, Value: parser.parseExpression(parser.Pop())}
	case SHOW:
		statement = Print{Value: parser.parseExpression(parser.Pop())}
	}

	parser.ast.Push(statement)
	return parser.parseStatements()
}

func (parser *Parser) peekTokenIsOperator() bool {
	if tok, ok := parser.Peek(); ok {
		switch tok.Type {
		case PLUS, MINUS, ASTERISK, SLASH, MODULO:
			return true
		}
	}
	return false
}

func (parser *Parser) parseExpression(current token, ok bool) Expression {
	if !ok {
		fmt.Println("No expression to parse")
		return NumberLiteral{Value: 0}
	}

	var expr Expression
	switch current.Type {
	case NUMBER:
		expr = NumberLiteral{Value: Atoi(current.Literal)}
	case STRING:
		expr = StringLiteral{Value: current.Literal}
	case IDENT:
		expr = Variable{Value: current.Literal}
	case LPAREN:
		expr = parser.parseParanthesis()
	}

	if parser.peekTokenIsOperator() {
		operator, _ := parser.Pop()
		next := parser.parseExpression(parser.Pop())
		expr = BinaryExpression{Lhs: expr, Operator: rune(operator.Literal[0]), Rhs: next}
	}

	return expr
}

func (parser *Parser) parseParanthesis() Expression {
	current, ok := parser.Pop()
	if !ok || current.Type == RPAREN {
		fmt.Println("No expression to parse")
		return NumberLiteral{Value: 0}
	}

	expr := parser.parseExpression(current, ok)

	if next, ok := parser.Peek(); ok && next.Type == RPAREN {
		parser.Pop()
	} else {
		fmt.Println("Failed to terminate paranthesis")
	}

	return expr
}
