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
	if scope, ok := ctx.currentLoop(); ok {
		fasm := ""
		for _, s := range ctx.scopesInScope(scope) {
			fasm += s.releaseScopeMemory()
		}
		fasm += fmt.Sprintf("  jmp .breakwhile%d\n", scope.labelCount)
		return fasm
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
	if scope, ok := ctx.currentLoop(); ok {
		fasm := ""
		for _, s := range ctx.scopesInScope(scope) {
			fasm += s.releaseScopeMemory()
		}
		fasm += fmt.Sprintf("  jmp .continuewhile%d\n", scope.labelCount)
		return fasm
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

	scope := ctx.addScope(true)
	fasm := fmt.Sprintf(".checkwhile%d:\n", scope.labelCount)
	fasm += exp
	fasm += "  test rax,rax\n"
	fasm += fmt.Sprintf("  jz .finishwhile%d\n", scope.labelCount)

	fasm += scope.reserveScopeMemory()

	for _, statement := range s.statements {
		fasm += statement.Generate(ctx)
	}

	ctx.popScope()
	fasm += fmt.Sprintf(".continuewhile%d:\n", scope.labelCount)
	fasm += scope.releaseScopeMemory()
	fasm += fmt.Sprintf("  jmp .checkwhile%d\n", scope.labelCount)
	fasm += fmt.Sprintf(".breakwhile%d:\n", scope.labelCount)
	fasm += scope.releaseScopeMemory()
	fasm += fmt.Sprintf(".finishwhile%d:\n", scope.labelCount)

	return fasm
}
