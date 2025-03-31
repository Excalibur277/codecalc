; fasm demonstration of printing int

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


macro CallFunc funcname {
  ; Setup Call Stack Frame
  push rbp
  mov rbp, rsp

  ; Call Function
  push rcx
  call funcname
  add rsp, 8

  ; Reset Stack Frame
  pop rbp 
}

routine:
  mov rcx, 250

  .repeat:
  push rcx

  CallFunc printInt

  pop rcx
  loop .repeat
 

  mov rcx, -1230
  CallFunc printInt


  ret
