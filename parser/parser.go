package parser

import (
	"github.com/mdaisuke/gobantam/ast"
	"github.com/mdaisuke/gobantam/lexer"
	"github.com/mdaisuke/gobantam/token"
)

const (
	_ = iota
	ASSIGNMENT
	CONDITIONAL
	SUM
	PRODUCT
	EXPONENT
	PREFIX
	POSTFIX
	CALL
)

type (
	PrefixParselet (func() ast.Expression)
	InfixParselet  (func(left ast.Expression) ast.Expression)
)

type Parser struct {
	l *lexer.Lexer

	prefixParselets map[token.TokenType]PrefixParselet
	infixParselets  map[token.TokenType]InfixParselet

	read []token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:               l,
		prefixParselets: map[token.TokenType]PrefixParselet{},
		infixParselets:  map[token.TokenType]InfixParselet{},
		read:            []token.Token{},
	}

	return p
}

func (p *Parser) Parse() ast.Expression {
	return p.parseExpression(0)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	tok := p.consume()
	prefix := p.prefixParselets[tok.Type]

	left := prefix()

	for precedence < p.getPrecedence() {
		tok = p.consume()

		infix := p.infixParseLets[tok.Type]
		left = infix(left)
	}

	return left
}

func (p *Parser) registerPrefixParselet(tokenType token.TokenType, parselet PrefixParselet) {
	p.prefixParselets[tokenType] = parselet
}

func (p *Parser) registerInfixParselet(tokenType token.TokenType, parselet InfixParselet) {
	p.infixParselets[tokenType] = parselet
}

func (p *Parser) lookAhead(distance int) token.Token {
	for len(p.read) <= distance {
		p.read = append(p.read, p.l.NextToken())
	}
	return p.read[distance]
}

func (p *Parser) consume() token.Token {
	p.lookAhead(0)

	tok := p.read[0]
	p.read = p.read[1:]

	return tok
}

func (p *Parser) getPrecedence() {

}
