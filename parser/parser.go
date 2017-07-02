package parser

import (
	"go/token"

	"github.com/mdaisuke/gobantam/ast"
	"github.com/mdaisuke/gobantam/lexer"
)

type (
	prefixParselet func() ast.Expression
	infixParselet  func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) nameParselet() ast.Expression {
	return &NameExpression{Name: p.curToken.Text}
}

func (p *Parser) groupParselet() ast.Expression {
	p.nextToken()

	exp := p.parseExpression()

	p.nextToken()

	return exp
}

func (p *Parser) prefixOperatorParselet() ast.Expression {
	exp := &PrefixExpression{Operator: p.curToken}

	p.nextToken()

	right := p.parseExpression()
	exp.Right = right

	return exp
}

func (p *Parser) callParselet(left ast.Expression) ast.Expression {
	args := []ast.Expression{}
	p.nextToken()

	for p.curToken.Type != token.RPAREN {
		arg := p.parseExpression()
		args = append(args, arg)
		if p.curToken.Type == token.COMMA {
			p.nextToken()
		}
	}

	return &CallExpression{
		Function: left,
		Args:     args,
	}
}
