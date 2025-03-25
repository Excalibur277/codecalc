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
		fmt.Println("Assigning")
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

type AddExpression struct {
	BaseExpression
	lhs Expression
	rhs Expression
}

type SubExpression struct {
	BaseExpression
	lhs Expression
	rhs Expression
}

func NewBinaryExpression(lhs Expression, operator antlr.Token, rhs Expression) Expression {
	switch op := operator.GetText(); op {
	case "+":
		return &AddExpression{lhs: lhs, rhs: rhs}
	case "-":
		return &SubExpression{lhs: lhs, rhs: rhs}
	default:
		panic(fmt.Sprintf("Unknown operator: %s", op))
	}
}

func (e *AddExpression) Generate(ctx *context) string {
	fmt.Printf("Enter Sub Expression: %s\n", e)
	fmt.Printf("Checking RHS: %s\n", e.rhs)
	fasm := e.rhs.Generate(ctx)
	fasm += "  push rax\n"
	fmt.Printf("Checking LHS: %s\n", e.lhs)
	fasm += e.lhs.Generate(ctx)
	fasm += "  pop r8\n"
	fasm += "  add rax, r8\n"
	return fasm
}

func (s *AddExpression) String() string {
	return fmt.Sprintf("%s + %s", s.lhs, s.rhs)
}

func (e *SubExpression) Generate(ctx *context) string {
	fmt.Printf("Enter Sub Expression: %s\n", e)
	fmt.Printf("Checking RHS: %s\n", e.rhs)
	fasm := e.rhs.Generate(ctx)
	fasm += "  push rax\n"
	fmt.Printf("Checking LHS: %s\n", e.lhs)
	fasm += e.lhs.Generate(ctx)
	fasm += "  pop r8\n"
	fasm += "  sub rax, r8\n"
	return fasm
}

func (s *SubExpression) String() string {
	return fmt.Sprintf("%s - %s", s.lhs, s.rhs)
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

func NewUnaryExpression(operator antlr.Token, expression Expression) Expression {
	switch op := operator.GetText(); op {
	case "+":
		return &PosExpression{expression: expression}
	case "-":
		fmt.Println("Negative op")
		return &NegExpression{expression: expression}
	default:
		panic(fmt.Sprintf("Unknown operator: %s", op))
	}
}
