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
	Left    string
	Right   string
}

type PrintStatement struct {
	VarName string
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
		// Assignment -- currently only supports arithmetic assignment & very strict rules
		p.nextToken()
		if p.curToken.Type != token.IDENT {
			fmt.Printf("Expected identifier after assignment operator: %+v\n", p.curToken.Literal)
			return nil
		}
		left := p.curToken.Literal
		
		p.nextToken()
		// TODO: Add support for more operators & to end here with re-assignment if a variable is being re-assigned
		if p.curToken.Type != token.PLUS {
			fmt.Printf("Expected arithmetic operator after assignment operator: %+v\n", p.curToken.Literal)
			return nil
		}
		
		p.nextToken()
		if p.curToken.Type != token.IDENT {
			fmt.Printf("Expected identifier after arithmetic operator: %+v\n", p.curToken.Literal)
			return nil
		}
		right := p.curToken.Literal
		return &Assignment{VarName: identifier, Left: left, Right: right}
	}

	return nil
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
