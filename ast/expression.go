package ast

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Expression interface {
	Node
	IsExpression()
	Generate(ctx *context) string
}
type BaseExpression struct {
}

func (e *BaseExpression) IsExpression() {}

type LiteralExpression struct {
	BaseExpression
	literal int64
}

func NewLiteralExpression(literal antlr.Token) Expression {
	if literal, err := strconv.ParseInt(literal.GetText(), 10, 64); err == nil {
		return &LiteralExpression{
			literal: literal,
		}
	} else {
		panic(err)
	}
}

func (e *LiteralExpression) String() string {
	return strconv.FormatInt(e.literal, 10)
}

func (e *LiteralExpression) Generate(ctx *context) string {
	fasm := "  mov rax, " + strconv.FormatInt(e.literal, 10) + "\n"
	return fasm
}

type IdentifierExpression struct {
	BaseExpression
	identifier string
}

func NewIdentifierExpression(identifier antlr.Token) Expression {
	return &IdentifierExpression{
		identifier: identifier.GetText(),
	}
}

func (e *IdentifierExpression) String() string {
	return e.identifier
}

func (e *IdentifierExpression) Generate(ctx *context) string {
	addr, ok := ctx.getVar(e.identifier)
	if !ok {
		panic(fmt.Sprintf("Variable %s referenced before being assigned", e.identifier))
	}

	fasm := "  mov rax, [" + addr + "]\n"
	return fasm
}
