package token

const (
	LPAREN   = "("
	RPAREN   = ")"
	COMMA    = ","
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	CARET    = "^"
	TILDA    = "~"
	BANG     = "!"
	QUESTION = "?"
	COLON    = ":"

	NAME = "NAME"
	EOF  = "EOF"
)

type TokenType string

type Token struct {
	Type TokenType
	Text string
}

func (t *Token) String() string {
	return t.Text
}
