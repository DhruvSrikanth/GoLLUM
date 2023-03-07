	.arch armv8-a


	.text

	.type fact, %function
	.global fact
	.p2align 2
fact:
.L0:
	sub sp, sp, #112
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #104]
	b .L1

.L1:
	ldr x1, [sp, #104]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	mov x2, #1
	cmp x1, x2
	cset x3, le
	str x3, [sp, #32]
	ldr x1, [sp, #32]
	mov x2, #0
	cmp x1, x2
	b.eq .L3

.L2:
	mov x1, #1
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #40]
	ldr x1, [sp, #40]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #112
	ret

.L3:
	ldr x1, [sp, #104]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	mov x2, #1
	sub x3, x1, x2
	str x3, [sp, #56]
	ldr x1, [sp, #56]
	mov x0, x1
	bl fact
	str x0, [sp, #64]
	ldr x1, [sp, #104]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	ldr x2, [sp, #64]
	mul x3, x1, x2
	str x3, [sp, #80]
	ldr x1, [sp, #80]
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #112
	ret

.L4:
	b .L5

.L5:
	ldr x1, [sp, #16]
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #112
	ret

	.size fact, (.-fact)

	.type main, %function
	.global main
	.p2align 2
main:
.L6:
	sub sp, sp, #112
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	mov x1, #0
	str x1, [sp, #24]
	mov x1, #0
	str x1, [sp, #32]
	b .L7

.L7:
	ldr x1, [sp, #24]
	str x1, [sp, #56]
	mov x1, #1
	ldr x2, [sp, #56]
	eor x3, x1, x2
	str x3, [sp, #64]
	ldr x1, [sp, #64]
	mov x2, #1
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #72]
	ldr x1, [sp, #72]
	mov x2, #0
	cmp x1, x2
	b.eq .L13

.L8:
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #32
	ldr x1, [sp, #32]
	str x1, [sp, #80]
	ldr x1, [sp, #80]
	mov x0, x1
	bl fact
	str x0, [sp, #88]
	ldr x1, [sp, #88]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	mov x1, x1
	adrp x2, .FSTR2
	add x2, x2, :lo12:.FSTR2
	mov x0, x2
	bl printf
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #40
	b .L9

.L9:
	ldr x1, [sp, #40]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	mov x2, #0
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #112]
	ldr x1, [sp, #112]
	mov x2, #0
	cmp x1, x2
	b.eq .L11

.L10:
	mov x1, #1
	str x1, [sp, #24]
	b .L12

.L11:
	b .L12

.L12:
	b .L7

.L13:
	b .L14

.L14:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #120]
	ldr x1, [sp, #120]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #112
	ret

	.size main, (.-main)


.READ:
	.asciz	"%ld"
	.size	.READ, 4

.FSTR2:
	.asciz	"%ld"
	.size	.FSTR2, 4

