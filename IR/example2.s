	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 13, 0
	.globl	_compare                        ; -- Begin function compare
	.p2align	2
_compare:                               ; @compare
	.cfi_startproc
; %bb.0:                                ; %L0
	sub	sp, sp, #32
	.cfi_def_cfa_offset 32
	cmp	x0, x1
	stp	x1, x0, [sp, #8]
	b.ge	LBB0_2
; %bb.1:                                ; %L2
	mov	w0, #1
	b	LBB0_4
LBB0_2:                                 ; %L4
	ldp	x9, x8, [sp, #8]
	cmp	x8, x9
	b.le	LBB0_5
; %bb.3:                                ; %L5
	mov	x0, #-1
LBB0_4:                                 ; %L2
	str	x0, [sp, #24]
	add	sp, sp, #32
	ret
LBB0_5:                                 ; %L6
	mov	x0, xzr
	str	xzr, [sp, #24]
	add	sp, sp, #32
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_addNode                        ; -- Begin function addNode
	.p2align	2
_addNode:                               ; @addNode
	.cfi_startproc
; %bb.0:                                ; %L10
	sub	sp, sp, #80
	.cfi_def_cfa_offset 80
	stp	x20, x19, [sp, #48]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #64]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
Lloh0:
	adrp	x19, _.nilNode@GOTPAGE
Lloh1:
	ldr	x19, [x19, _.nilNode@GOTPAGEOFF]
	stp	x1, x0, [sp, #8]
	ldr	x8, [x19]
	cmp	x1, x8
	b.eq	LBB1_2
; %bb.1:                                ; %L13
	ldp	x8, x1, [sp, #8]
	ldr	x0, [x8]
	bl	_compare
	str	x0, [sp, #32]
	b	LBB1_3
LBB1_2:                                 ; %L12
	mov	w0, #24
	bl	_malloc
Lloh2:
	adrp	x9, _root@GOTPAGE
	ldr	x8, [sp, #16]
Lloh3:
	ldr	x9, [x9, _root@GOTPAGEOFF]
	str	x0, [sp, #24]
	str	x8, [x0]
Lloh4:
	str	x0, [x9]
LBB1_3:                                 ; %L14
	ldr	x8, [sp, #8]
	ldr	x9, [x19]
	ldr	x8, [x8, #8]
	cmp	x8, x9
	b.eq	LBB1_5
; %bb.4:                                ; %L18
	ldp	x8, x0, [sp, #8]
	ldr	x1, [x8, #8]
	bl	_addNode
LBB1_5:                                 ; %L17
                                        ; =>This Inner Loop Header: Depth=1
	mov	w0, #24
	bl	_malloc
	ldp	x9, x8, [sp, #8]
	str	x0, [sp, #24]
	str	x8, [x0]
	str	x0, [x9, #8]
	b	LBB1_5
	.loh AdrpLdrGot	Lloh0, Lloh1
	.loh AdrpLdrGotStr	Lloh2, Lloh3, Lloh4
	.cfi_endproc
                                        ; -- End function
	.globl	_printDepthTree                 ; -- Begin function printDepthTree
	.p2align	2
_printDepthTree:                        ; @printDepthTree
	.cfi_startproc
; %bb.0:                                ; %L32
	sub	sp, sp, #64
	.cfi_def_cfa_offset 64
	stp	x20, x19, [sp, #32]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #48]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
Lloh5:
	adrp	x19, _.nilNode@GOTPAGE
Lloh6:
	ldr	x19, [x19, _.nilNode@GOTPAGEOFF]
	str	x0, [sp, #8]
	ldr	x8, [x0, #8]
	ldr	x9, [x19]
	cmp	x8, x9
	b.eq	LBB2_2
LBB2_1:                                 ; %L36
	ldr	x8, [sp, #8]
	ldr	x0, [x8, #8]
	bl	_printDepthTree
LBB2_2:                                 ; %L38
	ldr	x8, [sp, #8]
Lloh7:
	adrp	x0, l_.fstr1@PAGE
Lloh8:
	add	x0, x0, l_.fstr1@PAGEOFF
	ldr	x8, [x8]
	str	x8, [sp, #16]
	str	x8, [sp]
	bl	_printf
	ldr	x8, [sp, #8]
	ldr	x9, [x19]
	ldr	x8, [x8, #16]
	cmp	x8, x9
	b.eq	LBB2_1
; %bb.3:                                ; %L40
	ldr	x8, [sp, #8]
	ldr	x0, [x8, #16]
	bl	_printDepthTree
	b	LBB2_1
	.loh AdrpLdrGot	Lloh5, Lloh6
	.loh AdrpAdd	Lloh7, Lloh8
	.cfi_endproc
                                        ; -- End function
	.globl	_deleteLeavesTree               ; -- Begin function deleteLeavesTree
	.p2align	2
_deleteLeavesTree:                      ; @deleteLeavesTree
	.cfi_startproc
; %bb.0:                                ; %L46
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x20, x19, [sp, #16]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
Lloh9:
	adrp	x19, _.nilNode@GOTPAGE
Lloh10:
	ldr	x19, [x19, _.nilNode@GOTPAGEOFF]
	str	x0, [sp]
	ldr	x8, [x0, #8]
	ldr	x9, [x19]
	cmp	x8, x9
	b.eq	LBB3_2
LBB3_1:                                 ; %L50
	ldr	x8, [sp]
	ldr	x0, [x8, #8]
	bl	_deleteLeavesTree
LBB3_2:                                 ; %L53
	ldr	x8, [sp]
	ldr	x9, [x19]
	ldr	x8, [x8, #16]
	cmp	x8, x9
	b.eq	LBB3_4
; %bb.3:                                ; %L54
	ldr	x8, [sp]
	ldr	x0, [x8, #16]
	bl	_deleteLeavesTree
LBB3_4:                                 ; %L56
	ldr	x0, [sp]
	bl	_free
	b	LBB3_1
	.loh AdrpLdrGot	Lloh9, Lloh10
	.cfi_endproc
                                        ; -- End function
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:                                ; %L60
	sub	sp, sp, #80
	.cfi_def_cfa_offset 80
	stp	x22, x21, [sp, #32]             ; 16-byte Folded Spill
	stp	x20, x19, [sp, #48]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #64]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	.cfi_offset w21, -40
	.cfi_offset w22, -48
Lloh11:
	adrp	x8, _.nilNode@GOTPAGE
Lloh12:
	adrp	x20, _root@GOTPAGE
	add	x21, sp, #16
Lloh13:
	adrp	x19, l_.read@PAGE
Lloh14:
	add	x19, x19, l_.read@PAGEOFF
Lloh15:
	ldr	x8, [x8, _.nilNode@GOTPAGEOFF]
Lloh16:
	ldr	x20, [x20, _root@GOTPAGEOFF]
	str	xzr, [sp, #16]
Lloh17:
	ldr	x8, [x8]
	str	x8, [x20]
LBB4_1:                                 ; %L61
                                        ; =>This Inner Loop Header: Depth=1
	mov	x0, x19
	str	x21, [sp]
	bl	_scanf
	ldr	x8, [sp, #16]
	cbz	x8, LBB4_3
; %bb.2:                                ; %L62
                                        ;   in Loop: Header=BB4_1 Depth=1
	ldr	x0, [sp, #16]
	ldr	x1, [x20]
	bl	_addNode
	b	LBB4_1
LBB4_3:                                 ; %L63
	ldr	x0, [x20]
	bl	_printDepthTree
	ldr	x0, [x20]
	bl	_deleteLeavesTree
	ldp	x29, x30, [sp, #64]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #24]
	ldp	x20, x19, [sp, #48]             ; 16-byte Folded Reload
	ldp	x22, x21, [sp, #32]             ; 16-byte Folded Reload
	add	sp, sp, #80
	ret
	.loh AdrpAdd	Lloh13, Lloh14
	.loh AdrpLdrGot	Lloh12, Lloh16
	.loh AdrpLdrGotLdr	Lloh11, Lloh15, Lloh17
	.cfi_endproc
                                        ; -- End function
	.comm	_.nilNode,8,3                   ; @.nilNode
	.comm	_root,8,3                       ; @root
	.comm	_x,8,3                          ; @x
	.comm	_y,8,3                          ; @y
	.section	__TEXT,__cstring,cstring_literals
l_.fstr1:                               ; @.fstr1
	.asciz	"%ld"

l_.read:                                ; @.read
	.asciz	"%ld"

.subsections_via_symbols
