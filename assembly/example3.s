	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 13, 0
	.globl	_timesone                       ; -- Begin function timesone
	.p2align	2
_timesone:                              ; @timesone
	.cfi_startproc
; %bb.0:                                ; %L0
	str	x0, [sp, #-16]!
	.cfi_def_cfa_offset 16
LBB0_1:                                 ; %L1
                                        ; =>This Inner Loop Header: Depth=1
	ldr	x8, [sp]
	cmp	x8, #1
	b.lt	LBB0_3
; %bb.2:                                ; %L2
                                        ;   in Loop: Header=BB0_1 Depth=1
	ldr	x8, [sp]
	sub	x8, x8, #1
	str	x8, [sp]
	b	LBB0_1
LBB0_3:                                 ; %L4
	mov	x0, xzr
	str	xzr, [sp, #8]
	add	sp, sp, #16
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_find                           ; -- Begin function find
	.p2align	2
_find:                                  ; @find
	.cfi_startproc
; %bb.0:                                ; %L5
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	ldr	x8, [x1, #8]
	stp	x1, x0, [sp, #8]
	cmp	x8, x0
	b.ne	LBB1_2
; %bb.1:                                ; %L9
	ldr	x8, [sp, #8]
	ldr	x0, [x8, #16]
	b	LBB1_3
LBB1_2:                                 ; %L10
	ldp	x8, x0, [sp, #8]
	ldr	x1, [x8]
	bl	_find
LBB1_3:                                 ; %common.ret
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	str	x0, [sp, #24]
	add	sp, sp, #48
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_add                            ; -- Begin function add
	.p2align	2
_add:                                   ; @add
	.cfi_startproc
; %bb.0:                                ; %L14
	sub	sp, sp, #64
	.cfi_def_cfa_offset 64
	stp	x20, x19, [sp, #32]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #48]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
Lloh0:
	adrp	x19, _.nilListNode@GOTPAGE
Lloh1:
	ldr	x19, [x19, _.nilListNode@GOTPAGEOFF]
	stp	x1, x0, [sp, #8]
	str	x2, [sp]
	ldr	x8, [x19]
	cmp	x2, x8
	b.eq	LBB2_2
; %bb.1:                                ; %L17
	ldp	x19, x1, [sp]
	ldr	x0, [sp, #16]
	ldr	x2, [x19]
	bl	_add
	str	x0, [x19]
	b	LBB2_3
LBB2_2:                                 ; %L16
	mov	w0, #24
	bl	_malloc
	ldp	x9, x8, [sp, #8]
	str	x0, [sp]
	stp	x8, x9, [x0, #8]
	ldr	x8, [x19]
	str	x8, [x0]
LBB2_3:                                 ; %L18
	ldp	x29, x30, [sp, #48]             ; 16-byte Folded Reload
	ldp	x20, x19, [sp, #32]             ; 16-byte Folded Reload
	ldr	x0, [sp]
	str	x0, [sp, #24]
	add	sp, sp, #64
	ret
	.loh AdrpLdrGot	Lloh0, Lloh1
	.cfi_endproc
                                        ; -- End function
	.globl	_factorial                      ; -- Begin function factorial
	.p2align	2
_factorial:                             ; @factorial
	.cfi_startproc
; %bb.0:                                ; %L19
	sub	sp, sp, #64
	.cfi_def_cfa_offset 64
	stp	x29, x30, [sp, #48]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	stp	x1, x0, [sp, #8]
	mov	w0, #100
	bl	_timesone
	ldr	x8, [sp, #16]
	cmp	x8, #1
	b.ne	LBB3_2
; %bb.1:                                ; %L21
	mov	w0, #1
	b	LBB3_7
LBB3_2:                                 ; %L22
	ldp	x1, x0, [sp, #8]
	bl	_find
	cmn	x0, #1
	str	x0, [sp, #32]
	b.eq	LBB3_4
; %bb.3:                                ; %L24
	ldr	x0, [sp, #32]
	b	LBB3_7
LBB3_4:                                 ; %L26
	ldp	x1, x8, [sp, #8]
	sub	x0, x8, #1
	bl	_factorial
	ldr	x8, [sp, #16]
	mov	x9, #6148914691236517205
	movk	x9, #21846
	mul	x8, x8, x0
	smulh	x9, x8, x9
	str	x8, [sp, #24]
	cmn	x9, x9, lsr #63
	b.ne	LBB3_6
; %bb.5:                                ; %L28
	ldp	x0, x1, [sp, #16]
	ldr	x2, [sp, #8]
	bl	_add
LBB3_6:                                 ; %L30
	ldr	x0, [sp, #24]
LBB3_7:                                 ; %common.ret
	ldp	x29, x30, [sp, #48]             ; 16-byte Folded Reload
	str	x0, [sp, #40]
	add	sp, sp, #64
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_maxfactorial                   ; -- Begin function maxfactorial
	.p2align	2
_maxfactorial:                          ; @maxfactorial
	.cfi_startproc
; %bb.0:                                ; %L33
	sub	sp, sp, #96
	.cfi_def_cfa_offset 96
	stp	x20, x19, [sp, #64]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #80]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
Lloh2:
	adrp	x8, _.nilListNode@GOTPAGE
Lloh3:
	adrp	x9, _matrix@GOTPAGE
Lloh4:
	adrp	x19, _.nilRow@GOTPAGE
Lloh5:
	adrp	x20, _.nilCell@GOTPAGE
Lloh6:
	ldr	x8, [x8, _.nilListNode@GOTPAGEOFF]
Lloh7:
	ldr	x8, [x8]
Lloh8:
	ldr	x9, [x9, _matrix@GOTPAGEOFF]
	str	x0, [sp, #24]
	str	x8, [x0]
Lloh9:
	ldr	x8, [x9]
Lloh10:
	ldr	x19, [x19, _.nilRow@GOTPAGEOFF]
Lloh11:
	ldr	x20, [x20, _.nilCell@GOTPAGEOFF]
	str	x1, [sp, #16]
	str	x8, [sp, #48]
LBB4_1:                                 ; %L34
                                        ; =>This Loop Header: Depth=1
                                        ;     Child Loop BB4_3 Depth 2
	ldr	x8, [sp, #48]
	ldr	x9, [x19]
	cmp	x8, x9
	b.eq	LBB4_6
; %bb.2:                                ; %L35
                                        ;   in Loop: Header=BB4_1 Depth=1
	ldr	x8, [sp, #48]
	ldp	x8, x9, [x8]
	stp	x9, x8, [sp, #40]
LBB4_3:                                 ; %L36
                                        ;   Parent Loop BB4_1 Depth=1
                                        ; =>  This Inner Loop Header: Depth=2
	ldr	x8, [sp, #40]
	ldr	x9, [x20]
	cmp	x8, x9
	b.eq	LBB4_1
; %bb.4:                                ; %L37
                                        ;   in Loop: Header=BB4_3 Depth=2
	ldr	x8, [sp, #40]
	ldr	x1, [sp, #24]
	ldr	x0, [x8, #8]
	bl	_factorial
	ldr	x8, [sp, #40]
	ldr	x9, [sp, #16]
	ldr	x8, [x8]
	cmp	x0, x9
	stp	x0, x8, [sp, #32]
	b.le	LBB4_3
; %bb.5:                                ; %L39
                                        ;   in Loop: Header=BB4_3 Depth=2
	ldr	x8, [sp, #32]
	str	x8, [sp, #16]
	b	LBB4_3
LBB4_6:                                 ; %L43
	ldr	x8, [sp, #16]
Lloh12:
	adrp	x0, l_.fstr1@PAGE
Lloh13:
	add	x0, x0, l_.fstr1@PAGEOFF
	str	x8, [sp]
	bl	_printf
	ldp	x29, x30, [sp, #80]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #56]
	ldp	x20, x19, [sp, #64]             ; 16-byte Folded Reload
	add	sp, sp, #96
	ret
	.loh AdrpLdrGot	Lloh5, Lloh11
	.loh AdrpLdrGot	Lloh4, Lloh10
	.loh AdrpLdrGotLdr	Lloh3, Lloh8, Lloh9
	.loh AdrpLdrGotLdr	Lloh2, Lloh6, Lloh7
	.loh AdrpAdd	Lloh12, Lloh13
	.cfi_endproc
                                        ; -- End function
	.globl	_newvalue                       ; -- Begin function newvalue
	.p2align	2
_newvalue:                              ; @newvalue
	.cfi_startproc
; %bb.0:                                ; %L45
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	stp	x1, x0, [sp]
	mov	x0, xzr
	bl	_timesone
Lloh14:
	adrp	x10, _maxrange@GOTPAGE
	ldp	x8, x9, [sp]
Lloh15:
	ldr	x10, [x10, _maxrange@GOTPAGEOFF]
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	mul	x8, x9, x8
Lloh16:
	ldr	x10, [x10]
	sdiv	x10, x10, x8
	add	x0, x10, x9
	stp	x8, x0, [sp, #16]
	add	sp, sp, #48
	ret
	.loh AdrpLdrGotLdr	Lloh14, Lloh15, Lloh16
	.cfi_endproc
                                        ; -- End function
	.globl	_newcell                        ; -- Begin function newcell
	.p2align	2
_newcell:                               ; @newcell
	.cfi_startproc
; %bb.0:                                ; %L46
	sub	sp, sp, #64
	.cfi_def_cfa_offset 64
	stp	x20, x19, [sp, #32]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #48]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	mov	x19, x0
	stp	x1, x0, [sp, #8]
	mov	x0, x1
	mov	x1, x2
	str	x2, [sp]
	bl	_newvalue
	ldr	x8, [sp]
	str	x0, [x19, #8]
	ldr	x20, [sp, #16]
	cmp	x8, #2
	b.lt	LBB6_2
; %bb.1:                                ; %L48
	mov	w0, #16
	bl	_malloc
	ldp	x8, x1, [sp]
	sub	x2, x8, #1
	bl	_newcell
	mov	x8, x0
	b	LBB6_3
LBB6_2:                                 ; %L49
Lloh17:
	adrp	x8, _.nilCell@GOTPAGE
Lloh18:
	ldr	x8, [x8, _.nilCell@GOTPAGEOFF]
Lloh19:
	ldr	x8, [x8]
LBB6_3:                                 ; %L50
	str	x8, [x20]
	ldr	x0, [sp, #16]
	ldp	x29, x30, [sp, #48]             ; 16-byte Folded Reload
	ldp	x20, x19, [sp, #32]             ; 16-byte Folded Reload
	str	x0, [sp, #24]
	add	sp, sp, #64
	ret
	.loh AdrpLdrGotLdr	Lloh17, Lloh18, Lloh19
	.cfi_endproc
                                        ; -- End function
	.globl	_newrow                         ; -- Begin function newrow
	.p2align	2
_newrow:                                ; @newrow
	.cfi_startproc
; %bb.0:                                ; %L51
	sub	sp, sp, #64
	.cfi_def_cfa_offset 64
	stp	x20, x19, [sp, #32]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #48]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	mov	x19, x0
	stp	x1, x0, [sp, #8]
	mov	w0, #16
	str	x2, [sp]
	bl	_malloc
	ldp	x2, x1, [sp]
	bl	_newcell
	ldp	x8, x20, [sp, #8]
	str	x0, [x19, #8]
	cmp	x8, #2
	b.lt	LBB7_2
; %bb.1:                                ; %L53
	mov	w0, #16
	bl	_malloc
	ldp	x2, x8, [sp]
	sub	x1, x8, #1
	bl	_newrow
	mov	x8, x0
	b	LBB7_3
LBB7_2:                                 ; %L54
Lloh20:
	adrp	x8, _.nilRow@GOTPAGE
Lloh21:
	ldr	x8, [x8, _.nilRow@GOTPAGEOFF]
Lloh22:
	ldr	x8, [x8]
LBB7_3:                                 ; %L55
	str	x8, [x20]
	ldr	x0, [sp, #16]
	ldp	x29, x30, [sp, #48]             ; 16-byte Folded Reload
	ldp	x20, x19, [sp, #32]             ; 16-byte Folded Reload
	str	x0, [sp, #24]
	add	sp, sp, #64
	ret
	.loh AdrpLdrGotLdr	Lloh20, Lloh21, Lloh22
	.cfi_endproc
                                        ; -- End function
	.globl	_newmatrix                      ; -- Begin function newmatrix
	.p2align	2
_newmatrix:                             ; @newmatrix
	.cfi_startproc
; %bb.0:                                ; %L56
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	stp	x1, x0, [sp, #8]
	mov	w0, #16
	bl	_malloc
	ldp	x2, x1, [sp, #8]
	bl	_newrow
Lloh23:
	adrp	x9, _matrix@GOTPAGE
	mov	x8, x0
	mov	x0, xzr
Lloh24:
	ldr	x9, [x9, _matrix@GOTPAGEOFF]
	str	xzr, [sp, #24]
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
Lloh25:
	str	x8, [x9]
	add	sp, sp, #48
	ret
	.loh AdrpLdrGotStr	Lloh23, Lloh24, Lloh25
	.cfi_endproc
                                        ; -- End function
	.globl	_getmatrixsize                  ; -- Begin function getmatrixsize
	.p2align	2
_getmatrixsize:                         ; @getmatrixsize
	.cfi_startproc
; %bb.0:                                ; %L58
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	cmp	x0, #0
	str	x0, [sp, #16]
	b.gt	LBB9_2
; %bb.1:                                ; %L60
Lloh26:
	adrp	x0, l_.read@PAGE
Lloh27:
	add	x0, x0, l_.read@PAGEOFF
	add	x8, sp, #16
	str	x8, [sp]
	bl	_scanf
	ldr	x0, [sp, #16]
	bl	_getmatrixsize
LBB9_2:                                 ; %common.ret
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	ldr	x0, [sp, #16]
	str	x0, [sp, #24]
	add	sp, sp, #48
	ret
	.loh AdrpAdd	Lloh26, Lloh27
	.cfi_endproc
                                        ; -- End function
	.globl	_getmaxrange                    ; -- Begin function getmaxrange
	.p2align	2
_getmaxrange:                           ; @getmaxrange
	.cfi_startproc
; %bb.0:                                ; %L63
	sub	sp, sp, #64
	.cfi_def_cfa_offset 64
	stp	x20, x19, [sp, #32]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #48]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
Lloh28:
	adrp	x19, _maxrange@GOTPAGE
Lloh29:
	ldr	x19, [x19, _maxrange@GOTPAGEOFF]
	str	x0, [sp, #16]
	ldr	x8, [x19]
	cmp	x8, #1
	b.gt	LBB10_2
; %bb.1:                                ; %L65
Lloh30:
	adrp	x0, l_.read@PAGE
Lloh31:
	add	x0, x0, l_.read@PAGEOFF
	str	x19, [sp]
	bl	_scanf
	ldr	x0, [x19]
	bl	_getmaxrange
LBB10_2:                                ; %common.ret
	ldr	x0, [x19]
	ldp	x29, x30, [sp, #48]             ; 16-byte Folded Reload
	ldp	x20, x19, [sp, #32]             ; 16-byte Folded Reload
	str	x0, [sp, #24]
	add	sp, sp, #64
	ret
	.loh AdrpLdrGot	Lloh28, Lloh29
	.loh AdrpAdd	Lloh30, Lloh31
	.cfi_endproc
                                        ; -- End function
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:                                ; %L68
	sub	sp, sp, #64
	.cfi_def_cfa_offset 64
	stp	x20, x19, [sp, #32]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #48]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
Lloh32:
	adrp	x19, _maxrange@GOTPAGE
	mov	x0, xzr
Lloh33:
	ldr	x19, [x19, _maxrange@GOTPAGEOFF]
	stp	xzr, xzr, [sp, #8]
	str	xzr, [x19]
	bl	_getmatrixsize
	ldr	x8, [x19]
	stp	x0, x0, [sp, #8]
	mov	x0, x8
	bl	_getmaxrange
	ldp	x1, x8, [sp, #8]
	str	x0, [x19]
	mov	x0, x8
	bl	_newmatrix
	mov	w0, #24
	bl	_malloc
	mov	x1, xzr
	bl	_maxfactorial
	ldp	x29, x30, [sp, #48]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #24]
	ldp	x20, x19, [sp, #32]             ; 16-byte Folded Reload
	add	sp, sp, #64
	ret
	.loh AdrpLdrGot	Lloh32, Lloh33
	.cfi_endproc
                                        ; -- End function
	.comm	_.nilCell,8,3                   ; @.nilCell
	.comm	_.nilRow,8,3                    ; @.nilRow
	.comm	_.nilListNode,8,3               ; @.nilListNode
	.comm	_matrix,8,3                     ; @matrix
	.comm	_maxrange,8,3                   ; @maxrange
	.section	__TEXT,__cstring,cstring_literals
l_.fstr1:                               ; @.fstr1
	.asciz	"Max Factorial=%ld"

l_.read:                                ; @.read
	.asciz	"%ld"

.subsections_via_symbols
