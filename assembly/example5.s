	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 13, 0
	.globl	_tailrecursive                  ; -- Begin function tailrecursive
	.p2align	2
_tailrecursive:                         ; @tailrecursive
	.cfi_startproc
; %bb.0:                                ; %L0
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	cmp	x0, #1
	str	x0, [sp, #8]
	b.lt	LBB0_2
; %bb.1:                                ; %L4
	mov	w0, #24
	bl	_malloc
Lloh0:
	adrp	x9, _unusedGlobal@GOTPAGE
	ldr	x8, [sp, #8]
Lloh1:
	ldr	x9, [x9, _unusedGlobal@GOTPAGEOFF]
	sub	x8, x8, #1
	str	x0, [sp, #16]
Lloh2:
	str	x0, [x9]
	mov	x0, x8
	bl	_tailrecursive
LBB0_2:                                 ; %common.ret
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	str	xzr, [sp, #24]
	mov	x0, xzr
	add	sp, sp, #48
	ret
	.loh AdrpLdrGotStr	Lloh0, Lloh1, Lloh2
	.cfi_endproc
                                        ; -- End function
	.globl	_add                            ; -- Begin function add
	.p2align	2
_add:                                   ; @add
	.cfi_startproc
; %bb.0:                                ; %L6
	sub	sp, sp, #32
	.cfi_def_cfa_offset 32
	mov	x8, x0
	add	x0, x0, x1
	stp	x1, x8, [sp, #8]
	str	x0, [sp, #24]
	add	sp, sp, #32
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_domath                         ; -- Begin function domath
	.p2align	2
_domath:                                ; @domath
	.cfi_startproc
; %bb.0:                                ; %L7
	sub	sp, sp, #80
	.cfi_def_cfa_offset 80
	stp	x20, x19, [sp, #48]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #64]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	str	x0, [sp, #8]
	mov	w0, #24
	bl	_malloc
	mov	x19, x0
	str	x0, [sp, #32]
	mov	w0, #8
	bl	_malloc
	str	x0, [x19, #16]
	mov	w0, #24
	bl	_malloc
	mov	x19, x0
	str	x0, [sp, #24]
	mov	w0, #8
	bl	_malloc
	ldp	x10, x9, [sp, #24]
	mov	w11, #3
	str	x0, [x19, #16]
	ldr	x8, [sp, #8]
	str	x8, [x9]
	str	x11, [x10]
	ldr	x8, [x9, #16]
	ldr	x9, [x9]
	str	x9, [x8]
	ldr	x8, [x10, #16]
	ldr	x9, [x10]
	str	x9, [x8]
LBB2_1:                                 ; %L8
                                        ; =>This Inner Loop Header: Depth=1
	ldr	x8, [sp, #8]
	ldr	x0, [sp, #32]
	cmp	x8, #1
	b.lt	LBB2_3
; %bb.2:                                ; %L9
                                        ;   in Loop: Header=BB2_1 Depth=1
	ldp	x8, x9, [sp, #24]
	ldr	x10, [x0]
	ldr	x11, [x8]
	ldr	x12, [x9, #16]
	ldr	x8, [x8, #16]
	mul	x10, x10, x11
	ldr	x1, [x9]
	ldr	x12, [x12]
	ldr	x0, [x8]
	mul	x10, x10, x12
	sdiv	x10, x10, x11
	str	x10, [sp, #16]
	bl	_add
	ldp	x8, x9, [sp, #24]
	ldr	x10, [sp, #8]
	ldr	x8, [x8]
	ldr	x9, [x9]
	sub	x8, x8, x9
	sub	x9, x10, #1
	stp	x9, x8, [sp, #8]
	b	LBB2_1
LBB2_3:                                 ; %L10
	bl	_free
	ldr	x0, [sp, #24]
	bl	_free
	ldp	x29, x30, [sp, #64]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #40]
	ldp	x20, x19, [sp, #48]             ; 16-byte Folded Reload
	add	sp, sp, #80
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_objinstantiation               ; -- Begin function objinstantiation
	.p2align	2
_objinstantiation:                      ; @objinstantiation
	.cfi_startproc
; %bb.0:                                ; %L12
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	str	x0, [sp, #8]
LBB3_1:                                 ; %L13
                                        ; =>This Inner Loop Header: Depth=1
	ldr	x8, [sp, #8]
	cmp	x8, #1
	b.lt	LBB3_3
; %bb.2:                                ; %L14
                                        ;   in Loop: Header=BB3_1 Depth=1
	mov	w0, #24
	bl	_malloc
	str	x0, [sp, #16]
	bl	_free
	ldr	x8, [sp, #8]
	sub	x8, x8, #1
	str	x8, [sp, #8]
	b	LBB3_1
LBB3_3:                                 ; %L16
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #24]
	add	sp, sp, #48
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_ackermann                      ; -- Begin function ackermann
	.p2align	2
_ackermann:                             ; @ackermann
	.cfi_startproc
; %bb.0:                                ; %L17
	sub	sp, sp, #64
	.cfi_def_cfa_offset 64
	stp	x20, x19, [sp, #32]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #48]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	stp	x1, x0, [sp, #8]
	cbz	x0, LBB4_3
; %bb.1:                                ; %L22
	ldr	x8, [sp, #16]
	sub	x19, x8, #1
	cbz	x1, LBB4_4
; %bb.2:                                ; %L24
	ldp	x8, x0, [sp, #8]
	sub	x1, x8, #1
	bl	_ackermann
	mov	x1, x0
	mov	x0, x19
	b	LBB4_5
LBB4_3:                                 ; %L19
	add	x0, x1, #1
	b	LBB4_6
LBB4_4:                                 ; %L23
	mov	x0, x19
	mov	w1, #1
LBB4_5:                                 ; %common.ret
	bl	_ackermann
LBB4_6:                                 ; %common.ret
	ldp	x29, x30, [sp, #48]             ; 16-byte Folded Reload
	str	x0, [sp, #24]
	ldp	x20, x19, [sp, #32]             ; 16-byte Folded Reload
	add	sp, sp, #64
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:                                ; %L27
	sub	sp, sp, #128
	.cfi_def_cfa_offset 128
	stp	x20, x19, [sp, #96]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #112]            ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
Lloh3:
	adrp	x19, l_.read@PAGE
Lloh4:
	add	x19, x19, l_.read@PAGEOFF
	add	x8, sp, #80
	mov	x0, x19
	str	x8, [sp]
	bl	_scanf
	add	x8, sp, #72
	mov	x0, x19
	str	x8, [sp]
	bl	_scanf
	add	x8, sp, #64
	mov	x0, x19
	str	x8, [sp]
	bl	_scanf
	add	x8, sp, #56
	mov	x0, x19
	str	x8, [sp]
	bl	_scanf
	add	x8, sp, #48
	mov	x0, x19
	str	x8, [sp]
	bl	_scanf
	ldr	x0, [sp, #80]
	bl	_tailrecursive
	ldr	x0, [sp, #72]
	bl	_domath
	ldr	x0, [sp, #64]
	bl	_objinstantiation
	ldp	x1, x0, [sp, #48]
	bl	_ackermann
	ldr	x11, [sp, #64]
Lloh5:
	adrp	x8, l_.fstr2@PAGE
Lloh6:
	add	x8, x8, l_.fstr2@PAGEOFF
	str	x0, [sp, #40]
	ldp	x10, x9, [sp, #72]
	stp	x11, x0, [sp, #16]
	mov	x0, x8
	stp	x9, x10, [sp]
	bl	_printf
	ldp	x29, x30, [sp, #112]            ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #88]
	ldp	x20, x19, [sp, #96]             ; 16-byte Folded Reload
	add	sp, sp, #128
	ret
	.loh AdrpAdd	Lloh5, Lloh6
	.loh AdrpAdd	Lloh3, Lloh4
	.cfi_endproc
                                        ; -- End function
	.comm	_.nilsimple,8,3                 ; @.nilsimple
	.comm	_.nilfoo,8,3                    ; @.nilfoo
	.comm	_globalfoo,8,3                  ; @globalfoo
	.comm	_unusedGlobal,8,3               ; @unusedGlobal
	.section	__TEXT,__cstring,cstring_literals
l_.read:                                ; @.read
	.asciz	"%ld"

l_.fstr2:                               ; @.fstr2
	.asciz	"a=%ld\nb=%ld\nc=%ld\n,temp=%ld\n"

.subsections_via_symbols
