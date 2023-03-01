	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 13, 0
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:                                ; %L0
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	str	xzr, [sp, #16]
LBB0_1:                                 ; %L1
                                        ; =>This Inner Loop Header: Depth=1
	ldr	x8, [sp, #16]
	cmp	x8, #9
	b.gt	LBB0_3
; %bb.2:                                ; %L2
                                        ;   in Loop: Header=BB0_1 Depth=1
	add	x8, x8, #1
	str	x8, [sp, #16]
	b	LBB0_1
LBB0_3:                                 ; %L3
Lloh0:
	adrp	x0, l_.fstr1@PAGE
Lloh1:
	add	x0, x0, l_.fstr1@PAGEOFF
	str	x8, [sp]
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
	.asciz	"a=%ld\n"

.subsections_via_symbols
