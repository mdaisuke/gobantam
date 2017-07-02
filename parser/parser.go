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
