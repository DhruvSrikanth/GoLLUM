	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 13, 0
	.globl	_Init                           ; -- Begin function Init
	.p2align	2
_Init:                                  ; @Init
	.cfi_startproc
; %bb.0:                                ; %L0
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
Lloh0:
	adrp	x8, _.nilPoint2D@GOTPAGE
	cmp	x0, #0
Lloh1:
	ldr	x8, [x8, _.nilPoint2D@GOTPAGEOFF]
Lloh2:
	ldr	x8, [x8]
	stp	x0, x8, [sp, #8]
	b.le	LBB0_2
; %bb.1:                                ; %L2
	mov	w0, #16
	bl	_malloc
	ldr	x8, [sp, #8]
	str	x0, [sp, #16]
	stp	x8, x8, [x0]
	b	LBB0_3
LBB0_2:                                 ; %L4
	ldr	x0, [sp, #16]
LBB0_3:                                 ; %common.ret
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	str	x0, [sp, #24]
	add	sp, sp, #48
	ret
	.loh AdrpLdrGotLdr	Lloh0, Lloh1, Lloh2
	.cfi_endproc
                                        ; -- End function
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:                                ; %L5
	sub	sp, sp, #96
	.cfi_def_cfa_offset 96
	stp	x20, x19, [sp, #64]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #80]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	mov	w8, #5
	mov	w9, #36
	mov	w0, #16
	stp	x9, x8, [sp, #40]
	bl	_malloc
	ldp	x9, x8, [sp, #40]
Lloh3:
	adrp	x19, _globalInit@GOTPAGE
	str	x0, [sp, #32]
	stp	x8, x9, [x0]
Lloh4:
	adrp	x0, l_.read@PAGE
Lloh5:
	add	x0, x0, l_.read@PAGEOFF
Lloh6:
	ldr	x19, [x19, _globalInit@GOTPAGEOFF]
	str	x19, [sp]
	bl	_scanf
	ldr	x0, [x19]
	bl	_Init
	ldp	x9, x10, [x0]
	ldr	x8, [x19]
	stp	x10, x0, [sp, #16]
Lloh7:
	adrp	x0, l_.fstr2@PAGE
Lloh8:
	add	x0, x0, l_.fstr2@PAGEOFF
	stp	x8, x9, [sp]
	bl	_printf
	ldr	x0, [sp, #32]
	bl	_free
	ldr	x0, [sp, #24]
	bl	_free
	ldp	x29, x30, [sp, #80]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #56]
	ldp	x20, x19, [sp, #64]             ; 16-byte Folded Reload
	add	sp, sp, #96
	ret
	.loh AdrpAdd	Lloh7, Lloh8
	.loh AdrpAdd	Lloh4, Lloh5
	.loh AdrpLdrGot	Lloh3, Lloh6
	.cfi_endproc
                                        ; -- End function
	.comm	_.nilPoint2D,8,3                ; @.nilPoint2D
	.comm	_globalInit,8,3                 ; @globalInit
	.section	__TEXT,__cstring,cstring_literals
l_.read:                                ; @.read
	.asciz	"%ld"

l_.fstr2:                               ; @.fstr2
	.asciz	"offset=%ld\npt2.x=%ld\npt2.y=%ld\n"

.subsections_via_symbols
