; Module

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
  mov rax, 5
  push rax
  mov rax, [rbp-16]
  push rax
  mov rax, 2
  push rax
  pop r8
  pop r9
  xor rax, rax
  cmp r9, r8
  setg al
  test rax,rax
  jz .else1
  push rbp
  mov rbp, rsp
  mov rax, 2
  ; Setup Call Stack Frame
  push rbp
  mov rbp, rsp
  
  ; Call Function
  push rax
  call printInt
  add rsp, 8
  
  ; Reset Stack Frame
  mov rsp, rbp
  pop rbp
  add rsp, 0*8
  mov rsp, rbp
  pop rbp
  jmp .endif1
.else1:
  mov rax, [rbp-16]
  push rax
  mov rax, 1
  push rax
  pop r8
  pop r9
  xor rax, rax
  cmp r9, r8
  setg al
  test rax,rax
  jz .else2
  push rbp
  mov rbp, rsp
  mov rax, 1
  ; Setup Call Stack Frame
  push rbp
  mov rbp, rsp
  
  ; Call Function
  push rax
  call printInt
  add rsp, 8
  
  ; Reset Stack Frame
  mov rsp, rbp
  pop rbp
  add rsp, 0*8
  mov rsp, rbp
  pop rbp
  jmp .endif2
.else2:
  push rbp
  mov rbp, rsp
  mov rax, 0
  ; Setup Call Stack Frame
  push rbp
  mov rbp, rsp
  
  ; Call Function
  push rax
  call printInt
  add rsp, 8
  
  ; Reset Stack Frame
  mov rsp, rbp
  pop rbp
  add rsp, 0*8
  mov rsp, rbp
  pop rbp
.endif2:
.endif1:
  mov rax, 5
  push rax
  mov rax, [rbp-16]
  push rax
  mov rax, 2
  push rax
  pop r8
  pop r9
  xor rax, rax
  cmp r9, r8
  setg al
  test rax,rax
  jz .else3
  push rbp
  mov rbp, rsp
  mov rax, 4
  mov [rbp+8], rax
  add rsp, 0*8
  mov rsp, rbp
  pop rbp
  jmp .endif3
.else3:
.endif3:
  mov rax, [rbp-24]
  ; Setup Call Stack Frame
  push rbp
  mov rbp, rsp
  
  ; Call Function
  push rax
  call printInt
  add rsp, 8
  
  ; Reset Stack Frame
  mov rsp, rbp
  pop rbp
  add rsp, 16
  ret