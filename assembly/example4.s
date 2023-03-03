	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 13, 0
	.globl	_Add                            ; -- Begin function Add
	.p2align	2
_Add:                                   ; @Add
	.cfi_startproc
; %bb.0:                                ; %L0
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	str	x0, [sp, #8]
	mov	w0, #16
	bl	_malloc
Lloh0:
	adrp	x9, _.nilNode@GOTPAGE
	ldr	x8, [sp, #8]
Lloh1:
	adrp	x10, _head@GOTPAGE
Lloh2:
	ldr	x9, [x9, _.nilNode@GOTPAGEOFF]
	str	x8, [x0]
	ldr	x8, [x9]
Lloh3:
	ldr	x10, [x10, _head@GOTPAGEOFF]
	str	x8, [x0, #8]
Lloh4:
	adrp	x8, _tail@GOTPAGE
	ldr	x11, [x10]
	ldr	x9, [x9]
Lloh5:
	ldr	x8, [x8, _tail@GOTPAGEOFF]
	str	x0, [sp, #16]
	cmp	x11, x9
	b.eq	LBB0_2
; %bb.1:                                ; %L3
	ldr	x10, [x8]
	ldr	x9, [sp, #16]
	str	x9, [x10, #8]
	b	LBB0_3
LBB0_2:                                 ; %L2
	ldr	x9, [sp, #16]
	str	x9, [x10]
LBB0_3:                                 ; %L4
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	x9, [x8]
	str	xzr, [sp, #24]
	add	sp, sp, #48
	ret
	.loh AdrpLdrGot	Lloh4, Lloh5
	.loh AdrpLdrGot	Lloh1, Lloh3
	.loh AdrpLdrGot	Lloh0, Lloh2
	.cfi_endproc
                                        ; -- End function
	.globl	_PrintList                      ; -- Begin function PrintList
	.p2align	2
_PrintList:                             ; @PrintList
	.cfi_startproc
; %bb.0:                                ; %L6
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
Lloh6:
	adrp	x8, _tail@GOTPAGE
Lloh7:
	ldr	x8, [x8, _tail@GOTPAGEOFF]
Lloh8:
	ldr	x9, [x8]
	ldr	x8, [x0]
	cmp	x0, x9
	stp	x0, x8, [sp, #8]
	b.eq	LBB1_2
; %bb.1:                                ; %L9
Lloh9:
	adrp	x0, l_.fstr2@PAGE
Lloh10:
	add	x0, x0, l_.fstr2@PAGEOFF
	str	x8, [sp]
	bl	_printf
	ldr	x8, [sp, #8]
	ldr	x0, [x8, #8]
	bl	_PrintList
	b	LBB1_3
LBB1_2:                                 ; %L8
Lloh11:
	adrp	x0, l_.fstr1@PAGE
Lloh12:
	add	x0, x0, l_.fstr1@PAGEOFF
	str	x8, [sp]
	bl	_printf
LBB1_3:                                 ; %L11
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #24]
	add	sp, sp, #48
	ret
	.loh AdrpLdrGotLdr	Lloh6, Lloh7, Lloh8
	.loh AdrpAdd	Lloh9, Lloh10
	.loh AdrpAdd	Lloh11, Lloh12
	.cfi_endproc
                                        ; -- End function
	.globl	_Del                            ; -- Begin function Del
	.p2align	2
_Del:                                   ; @Del
	.cfi_startproc
; %bb.0:                                ; %L12
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
Lloh13:
	adrp	x8, _head@GOTPAGE
Lloh14:
	ldr	x8, [x8, _head@GOTPAGEOFF]
	stp	x1, x0, [sp]
	ldr	x9, [x8]
	ldr	x9, [x9]
	cmp	x9, x1
	b.ne	LBB2_2
; %bb.1:                                ; %L17
	ldr	x0, [x8]
	ldr	x9, [x0, #8]
	str	x0, [sp, #16]
	str	x9, [x8]
	bl	_free
LBB2_2:                                 ; %L19
Lloh15:
	adrp	x8, _tail@GOTPAGE
	ldr	x9, [sp, #8]
Lloh16:
	ldr	x8, [x8, _tail@GOTPAGEOFF]
	ldr	x9, [x9, #8]
	ldr	x10, [x8]
	cmp	x9, x10
	b.ne	LBB2_5
; %bb.3:                                ; %L20
Lloh17:
	adrp	x9, _.nilNode@GOTPAGE
Lloh18:
	ldr	x9, [x9, _.nilNode@GOTPAGEOFF]
	ldr	x0, [x8]
	ldr	x10, [sp, #8]
Lloh19:
	ldr	x9, [x9]
	str	x0, [sp, #16]
	str	x10, [x8]
	str	x9, [x10, #8]
LBB2_4:                                 ; %L22
	bl	_free
LBB2_5:                                 ; %L22
                                        ; =>This Inner Loop Header: Depth=1
	ldp	x9, x8, [sp]
	ldr	x0, [x8, #8]
	ldr	x8, [x0]
	cmp	x8, x9
	b.eq	LBB2_7
; %bb.6:                                ; %L24
                                        ;   in Loop: Header=BB2_5 Depth=1
	ldr	x1, [sp]
	bl	_Del
	b	LBB2_5
LBB2_7:                                 ; %L23
	ldr	x8, [sp, #8]
	str	x0, [sp, #16]
	ldr	x9, [x8, #8]
	ldr	x9, [x9, #8]
	str	x9, [x8, #8]
	b	LBB2_4
	.loh AdrpLdrGot	Lloh13, Lloh14
	.loh AdrpLdrGot	Lloh15, Lloh16
	.loh AdrpLdrGotLdr	Lloh17, Lloh18, Lloh19
	.cfi_endproc
                                        ; -- End function
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:                                ; %L30
	sub	sp, sp, #80
	.cfi_def_cfa_offset 80
	stp	x20, x19, [sp, #48]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #64]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
Lloh20:
	adrp	x19, l_.read@PAGE
Lloh21:
	add	x19, x19, l_.read@PAGEOFF
	add	x8, sp, #32
	mov	x0, x19
	mov	w20, #61568
	movk	w20, #762, lsl #16
	str	x8, [sp]
	bl	_scanf
	add	x8, sp, #24
	mov	x0, x19
	str	x8, [sp]
	bl	_scanf
	mov	w0, #1
	bl	_Add
	mov	w0, #10
	bl	_Add
	mov	w0, #3
	bl	_Add
	mov	w0, #4
	bl	_Add
	ldr	x0, [sp, #32]
	bl	_Add
Lloh22:
	adrp	x19, _head@GOTPAGE
Lloh23:
	ldr	x19, [x19, _head@GOTPAGEOFF]
	ldr	x0, [x19]
	bl	_PrintList
	str	xzr, [sp, #16]
LBB3_1:                                 ; %L31
                                        ; =>This Inner Loop Header: Depth=1
	ldr	x8, [sp, #16]
	cmp	x8, x20
	b.ge	LBB3_3
; %bb.2:                                ; %L32
                                        ;   in Loop: Header=BB3_1 Depth=1
	ldr	x0, [sp, #16]
	bl	_Add
	ldr	x8, [sp, #16]
	add	x8, x8, #1
	str	x8, [sp, #16]
	b	LBB3_1
LBB3_3:                                 ; %L33
	str	xzr, [sp, #16]
LBB3_4:                                 ; %L34
                                        ; =>This Inner Loop Header: Depth=1
	ldr	x8, [sp, #16]
	ldr	x0, [x19]
	cmp	x8, x20
	b.ge	LBB3_6
; %bb.5:                                ; %L35
                                        ;   in Loop: Header=BB3_4 Depth=1
	ldr	x1, [sp, #16]
	bl	_Del
	ldr	x8, [sp, #16]
	add	x8, x8, #1
	str	x8, [sp, #16]
	b	LBB3_4
LBB3_6:                                 ; %L36
	ldr	x1, [sp, #24]
	bl	_Del
	ldr	x0, [x19]
	bl	_PrintList
	ldp	x29, x30, [sp, #64]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #40]
	ldp	x20, x19, [sp, #48]             ; 16-byte Folded Reload
	add	sp, sp, #80
	ret
	.loh AdrpLdrGot	Lloh22, Lloh23
	.loh AdrpAdd	Lloh20, Lloh21
	.cfi_endproc
                                        ; -- End function
	.comm	_.nilNode,8,3                   ; @.nilNode
	.comm	_head,8,3                       ; @head
	.comm	_tail,8,3                       ; @tail
	.section	__TEXT,__cstring,cstring_literals
l_.fstr1:                               ; @.fstr1
	.asciz	"%ld"

l_.fstr2:                               ; @.fstr2
	.asciz	"%ld"

l_.read:                                ; @.read
	.asciz	"%ld"

.subsections_via_symbols
