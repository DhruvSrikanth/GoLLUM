	.arch armv8-a


	.text

	.type isqrt, %function
	.global isqrt
	.p2align 2
isqrt:
.L0:
	sub sp, sp, #144
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #136]
	mov x1, #1
	str x1, [sp, #24]
	mov x1, #3
	str x1, [sp, #32]
	b .L1

.L1:
	ldr x1, [sp, #24]
	str x1, [sp, #40]
	ldr x1, [sp, #136]
	str x1, [sp, #48]
	ldr x1, [sp, #40]
	ldr x2, [sp, #48]
	cmp x1, x2
	cset x3, le
	str x3, [sp, #56]
	ldr x1, [sp, #56]
	mov x2, #0
	cmp x1, x2
	b.eq .L3

.L2:
	ldr x1, [sp, #24]
	str x1, [sp, #64]
	ldr x1, [sp, #32]
	str x1, [sp, #72]
	ldr x1, [sp, #64]
	ldr x2, [sp, #72]
	add x3, x1, x2
	str x3, [sp, #80]
	ldr x1, [sp, #80]
	str x1, [sp, #24]
	ldr x1, [sp, #32]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	mov x2, #2
	add x3, x1, x2
	str x3, [sp, #96]
	ldr x1, [sp, #96]
	str x1, [sp, #32]
	b .L1

.L3:
	ldr x1, [sp, #32]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	mov x2, #2
	sdiv x3, x1, x2
	str x3, [sp, #112]
	ldr x1, [sp, #112]
	mov x2, #1
	sub x3, x1, x2
	str x3, [sp, #120]
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

	.size isqrt, (.-isqrt)

	.type prime, %function
	.global prime
	.p2align 2
prime:
.L4:
	sub sp, sp, #224
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #224]
	b .L5

.L5:
	ldr x1, [sp, #224]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	mov x2, #2
	cmp x1, x2
	cset x3, lt
	str x3, [sp, #56]
	ldr x1, [sp, #56]
	mov x2, #0
	cmp x1, x2
	b.eq .L7

.L6:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #224
	ret

.L7:
	ldr x1, [sp, #224]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	mov x0, x1
	bl isqrt
	str x0, [sp, #80]
	ldr x1, [sp, #80]
	str x1, [sp, #24]
	mov x1, #2
	str x1, [sp, #32]
	b .L8

.L8:
	ldr x1, [sp, #32]
	str x1, [sp, #88]
	ldr x1, [sp, #24]
	str x1, [sp, #96]
	ldr x1, [sp, #88]
	ldr x2, [sp, #96]
	cmp x1, x2
	cset x3, le
	str x3, [sp, #104]
	ldr x1, [sp, #104]
	mov x2, #0
	cmp x1, x2
	b.eq .L14

.L9:
	ldr x1, [sp, #224]
	str x1, [sp, #112]
	ldr x1, [sp, #32]
	str x1, [sp, #120]
	ldr x1, [sp, #112]
	ldr x2, [sp, #120]
	sdiv x3, x1, x2
	str x3, [sp, #128]
	ldr x1, [sp, #32]
	str x1, [sp, #136]
	ldr x1, [sp, #128]
	ldr x2, [sp, #136]
	mul x3, x1, x2
	str x3, [sp, #144]
	ldr x1, [sp, #224]
	str x1, [sp, #152]
	ldr x1, [sp, #152]
	ldr x2, [sp, #144]
	sub x3, x1, x2
	str x3, [sp, #160]
	ldr x1, [sp, #160]
	str x1, [sp, #40]
	b .L10

.L10:
	ldr x1, [sp, #40]
	str x1, [sp, #168]
	ldr x1, [sp, #168]
	mov x2, #0
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #176]
	ldr x1, [sp, #176]
	mov x2, #0
	cmp x1, x2
	b.eq .L12

.L11:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #184]
	ldr x1, [sp, #184]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #224
	ret

.L12:
	b .L13

.L13:
	ldr x1, [sp, #32]
	str x1, [sp, #192]
	ldr x1, [sp, #192]
	mov x2, #1
	add x3, x1, x2
	str x3, [sp, #200]
	ldr x1, [sp, #200]
	str x1, [sp, #32]
	b .L8

.L14:
	mov x1, #1
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #208]
	ldr x1, [sp, #208]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #224
	ret

.L15:
	b .L16

.L16:
	ldr x1, [sp, #16]
	str x1, [sp, #216]
	ldr x1, [sp, #216]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #224
	ret

	.size prime, (.-prime)

	.type main, %function
	.global main
	.p2align 2
main:
.L17:
	sub sp, sp, #112
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #24
	bl scanf
	mov x1, #0
	str x1, [sp, #32]
	b .L18

.L18:
	ldr x1, [sp, #32]
	str x1, [sp, #40]
	ldr x1, [sp, #24]
	str x1, [sp, #48]
	ldr x1, [sp, #40]
	ldr x2, [sp, #48]
	cmp x1, x2
	cset x3, le
	str x3, [sp, #56]
	ldr x1, [sp, #56]
	mov x2, #0
	cmp x1, x2
	b.eq .L24

.L19:
	b .L20

.L20:
	ldr x1, [sp, #32]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	mov x0, x1
	bl prime
	str x0, [sp, #72]
	ldr x1, [sp, #72]
	mov x2, #1
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #80]
	ldr x1, [sp, #80]
	mov x2, #0
	cmp x1, x2
	b.eq .L22

.L21:
	ldr x1, [sp, #32]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	mov x1, x1
	adrp x2, .FSTR2
	add x2, x2, :lo12:.FSTR2
	mov x0, x2
	bl printf
	b .L23

.L22:
	b .L23

.L23:
	ldr x1, [sp, #32]
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	mov x2, #1
	add x3, x1, x2
	str x3, [sp, #104]
	ldr x1, [sp, #104]
	str x1, [sp, #32]
	b .L18

.L24:
	b .L25

.L25:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
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

