package parser

import(
	"fmt"
	"zwaLang/src/lexer"
	"zwaLang/src/token"
)

// AST Node types
type Node interface{}

type VarDeclaration struct {
	Name  string
	Value int
}

type Assignment struct {
	VarName string
	Expression Node
}

type PrintStatement struct {
	VarName string
}

type Expression interface {
	Node
}

type NumberLiteral struct {
	Value int
}

type Variable struct {
	Name string
}

type BinaryExpression struct {
	Left Expression
	Operator string
	Right Expression
}

// Parser structure
type Parser struct {
	lex 		*lexer.Lexer
	curToken 	token.Token
	peekToken 	token.Token
}

// NewParser initializes a parser with the given lexer
func NewParser(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}
	p.nextToken()
	p.nextToken()
	return p
}

// -- Receiver Functions --

// NextToken advances the parser to the next token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) ParseProgram() []Node {
	var nodes []Node
	for p.curToken.Type != token.EOF {
		node := p.parseStatement()
		if node != nil {
			nodes = append(nodes, node)
		}
		p.nextToken()
	}
	return nodes
}

func (p *Parser) parseStatement() Node {
	switch p.curToken.Type {
		case token.IDENT:
			return p.parseAssignmentOrDeclaration()
		case token.SHOW:
			return p.parsePrintStatement()
		default:
			fmt.Printf("Unknown statement: %+v\n", p.curToken.Literal)
			return nil
	}
}

func (p *Parser) parseAssignmentOrDeclaration() Node {
	identifier := p.curToken.Literal
	p.nextToken()

	if p.curToken.Type == token.COLON {
		// Declaration
		p.nextToken()
		if p.curToken.Type != token.NATURAL_NUMBER_TYPE {
			fmt.Printf("Expected type of variable (Currently only supports type: number) after colon: %+v\n", p.curToken.Literal)
			return nil
		}
		p.nextToken()
		if p.curToken.Type != token.ASSIGN {
			fmt.Printf("Expected assignment operator after type declaration: %+v\n", p.curToken.Literal)
			return nil
		}
		p.nextToken()
		value := p.curToken.Literal
		return &VarDeclaration{Name: identifier, Value: atoi(value)}
	} else if p.curToken.Type == token.ASSIGN {
		p.nextToken()
		expression := p.parseExpression()
		return &Assignment{VarName: identifier, Expression: expression}
	}

	return nil
}

func (p *Parser) parseExpression() Expression {
	left := p.parseTerm()

	for p.curToken.Type == token.PLUS {
		operator := p.curToken.Literal
		p.nextToken()
		right := p.parseTerm()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left;
}


// parseTerm parses either a variable or an integer literal
func (p *Parser) parseTerm() Expression {
	switch p.curToken.Type {
		case token.IDENT:
			ident := p.curToken.Literal
			// if the next token is a plus, then we need to parse the next term
			if(p.peekTokenIs(token.PLUS)) {
				p.nextToken()
			}
			return &Variable{Name: ident}
		case token.NUMBER:
			value := atoi(p.curToken.Literal)
			// if the next token is a plus, then we need to parse the next term (needs to be upgraded eventually when we have more operations)
			if(p.peekTokenIs(token.PLUS)) {
				p.nextToken()
			}
			return &NumberLiteral{Value: value}
		default:
			fmt.Printf("Unexpected token: %+v\n", p.curToken.Literal)
			return nil
	}
}

func (p *Parser) parsePrintStatement() Node {
	p.nextToken()
	// TODO: Add support for printing more than just a variable
	if p.curToken.Type != token.IDENT {
		fmt.Printf("Expected identifier after print statement: %+v\n", p.curToken.Literal)
		return nil
	}
	return &PrintStatement{VarName: p.curToken.Literal}
}

// -- Helper Functions --

func atoi(str string) int {
	var result int
	fmt.Sscanf(str, "%d", &result)
	return result
}
