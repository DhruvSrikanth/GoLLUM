	.arch armv8-a

	.comm .nilnums,8,8

	.text

	.type main, %function
	.global main
	.p2align 2
main:
.L0:
	sub sp, sp, #128
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	mov x0, #16
	bl malloc

	str x0, [sp, #48]
	ldr x1, [sp, #48]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	str x1, [sp, #24]
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #32
	bl scanf

	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #40
	bl scanf

	ldr x1, [sp, #24]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	add x1, x1, #0
	str x1, [sp, #72]

	ldr x1, [sp, #32]
	str x1, [sp, #80]

	ldr x1, [sp, #80]
	ldr x3, [sp, #72]
	str x1, [x3]
	
	ldr x1, [sp, #24]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	add x1, x1, #8
	str x1, [sp, #96]
	ldr x1, [sp, #40]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	ldr x3, [sp, #96]
	str x1, [x3]
	ldr x1, [sp, #24]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	str x1, [sp, #120]
	ldr x1, [sp, #120]
	mov x0, x1
	bl free
	b .L1

.L1:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #128]
	ldr x1, [sp, #128]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #128
	ret

	.size main, (.-main)


.READ:
	.asciz	"%ld"
	.size	.READ, 4

