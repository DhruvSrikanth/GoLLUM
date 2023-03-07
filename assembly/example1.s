	.arch armv8-a

	.comm .nilnums,8,8

	.text

	.type fib1, %function
	.global fib1
	.p2align 2
fib1:
.L0:
	sub sp, sp, #128
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #128]
	b .L1

.L1:
	ldr x1, [sp, #128]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	mov x2, #2
	cmp x1, x2
	cset x3, lt
	str x3, [sp, #32]
	ldr x1, [sp, #32]
	mov x2, #0
	cmp x1, x2
	b.eq .L3

.L2:
	ldr x1, [sp, #128]
	str x1, [sp, #40]
	ldr x1, [sp, #40]
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #128
	ret

.L3:
	ldr x1, [sp, #128]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	mov x2, #1
	sub x3, x1, x2
	str x3, [sp, #64]
	ldr x1, [sp, #64]
	mov x0, x1
	bl fib1
	str x0, [sp, #72]
	ldr x1, [sp, #128]
	str x1, [sp, #80]
	ldr x1, [sp, #80]
	mov x2, #2
	sub x3, x1, x2
	str x3, [sp, #88]
	ldr x1, [sp, #88]
	mov x0, x1
	bl fib1
	str x0, [sp, #96]
	ldr x1, [sp, #72]
	ldr x2, [sp, #96]
	add x3, x1, x2
	str x3, [sp, #104]
	ldr x1, [sp, #104]
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #128
	ret

.L4:
	b .L5

.L5:
	ldr x1, [sp, #16]
	str x1, [sp, #120]
	ldr x1, [sp, #120]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #128
	ret

	.size fib1, (.-fib1)

	.type fib2, %function
	.global fib2
	.p2align 2
fib2:
.L6:
	sub sp, sp, #144
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #136]
	mov x1, #0
	str x1, [sp, #24]
	mov x1, #1
	str x1, [sp, #32]
	b .L7

.L7:
	ldr x1, [sp, #136]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	mov x2, #0
	cmp x1, x2
	cset x3, ne
	str x3, [sp, #56]
	ldr x1, [sp, #56]
	mov x2, #0
	cmp x1, x2
	b.eq .L9

.L8:
	ldr x1, [sp, #136]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	mov x2, #1
	sub x3, x1, x2
	str x3, [sp, #72]
	ldr x1, [sp, #72]
	str x1, [sp, #136]
	ldr x1, [sp, #24]
	str x1, [sp, #80]
	ldr x1, [sp, #32]
	str x1, [sp, #88]
	ldr x1, [sp, #80]
	ldr x2, [sp, #88]
	add x3, x1, x2
	str x3, [sp, #96]
	ldr x1, [sp, #96]
	str x1, [sp, #40]
	ldr x1, [sp, #32]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	str x1, [sp, #24]
	ldr x1, [sp, #40]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	str x1, [sp, #32]
	b .L7

.L9:
	ldr x1, [sp, #24]
	str x1, [sp, #120]
	ldr x1, [sp, #120]
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #128]
	ldr x1, [sp, #128]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #144
	ret

	.size fib2, (.-fib2)

	.type main, %function
	.global main
	.p2align 2
main:
.L10:
	sub sp, sp, #208
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
	add x1, sp, #40
	bl scanf
	ldr x1, [sp, #24]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	add x1, x1, #0
	str x1, [sp, #72]
	ldr x1, [sp, #40]
	str x1, [sp, #80]
	ldr x1, [sp, #80]
	str x1, [sp, #72]
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #40
	bl scanf
	ldr x1, [sp, #24]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	add x1, x1, #8
	str x1, [sp, #96]
	ldr x1, [sp, #40]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	str x1, [sp, #96]
	ldr x1, [sp, #24]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	add x1, x1, #0
	str x1, [sp, #120]
	ldr x1, [sp, #120]
	str x1, [sp, #128]
	ldr x1, [sp, #128]
	mov x0, x1
	bl fib1
	str x0, [sp, #136]
	ldr x1, [sp, #136]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	str x1, [sp, #144]
	ldr x1, [sp, #144]
	add x1, x1, #8
	str x1, [sp, #152]
	ldr x1, [sp, #152]
	str x1, [sp, #160]
	ldr x1, [sp, #160]
	mov x0, x1
	bl fib2
	str x0, [sp, #168]
	ldr x1, [sp, #168]
	str x1, [sp, #32]
	ldr x1, [sp, #24]
	str x1, [sp, #176]
	ldr x1, [sp, #176]
	str x1, [sp, #184]
	ldr x1, [sp, #184]
	mov x0, x1
	bl free
	ldr x1, [sp, #24]
	str x1, [sp, #192]
	ldr x1, [sp, #32]
	str x1, [sp, #200]
	ldr x1, [sp, #192]
	ldr x2, [sp, #200]
	mov x1, x1
	mov x2, x2
	adrp x3, .FSTR2
	add x3, x3, :lo12:.FSTR2
	mov x0, x3
	bl printf
	b .L11

.L11:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #208]
	ldr x1, [sp, #208]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #208
	ret

	.size main, (.-main)


.READ:
	.asciz	"%ld"
	.size	.READ, 4

.FSTR2:
	.asciz	"c=%ld\nd=%ld"
	.size	.FSTR2, 13

