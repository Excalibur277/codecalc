package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Statement interface {
	Node
	IsStatement()
	Generate(ctx *context) string
}

type BaseStatement struct {
}

func (s *BaseStatement) IsStatement() {}

type AssignStatement struct {
	BaseStatement
	identifier string
	expression Expression
}

func NewAssignStatement(identifier antlr.Token, expression Expression) Statement {
	return &AssignStatement{
		identifier: identifier.GetText(),
		expression: expression,
	}
}

func (s *AssignStatement) String() string {
	return fmt.Sprintf("%s = %s;", s.identifier, s.expression)
}

func (s *AssignStatement) Generate(ctx *context) string {
	fasm := s.expression.Generate(ctx)
	addr, dataType, ok := ctx.getVar(s.identifier)

	if dataType == Array {
		panic(fmt.Sprintf("Variable %s must be assigned through indexing, not directly", s.identifier))
	}

	if !ok {
		ctx.addVar(s.identifier, Integer)
		fasm += "  push rax\n"
	} else {
		fasm += "  mov [" + addr + "], rax\n"
	}

	return fasm
}

type ExpressionStatement struct {
	BaseStatement
	expression Expression
}

func NewExpressionStatement(expression Expression) Statement {
	return &ExpressionStatement{
		expression: expression,
	}
}

func (s *ExpressionStatement) String() string {
	return fmt.Sprintf("%s;", s.expression)
}

func (s *ExpressionStatement) Generate(ctx *context) string {
	fasm := s.expression.Generate(ctx)
	fasm += `  ; Setup Call Stack Frame
  push rbp
  mov rbp, rsp
  
  ; Call Function
  push rax
  call printInt
  add rsp, 8
  
  ; Reset Stack Frame
  mov rsp, rbp
  pop rbp
`
	return fasm
}
