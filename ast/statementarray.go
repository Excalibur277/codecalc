package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type ArrayStatement struct {
	BaseStatement
	identifier string
	expression Expression
}

func NewArrayStatement(identifier antlr.Token, expression Expression) Statement {
	return &ArrayStatement{
		identifier: identifier.GetText(),
		expression: expression,
	}
}

func (s *ArrayStatement) String() string {
	return fmt.Sprintf("array %s[%s];", s.identifier, s.expression)
}

func (s *ArrayStatement) Generate(ctx *context) string {
	fasm := s.expression.Generate(ctx)
	fasm += "  cmp rax,0\n"
	fasm += "  jle arraypanic\n"
	fasm += "  push rax\n"    // Save length
	fasm += "  inc rax\n"     // add val for length
	fasm += "  imul rax, 8\n" // 8 bytes each
	ctx.addVar(s.identifier, Array)

	// Allocate memory
	// Arguments rdi, rsi, rdx, r10, r8, r9
	fasm += "  mov rdi, 0\n"   // NULL addr
	fasm += "  mov rsi, rax\n" // length
	fasm += "  mov rdx, 3\n"   // Read,Write but no execution
	fasm += "  mov r10, 34\n"  // Private & Anonymous file
	fasm += "  mov r8, -1\n"
	fasm += "  mov r9, 0\n"
	fasm += "  mov rax, 9\n"
	fasm += "  syscall\n"
	fasm += "  pop r8\n"        // Get length
	fasm += "  mov [rax], r8\n" // Set length in array values
	fasm += "  push rax\n"
	// TODO - deallocate array
	return fasm
}

type AssignIndexStatement struct {
	BaseStatement
	identifier      string
	indexExpression Expression
	valueExpression Expression
}

func NewAssignIndexStatement(identifier antlr.Token, indexExpression, valueExpression Expression) Statement {
	return &AssignIndexStatement{
		identifier:      identifier.GetText(),
		indexExpression: indexExpression,
		valueExpression: valueExpression,
	}
}

func (s *AssignIndexStatement) String() string {
	return fmt.Sprintf("%s[%s] = %s;", s.identifier, s.indexExpression, s.valueExpression)
}

func (s *AssignIndexStatement) Generate(ctx *context) string {
	addr, dataType, ok := ctx.getVar(s.identifier)
	if !ok {
		panic(fmt.Sprintf("Variable %s indexed before being allocated", s.identifier))
	}

	if dataType != Array {
		panic(fmt.Sprintf("Variable %s is not an array and cannot be indexed", s.identifier))
	}

	fasm := s.valueExpression.Generate(ctx)
	fasm += "  push rax\n"
	fasm += s.indexExpression.Generate(ctx)
	fasm += "  mov r8, [" + addr + "]\n"

	// Check arround bounds
	fasm += "  cmp rax,0\n"
	fasm += "  jl negindexpanic\n"
	fasm += "  mov r10,[r8]\n"
	fasm += "  dec r10\n"
	fasm += "  cmp rax,r10\n"
	fasm += "  jg boundspanic\n"

	fasm += "  pop r9\n"
	fasm += "  mov [r8+rax*8+8], r9\n"
	return fasm
}
