package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type BracketExpression struct {
	BaseExpression
	expression Expression
}

func (e *BracketExpression) Generate(ctx *context) string {
	return e.expression.Generate(ctx)
}

func (s *BracketExpression) String() string {
	return fmt.Sprintf("(%s)", s.expression)
}

type PosExpression struct {
	BaseExpression
	expression Expression
}

func (e *PosExpression) Generate(ctx *context) string {
	return e.expression.Generate(ctx)
}

func (s *PosExpression) String() string {
	return fmt.Sprintf("+%s", s.expression)
}

type NegExpression struct {
	BaseExpression
	expression Expression
}

func (s *NegExpression) String() string {
	return fmt.Sprintf("-%s", s.expression)
}

func (e *NegExpression) Generate(ctx *context) string {
	fasm := e.expression.Generate(ctx)
	fasm += "  neg rax\n"
	return fasm
}

type NotExpression struct {
	BaseExpression
	expression Expression
}

func (s *NotExpression) String() string {
	return fmt.Sprintf("!%s", s.expression)
}

func (e *NotExpression) Generate(ctx *context) string {
	fasm := e.expression.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  xor rax, rax\n"
	fasm += "  test r8,r8\n"
	fasm += "  setz al\n"
	return fasm
}

func NewUnaryExpression(operator antlr.Token, expression Expression) Expression {
	switch op := operator.GetText(); op {
	case "(":
		return &BracketExpression{expression: expression}
	case "+":
		return &PosExpression{expression: expression}
	case "-":
		return &NegExpression{expression: expression}
	case "!":
		return &NotExpression{expression: expression}
	default:
		panic(fmt.Sprintf("Unknown operator: %s", op))
	}
}
