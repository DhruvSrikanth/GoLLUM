	.arch armv8-a


	.text

	.type main, %function
	.global main
	.p2align 2
main:
.L0:
	sub sp, sp, #96
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	mov x1, #129
	str x1, [sp, #24]
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #32
	ldr x1, [sp, #24]
	str x1, [sp, #48]
	ldr x1, [sp, #32]
	str x1, [sp, #56]
	ldr x1, [sp, #48]
	ldr x2, [sp, #56]
	add x3, x1, x2
	str x3, [sp, #64]
	ldr x1, [sp, #64]
	str x1, [sp, #40]
	ldr x1, [sp, #24]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	mov x1, x1
	adrp x2, .FSTR2
	add x2, x2, :lo12:.FSTR2
	mov x0, x2
	bl printf
	ldr x1, [sp, #32]
	str x1, [sp, #80]
	ldr x1, [sp, #80]
	mov x1, x1
	adrp x2, .FSTR3
	add x2, x2, :lo12:.FSTR3
	mov x0, x2
	bl printf
	ldr x1, [sp, #40]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	mov x1, x1
	adrp x2, .FSTR4
	add x2, x2, :lo12:.FSTR4
	mov x0, x2
	bl printf
	b .L1

.L1:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #96
	ret

	.size main, (.-main)


.READ:
	.asciz	"%ld"
	.size	.READ, 4

.FSTR2:
	.asciz	"%ld\n"
	.size	.FSTR2, 6

.FSTR3:
	.asciz	"%ld\n"
	.size	.FSTR3, 6

.FSTR4:
	.asciz	"%ld\n"
	.size	.FSTR4, 6

