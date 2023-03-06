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

.L1:
	ldr x1, [sp, #112]
	str x1, [sp, #24]
	ldr x1, [sp, #120]
	str x1, [sp, #32]

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

.L4:
	ldr x1, [sp, #112]
	str x1, [sp, #56]
	ldr x1, [sp, #120]
	str x1, [sp, #64]

.L5:
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

.L8:

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

.L11:
	ldr x1, [sp, #496]
	str x1, [sp, #40]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #48]

.L12:
	mov x0, #24
	bl malloc
	str x0, [sp, #64]
	ldr x1, [sp, #72]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	str x1, [sp, #80]
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

.L13:
	ldr x1, [sp, #496]
	str x1, [sp, #112]
	ldr x1, [sp, #120]
	str x1, [sp, #128]
	ldr x1, [sp, #488]
	str x1, [sp, #136]
	ldr x1, [sp#, 128]
	ldr x2, [sp#, 136]
	mov x0, x1
	mov x1, x2
	bl compare
	str x0, [sp#, 144]
	ldr x1, [sp, #144]
	str x1, [sp, #24]

.L14:
	ldr x1, [sp, #24]
	str x1, [sp, #160]

.L15:

.L16:
	ldr x1, [sp, #496]
	str x1, [sp, #176]
	ldr x1, [sp, #184]
	str x1, [sp, #192]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #200]

.L17:
	mov x0, #24
	bl malloc
	str x0, [sp, #216]
	ldr x1, [sp, #224]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	str x1, [sp, #232]
	ldr x1, [sp, #488]
	str x1, [sp, #248]
	ldr x1, [sp, #248]
	str x1, [sp, #240]
	ldr x1, [sp, #496]
	str x1, [sp, #256]
	ldr x1, [sp, #32]
	str x1, [sp, #272]
	ldr x1, [sp, #272]
	str x1, [sp, #264]

.L18:
	ldr x1, [sp, #488]
	str x1, [sp, #280]
	ldr x1, [sp, #496]
	str x1, [sp, #288]
	ldr x1, [sp, #296]
	str x1, [sp, #304]
	ldr x1, [sp#, 280]
	ldr x2, [sp#, 304]
	mov x0, x1
	mov x1, x2
	bl addNode
	str x0, [sp#, 312]

.L19:

.L20:

.L21:
	ldr x1, [sp, #24]
	str x1, [sp, #320]

.L22:

.L23:
	ldr x1, [sp, #496]
	str x1, [sp, #336]
	ldr x1, [sp, #344]
	str x1, [sp, #352]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #360]

.L24:
	mov x0, #24
	bl malloc
	str x0, [sp, #376]
	ldr x1, [sp, #384]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	str x1, [sp, #392]
	ldr x1, [sp, #488]
	str x1, [sp, #408]
	ldr x1, [sp, #408]
	str x1, [sp, #400]
	ldr x1, [sp, #496]
	str x1, [sp, #416]
	ldr x1, [sp, #32]
	str x1, [sp, #432]
	ldr x1, [sp, #432]
	str x1, [sp, #424]

.L25:
	ldr x1, [sp, #488]
	str x1, [sp, #440]
	ldr x1, [sp, #496]
	str x1, [sp, #448]
	ldr x1, [sp, #456]
	str x1, [sp, #464]
	ldr x1, [sp#, 440]
	ldr x2, [sp#, 464]
	mov x0, x1
	mov x1, x2
	bl addNode
	str x0, [sp#, 472]

.L26:

.L27:

.L28:

.L29:

.L30:

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

.L33:
	ldr x1, [sp, #240]
	str x1, [sp, #32]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #40]

.L34:

.L35:
	ldr x1, [sp, #240]
	str x1, [sp, #56]
	ldr x1, [sp, #64]
	str x1, [sp, #72]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #80]

.L36:
	ldr x1, [sp, #240]
	str x1, [sp, #96]
	ldr x1, [sp, #104]
	str x1, [sp, #112]
	ldr x1, [sp#, 112]
	mov x0, x1
	bl printDepthTree
	str x0, [sp#, 120]

.L37:

.L38:
	ldr x1, [sp, #240]
	str x1, [sp, #128]
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

.L39:
	ldr x1, [sp, #240]
	str x1, [sp, #160]
	ldr x1, [sp, #168]
	str x1, [sp, #176]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #184]

.L40:
	ldr x1, [sp, #240]
	str x1, [sp, #200]
	ldr x1, [sp, #208]
	str x1, [sp, #216]
	ldr x1, [sp#, 216]
	mov x0, x1
	bl printDepthTree
	str x0, [sp#, 224]

.L41:

.L42:

.L43:

.L44:

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

.L47:
	ldr x1, [sp, #216]
	str x1, [sp, #24]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #32]

.L48:

.L49:
	ldr x1, [sp, #216]
	str x1, [sp, #48]
	ldr x1, [sp, #56]
	str x1, [sp, #64]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #72]

.L50:
	ldr x1, [sp, #216]
	str x1, [sp, #88]
	ldr x1, [sp, #96]
	str x1, [sp, #104]
	ldr x1, [sp#, 104]
	mov x0, x1
	bl deleteLeavesTree
	str x0, [sp#, 112]

.L51:

.L52:

.L53:
	ldr x1, [sp, #216]
	str x1, [sp, #120]
	ldr x1, [sp, #128]
	str x1, [sp, #136]
	adrp x1, .nilNode
	add x1, x1, :lo12:.nilNode
	ldr x1, [x1]
	str x1, [sp, #144]

.L54:
	ldr x1, [sp, #216]
	str x1, [sp, #160]
	ldr x1, [sp, #168]
	str x1, [sp, #176]
	ldr x1, [sp#, 176]
	mov x0, x1
	bl deleteLeavesTree
	str x0, [sp#, 184]

.L55:

.L56:
	ldr x1, [sp, #216]
	str x1, [sp, #192]
	ldr x1, [sp, #200]
	mov x0, x1
	bl free

.L57:

.L58:

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

.L61:
	ldr x1, [sp, #24]
	str x1, [sp, #48]

.L62:
	ldr x1, [sp, #24]
	str x1, [sp, #64]
	adrp x1, root
	add x1, x1, :lo12:root
	ldr x1, [x1]
	str x1, [sp, #72]
	ldr x1, [sp#, 64]
	ldr x2, [sp#, 72]
	mov x0, x1
	mov x1, x2
	bl addNode
	str x0, [sp#, 80]
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #24

.L63:
	adrp x1, root
	add x1, x1, :lo12:root
	ldr x1, [x1]
	str x1, [sp, #88]
	ldr x1, [sp#, 88]
	mov x0, x1
	bl printDepthTree
	str x0, [sp#, 96]
	adrp x1, root
	add x1, x1, :lo12:root
	ldr x1, [x1]
	str x1, [sp, #104]
	ldr x1, [sp#, 104]
	mov x0, x1
	bl deleteLeavesTree
	str x0, [sp#, 112]

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

