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
	bl .L1

.L1:
	ldr x1, [sp, #112]
	str x1, [sp, #24]
	ldr x1, [sp, #120]
	str x1, [sp, #32]
	ldr x1, [sp, #24]
	ldr x2, [sp, #32]
	cmp x1, x2
	cset x3, lt
	str x3, [sp, #40]

.L2:
	mov x1, #1
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
	bl .L4

.L4:
	ldr x1, [sp, #112]
	str x1, [sp, #56]
	ldr x1, [sp, #120]
	str x1, [sp, #64]
	ldr x1, [sp, #56]
	ldr x2, [sp, #64]
	cmp x1, x2
	cset x3, gt
	str x3, [sp, #72]

.L5:
	mov x1, #0
	sub x1, x1, #1
	mov x2, #1
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
	add sp, sp, #128
	ret

.L6:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #128
	ret

.L7:
	bl .L8

.L8:
	bl .L9

.L9:
	ldr x1, [sp, #16]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	mov x0, x1
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
	bl .L11

.L11:
	ldr x1, [sp, #496]
	str x1, [sp, #40]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #48]
	ldr x1, [sp, #40]
	ldr x2, [sp, #48]
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #56]

.L12:
	mov x0, #24
	bl malloc
	str x0, [sp, #64]
	ldr x1, [sp, #64]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	str x1, [sp, #80]
	ldr x1, [sp, #80]
	add x1, x1, #0
	str x1, [sp, #88]
	ldr x1, [sp, #488]
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	str x1, [sp, #88]
	ldr x1, [sp, #32]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	adrp x2, root
	add x2, x2, :lo12:root
	str x1, [x2]
	bl .L30

.L13:
	ldr x1, [sp, #496]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	add x1, x1, #0
	str x1, [sp, #120]
	ldr x1, [sp, #120]
	str x1, [sp, #128]
	ldr x1, [sp, #488]
	str x1, [sp, #136]
	ldr x1, [sp, #128]
	ldr x2, [sp, #136]
	mov x0, x1
	mov x1, x2
	bl compare
	str x0, [sp, #144]
	ldr x1, [sp, #144]
	str x1, [sp, #24]
	bl .L14

.L14:
	mov x1, #0
	sub x1, x1, #1
	mov x2, #1
	mul x3, x1, x2
	str x3, [sp, #152]
	ldr x1, [sp, #24]
	str x1, [sp, #160]
	ldr x1, [sp, #160]
	ldr x2, [sp, #152]
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #168]

.L15:
	bl .L16

.L16:
	ldr x1, [sp, #496]
	str x1, [sp, #176]
	ldr x1, [sp, #176]
	add x1, x1, #8
	str x1, [sp, #184]
	ldr x1, [sp, #184]
	str x1, [sp, #192]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #200]
	ldr x1, [sp, #192]
	ldr x2, [sp, #200]
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #208]

.L17:
	mov x0, #24
	bl malloc
	str x0, [sp, #216]
	ldr x1, [sp, #216]
	str x1, [sp, #224]
	ldr x1, [sp, #224]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	str x1, [sp, #232]
	ldr x1, [sp, #232]
	add x1, x1, #0
	str x1, [sp, #240]
	ldr x1, [sp, #488]
	str x1, [sp, #248]
	ldr x1, [sp, #248]
	str x1, [sp, #240]
	ldr x1, [sp, #496]
	str x1, [sp, #256]
	ldr x1, [sp, #256]
	add x1, x1, #8
	str x1, [sp, #264]
	ldr x1, [sp, #32]
	str x1, [sp, #272]
	ldr x1, [sp, #272]
	str x1, [sp, #264]
	bl .L19

.L18:
	ldr x1, [sp, #488]
	str x1, [sp, #280]
	ldr x1, [sp, #496]
	str x1, [sp, #288]
	ldr x1, [sp, #288]
	add x1, x1, #8
	str x1, [sp, #296]
	ldr x1, [sp, #296]
	str x1, [sp, #304]
	ldr x1, [sp, #280]
	ldr x2, [sp, #304]
	mov x0, x1
	mov x1, x2
	bl addNode
	str x0, [sp, #312]
	bl .L19

.L19:
	bl .L29

.L20:
	bl .L21

.L21:
	ldr x1, [sp, #24]
	str x1, [sp, #320]
	ldr x1, [sp, #320]
	mov x2, #1
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #328]

.L22:
	bl .L23

.L23:
	ldr x1, [sp, #496]
	str x1, [sp, #336]
	ldr x1, [sp, #336]
	add x1, x1, #16
	str x1, [sp, #344]
	ldr x1, [sp, #344]
	str x1, [sp, #352]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #360]
	ldr x1, [sp, #352]
	ldr x2, [sp, #360]
	cmp x1, x2
	cset x3, eq
	str x3, [sp, #368]

.L24:
	mov x0, #24
	bl malloc
	str x0, [sp, #376]
	ldr x1, [sp, #376]
	str x1, [sp, #384]
	ldr x1, [sp, #384]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	str x1, [sp, #392]
	ldr x1, [sp, #392]
	add x1, x1, #0
	str x1, [sp, #400]
	ldr x1, [sp, #488]
	str x1, [sp, #408]
	ldr x1, [sp, #408]
	str x1, [sp, #400]
	ldr x1, [sp, #496]
	str x1, [sp, #416]
	ldr x1, [sp, #416]
	add x1, x1, #16
	str x1, [sp, #424]
	ldr x1, [sp, #32]
	str x1, [sp, #432]
	ldr x1, [sp, #432]
	str x1, [sp, #424]
	bl .L26

.L25:
	ldr x1, [sp, #488]
	str x1, [sp, #440]
	ldr x1, [sp, #496]
	str x1, [sp, #448]
	ldr x1, [sp, #448]
	add x1, x1, #16
	str x1, [sp, #456]
	ldr x1, [sp, #456]
	str x1, [sp, #464]
	ldr x1, [sp, #440]
	ldr x2, [sp, #464]
	mov x0, x1
	mov x1, x2
	bl addNode
	str x0, [sp, #472]
	bl .L26

.L26:
	bl .L28

.L27:
	bl .L28

.L28:
	bl .L29

.L29:
	bl .L30

.L30:
	bl .L31

.L31:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #480]
	ldr x1, [sp, #480]
	mov x0, x1
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
	bl .L33

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

.L34:
	bl .L35

.L35:
	ldr x1, [sp, #240]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	add x1, x1, #8
	str x1, [sp, #64]
	ldr x1, [sp, #64]
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

.L36:
	ldr x1, [sp, #240]
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	add x1, x1, #8
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	mov x0, x1
	bl printDepthTree
	str x0, [sp, #120]
	bl .L38

.L37:
	bl .L38

.L38:
	ldr x1, [sp, #240]
	str x1, [sp, #128]
	ldr x1, [sp, #128]
	add x1, x1, #0
	str x1, [sp, #136]
	ldr x1, [sp, #136]
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
	bl .L39

.L39:
	ldr x1, [sp, #240]
	str x1, [sp, #160]
	ldr x1, [sp, #160]
	add x1, x1, #16
	str x1, [sp, #168]
	ldr x1, [sp, #168]
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

.L40:
	ldr x1, [sp, #240]
	str x1, [sp, #200]
	ldr x1, [sp, #200]
	add x1, x1, #16
	str x1, [sp, #208]
	ldr x1, [sp, #208]
	str x1, [sp, #216]
	ldr x1, [sp, #216]
	mov x0, x1
	bl printDepthTree
	str x0, [sp, #224]
	bl .L42

.L41:
	bl .L42

.L42:
	bl .L44

.L43:
	bl .L44

.L44:
	bl .L45

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
	bl .L47

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

.L48:
	bl .L49

.L49:
	ldr x1, [sp, #216]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	add x1, x1, #8
	str x1, [sp, #56]
	ldr x1, [sp, #56]
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

.L50:
	ldr x1, [sp, #216]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	add x1, x1, #8
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	mov x0, x1
	bl deleteLeavesTree
	str x0, [sp, #112]
	bl .L52

.L51:
	bl .L52

.L52:
	bl .L53

.L53:
	ldr x1, [sp, #216]
	str x1, [sp, #120]
	ldr x1, [sp, #120]
	add x1, x1, #16
	str x1, [sp, #128]
	ldr x1, [sp, #128]
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

.L54:
	ldr x1, [sp, #216]
	str x1, [sp, #160]
	ldr x1, [sp, #160]
	add x1, x1, #16
	str x1, [sp, #168]
	ldr x1, [sp, #168]
	str x1, [sp, #176]
	ldr x1, [sp, #176]
	mov x0, x1
	bl deleteLeavesTree
	str x0, [sp, #184]
	bl .L56

.L55:
	bl .L56

.L56:
	ldr x1, [sp, #216]
	str x1, [sp, #192]
	ldr x1, [sp, #192]
	str x1, [sp, #200]
	ldr x1, [sp, #200]
	mov x0, x1
	bl free
	bl .L58

.L57:
	bl .L58

.L58:
	bl .L59

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
	bl .L61

.L61:
	ldr x1, [sp, #24]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	mov x2, #0
	cmp x1, x2
	cset x3, ne
	str x3, [sp, #56]

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
	bl .L61

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
	bl .L64

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
	.asciz	"%ld"
	.size	.FSTR1, 4

.READ:
	.asciz	"%ld"
	.size	.READ, 4

