	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 13, 0
	.globl	_fact                           ; -- Begin function fact
	.p2align	2
_fact:                                  ; @fact
	.cfi_startproc
; %bb.0:                                ; %L0
	sub	sp, sp, #32
	.cfi_def_cfa_offset 32
	stp	x29, x30, [sp, #16]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	cmp	x0, #2
	str	x0, [sp]
	b.ge	LBB0_2
; %bb.1:                                ; %L2
	mov	w0, #1
	b	LBB0_3
LBB0_2:                                 ; %L3
	ldr	x8, [sp]
	sub	x0, x8, #1
	bl	_fact
	ldr	x8, [sp]
	mul	x0, x8, x0
LBB0_3:                                 ; %common.ret
	ldp	x29, x30, [sp, #16]             ; 16-byte Folded Reload
	str	x0, [sp, #8]
	add	sp, sp, #32
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:                                ; %L6
	sub	sp, sp, #112
	.cfi_def_cfa_offset 112
	stp	x24, x23, [sp, #48]             ; 16-byte Folded Spill
	stp	x22, x21, [sp, #64]             ; 16-byte Folded Spill
	stp	x20, x19, [sp, #80]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #96]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	.cfi_offset w21, -40
	.cfi_offset w22, -48
	.cfi_offset w23, -56
	.cfi_offset w24, -64
	add	x21, sp, #24
	add	x22, sp, #16
	mov	w23, #1
Lloh0:
	adrp	x19, l_.read@PAGE
Lloh1:
	add	x19, x19, l_.read@PAGEOFF
Lloh2:
	adrp	x20, l_.fstr2@PAGE
Lloh3:
	add	x20, x20, l_.fstr2@PAGEOFF
	stp	xzr, xzr, [sp, #24]
LBB1_1:                                 ; %L7
                                        ; =>This Inner Loop Header: Depth=1
	ldr	x8, [sp, #32]
	cbnz	x8, LBB1_4
; %bb.2:                                ; %L8
                                        ;   in Loop: Header=BB1_1 Depth=1
	mov	x0, x19
	str	x21, [sp]
	bl	_scanf
	ldr	x0, [sp, #24]
	bl	_fact
	mov	x8, x0
	stp	x8, x0, [sp]
	mov	x0, x20
	bl	_printf
	mov	x0, x19
	str	x22, [sp]
	bl	_scanf
	ldr	x8, [sp, #16]
	cbnz	x8, LBB1_1
; %bb.3:                                ; %L10
                                        ;   in Loop: Header=BB1_1 Depth=1
	str	x23, [sp, #32]
	b	LBB1_1
LBB1_4:                                 ; %L14
	ldp	x29, x30, [sp, #96]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #40]
	ldp	x20, x19, [sp, #80]             ; 16-byte Folded Reload
	ldp	x22, x21, [sp, #64]             ; 16-byte Folded Reload
	ldp	x24, x23, [sp, #48]             ; 16-byte Folded Reload
	add	sp, sp, #112
	ret
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
