	.arch armv8-a

	.comm root,8,8

	.text

	.type compare, %function
	.global compare
	.p2align 2
compare:
L0:
	sub sp, sp, 128
	sub sp, sp, 16
	stp x29, x30, [sp]
	mov fp, sp

L1:

L2:
	Ldp x29, x30, [sp]
	add sp, sp, 16
	add sp, sp, 128
	ret

L3:

L4:

L5:
	Ldp x29, x30, [sp]
	add sp, sp, 16
	add sp, sp, 128
	ret

L6:
	Ldp x29, x30, [sp]
	add sp, sp, 16
	add sp, sp, 128
	ret

L7:

L8:

L9:
	Ldp x29, x30, [sp]
	add sp, sp, 16
	add sp, sp, 128
	ret

	.size compare, (.-compare)

	.type addNode, %function
	.global addNode
	.p2align 2
addNode:
L10:
	sub sp, sp, 512
	sub sp, sp, 16
	stp x29, x30, [sp]
	mov fp, sp

L11:

L12:
	mov x1, x0
	mov x0, #24
	bl malloc
	str x0, [sp, #64]
	mov x0, x1

L13:

L14:

L15:

L16:

L17:
	mov x1, x0
	mov x0, #24
	bl malloc
	str x0, [sp, #216]
	mov x0, x1

L18:

L19:

L20:

L21:

L22:

L23:

L24:
	mov x1, x0
	mov x0, #24
	bl malloc
	str x0, [sp, #376]
	mov x0, x1

L25:

L26:

L27:

L28:

L29:

L30:

L31:
	Ldp x29, x30, [sp]
	add sp, sp, 16
	add sp, sp, 512
	ret

	.size addNode, (.-addNode)

	.type printDepthTree, %function
	.global printDepthTree
	.p2align 2
printDepthTree:
L32:
	sub sp, sp, 240
	sub sp, sp, 16
	stp x29, x30, [sp]
	mov fp, sp

L33:

L34:

L35:

L36:

L37:

L38:

L39:

L40:

L41:

L42:

L43:

L44:

L45:
	Ldp x29, x30, [sp]
	add sp, sp, 16
	add sp, sp, 240
	ret

	.size printDepthTree, (.-printDepthTree)

	.type deleteLeavesTree, %function
	.global deleteLeavesTree
	.p2align 2
deleteLeavesTree:
L46:
	sub sp, sp, 224
	sub sp, sp, 16
	stp x29, x30, [sp]
	mov fp, sp

L47:

L48:

L49:

L50:

L51:

L52:

L53:

L54:

L55:

L56:

L57:

L58:

L59:
	Ldp x29, x30, [sp]
	add sp, sp, 16
	add sp, sp, 224
	ret

	.size deleteLeavesTree, (.-deleteLeavesTree)

	.type main, %function
	.global main
	.p2align 2
main:
L60:
	sub sp, sp, 112
	sub sp, sp, 16
	stp x29, x30, [sp]
	mov fp, sp

L61:

L62:

L63:

L64:
	Ldp x29, x30, [sp]
	add sp, sp, 16
	add sp, sp, 112
	ret

	.size main, (.-main)


