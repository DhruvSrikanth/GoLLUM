	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 13, 0
	.globl	_something                      ; -- Begin function something
	.p2align	2
_something:                             ; @something
	.cfi_startproc
; %bb.0:                                ; %L0
	sub	sp, sp, #16
	.cfi_def_cfa_offset 16
	add	x8, x0, #1
	stp	x0, x8, [sp], #16
	mov	x0, x8
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:                                ; %L2
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	mov	w8, #1
	mov	w0, #1
	str	x8, [sp, #16]
	bl	_something
	mov	x8, x0
	stp	x8, x0, [sp]
Lloh0:
	adrp	x0, l_.fstr1@PAGE
Lloh1:
	add	x0, x0, l_.fstr1@PAGEOFF
	bl	_printf
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #24]
	add	sp, sp, #48
	ret
	.loh AdrpAdd	Lloh0, Lloh1
	.cfi_endproc
                                        ; -- End function
	.section	__TEXT,__cstring,cstring_literals
l_.fstr1:                               ; @.fstr1
	.asciz	"%ld\n"

.subsections_via_symbols
