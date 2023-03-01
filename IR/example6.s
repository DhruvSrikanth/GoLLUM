	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 13, 0
	.globl	_isqrt                          ; -- Begin function isqrt
	.p2align	2
_isqrt:                                 ; @isqrt
	.cfi_startproc
; %bb.0:                                ; %L0
	mov	w8, #1
	mov	w9, #3
	str	x0, [sp, #-32]!
	.cfi_def_cfa_offset 32
LBB0_1:                                 ; %L1
                                        ; =>This Inner Loop Header: Depth=1
	stp	x9, x8, [sp, #8]
	ldr	x9, [sp]
	cmp	x8, x9
	b.gt	LBB0_3
; %bb.2:                                ; %L2
                                        ;   in Loop: Header=BB0_1 Depth=1
	ldp	x9, x8, [sp, #8]
	add	x8, x8, x9
	add	x9, x9, #2
	b	LBB0_1
LBB0_3:                                 ; %L3
	ldr	x8, [sp, #8]
	cmp	x8, #0
	cinc	x8, x8, lt
	asr	x8, x8, #1
	sub	x0, x8, #1
	str	x0, [sp, #24]
	add	sp, sp, #32
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_prime                          ; -- Begin function prime
	.p2align	2
_prime:                                 ; @prime
	.cfi_startproc
; %bb.0:                                ; %L5
	sub	sp, sp, #64
	.cfi_def_cfa_offset 64
	stp	x29, x30, [sp, #48]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	cmp	x0, #2
	str	x0, [sp, #8]
	b.ge	LBB1_2
LBB1_1:                                 ; %L12
	ldp	x29, x30, [sp, #48]             ; 16-byte Folded Reload
	str	xzr, [sp, #40]
	mov	x0, xzr
	add	sp, sp, #64
	ret
LBB1_2:                                 ; %L8
	ldr	x0, [sp, #8]
	bl	_isqrt
	mov	w8, #2
	str	x0, [sp, #32]
LBB1_3:                                 ; %L9
                                        ; =>This Inner Loop Header: Depth=1
	ldr	x9, [sp, #32]
	str	x8, [sp, #24]
	cmp	x8, x9
	b.gt	LBB1_5
; %bb.4:                                ; %L10
                                        ;   in Loop: Header=BB1_3 Depth=1
	ldr	x8, [sp, #8]
	ldr	x9, [sp, #24]
	sdiv	x10, x8, x9
	msub	x8, x10, x9, x8
	str	x8, [sp, #16]
LBB1_5:                                 ; %L11
                                        ;   in Loop: Header=BB1_3 Depth=1
	ldr	x8, [sp, #16]
	cbz	x8, LBB1_1
; %bb.6:                                ; %L14
                                        ;   in Loop: Header=BB1_3 Depth=1
	ldr	x8, [sp, #24]
	add	x8, x8, #1
	b	LBB1_3
	.cfi_endproc
                                        ; -- End function
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:                                ; %L18
	sub	sp, sp, #64
	.cfi_def_cfa_offset 64
	stp	x20, x19, [sp, #32]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #48]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
Lloh0:
	adrp	x0, l_.read@PAGE
Lloh1:
	add	x0, x0, l_.read@PAGEOFF
	add	x8, sp, #16
	str	x8, [sp]
	bl	_scanf
Lloh2:
	adrp	x19, l_.fstr2@PAGE
Lloh3:
	add	x19, x19, l_.fstr2@PAGEOFF
	str	xzr, [sp, #8]
	b	LBB2_2
LBB2_1:                                 ; %L24
                                        ;   in Loop: Header=BB2_2 Depth=1
	ldr	x8, [sp, #8]
	add	x8, x8, #1
	str	x8, [sp, #8]
LBB2_2:                                 ; %L19
                                        ; =>This Inner Loop Header: Depth=1
	ldr	x0, [sp, #8]
	bl	_prime
	cmp	x0, #1
	b.ne	LBB2_1
; %bb.3:                                ; %L22
                                        ;   in Loop: Header=BB2_2 Depth=1
	ldr	x8, [sp, #8]
	mov	x0, x19
	str	x8, [sp]
	bl	_printf
	b	LBB2_1
	.loh AdrpAdd	Lloh2, Lloh3
	.loh AdrpAdd	Lloh0, Lloh1
	.cfi_endproc
                                        ; -- End function
	.section	__TEXT,__cstring,cstring_literals
l_.read:                                ; @.read
	.asciz	"%ld"

l_.fstr2:                               ; @.fstr2
	.asciz	"%ld"

.subsections_via_symbols
