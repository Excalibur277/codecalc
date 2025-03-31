package ast

import (
	"fmt"
	"strings"
)

type IfStatement struct {
	BaseStatement
	expression    Expression
	statements    []Statement
	elseStatement Statement
}

func NewIfStatement(expression Expression, statements []Statement) Statement {
	return &IfStatement{expression: expression, statements: statements}
}

func NewIfElseStatement(expression Expression, statements []Statement, elseStatement Statement) Statement {
	return &IfStatement{expression: expression, statements: statements, elseStatement: elseStatement}
}

func (s *IfStatement) String() string {
	val := fmt.Sprintf("if %s {\n\t%s\n}", s.expression, strings.ReplaceAll(SliceToString(s.statements, "\n"), "\n", "\n\t"))
	if s.elseStatement != nil {
		val += " " + s.elseStatement.String()
	}
	return val
}

func (s *IfStatement) Generate(ctx *context) string {

	fasm := s.expression.Generate(ctx)

	labelCount := ctx.addScope(false)

	fasm += "  test rax,rax\n"
	fasm += fmt.Sprintf("  jz .else%d\n", labelCount)

	fasm += "  push rbp\n"
	fasm += "  mov rbp, rsp\n"

	for _, statement := range s.statements {
		fasm += statement.Generate(ctx)
	}

	assigned := ctx.popScope()

	fasm += fmt.Sprintf("  add rsp, %d*8\n", assigned)
	fasm += "  mov rsp, rbp\n"
	fasm += "  pop rbp\n"
	fasm += fmt.Sprintf("  jmp .endif%d\n", labelCount)
	fasm += fmt.Sprintf(".else%d:\n", labelCount)

	if s.elseStatement != nil {
		fasm += s.elseStatement.Generate(ctx)
	}

	fasm += fmt.Sprintf(".endif%d:\n", labelCount)

	return fasm
}

type ElseStatement struct {
	BaseStatement
	statements []Statement
}

func NewElseStatement(statements []Statement) Statement {
	return &ElseStatement{statements: statements}
}

func (s *ElseStatement) String() string {
	return fmt.Sprintf("else {\n\t%s\n}", strings.ReplaceAll(SliceToString(s.statements, "\n"), "\n", "\n\t"))
}

func (s *ElseStatement) Generate(ctx *context) string {
	fasm := "  push rbp\n"
	fasm += "  mov rbp, rsp\n"

	ctx.addScope(false)

	for _, statement := range s.statements {
		fasm += statement.Generate(ctx)
	}

	ctx.popScope()
	fasm += "  mov rsp, rbp\n"
	fasm += "  pop rbp\n"

	return fasm
}
