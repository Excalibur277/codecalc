package ast

import (
	"fmt"
	"strconv"
	"strings"
)

type BreakStatement struct {
	BaseStatement
}

func NewBreakStatement() Statement {
	return &BreakStatement{}
}

func (s *BreakStatement) String() string {
	return "break;"
}

func (s *BreakStatement) Generate(ctx *context) string {
	if labelCount, offset, ok := ctx.currentLoop(); ok {
		return fmt.Sprintf("  add rbp, %s\n  jmp .breakwhile%d\n", strconv.FormatInt((offset)*8, 10), labelCount)
	} else {
		panic("Cannot 'break' outside of a while statement")
	}
}

type ContinueStatement struct {
	BaseStatement
}

func NewContinueStatement() Statement {
	return &ContinueStatement{}
}

func (s *ContinueStatement) String() string {
	return "continue;"
}

func (s *ContinueStatement) Generate(ctx *context) string {
	if labelCount, offset, ok := ctx.currentLoop(); ok {
		return fmt.Sprintf("  add rbp, %s\n  jmp .continuewhile%d\n", strconv.FormatInt((offset)*8, 10), labelCount)
	} else {
		panic("Cannot 'continue' outside of a while statement")
	}
}

type WhileStatement struct {
	BaseStatement
	expression Expression
	statements []Statement
}

func NewWhileStatement(expression Expression, statements []Statement) Statement {
	return &WhileStatement{expression: expression, statements: statements}
}

func (s *WhileStatement) String() string {
	return fmt.Sprintf("while %s {\n\t%s\n}", s.expression, strings.ReplaceAll(SliceToString(s.statements, "\n"), "\n", "\n\t"))
}

func (s *WhileStatement) Generate(ctx *context) string {
	exp := s.expression.Generate(ctx)

	labelCount := ctx.addScope(true)
	fasm := fmt.Sprintf(".checkwhile%d:\n", labelCount)
	fasm += exp
	fasm += "  test rax,rax\n"
	fasm += fmt.Sprintf("  jz .finishwhile%d\n", labelCount)

	fasm += "  push rbp\n"
	fasm += "  mov rbp, rsp\n"

	for _, statement := range s.statements {
		fasm += statement.Generate(ctx)
	}

	ctx.popScope()
	fasm += fmt.Sprintf(".continuewhile%d:\n", labelCount)
	fasm += "  mov rsp, rbp\n"
	fasm += "  pop rbp\n"
	fasm += fmt.Sprintf("  jmp .checkwhile%d\n", labelCount)
	fasm += fmt.Sprintf(".breakwhile%d:\n", labelCount)
	fasm += "  mov rsp, rbp\n"
	fasm += "  pop rbp\n"
	fasm += fmt.Sprintf(".finishwhile%d:\n", labelCount)

	return fasm
}
