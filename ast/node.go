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

func (ctx *context) getVar(identifier string) (stackAddress string, dataType dataType, ok bool) {
	offset := int64(0)
	for i := len(ctx.scopes) - 1; i >= 0; i-- {
		if i != len(ctx.scopes)-1 {
			offset += ctx.scopes[i].stackOffset + 1
		}
		if v, ok := ctx.scopes[i].variables[identifier]; ok {
			if i == len(ctx.scopes)-1 {
				return "rbp-" + strconv.FormatInt((v.index+1)*8, 10), v.dataType, true
			} else {
				return "rbp+" + strconv.FormatInt((offset-v.index)*8, 10), v.dataType, true
			}
		}
	}
	return "", Integer, false
}

func (ctx *context) addVar(identifier string, dataType dataType) {
	for i := len(ctx.scopes) - 1; i >= 0; i-- {
		if _, ok := ctx.scopes[i].variables[identifier]; ok {
			panic(fmt.Sprintf("Variable %s has already been allocated. It should not be reallocated.", identifier))
		}
	}
	scope := ctx.currentScope()
	scope.stackOffset++
	scope.variables[identifier] = variable{scope.stackOffset, dataType}
}

func (ctx *context) addScope() {
	ctx.scopes = append(ctx.scopes, &scope{0, map[string]variable{}})
}

func (ctx *context) popScope() int64 {
	scope := ctx.currentScope()
	ctx.scopes = ctx.scopes[:len(ctx.scopes)-1]
	return scope.stackOffset
}

type dataType uint8

const (
	Integer dataType = iota
	Array
)

type variable struct {
	index    int64
	dataType dataType
}
type scope struct {
	stackOffset int64
	variables   map[string]variable
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

segment readable writeable

panic_msg db 'Panic! Attempting to divide or mod by 0, exiting early.',0xA
panic_msg_size = $-panic_msg
arraypanic_msg db 'Panic! Attempting to initilize array with size less than or equal to 0, exiting early.',0xA
arraypanic_msg_size = $-arraypanic_msg

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
  lea rsi, [rsp+rcx]
  mov rdi, 1 ; STDOUT
  mov rax, 1 ; sys_write
  syscall

  ; Clear up String Buffer
  add rsp, 21

  ret


exit:
  mov rax, 60 ; exit
  mov rdi, 0  ; no error 
  syscall

panic:
  mov rdx, panic_msg_size
  lea rsi, [panic_msg]
  mov rdi, 2 ; STDERR
  mov rax, 1 ; sys_write
  syscall

  mov rax, 60 ; exit
  mov rdi, 0  ; error 
  syscall

arraypanic:
  mov rdx, arraypanic_msg_size
  lea rsi, [arraypanic_msg]
  mov rdi, 2 ; STDERR
  mov rax, 1 ; sys_write
  syscall

  mov rax, 60 ; exit
  mov rdi, 0  ; error 
  syscall

routine:
`
	ctx := &context{[]*scope{{0, map[string]variable{}}}, 0, []uint64{}}
	for _, s := range m.statements {
		fasm += s.Generate(ctx)
	}
	fasm += "  add rsp, " + strconv.FormatInt((ctx.stackOffset())*8, 10) + "\n"
	fasm += "  ret"
	return fasm
}

func (m *ModuleBase) IsModule() {}
