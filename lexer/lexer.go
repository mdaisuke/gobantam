package lexer

import (
	"github.com/mdaisuke/gobantam/token"
)

type Lexer struct {
	text  string
	index int
}

func New(text string) *Lexer {
	l := &Lexer{
		text:  text,
		index: 0,
	}
	return l
}

func (l *Lexer) NextToken() token.Token {
	for l.index < len(l.text) {
		ch := l.text[l.index]
		l.index += 1

		var tok token.Token
		switch ch {
		case '(':
			tok = token.Token{Type: token.LPAREN, Text: string(ch)}
		case ')':
			tok = token.Token{Type: token.RPAREN, Text: string(ch)}
		case '.':
			tok = token.Token{Type: token.COMMA, Text: string(ch)}
		case '=':
			tok = token.Token{Type: token.ASSIGN, Text: string(ch)}
		case '+':
			tok = token.Token{Type: token.PLUS, Text: string(ch)}
		case '-':
			tok = token.Token{Type: token.MINUS, Text: string(ch)}
		case '*':
			tok = token.Token{Type: token.ASTERISK, Text: string(ch)}
		case '/':
			tok = token.Token{Type: token.SLASH, Text: string(ch)}
		case '^':
			tok = token.Token{Type: token.CARET, Text: string(ch)}
		case '~':
			tok = token.Token{Type: token.TILDA, Text: string(ch)}
		case '!':
			tok = token.Token{Type: token.BANG, Text: string(ch)}
		case '?':
			tok = token.Token{Type: token.QUESTION, Text: string(ch)}
		case ':':
			tok = token.Token{Type: token.COLON, Text: string(ch)}
		default:
			if isLetter(ch) {
				return token.Token{Type: token.NAME, Text: l.readName()}
			}
		}
		if len(tok.Text) > 0 {
			return tok
		}

	}

	return token.Token{Type: token.EOF, Text: ""}
}

func isLetter(ch byte) bool {
	if 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' {
		return true
	}
	return false
}

func (l *Lexer) readName() string {
	l.index -= 1
	index := l.index
	for l.index < len(l.text) && isLetter(l.text[l.index]) {
		l.index += 1
	}
	return l.text[index:l.index]
}
