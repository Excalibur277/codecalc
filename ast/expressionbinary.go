package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type BaseBinaryExpression struct {
	BaseExpression
	lhs Expression
	rhs Expression
}

type AddExpression struct {
	BaseBinaryExpression
}

type SubExpression struct {
	BaseBinaryExpression
}

type MulExpression struct {
	BaseBinaryExpression
}

type DivExpression struct {
	BaseBinaryExpression
}
type ModExpression struct {
	BaseBinaryExpression
}
type LSTExpression struct {
	BaseBinaryExpression
}
type LTEExpression struct {
	BaseBinaryExpression
}
type GRTExpression struct {
	BaseBinaryExpression
}
type GTEExpression struct {
	BaseBinaryExpression
}
type EQUExpression struct {
	BaseBinaryExpression
}
type NEQExpression struct {
	BaseBinaryExpression
}
type AndExpression struct {
	BaseBinaryExpression
}
type OrExpression struct {
	BaseBinaryExpression
}

func NewBinaryExpression(lhs Expression, operator antlr.Token, rhs Expression) Expression {
	switch op := operator.GetText(); op {
	case "+":
		return &AddExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case "-":
		return &SubExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case "*":
		return &MulExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case "/":
		return &DivExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case "%":
		return &ModExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case "<":
		return &LSTExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case "<=":
		return &LTEExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case ">":
		return &GRTExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case ">=":
		return &GTEExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case "==":
		return &EQUExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case "!=":
		return &NEQExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case "&":
		return &AndExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	case "|":
		return &OrExpression{BaseBinaryExpression{lhs: lhs, rhs: rhs}}
	default:
		panic(fmt.Sprintf("Unknown operator: %s", op))
	}
}

func (e *AddExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop rax\n"
	fasm += "  add rax, r8\n"
	return fasm
}

func (s *AddExpression) String() string {
	return fmt.Sprintf("%s + %s", s.lhs, s.rhs)
}

func (e *SubExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop rax\n"
	fasm += "  sub rax, r8\n"
	return fasm
}

func (s *SubExpression) String() string {
	return fmt.Sprintf("%s - %s", s.lhs, s.rhs)
}

func (e *MulExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop rax\n"
	fasm += "  imul rax, r8\n"
	return fasm
}

func (s *MulExpression) String() string {
	return fmt.Sprintf("%s * %s", s.lhs, s.rhs)
}

func (e *DivExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop rax\n"
	fasm += "  test r8,r8\n"
	fasm += "  jz panic\n"
	fasm += "  xor rdx, rdx\n"
	fasm += "  idiv qword r8\n"
	return fasm
}

func (s *DivExpression) String() string {
	return fmt.Sprintf("%s / %s", s.lhs, s.rhs)
}

func (e *ModExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop rax\n"
	fasm += "  test r8,r8\n"
	fasm += "  jz panic\n"
	fasm += "  xor rdx, rdx\n"
	fasm += "  idiv qword r8\n"
	fasm += "  mov rax, rdx\n"
	return fasm
}

func (s *ModExpression) String() string {
	return fmt.Sprintf("%s %% %s", s.lhs, s.rhs)
}

func (e *LSTExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop r9\n"
	fasm += "  xor rax, rax\n"
	fasm += "  cmp r9, r8\n"
	fasm += "  setl al\n"
	return fasm
}

func (s *LSTExpression) String() string {
	return fmt.Sprintf("%s < %s", s.lhs, s.rhs)
}

func (e *LTEExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop r9\n"
	fasm += "  xor rax, rax\n"
	fasm += "  cmp r9, r8\n"
	fasm += "  setle al\n"
	return fasm
}

func (s *LTEExpression) String() string {
	return fmt.Sprintf("%s <= %s", s.lhs, s.rhs)
}

func (e *GRTExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop r9\n"
	fasm += "  xor rax, rax\n"
	fasm += "  cmp r9, r8\n"
	fasm += "  setg al\n"
	return fasm
}

func (s *GRTExpression) String() string {
	return fmt.Sprintf("%s > %s", s.lhs, s.rhs)
}

func (e *GTEExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop r9\n"
	fasm += "  xor rax, rax\n"
	fasm += "  cmp r9, r8\n"
	fasm += "  setge al\n"
	return fasm
}

func (s *GTEExpression) String() string {
	return fmt.Sprintf("%s >= %s", s.lhs, s.rhs)
}

func (e *EQUExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop r9\n"
	fasm += "  xor rax, rax\n"
	fasm += "  cmp r9, r8\n"
	fasm += "  sete al\n"
	return fasm
}

func (s *EQUExpression) String() string {
	return fmt.Sprintf("%s == %s", s.lhs, s.rhs)
}

func (e *NEQExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop r9\n"
	fasm += "  xor rax, rax\n"
	fasm += "  cmp r9, r8\n"
	fasm += "  setne al\n"
	return fasm
}

func (s *NEQExpression) String() string {
	return fmt.Sprintf("%s != %s", s.lhs, s.rhs)
}

func (e *AndExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop r9\n"
	fasm += "  xor rax, rax\n"
	fasm += "  xor rbx, rbx\n"
	fasm += "  test r8,r8\n"
	fasm += "  setnz al\n"
	fasm += "  test r9,r9\n"
	fasm += "  setnz bl\n"
	fasm += "  and rax, rbx\n"
	return fasm
}

func (s *AndExpression) String() string {
	return fmt.Sprintf("%s & %s", s.lhs, s.rhs)
}

func (e *OrExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  mov r8, rax\n"
	fasm += "  pop r9\n"
	fasm += "  or r8, r9\n"
	fasm += "  xor rax, rax\n"
	fasm += "  test r8,r8\n"
	fasm += "  setnz al\n"
	return fasm
}

func (s *OrExpression) String() string {
	return fmt.Sprintf("%s | %s", s.lhs, s.rhs)
}
