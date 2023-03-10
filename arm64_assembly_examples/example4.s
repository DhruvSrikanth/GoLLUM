	.arch armv8-a

	.comm .nilNode,8,8
	.comm head,8,8
	.comm tail,8,8

	.text

	.type Add, %function
	.global Add
	.p2align 2
Add:
.L0:
	sub sp, sp, #176
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #176]
	mov x0, #16
	bl malloc
	str x0, [sp, #32]
	ldr x1, [sp, #32]
	str x1, [sp, #40]
	ldr x1, [sp, #40]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	add x1, x1, #0
	str x1, [sp, #56]
	ldr x1, [sp, #176]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	ldr x3, [sp, #56]
	str x1, [x3]
	ldr x1, [sp, #24]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	add x1, x1, #8
	str x1, [sp, #80]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	ldr x3, [sp, #80]
	str x1, [x3]
	b .L1

.L1:
	adrp x1, head
	add x1, x1, :lo12:head
	ldr x1, [x1]
	str x1, [sp, #96]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #104]
	ldr x1, [sp, #96]
	ldr x2, [sp, #104]
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #112]
	ldr x1, [sp, #112]
	mov x2, #0
	cmp x1, x2
	b.eq .L3

.L2:
	ldr x1, [sp, #24]
	str x1, [sp, #120]
	ldr x1, [sp, #120]
	adrp x2, head
	add x2, x2, :lo12:head
	str x1, [x2]
	ldr x1, [sp, #24]
	str x1, [sp, #128]
	ldr x1, [sp, #128]
	adrp x2, tail
	add x2, x2, :lo12:tail
	str x1, [x2]
	b .L4

.L3:
	adrp x1, tail
	add x1, x1, :lo12:tail
	ldr x1, [x1]
	str x1, [sp, #136]
	ldr x1, [sp, #136]
	add x1, x1, #8
	str x1, [sp, #144]
	ldr x1, [sp, #24]
	str x1, [sp, #152]
	ldr x1, [sp, #152]
	ldr x3, [sp, #144]
	str x1, [x3]
	ldr x1, [sp, #24]
	str x1, [sp, #160]
	ldr x1, [sp, #160]
	adrp x2, tail
	add x2, x2, :lo12:tail
	str x1, [x2]
	b .L4

.L4:
	b .L5

.L5:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #168]
	ldr x1, [sp, #168]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #176
	ret

	.size Add, (.-Add)

	.type PrintList, %function
	.global PrintList
	.p2align 2
PrintList:
.L6:
	sub sp, sp, #160
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #160]
	b .L7

.L7:
	ldr x1, [sp, #160]
	str x1, [sp, #32]
	adrp x1, tail
	add x1, x1, :lo12:tail
	ldr x1, [x1]
	str x1, [sp, #40]
	ldr x1, [sp, #32]
	ldr x2, [sp, #40]
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #48]
	ldr x1, [sp, #48]
	mov x2, #0
	cmp x1, x2
	b.eq .L9

.L8:
	ldr x1, [sp, #160]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	add x1, x1, #0
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	ldr x1, [x1]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	str x1, [sp, #80]
	ldr x1, [sp, #80]
	mov x1, x1
	adrp x2, .FSTR1
	add x2, x2, :lo12:.FSTR1
	mov x0, x2
	bl printf
	b .L10

.L9:
	ldr x1, [sp, #160]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	add x1, x1, #0
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	ldr x1, [x1]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	mov x1, x1
	adrp x2, .FSTR2
	add x2, x2, :lo12:.FSTR2
	mov x0, x2
	bl printf
	ldr x1, [sp, #160]
	str x1, [sp, #120]
	ldr x1, [sp, #120]
	add x1, x1, #8
	str x1, [sp, #128]
	ldr x1, [sp, #128]
	ldr x1, [x1]
	str x1, [sp, #136]
	ldr x1, [sp, #136]
	mov x0, x1
	bl PrintList
	str x0, [sp, #144]
	b .L10

.L10:
	b .L11

.L11:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #152]
	ldr x1, [sp, #152]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #160
	ret

	.size PrintList, (.-PrintList)

	.type Del, %function
	.global Del
	.p2align 2
Del:
.L12:
	sub sp, sp, #464
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #440]
	str x1, [sp, #448]
	b .L13

.L13:
	ldr x2, [sp, #440]
	str x2, [sp, #32]
	adrp x2, .nilNode
	add x2, x2, :lo12:.nilNode
	ldr x2, [x2]
	str x2, [sp, #40]
	ldr x2, [sp, #32]
	ldr x3, [sp, #40]
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #48]
	ldr x2, [sp, #48]
	mov x3, #0
	cmp x2, x3
	b.eq .L15

.L14:
	b .L28

.L15:
	b .L16

.L16:
	adrp x2, head
	add x2, x2, :lo12:head
	ldr x2, [x2]
	str x2, [sp, #56]
	ldr x2, [sp, #56]
	add x2, x2, #0
	str x2, [sp, #64]
	ldr x2, [sp, #64]
	ldr x2, [x2]
	str x2, [sp, #72]
	ldr x2, [sp, #448]
	str x2, [sp, #80]
	ldr x2, [sp, #72]
	ldr x3, [sp, #80]
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #88]
	ldr x2, [sp, #88]
	mov x3, #0
	cmp x2, x3
	b.eq .L18

.L17:
	adrp x2, head
	add x2, x2, :lo12:head
	ldr x2, [x2]
	str x2, [sp, #96]
	ldr x2, [sp, #96]
	str x2, [sp, #24]
	adrp x2, head
	add x2, x2, :lo12:head
	ldr x2, [x2]
	str x2, [sp, #104]
	ldr x2, [sp, #104]
	add x2, x2, #8
	str x2, [sp, #112]
	ldr x2, [sp, #112]
	ldr x2, [x2]
	str x2, [sp, #120]
	ldr x2, [sp, #120]
	adrp x3, head
	add x3, x3, :lo12:head
	str x2, [x3]
	ldr x2, [sp, #24]
	str x2, [sp, #128]
	ldr x2, [sp, #128]
	str x2, [sp, #136]
	ldr x2, [sp, #136]
	mov x0, x2
	bl free
	b .L27

.L18:
	b .L19

.L19:
	ldr x2, [sp, #440]
	str x2, [sp, #144]
	ldr x2, [sp, #144]
	add x2, x2, #8
	str x2, [sp, #152]
	ldr x2, [sp, #152]
	ldr x2, [x2]
	str x2, [sp, #160]
	adrp x2, tail
	add x2, x2, :lo12:tail
	ldr x2, [x2]
	str x2, [sp, #168]
	ldr x2, [sp, #160]
	ldr x3, [sp, #168]
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #176]
	ldr x2, [sp, #176]
	mov x3, #0
	cmp x2, x3
	b.eq .L21

.L20:
	adrp x2, tail
	add x2, x2, :lo12:tail
	ldr x2, [x2]
	str x2, [sp, #184]
	ldr x2, [sp, #184]
	str x2, [sp, #24]
	ldr x2, [sp, #440]
	str x2, [sp, #192]
	ldr x2, [sp, #192]
	adrp x3, tail
	add x3, x3, :lo12:tail
	str x2, [x3]
	adrp x2, tail
	add x2, x2, :lo12:tail
	ldr x2, [x2]
	str x2, [sp, #200]
	ldr x2, [sp, #200]
	add x2, x2, #8
	str x2, [sp, #208]
	adrp x2, .nilNode
	add x2, x2, :lo12:.nilNode
	ldr x2, [x2]
	str x2, [sp, #216]
	ldr x2, [sp, #216]
	ldr x4, [sp, #208]
	str x2, [x4]
	ldr x2, [sp, #24]
	str x2, [sp, #224]
	ldr x2, [sp, #224]
	str x2, [sp, #232]
	ldr x2, [sp, #232]
	mov x0, x2
	bl free
	b .L26

.L21:
	b .L22

.L22:
	ldr x2, [sp, #440]
	str x2, [sp, #240]
	ldr x2, [sp, #240]
	add x2, x2, #8
	str x2, [sp, #248]
	ldr x2, [sp, #248]
	ldr x2, [x2]
	str x2, [sp, #256]
	ldr x2, [sp, #256]
	add x2, x2, #0
	str x2, [sp, #264]
	ldr x2, [sp, #264]
	ldr x2, [x2]
	str x2, [sp, #272]
	ldr x2, [sp, #448]
	str x2, [sp, #280]
	ldr x2, [sp, #272]
	ldr x3, [sp, #280]
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #288]
	ldr x2, [sp, #288]
	mov x3, #0
	cmp x2, x3
	b.eq .L24

.L23:
	ldr x2, [sp, #440]
	str x2, [sp, #296]
	ldr x2, [sp, #296]
	add x2, x2, #8
	str x2, [sp, #304]
	ldr x2, [sp, #304]
	ldr x2, [x2]
	str x2, [sp, #312]
	ldr x2, [sp, #312]
	str x2, [sp, #24]
	ldr x2, [sp, #440]
	str x2, [sp, #320]
	ldr x2, [sp, #320]
	add x2, x2, #8
	str x2, [sp, #328]
	ldr x2, [sp, #440]
	str x2, [sp, #336]
	ldr x2, [sp, #336]
	add x2, x2, #8
	str x2, [sp, #344]
	ldr x2, [sp, #344]
	ldr x2, [x2]
	str x2, [sp, #352]
	ldr x2, [sp, #352]
	add x2, x2, #8
	str x2, [sp, #360]
	ldr x2, [sp, #360]
	ldr x2, [x2]
	str x2, [sp, #368]
	ldr x2, [sp, #368]
	ldr x4, [sp, #328]
	str x2, [x4]
	ldr x2, [sp, #24]
	str x2, [sp, #376]
	ldr x2, [sp, #376]
	str x2, [sp, #384]
	ldr x2, [sp, #384]
	mov x0, x2
	bl free
	b .L25

.L24:
	ldr x2, [sp, #440]
	str x2, [sp, #392]
	ldr x2, [sp, #392]
	add x2, x2, #8
	str x2, [sp, #400]
	ldr x2, [sp, #400]
	ldr x2, [x2]
	str x2, [sp, #408]
	ldr x2, [sp, #448]
	str x2, [sp, #416]
	ldr x2, [sp, #408]
	ldr x3, [sp, #416]
	mov x0, x2
	mov x1, x3
	bl Del
	str x0, [sp, #424]
	b .L25

.L25:
	b .L26

.L26:
	b .L27

.L27:
	b .L28

.L28:
	b .L29

.L29:
	mov x2, #0
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #432]
	ldr x2, [sp, #432]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #464
	ret

	.size Del, (.-Del)

	.type main, %function
	.global main
	.p2align 2
main:
.L30:
	sub sp, sp, #256
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #24
	bl scanf
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #32
	bl scanf
	mov x1, #1
	mov x0, x1
	bl Add
	str x0, [sp, #48]
	mov x1, #10
	mov x0, x1
	bl Add
	str x0, [sp, #56]
	mov x1, #3
	mov x0, x1
	bl Add
	str x0, [sp, #64]
	mov x1, #4
	mov x0, x1
	bl Add
	str x0, [sp, #72]
	ldr x1, [sp, #24]
	str x1, [sp, #80]
	ldr x1, [sp, #80]
	mov x0, x1
	bl Add
	str x0, [sp, #88]
	adrp x1, head
	add x1, x1, :lo12:head
	ldr x1, [x1]
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	mov x0, x1
	bl PrintList
	str x0, [sp, #104]
	mov x1, #0
	str x1, [sp, #40]
	b .L31

.L31:
	ldr x1, [sp, #40]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	movz x2, #0x0000, lsl #48
	movk x2, #0x0000, lsl #32
	movk x2, #0x02fa, lsl #16
	movk x2, #0xf080, lsl #0
	cmp x1, x2
	cset x3, lt
	str x3, [sp, #120]
	ldr x1, [sp, #120]
	mov x2, #0
	cmp x1, x2
	b.eq .L33

.L32:
	ldr x1, [sp, #40]
	str x1, [sp, #128]
	ldr x1, [sp, #128]
	mov x0, x1
	bl Add
	str x0, [sp, #136]
	ldr x1, [sp, #40]
	str x1, [sp, #144]
	ldr x1, [sp, #144]
	mov x2, #1
	add x3, x1, x2
	str x3, [sp, #152]
	ldr x1, [sp, #152]
	str x1, [sp, #40]
	b .L31

.L33:
	mov x1, #0
	str x1, [sp, #40]
	b .L34

.L34:
	ldr x1, [sp, #40]
	str x1, [sp, #160]
	ldr x1, [sp, #160]
	movz x2, #0x0000, lsl #48
	movk x2, #0x0000, lsl #32
	movk x2, #0x02fa, lsl #16
	movk x2, #0xf080, lsl #0
	cmp x1, x2
	cset x3, lt
	str x3, [sp, #168]
	ldr x1, [sp, #168]
	mov x2, #0
	cmp x1, x2
	b.eq .L36

.L35:
	adrp x1, head
	add x1, x1, :lo12:head
	ldr x1, [x1]
	str x1, [sp, #176]
	ldr x1, [sp, #40]
	str x1, [sp, #184]
	ldr x1, [sp, #176]
	ldr x2, [sp, #184]
	mov x0, x1
	mov x1, x2
	bl Del
	str x0, [sp, #192]
	ldr x1, [sp, #40]
	str x1, [sp, #200]
	ldr x1, [sp, #200]
	mov x2, #1
	add x3, x1, x2
	str x3, [sp, #208]
	ldr x1, [sp, #208]
	str x1, [sp, #40]
	b .L34

.L36:
	adrp x1, head
	add x1, x1, :lo12:head
	ldr x1, [x1]
	str x1, [sp, #216]
	ldr x1, [sp, #32]
	str x1, [sp, #224]
	ldr x1, [sp, #216]
	ldr x2, [sp, #224]
	mov x0, x1
	mov x1, x2
	bl Del
	str x0, [sp, #232]
	adrp x1, head
	add x1, x1, :lo12:head
	ldr x1, [x1]
	str x1, [sp, #240]
	ldr x1, [sp, #240]
	mov x0, x1
	bl PrintList
	str x0, [sp, #248]
	b .L37

.L37:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #256]
	ldr x1, [sp, #256]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #256
	ret

	.size main, (.-main)


.FSTR1:
	.asciz	"%ld\n"
	.size	.FSTR1, 6

.FSTR2:
	.asciz	"%ld\n"
	.size	.FSTR2, 6

.READ:
	.asciz	"%ld"
	.size	.READ, 4

