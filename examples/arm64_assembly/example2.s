	.arch armv8-a

	.comm .nilNode,8,8
	.comm root,8,8

	.text

	.type compare, %function
	.global compare
	.p2align 2
compare:
.L0:
	sub sp, sp, #128
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #112]
	str x1, [sp, #120]
	b .L1

.L1:
	ldr x2, [sp, #112]
	str x2, [sp, #24]
	ldr x2, [sp, #120]
	str x2, [sp, #32]
	ldr x2, [sp, #24]
	ldr x3, [sp, #32]
	cmp x2, x3
	cset x4, lt
	str x4, [sp, #40]
	ldr x2, [sp, #40]
	mov x3, #0
	cmp x2, x3
	b.eq .L3

.L2:
	mov x2, #1
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #48]
	ldr x2, [sp, #48]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #128
	ret

.L3:
	b .L4

.L4:
	ldr x2, [sp, #112]
	str x2, [sp, #56]
	ldr x2, [sp, #120]
	str x2, [sp, #64]
	ldr x2, [sp, #56]
	ldr x3, [sp, #64]
	cmp x2, x3
	cset x4, gt
	str x4, [sp, #72]
	ldr x2, [sp, #72]
	mov x3, #0
	cmp x2, x3
	b.eq .L6

.L5:
	mov x2, #0
	movz x3, #0x0000, lsl #48
	movk x3, #0x0000, lsl #32
	movk x3, #0x0000, lsl #16
	movk x3, #0x0001, lsl #0
	sub x2, x2, x3
	mov x4, #1
	mul x5, x2, x4
	str x5, [sp, #80]
	ldr x2, [sp, #80]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #88]
	ldr x2, [sp, #88]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #128
	ret

.L6:
	mov x2, #0
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #96]
	ldr x2, [sp, #96]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #128
	ret

.L7:
	b .L8

.L8:
	b .L9

.L9:
	ldr x2, [sp, #16]
	str x2, [sp, #104]
	ldr x2, [sp, #104]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #128
	ret

	.size compare, (.-compare)

	.type addNode, %function
	.global addNode
	.p2align 2
addNode:
.L10:
	sub sp, sp, #512
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #488]
	str x1, [sp, #496]
	b .L11

.L11:
	ldr x2, [sp, #496]
	str x2, [sp, #40]
	adrp x2, .nilNode
	add x2, x2, :lo12:.nilNode
	ldr x2, [x2]
	str x2, [sp, #48]
	ldr x2, [sp, #40]
	ldr x3, [sp, #48]
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #56]
	ldr x2, [sp, #56]
	mov x3, #0
	cmp x2, x3
	b.eq .L13

.L12:
	mov x0, #24
	bl malloc
	str x0, [sp, #64]
	ldr x2, [sp, #64]
	str x2, [sp, #72]
	ldr x2, [sp, #72]
	str x2, [sp, #32]
	ldr x2, [sp, #32]
	str x2, [sp, #80]
	ldr x2, [sp, #80]
	add x2, x2, #0
	str x2, [sp, #88]
	ldr x2, [sp, #488]
	str x2, [sp, #96]
	ldr x2, [sp, #96]
	ldr x4, [sp, #88]
	str x2, [x4]
	ldr x2, [sp, #32]
	str x2, [sp, #104]
	ldr x2, [sp, #104]
	adrp x3, root
	add x3, x3, :lo12:root
	str x2, [x3]
	b .L30

.L13:
	ldr x2, [sp, #496]
	str x2, [sp, #112]
	ldr x2, [sp, #112]
	add x2, x2, #0
	str x2, [sp, #120]
	ldr x2, [sp, #120]
	ldr x2, [x2]
	str x2, [sp, #128]
	ldr x2, [sp, #488]
	str x2, [sp, #136]
	ldr x2, [sp, #128]
	ldr x3, [sp, #136]
	mov x0, x2
	mov x1, x3
	bl compare
	str x0, [sp, #144]
	ldr x2, [sp, #144]
	str x2, [sp, #24]
	b .L14

.L14:
	mov x2, #0
	movz x3, #0x0000, lsl #48
	movk x3, #0x0000, lsl #32
	movk x3, #0x0000, lsl #16
	movk x3, #0x0001, lsl #0
	sub x2, x2, x3
	mov x4, #1
	mul x5, x2, x4
	str x5, [sp, #152]
	ldr x2, [sp, #24]
	str x2, [sp, #160]
	ldr x2, [sp, #160]
	ldr x3, [sp, #152]
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #168]
	ldr x2, [sp, #168]
	mov x3, #0
	cmp x2, x3
	b.eq .L20

.L15:
	b .L16

.L16:
	ldr x2, [sp, #496]
	str x2, [sp, #176]
	ldr x2, [sp, #176]
	add x2, x2, #8
	str x2, [sp, #184]
	ldr x2, [sp, #184]
	ldr x2, [x2]
	str x2, [sp, #192]
	adrp x2, .nilNode
	add x2, x2, :lo12:.nilNode
	ldr x2, [x2]
	str x2, [sp, #200]
	ldr x2, [sp, #192]
	ldr x3, [sp, #200]
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #208]
	ldr x2, [sp, #208]
	mov x3, #0
	cmp x2, x3
	b.eq .L18

.L17:
	mov x0, #24
	bl malloc
	str x0, [sp, #216]
	ldr x2, [sp, #216]
	str x2, [sp, #224]
	ldr x2, [sp, #224]
	str x2, [sp, #32]
	ldr x2, [sp, #32]
	str x2, [sp, #232]
	ldr x2, [sp, #232]
	add x2, x2, #0
	str x2, [sp, #240]
	ldr x2, [sp, #488]
	str x2, [sp, #248]
	ldr x2, [sp, #248]
	ldr x4, [sp, #240]
	str x2, [x4]
	ldr x2, [sp, #496]
	str x2, [sp, #256]
	ldr x2, [sp, #256]
	add x2, x2, #8
	str x2, [sp, #264]
	ldr x2, [sp, #32]
	str x2, [sp, #272]
	ldr x2, [sp, #272]
	ldr x4, [sp, #264]
	str x2, [x4]
	b .L19

.L18:
	ldr x2, [sp, #488]
	str x2, [sp, #280]
	ldr x2, [sp, #496]
	str x2, [sp, #288]
	ldr x2, [sp, #288]
	add x2, x2, #8
	str x2, [sp, #296]
	ldr x2, [sp, #296]
	ldr x2, [x2]
	str x2, [sp, #304]
	ldr x2, [sp, #280]
	ldr x3, [sp, #304]
	mov x0, x2
	mov x1, x3
	bl addNode
	str x0, [sp, #312]
	b .L19

.L19:
	b .L29

.L20:
	b .L21

.L21:
	ldr x2, [sp, #24]
	str x2, [sp, #320]
	ldr x2, [sp, #320]
	mov x3, #1
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #328]
	ldr x2, [sp, #328]
	mov x3, #0
	cmp x2, x3
	b.eq .L27

.L22:
	b .L23

.L23:
	ldr x2, [sp, #496]
	str x2, [sp, #336]
	ldr x2, [sp, #336]
	add x2, x2, #16
	str x2, [sp, #344]
	ldr x2, [sp, #344]
	ldr x2, [x2]
	str x2, [sp, #352]
	adrp x2, .nilNode
	add x2, x2, :lo12:.nilNode
	ldr x2, [x2]
	str x2, [sp, #360]
	ldr x2, [sp, #352]
	ldr x3, [sp, #360]
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #368]
	ldr x2, [sp, #368]
	mov x3, #0
	cmp x2, x3
	b.eq .L25

.L24:
	mov x0, #24
	bl malloc
	str x0, [sp, #376]
	ldr x2, [sp, #376]
	str x2, [sp, #384]
	ldr x2, [sp, #384]
	str x2, [sp, #32]
	ldr x2, [sp, #32]
	str x2, [sp, #392]
	ldr x2, [sp, #392]
	add x2, x2, #0
	str x2, [sp, #400]
	ldr x2, [sp, #488]
	str x2, [sp, #408]
	ldr x2, [sp, #408]
	ldr x4, [sp, #400]
	str x2, [x4]
	ldr x2, [sp, #496]
	str x2, [sp, #416]
	ldr x2, [sp, #416]
	add x2, x2, #16
	str x2, [sp, #424]
	ldr x2, [sp, #32]
	str x2, [sp, #432]
	ldr x2, [sp, #432]
	ldr x4, [sp, #424]
	str x2, [x4]
	b .L26

.L25:
	ldr x2, [sp, #488]
	str x2, [sp, #440]
	ldr x2, [sp, #496]
	str x2, [sp, #448]
	ldr x2, [sp, #448]
	add x2, x2, #16
	str x2, [sp, #456]
	ldr x2, [sp, #456]
	ldr x2, [x2]
	str x2, [sp, #464]
	ldr x2, [sp, #440]
	ldr x3, [sp, #464]
	mov x0, x2
	mov x1, x3
	bl addNode
	str x0, [sp, #472]
	b .L26

.L26:
	b .L28

.L27:
	b .L28

.L28:
	b .L29

.L29:
	b .L30

.L30:
	b .L31

.L31:
	mov x2, #0
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #480]
	ldr x2, [sp, #480]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #512
	ret

	.size addNode, (.-addNode)

	.type printDepthTree, %function
	.global printDepthTree
	.p2align 2
printDepthTree:
.L32:
	sub sp, sp, #240
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #240]
	b .L33

.L33:
	ldr x1, [sp, #240]
	str x1, [sp, #32]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #40]
	ldr x1, [sp, #32]
	ldr x2, [sp, #40]
	cmp x1, x2
	cset x3, ne
	str x3, [sp, #48]
	ldr x1, [sp, #48]
	mov x2, #0
	cmp x1, x2
	b.eq .L43

.L34:
	b .L35

.L35:
	ldr x1, [sp, #240]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	add x1, x1, #8
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	ldr x1, [x1]
	str x1, [sp, #72]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #80]
	ldr x1, [sp, #72]
	ldr x2, [sp, #80]
	cmp x1, x2
	cset x3, ne
	str x3, [sp, #88]
	ldr x1, [sp, #88]
	mov x2, #0
	cmp x1, x2
	b.eq .L37

.L36:
	ldr x1, [sp, #240]
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	add x1, x1, #8
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	ldr x1, [x1]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	mov x0, x1
	bl printDepthTree
	str x0, [sp, #120]
	b .L38

.L37:
	b .L38

.L38:
	ldr x1, [sp, #240]
	str x1, [sp, #128]
	ldr x1, [sp, #128]
	add x1, x1, #0
	str x1, [sp, #136]
	ldr x1, [sp, #136]
	ldr x1, [x1]
	str x1, [sp, #144]
	ldr x1, [sp, #144]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	str x1, [sp, #152]
	ldr x1, [sp, #152]
	mov x1, x1
	adrp x2, .FSTR1
	add x2, x2, :lo12:.FSTR1
	mov x0, x2
	bl printf
	b .L39

.L39:
	ldr x1, [sp, #240]
	str x1, [sp, #160]
	ldr x1, [sp, #160]
	add x1, x1, #16
	str x1, [sp, #168]
	ldr x1, [sp, #168]
	ldr x1, [x1]
	str x1, [sp, #176]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #184]
	ldr x1, [sp, #176]
	ldr x2, [sp, #184]
	cmp x1, x2
	cset x3, ne
	str x3, [sp, #192]
	ldr x1, [sp, #192]
	mov x2, #0
	cmp x1, x2
	b.eq .L41

.L40:
	ldr x1, [sp, #240]
	str x1, [sp, #200]
	ldr x1, [sp, #200]
	add x1, x1, #16
	str x1, [sp, #208]
	ldr x1, [sp, #208]
	ldr x1, [x1]
	str x1, [sp, #216]
	ldr x1, [sp, #216]
	mov x0, x1
	bl printDepthTree
	str x0, [sp, #224]
	b .L42

.L41:
	b .L42

.L42:
	b .L44

.L43:
	b .L44

.L44:
	b .L45

.L45:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #232]
	ldr x1, [sp, #232]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #240
	ret

	.size printDepthTree, (.-printDepthTree)

	.type deleteLeavesTree, %function
	.global deleteLeavesTree
	.p2align 2
deleteLeavesTree:
.L46:
	sub sp, sp, #224
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #216]
	b .L47

.L47:
	ldr x1, [sp, #216]
	str x1, [sp, #24]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #32]
	ldr x1, [sp, #24]
	ldr x2, [sp, #32]
	cmp x1, x2
	cset x3, ne
	str x3, [sp, #40]
	ldr x1, [sp, #40]
	mov x2, #0
	cmp x1, x2
	b.eq .L57

.L48:
	b .L49

.L49:
	ldr x1, [sp, #216]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	add x1, x1, #8
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	ldr x1, [x1]
	str x1, [sp, #64]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #72]
	ldr x1, [sp, #64]
	ldr x2, [sp, #72]
	cmp x1, x2
	cset x3, ne
	str x3, [sp, #80]
	ldr x1, [sp, #80]
	mov x2, #0
	cmp x1, x2
	b.eq .L51

.L50:
	ldr x1, [sp, #216]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	add x1, x1, #8
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	ldr x1, [x1]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	mov x0, x1
	bl deleteLeavesTree
	str x0, [sp, #112]
	b .L52

.L51:
	b .L52

.L52:
	b .L53

.L53:
	ldr x1, [sp, #216]
	str x1, [sp, #120]
	ldr x1, [sp, #120]
	add x1, x1, #16
	str x1, [sp, #128]
	ldr x1, [sp, #128]
	ldr x1, [x1]
	str x1, [sp, #136]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #144]
	ldr x1, [sp, #136]
	ldr x2, [sp, #144]
	cmp x1, x2
	cset x3, ne
	str x3, [sp, #152]
	ldr x1, [sp, #152]
	mov x2, #0
	cmp x1, x2
	b.eq .L55

.L54:
	ldr x1, [sp, #216]
	str x1, [sp, #160]
	ldr x1, [sp, #160]
	add x1, x1, #16
	str x1, [sp, #168]
	ldr x1, [sp, #168]
	ldr x1, [x1]
	str x1, [sp, #176]
	ldr x1, [sp, #176]
	mov x0, x1
	bl deleteLeavesTree
	str x0, [sp, #184]
	b .L56

.L55:
	b .L56

.L56:
	ldr x1, [sp, #216]
	str x1, [sp, #192]
	ldr x1, [sp, #192]
	str x1, [sp, #200]
	ldr x1, [sp, #200]
	mov x0, x1
	bl free
	b .L58

.L57:
	b .L58

.L58:
	b .L59

.L59:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #208]
	ldr x1, [sp, #208]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #224
	ret

	.size deleteLeavesTree, (.-deleteLeavesTree)

	.type main, %function
	.global main
	.p2align 2
main:
.L60:
	sub sp, sp, #112
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #40]
	ldr x1, [sp, #40]
	adrp x2, root
	add x2, x2, :lo12:root
	str x1, [x2]
	mov x1, #0
	str x1, [sp, #24]
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #24
	bl scanf
	b .L61

.L61:
	ldr x1, [sp, #24]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	mov x2, #0
	cmp x1, x2
	cset x3, ne
	str x3, [sp, #56]
	ldr x1, [sp, #56]
	mov x2, #0
	cmp x1, x2
	b.eq .L63

.L62:
	ldr x1, [sp, #24]
	str x1, [sp, #64]
	adrp x1, root
	add x1, x1, :lo12:root
	ldr x1, [x1]
	str x1, [sp, #72]
	ldr x1, [sp, #64]
	ldr x2, [sp, #72]
	mov x0, x1
	mov x1, x2
	bl addNode
	str x0, [sp, #80]
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #24
	bl scanf
	b .L61

.L63:
	adrp x1, root
	add x1, x1, :lo12:root
	ldr x1, [x1]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	mov x0, x1
	bl printDepthTree
	str x0, [sp, #96]
	adrp x1, root
	add x1, x1, :lo12:root
	ldr x1, [x1]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	mov x0, x1
	bl deleteLeavesTree
	str x0, [sp, #112]
	b .L64

.L64:
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


.FSTR1:
	.asciz	"%ld\n"
	.size	.FSTR1, 6

.READ:
	.asciz	"%ld"
	.size	.READ, 4

