package ast

import (
	"fmt"
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
	if len(ctx.whiles) > 0 {
		return fmt.Sprintf("  jmp .breakwhile%d\n", ctx.whiles[len(ctx.whiles)-1])
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
	if len(ctx.whiles) > 0 {
		return fmt.Sprintf("  jmp .continuewhile%d\n", ctx.whiles[len(ctx.whiles)-1])
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
	ctx.labelCounter++
	labelCount := ctx.labelCounter
	ctx.whiles = append(ctx.whiles, labelCount)

	fasm := fmt.Sprintf(".checkwhile%d:\n", labelCount)

	fasm += s.expression.Generate(ctx)

	fasm += "  test rax,rax\n"
	fasm += fmt.Sprintf("  jz .finishwhile%d\n", labelCount)

	fasm += "  push rbp\n"
	fasm += "  mov rbp, rsp\n"

	ctx.addScope()

	for _, statement := range s.statements {
		fasm += statement.Generate(ctx)
	}

	assigned := ctx.popScope()

	fasm += fmt.Sprintf(".continuewhile%d:\n", labelCount)
	fasm += fmt.Sprintf("  add rsp, %d*8\n", assigned)
	fasm += "  mov rsp, rbp\n"
	fasm += "  pop rbp\n"
	fasm += fmt.Sprintf("  jmp .checkwhile%d\n", labelCount)
	fasm += fmt.Sprintf(".breakwhile%d:\n", labelCount)
	fasm += fmt.Sprintf("  add rsp, %d*8\n", assigned)
	fasm += "  mov rsp, rbp\n"
	fasm += "  pop rbp\n"
	fasm += fmt.Sprintf(".finishwhile%d:\n", labelCount)

	return fasm
}
