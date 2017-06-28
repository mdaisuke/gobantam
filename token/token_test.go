package token

import (
	"testing"
)

func TestPunctuator(t *testing.T) {
	tt := COMMA
	if tt != 3 {

	}
	c := tt.Punctuator()
	if c != "." {
		t.Errorf("commas is not '.', got=%s", c)
	}
}

func TestToken(t *testing.T) {
	tt := &Token{Type: COMMA, Text: "."}
	if tt.String() != tt.Type.Punctuator() {
		t.Errorf("type.String does not match tt.type.punctuator. got=%s", tt.String())
	}
}
