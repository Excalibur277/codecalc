package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Node interface {
	String() string
}

func MapString(nodes []Node) []string {
	stringslice := make([]string, len(nodes))
	for i, n := range nodes {
		stringslice[i] = n.String()
	}
	return stringslice
}

func SliceToString[T Node](nodes []T, sep string) string {
	stringslice := make([]string, len(nodes))
	for i, n := range nodes {
		stringslice[i] = n.String()
	}
	return strings.Join(stringslice, sep)
}

type context struct {
	stackOffset int64
	variables   map[string]int64
}

type Module interface {
	Node
	IsModule()
	Generate() string
}

type ModuleBase struct {
	statements []Statement
}

func NewModule(statements []Statement) Module {
	return &ModuleBase{
		statements: statements,
	}
}

func (m *ModuleBase) String() string {
	return SliceToString(m.statements, "\n")
}

func (m *ModuleBase) Generate() string {
	fasm := `; Module

use64

format ELF64 executable 3
entry main

segment readable executable

main: 
  ; Setup Call Stack Frame
  push rbp
  mov rbp, rsp

  ; Call Function
  call routine

  ; Reset Stack Frame
  pop rbp


  jmp exit

printInt:
  ; Move parameter to rax
  mov rax, [rsp+8]
 
  ; Create String Buffer
  sub rsp, 21

  ; Move divisor to rbx
  mov rbx, 10


  ; Add Newline to end of msg
  mov rcx, 21
  mov byte [rsp+rcx-1], 10
  dec rcx
  mov r14,1
 
  xor r15, r15
  ; Test if the value is positive, else make it negative
  test rax, rax
  jns .next 
  neg rax 
  mov r15, 1
  .next:
 
  .repeat:
  xor rdx, rdx ; 0 out rdx
  div rbx ; RDX:RAX / RBX
  add rdx, '0' ; Remainder [0,9] -> ["0","9"]

  mov byte [rsp+rcx-1], dl 
  inc r14
  dec rcx

  cmp rax, 0
  je .finish
 
  cmp rcx, 2
  jae .repeat

  .finish:

  test r15, r15
  jz .skipMinus
  mov byte [rsp+rcx-1], '-'
  inc r14
  dec rcx
  .skipMinus:

 
  mov rdx, r14
  lea rsi,[rsp+rcx]
  mov rdi,1 ; STDOUT
  mov rax,1 ; sys_write
  syscall

  ; Clear up String Buffer
  add rsp, 21

  ret


exit:
  mov eax,1
  xor ebx,ebx
  int 0x80

panic:
  mov eax,1
  mov ebx,1
  int 0x80

routine:
`
	ctx := &context{0, map[string]int64{}}
	for _, s := range m.statements {
		fasm += s.Generate(ctx)
	}
	fasm += "  add rsp, " + strconv.FormatInt((ctx.stackOffset)*8, 10) + "\n"
	fasm += "  ret"
	return fasm
}

func (m *ModuleBase) IsModule() {}

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
	index, ok := ctx.variables[s.identifier]
	if !ok {
		ctx.stackOffset++
		ctx.variables[s.identifier] = ctx.stackOffset
		fasm += "  push rax\n"
	} else {
		fasm += "  mov [rbp-" + strconv.FormatInt((index+1)*8, 10) + "], rax\n"
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
  pop rbp
`
	return fasm
}

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
	index, ok := ctx.variables[e.identifier]
	if !ok {
		panic(fmt.Sprintf("Variable %s referenced before being assigned", e.identifier))
	}

	fasm := "  mov rax, [rbp-" + strconv.FormatInt((index+1)*8, 10) + "]\n"
	return fasm
}

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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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

func (e *LSTExpression) Generate(ctx *context) string {
	fasm := e.lhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += e.rhs.Generate(ctx)
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
	fasm += "  push rax\n"
	fasm += "  pop r8\n"
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
