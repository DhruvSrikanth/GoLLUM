	.arch armv8-a

	.comm .nilCell,8,8
	.comm .nilRow,8,8
	.comm .nilListNode,8,8
	.comm matrix,8,8
	.comm maxrange,8,8

	.text

	.type timesone, %function
	.global timesone
	.p2align 2
timesone:
.L0:
	sub sp, sp, #80
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #80]
	b .L1

.L1:
	ldr x1, [sp, #80]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	mov x2, #0
	cmp x1, x2
	cset x3, gt
	str x3, [sp, #32]
	ldr x1, [sp, #32]
	mov x2, #0
	cmp x1, x2
	b.eq .L3

.L2:
	adrp x1, maxrange
	add x1, x1, :lo12:maxrange
	ldr x1, [x1]
	str x1, [sp, #40]
	ldr x1, [sp, #40]
	mov x2, #1
	mul x3, x1, x2
	str x3, [sp, #48]
	ldr x1, [sp, #48]
	adrp x2, maxrange
	add x2, x2, :lo12:maxrange
	str x1, [x2]
	ldr x1, [sp, #80]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	mov x2, #1
	sub x3, x1, x2
	str x3, [sp, #64]
	ldr x1, [sp, #64]
	str x1, [sp, #80]
	b .L1

.L3:
	b .L4

.L4:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #80
	ret

	.size timesone, (.-timesone)

	.type find, %function
	.global find
	.p2align 2
find:
.L5:
	sub sp, sp, #208
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #184]
	str x1, [sp, #192]
	b .L6

.L6:
	ldr x2, [sp, #192]
	str x2, [sp, #24]
	adrp x2, .nilListNode
	add x2, x2, :lo12:.nilListNode
	ldr x2, [x2]
	str x2, [sp, #32]
	ldr x2, [sp, #24]
	ldr x3, [sp, #32]
	cmp x2, x3
	cset x4, ne
	str x4, [sp, #40]
	ldr x2, [sp, #40]
	mov x3, #0
	cmp x2, x3
	b.eq .L12

.L7:
	b .L8

.L8:
	ldr x2, [sp, #192]
	str x2, [sp, #48]
	ldr x2, [sp, #48]
	add x2, x2, #8
	str x2, [sp, #56]
	ldr x2, [sp, #56]
	str x2, [sp, #64]
	ldr x2, [sp, #184]
	str x2, [sp, #72]
	ldr x2, [sp, #64]
	ldr x3, [sp, #72]
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #80]
	ldr x2, [sp, #80]
	mov x3, #0
	cmp x2, x3
	b.eq .L10

.L9:
	ldr x2, [sp, #192]
	str x2, [sp, #88]
	ldr x2, [sp, #88]
	add x2, x2, #16
	str x2, [sp, #96]
	ldr x2, [sp, #96]
	str x2, [sp, #104]
	ldr x2, [sp, #104]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #112]
	ldr x2, [sp, #112]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #208
	ret

.L10:
	ldr x2, [sp, #184]
	str x2, [sp, #120]
	ldr x2, [sp, #192]
	str x2, [sp, #128]
	ldr x2, [sp, #128]
	add x2, x2, #0
	str x2, [sp, #136]
	ldr x2, [sp, #136]
	str x2, [sp, #144]
	ldr x2, [sp, #120]
	ldr x3, [sp, #144]
	mov x0, x2
	mov x1, x3
	bl find
	str x0, [sp, #152]
	ldr x2, [sp, #152]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #160]
	ldr x2, [sp, #160]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #208
	ret

.L11:
	b .L13

.L12:
	b .L13

.L13:
	mov x2, #0
	sub x2, x2, #1
	mov x3, #1
	mul x4, x2, x3
	str x4, [sp, #168]
	ldr x2, [sp, #168]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #176]
	ldr x2, [sp, #176]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #208
	ret

	.size find, (.-find)

	.type add, %function
	.global add
	.p2align 2
add:
.L14:
	sub sp, sp, #256
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #216]
	str x1, [sp, #224]
	str x2, [sp, #232]
	b .L15

.L15:
	ldr x3, [sp, #232]
	str x3, [sp, #24]
	adrp x3, .nilListNode
	add x3, x3, :lo12:.nilListNode
	ldr x3, [x3]
	str x3, [sp, #32]
	ldr x3, [sp, #24]
	ldr x4, [sp, #32]
	cmp x3, x4
	cset x5, eq
	str x5, [sp, #40]
	ldr x3, [sp, #40]
	mov x4, #0
	cmp x3, x4
	b.eq .L17

.L16:
	mov x0, #24
	bl malloc
	str x0, [sp, #48]
	ldr x3, [sp, #48]
	str x3, [sp, #56]
	ldr x3, [sp, #56]
	str x3, [sp, #232]
	ldr x3, [sp, #232]
	str x3, [sp, #64]
	ldr x3, [sp, #64]
	add x3, x3, #8
	str x3, [sp, #72]
	ldr x3, [sp, #216]
	str x3, [sp, #80]
	ldr x3, [sp, #80]
	str x3, [sp, #72]
	ldr x3, [sp, #232]
	str x3, [sp, #88]
	ldr x3, [sp, #88]
	add x3, x3, #16
	str x3, [sp, #96]
	ldr x3, [sp, #224]
	str x3, [sp, #104]
	ldr x3, [sp, #104]
	str x3, [sp, #96]
	ldr x3, [sp, #232]
	str x3, [sp, #112]
	ldr x3, [sp, #112]
	add x3, x3, #0
	str x3, [sp, #120]
	adrp x3, .nilListNode
	add x3, x3, :lo12:.nilListNode
	ldr x3, [x3]
	str x3, [sp, #128]
	ldr x3, [sp, #128]
	str x3, [sp, #120]
	b .L18

.L17:
	ldr x3, [sp, #232]
	str x3, [sp, #136]
	ldr x3, [sp, #136]
	add x3, x3, #0
	str x3, [sp, #144]
	ldr x3, [sp, #216]
	str x3, [sp, #152]
	ldr x3, [sp, #224]
	str x3, [sp, #160]
	ldr x3, [sp, #232]
	str x3, [sp, #168]
	ldr x3, [sp, #168]
	add x3, x3, #0
	str x3, [sp, #176]
	ldr x3, [sp, #176]
	str x3, [sp, #184]
	ldr x3, [sp, #152]
	ldr x4, [sp, #160]
	ldr x5, [sp, #184]
	mov x0, x3
	mov x1, x4
	mov x2, x5
	bl add
	str x0, [sp, #192]
	ldr x3, [sp, #192]
	str x3, [sp, #144]
	b .L18

.L18:
	ldr x3, [sp, #232]
	str x3, [sp, #200]
	ldr x3, [sp, #200]
	str x3, [sp, #16]
	ldr x3, [sp, #16]
	str x3, [sp, #208]
	ldr x3, [sp, #208]
	mov x0, x3
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #256
	ret

	.size add, (.-add)

	.type factorial, %function
	.global factorial
	.p2align 2
factorial:
.L19:
	sub sp, sp, #288
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #264]
	str x1, [sp, #272]
	mov x2, #100
	mov x0, x2
	bl timesone
	str x0, [sp, #40]
	b .L20

.L20:
	ldr x2, [sp, #264]
	str x2, [sp, #48]
	ldr x2, [sp, #48]
	mov x3, #1
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #56]
	ldr x2, [sp, #56]
	mov x3, #0
	cmp x2, x3
	b.eq .L22

.L21:
	mov x2, #1
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #64]
	ldr x2, [sp, #64]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #288
	ret

.L22:
	ldr x2, [sp, #264]
	str x2, [sp, #72]
	ldr x2, [sp, #272]
	str x2, [sp, #80]
	ldr x2, [sp, #72]
	ldr x3, [sp, #80]
	mov x0, x2
	mov x1, x3
	bl find
	str x0, [sp, #88]
	ldr x2, [sp, #88]
	str x2, [sp, #24]
	b .L23

.L23:
	mov x2, #0
	sub x2, x2, #1
	mov x3, #1
	mul x4, x2, x3
	str x4, [sp, #96]
	ldr x2, [sp, #24]
	str x2, [sp, #104]
	ldr x2, [sp, #104]
	ldr x3, [sp, #96]
	cmp x2, x3
	cset x4, ne
	str x4, [sp, #112]
	ldr x2, [sp, #112]
	mov x3, #0
	cmp x2, x3
	b.eq .L25

.L24:
	ldr x2, [sp, #24]
	str x2, [sp, #120]
	ldr x2, [sp, #120]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #128]
	ldr x2, [sp, #128]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #288
	ret

.L25:
	b .L26

.L26:
	ldr x2, [sp, #264]
	str x2, [sp, #136]
	ldr x2, [sp, #136]
	mov x3, #1
	sub x4, x2, x3
	str x4, [sp, #144]
	ldr x2, [sp, #272]
	str x2, [sp, #152]
	ldr x2, [sp, #144]
	ldr x3, [sp, #152]
	mov x0, x2
	mov x1, x3
	bl factorial
	str x0, [sp, #160]
	ldr x2, [sp, #264]
	str x2, [sp, #168]
	ldr x2, [sp, #168]
	ldr x3, [sp, #160]
	mul x4, x2, x3
	str x4, [sp, #176]
	ldr x2, [sp, #176]
	str x2, [sp, #32]
	b .L27

.L27:
	ldr x2, [sp, #32]
	str x2, [sp, #184]
	ldr x2, [sp, #184]
	mov x3, #3
	sdiv x4, x2, x3
	str x4, [sp, #192]
	ldr x2, [sp, #192]
	mov x3, #0
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #200]
	ldr x2, [sp, #200]
	mov x3, #0
	cmp x2, x3
	b.eq .L29

.L28:
	ldr x2, [sp, #264]
	str x2, [sp, #208]
	ldr x2, [sp, #32]
	str x2, [sp, #216]
	ldr x2, [sp, #272]
	str x2, [sp, #224]
	ldr x2, [sp, #208]
	ldr x3, [sp, #216]
	ldr x4, [sp, #224]
	mov x0, x2
	mov x1, x3
	mov x2, x4
	bl add
	str x0, [sp, #232]
	b .L30

.L29:
	b .L30

.L30:
	ldr x2, [sp, #32]
	str x2, [sp, #240]
	ldr x2, [sp, #240]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #248]
	ldr x2, [sp, #248]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #288
	ret

.L31:
	b .L32

.L32:
	ldr x2, [sp, #16]
	str x2, [sp, #256]
	ldr x2, [sp, #256]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #288
	ret

	.size factorial, (.-factorial)

	.type maxfactorial, %function
	.global maxfactorial
	.p2align 2
maxfactorial:
.L33:
	sub sp, sp, #304
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #280]
	str x1, [sp, #288]
	ldr x2, [sp, #280]
	str x2, [sp, #40]
	ldr x2, [sp, #40]
	add x2, x2, #0
	str x2, [sp, #48]
	adrp x2, .nilListNode
	add x2, x2, :lo12:.nilListNode
	ldr x2, [x2]
	str x2, [sp, #56]
	ldr x2, [sp, #56]
	str x2, [sp, #48]
	adrp x2, matrix
	add x2, x2, :lo12:matrix
	ldr x2, [x2]
	str x2, [sp, #64]
	ldr x2, [sp, #64]
	str x2, [sp, #16]
	b .L34

.L34:
	ldr x2, [sp, #16]
	str x2, [sp, #72]
	adrp x2, .nilRow
	add x2, x2, :lo12:.nilRow
	ldr x2, [x2]
	str x2, [sp, #80]
	ldr x2, [sp, #72]
	ldr x3, [sp, #80]
	cmp x2, x3
	cset x4, ne
	str x4, [sp, #88]
	ldr x2, [sp, #88]
	mov x3, #0
	cmp x2, x3
	b.eq .L43

.L35:
	ldr x2, [sp, #16]
	str x2, [sp, #96]
	ldr x2, [sp, #96]
	add x2, x2, #8
	str x2, [sp, #104]
	ldr x2, [sp, #104]
	str x2, [sp, #112]
	ldr x2, [sp, #112]
	str x2, [sp, #24]
	ldr x2, [sp, #16]
	str x2, [sp, #120]
	ldr x2, [sp, #120]
	add x2, x2, #0
	str x2, [sp, #128]
	ldr x2, [sp, #128]
	str x2, [sp, #136]
	ldr x2, [sp, #136]
	str x2, [sp, #16]
	b .L36

.L36:
	ldr x2, [sp, #24]
	str x2, [sp, #144]
	adrp x2, .nilCell
	add x2, x2, :lo12:.nilCell
	ldr x2, [x2]
	str x2, [sp, #152]
	ldr x2, [sp, #144]
	ldr x3, [sp, #152]
	cmp x2, x3
	cset x4, ne
	str x4, [sp, #160]
	ldr x2, [sp, #160]
	mov x3, #0
	cmp x2, x3
	b.eq .L42

.L37:
	ldr x2, [sp, #24]
	str x2, [sp, #168]
	ldr x2, [sp, #168]
	add x2, x2, #8
	str x2, [sp, #176]
	ldr x2, [sp, #176]
	str x2, [sp, #184]
	ldr x2, [sp, #280]
	str x2, [sp, #192]
	ldr x2, [sp, #184]
	ldr x3, [sp, #192]
	mov x0, x2
	mov x1, x3
	bl factorial
	str x0, [sp, #200]
	ldr x2, [sp, #200]
	str x2, [sp, #32]
	ldr x2, [sp, #24]
	str x2, [sp, #208]
	ldr x2, [sp, #208]
	add x2, x2, #0
	str x2, [sp, #216]
	ldr x2, [sp, #216]
	str x2, [sp, #224]
	ldr x2, [sp, #224]
	str x2, [sp, #24]
	b .L38

.L38:
	ldr x2, [sp, #32]
	str x2, [sp, #232]
	ldr x2, [sp, #288]
	str x2, [sp, #240]
	ldr x2, [sp, #232]
	ldr x3, [sp, #240]
	cmp x2, x3
	cset x4, gt
	str x4, [sp, #248]
	ldr x2, [sp, #248]
	mov x3, #0
	cmp x2, x3
	b.eq .L40

.L39:
	ldr x2, [sp, #32]
	str x2, [sp, #256]
	ldr x2, [sp, #256]
	str x2, [sp, #288]
	b .L41

.L40:
	b .L41

.L41:
	b .L36

.L42:
	b .L34

.L43:
	ldr x2, [sp, #288]
	str x2, [sp, #264]
	ldr x2, [sp, #264]
	mov x1, x2
	adrp x3, .FSTR1
	add x3, x3, :lo12:.FSTR1
	mov x0, x3
	bl printf
	b .L44

.L44:
	mov x2, #0
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #272]
	ldr x2, [sp, #272]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #304
	ret

	.size maxfactorial, (.-maxfactorial)

	.type newvalue, %function
	.global newvalue
	.p2align 2
newvalue:
.L45:
	sub sp, sp, #128
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #112]
	str x1, [sp, #120]
	mov x2, #0
	mov x0, x2
	bl timesone
	str x0, [sp, #32]
	ldr x2, [sp, #112]
	str x2, [sp, #40]
	ldr x2, [sp, #120]
	str x2, [sp, #48]
	ldr x2, [sp, #40]
	ldr x3, [sp, #48]
	mul x4, x2, x3
	str x4, [sp, #56]
	ldr x2, [sp, #56]
	str x2, [sp, #24]
	adrp x2, maxrange
	add x2, x2, :lo12:maxrange
	ldr x2, [x2]
	str x2, [sp, #64]
	ldr x2, [sp, #24]
	str x2, [sp, #72]
	ldr x2, [sp, #64]
	ldr x3, [sp, #72]
	sdiv x4, x2, x3
	str x4, [sp, #80]
	ldr x2, [sp, #112]
	str x2, [sp, #88]
	ldr x2, [sp, #80]
	ldr x3, [sp, #88]
	add x4, x2, x3
	str x4, [sp, #96]
	ldr x2, [sp, #96]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #104]
	ldr x2, [sp, #104]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #128
	ret

	.size newvalue, (.-newvalue)

	.type newcell, %function
	.global newcell
	.p2align 2
newcell:
.L46:
	sub sp, sp, #224
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #184]
	str x1, [sp, #192]
	str x2, [sp, #200]
	ldr x3, [sp, #184]
	str x3, [sp, #24]
	ldr x3, [sp, #24]
	add x3, x3, #8
	str x3, [sp, #32]
	ldr x3, [sp, #192]
	str x3, [sp, #40]
	ldr x3, [sp, #200]
	str x3, [sp, #48]
	ldr x3, [sp, #40]
	ldr x4, [sp, #48]
	mov x0, x3
	mov x1, x4
	bl newvalue
	str x0, [sp, #56]
	ldr x3, [sp, #56]
	str x3, [sp, #32]
	b .L47

.L47:
	ldr x3, [sp, #200]
	str x3, [sp, #64]
	ldr x3, [sp, #64]
	mov x4, #1
	cmp x3, x4
	cset x5, gt
	str x5, [sp, #72]
	ldr x3, [sp, #72]
	mov x4, #0
	cmp x3, x4
	b.eq .L49

.L48:
	ldr x3, [sp, #184]
	str x3, [sp, #80]
	ldr x3, [sp, #80]
	add x3, x3, #0
	str x3, [sp, #88]
	mov x0, #16
	bl malloc
	str x0, [sp, #96]
	ldr x3, [sp, #96]
	str x3, [sp, #104]
	ldr x3, [sp, #192]
	str x3, [sp, #112]
	ldr x3, [sp, #200]
	str x3, [sp, #120]
	ldr x3, [sp, #120]
	mov x4, #1
	sub x5, x3, x4
	str x5, [sp, #128]
	ldr x3, [sp, #104]
	ldr x4, [sp, #112]
	ldr x5, [sp, #128]
	mov x0, x3
	mov x1, x4
	mov x2, x5
	bl newcell
	str x0, [sp, #136]
	ldr x3, [sp, #136]
	str x3, [sp, #88]
	b .L50

.L49:
	ldr x3, [sp, #184]
	str x3, [sp, #144]
	ldr x3, [sp, #144]
	add x3, x3, #0
	str x3, [sp, #152]
	adrp x3, .nilCell
	add x3, x3, :lo12:.nilCell
	ldr x3, [x3]
	str x3, [sp, #160]
	ldr x3, [sp, #160]
	str x3, [sp, #152]
	b .L50

.L50:
	ldr x3, [sp, #184]
	str x3, [sp, #168]
	ldr x3, [sp, #168]
	str x3, [sp, #16]
	ldr x3, [sp, #16]
	str x3, [sp, #176]
	ldr x3, [sp, #176]
	mov x0, x3
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #224
	ret

	.size newcell, (.-newcell)

	.type newrow, %function
	.global newrow
	.p2align 2
newrow:
.L51:
	sub sp, sp, #240
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #200]
	str x1, [sp, #208]
	str x2, [sp, #216]
	ldr x3, [sp, #200]
	str x3, [sp, #24]
	ldr x3, [sp, #24]
	add x3, x3, #8
	str x3, [sp, #32]
	mov x0, #16
	bl malloc
	str x0, [sp, #40]
	ldr x3, [sp, #40]
	str x3, [sp, #48]
	ldr x3, [sp, #208]
	str x3, [sp, #56]
	ldr x3, [sp, #216]
	str x3, [sp, #64]
	ldr x3, [sp, #48]
	ldr x4, [sp, #56]
	ldr x5, [sp, #64]
	mov x0, x3
	mov x1, x4
	mov x2, x5
	bl newcell
	str x0, [sp, #72]
	ldr x3, [sp, #72]
	str x3, [sp, #32]
	b .L52

.L52:
	ldr x3, [sp, #208]
	str x3, [sp, #80]
	ldr x3, [sp, #80]
	mov x4, #1
	cmp x3, x4
	cset x5, gt
	str x5, [sp, #88]
	ldr x3, [sp, #88]
	mov x4, #0
	cmp x3, x4
	b.eq .L54

.L53:
	ldr x3, [sp, #200]
	str x3, [sp, #96]
	ldr x3, [sp, #96]
	add x3, x3, #0
	str x3, [sp, #104]
	mov x0, #16
	bl malloc
	str x0, [sp, #112]
	ldr x3, [sp, #112]
	str x3, [sp, #120]
	ldr x3, [sp, #208]
	str x3, [sp, #128]
	ldr x3, [sp, #128]
	mov x4, #1
	sub x5, x3, x4
	str x5, [sp, #136]
	ldr x3, [sp, #216]
	str x3, [sp, #144]
	ldr x3, [sp, #120]
	ldr x4, [sp, #136]
	ldr x5, [sp, #144]
	mov x0, x3
	mov x1, x4
	mov x2, x5
	bl newrow
	str x0, [sp, #152]
	ldr x3, [sp, #152]
	str x3, [sp, #104]
	b .L55

.L54:
	ldr x3, [sp, #200]
	str x3, [sp, #160]
	ldr x3, [sp, #160]
	add x3, x3, #0
	str x3, [sp, #168]
	adrp x3, .nilRow
	add x3, x3, :lo12:.nilRow
	ldr x3, [x3]
	str x3, [sp, #176]
	ldr x3, [sp, #176]
	str x3, [sp, #168]
	b .L55

.L55:
	ldr x3, [sp, #200]
	str x3, [sp, #184]
	ldr x3, [sp, #184]
	str x3, [sp, #16]
	ldr x3, [sp, #16]
	str x3, [sp, #192]
	ldr x3, [sp, #192]
	mov x0, x3
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #240
	ret

	.size newrow, (.-newrow)

	.type newmatrix, %function
	.global newmatrix
	.p2align 2
newmatrix:
.L56:
	sub sp, sp, #96
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #64]
	str x1, [sp, #72]
	mov x0, #16
	bl malloc
	str x0, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #24]
	ldr x2, [sp, #64]
	str x2, [sp, #32]
	ldr x2, [sp, #72]
	str x2, [sp, #40]
	ldr x2, [sp, #24]
	ldr x3, [sp, #32]
	ldr x4, [sp, #40]
	mov x0, x2
	mov x1, x3
	mov x2, x4
	bl newrow
	str x0, [sp, #48]
	ldr x2, [sp, #48]
	adrp x3, matrix
	add x3, x3, :lo12:matrix
	str x2, [x3]
	b .L57

.L57:
	mov x2, #0
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #56]
	ldr x2, [sp, #56]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #96
	ret

	.size newmatrix, (.-newmatrix)

	.type getmatrixsize, %function
	.global getmatrixsize
	.p2align 2
getmatrixsize:
.L58:
	sub sp, sp, #96
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #80]
	b .L59

.L59:
	ldr x1, [sp, #80]
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	mov x2, #0
	cmp x1, x2
	cset x3, le
	str x3, [sp, #24]
	ldr x1, [sp, #24]
	mov x2, #0
	cmp x1, x2
	b.eq .L61

.L60:
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #80
	bl scanf
	ldr x1, [sp, #80]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	mov x0, x1
	bl getmatrixsize
	str x0, [sp, #40]
	b .L62

.L61:
	ldr x1, [sp, #80]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #96
	ret

.L62:
	ldr x1, [sp, #80]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #96
	ret

	.size getmatrixsize, (.-getmatrixsize)

	.type getmaxrange, %function
	.global getmaxrange
	.p2align 2
getmaxrange:
.L63:
	sub sp, sp, #96
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #80]
	b .L64

.L64:
	adrp x1, maxrange
	add x1, x1, :lo12:maxrange
	ldr x1, [x1]
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	mov x2, #1
	cmp x1, x2
	cset x3, le
	str x3, [sp, #24]
	ldr x1, [sp, #24]
	mov x2, #0
	cmp x1, x2
	b.eq .L66

.L65:
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	adrp x2, maxrange
	add x2, x2, :lo12:maxrange
	mov x1, x2
	bl scanf
	adrp x1, maxrange
	add x1, x1, :lo12:maxrange
	ldr x1, [x1]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	mov x0, x1
	bl getmaxrange
	str x0, [sp, #40]
	b .L67

.L66:
	adrp x1, maxrange
	add x1, x1, :lo12:maxrange
	ldr x1, [x1]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #96
	ret

.L67:
	adrp x1, maxrange
	add x1, x1, :lo12:maxrange
	ldr x1, [x1]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #96
	ret

	.size getmaxrange, (.-getmaxrange)

	.type main, %function
	.global main
	.p2align 2
main:
.L68:
	sub sp, sp, #128
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	mov x1, #0
	str x1, [sp, #24]
	mov x1, #0
	str x1, [sp, #32]
	mov x1, #0
	adrp x2, maxrange
	add x2, x2, :lo12:maxrange
	str x1, [x2]
	ldr x1, [sp, #24]
	str x1, [sp, #40]
	ldr x1, [sp, #40]
	mov x0, x1
	bl getmatrixsize
	str x0, [sp, #48]
	ldr x1, [sp, #48]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	str x1, [sp, #32]
	adrp x1, maxrange
	add x1, x1, :lo12:maxrange
	ldr x1, [x1]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	mov x0, x1
	bl getmaxrange
	str x0, [sp, #72]
	ldr x1, [sp, #72]
	adrp x2, maxrange
	add x2, x2, :lo12:maxrange
	str x1, [x2]
	ldr x1, [sp, #24]
	str x1, [sp, #80]
	ldr x1, [sp, #32]
	str x1, [sp, #88]
	ldr x1, [sp, #80]
	ldr x2, [sp, #88]
	mov x0, x1
	mov x1, x2
	bl newmatrix
	str x0, [sp, #96]
	mov x0, #24
	bl malloc
	str x0, [sp, #104]
	ldr x1, [sp, #104]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	mov x2, #0
	mov x0, x1
	mov x1, x2
	bl maxfactorial
	str x0, [sp, #120]
	b .L69

.L69:
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


.FSTR1:
	.asciz	"Max Factorial=%ld"
	.size	.FSTR1, 18

.READ:
	.asciz	"%ld"
	.size	.READ, 4

