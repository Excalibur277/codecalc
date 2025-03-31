package ast

import (
	"fmt"
	"strconv"
	"strings"
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
	scopes       []*scope
	labelCounter uint64
	whiles       []uint64
}

func (ctx *context) currentScope() *scope {
	return ctx.scopes[len(ctx.scopes)-1]
}

func (ctx *context) stackOffset() int64 {
	return ctx.currentScope().stackOffset
}

func (ctx *context) getVar(identifier string) (string, bool) {
	offset := int64(0)
	for i := len(ctx.scopes) - 1; i >= 0; i-- {
		if i != len(ctx.scopes)-1 {
			offset += ctx.scopes[i].stackOffset + 1
		}
		if index, ok := ctx.scopes[i].variables[identifier]; ok {
			if i == len(ctx.scopes)-1 {
				return "rbp-" + strconv.FormatInt((index+1)*8, 10), true
			} else {
				fmt.Println(offset, index)
				return "rbp+" + strconv.FormatInt((offset-index)*8, 10), true
			}
		}
	}
	return "", false
}

func (ctx *context) addVar(identifier string) {
	scope := ctx.currentScope()
	scope.stackOffset++
	scope.variables[identifier] = scope.stackOffset
}

func (ctx *context) addScope() {
	ctx.scopes = append(ctx.scopes, &scope{0, map[string]int64{}})
}

func (ctx *context) popScope() int64 {
	scope := ctx.currentScope()
	ctx.scopes = ctx.scopes[:len(ctx.scopes)-1]
	return scope.stackOffset
}

type scope struct {
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
  mov rsp, rbp
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
	ctx := &context{[]*scope{{0, map[string]int64{}}}, 0, []uint64{}}
	for _, s := range m.statements {
		fasm += s.Generate(ctx)
	}
	fasm += "  add rsp, " + strconv.FormatInt((ctx.stackOffset())*8, 10) + "\n"
	fasm += "  ret"
	return fasm
}

func (m *ModuleBase) IsModule() {}
