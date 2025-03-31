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
  mov rax, 1
  push rax
  mov rax, 5
  push rax
  pop r8
  pop rax
  add rax, r8
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
  mov rax, 43
  push rax
  mov rax, [rbp-16]
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
  mov rax, 21334
  mov [rbp-16], rax
  mov rax, [rbp-16]
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
  mov rax, 2
  mov [rbp-16], rax
  mov rax, 5
  push rax
  mov rax, [rbp-16]
  push rax
  pop r8
  pop rax
  add rax, r8
  push rax
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
  mov rax, [rbp-16]
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
  mov rax, [rbp-16]
  push rax
  mov rax, 34
  push rax
  pop r8
  pop rax
  sub rax, r8
  push rax
  mov rax, [rbp-24]
  push rax
  pop r8
  pop rax
  add rax, r8
  mov [rbp-16], rax
  mov rax, [rbp-16]
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
  mov rax, [rbp-24]
  push rax
  mov rax, [rbp-16]
  push rax
  pop r8
  pop rax
  imul rax, r8
  push rax
  mov rax, [rbp-32]
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
  mov rax, 1
  push rax
  mov rax, 5
  push rax
  pop r8
  pop rax
  add rax, r8
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
  mov rax, 43
  mov [rbp-16], rax
  mov rax, [rbp-16]
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
  mov rax, 21334
  mov [rbp-16], rax
  mov rax, [rbp-16]
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
  mov rax, 2
  mov [rbp-16], rax
  mov rax, 5
  push rax
  mov rax, [rbp-16]
  push rax
  pop r8
  pop rax
  add rax, r8
  mov [rbp-24], rax
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
  mov rax, [rbp-16]
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
  mov rax, [rbp-16]
  push rax
  mov rax, 34
  push rax
  pop r8
  pop rax
  sub rax, r8
  push rax
  mov rax, [rbp-24]
  push rax
  pop r8
  pop rax
  add rax, r8
  mov [rbp-16], rax
  mov rax, [rbp-16]
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
  mov rax, [rbp-24]
  push rax
  mov rax, [rbp-16]
  push rax
  pop r8
  pop rax
  imul rax, r8
  mov [rbp-32], rax
  mov rax, [rbp-32]
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
  mov rax, 5
  push rax
  mov rax, 0
  push rax
  pop r8
  pop rax
  test r8,r8
  jz panic
  xor rdx, rdx
  idiv qword r8
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
  mov rax, 5
  push rax
  mov rax, 2
  neg rax
  push rax
  pop r8
  pop rax
  test r8,r8
  jz panic
  xor rdx, rdx
  idiv qword r8
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
  mov rax, 2
  push rax
  mov rax, 4
  push rax
  pop r8
  pop rax
  test r8,r8
  jz panic
  xor rdx, rdx
  idiv qword r8
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
  add rsp, 24
  ret