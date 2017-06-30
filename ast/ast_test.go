package ast

import (
	"testing"

	"github.com/mdaisuke/gobantam/token"
)

func TestString(t *testing.T) {
	exp := &AssignExpression{
		Name: "a",
		Right: &OperatorExpression{
			Operator: token.Token{Type: token.PLUS, Text: "+"},
			Left: &NameExpression{
				Name: "b",
			},
			Right: &NameExpression{
				Name: "c",
			},
		},
	}
	expected := "(a = (b + c))"
	if exp.String() != expected {
		t.Errorf("expected=%s, got=%s", expected, exp.String())
	}
}
