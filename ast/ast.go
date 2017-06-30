package ast

import (
	"bytes"

	"github.com/mdaisuke/gobantam/token"
)

type Expression interface {
	String() string
}

type CallExpression struct {
	Function Expression
	Args     []Expression
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	for i, arg := range ce.Args {
		out.WriteString(arg.String())
		if i < len(ce.Args)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")

	return out.String()
}

type NameExpression struct {
	Name string
}

func (ne *NameExpression) String() string {
	return ne.Name
}

type AssignExpression struct {
	Name  string
	Right Expression
}

func (ae *AssignExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ae.Name)
	out.WriteString(" = ")
	out.WriteString(ae.Right.String())
	out.WriteString(")")

	return out.String()
}

type PrefixExpression struct {
	Operator token.Token
	Right    Expression
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator.String())
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type PostfixExpression struct {
	Operator token.Token
	Left     Expression
}

func (pe *PostfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Left.String())
	out.WriteString(pe.Operator.String())
	out.WriteString(")")

	return out.String()
}

type OperatorExpression struct {
	Operator token.Token
	Left     Expression
	Right    Expression
}

func (oe *OperatorExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" ")
	out.WriteString(oe.Operator.String())
	out.WriteString(" ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

type ConditionalExpression struct {
	Condition Expression
	ThenArm   Expression
	ElseArm   Expression
}

func (ce *ConditionalExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ce.Condition.String())
	out.WriteString(" ? ")
	out.WriteString(ce.ThenArm.String())
	out.WriteString(" : ")
	out.WriteString(ce.ElseArm.String())
	out.WriteString(")")

	return out.String()
}
