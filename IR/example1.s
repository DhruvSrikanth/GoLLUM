	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 13, 0
	.globl	_fib1                           ; -- Begin function fib1
	.p2align	2
_fib1:                                  ; @fib1
	.cfi_startproc
; %bb.0:                                ; %L0
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	stp	x20, x19, [sp, #16]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	cmp	x0, #2
	str	x0, [sp]
	b.lt	LBB0_2
; %bb.1:                                ; %L3
	sub	x0, x0, #1
	bl	_fib1
	ldr	x8, [sp]
	mov	x19, x0
	sub	x0, x8, #2
	bl	_fib1
	add	x0, x19, x0
LBB0_2:                                 ; %common.ret
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	str	x0, [sp, #8]
	ldp	x20, x19, [sp, #16]             ; 16-byte Folded Reload
	add	sp, sp, #48
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_fib2                           ; -- Begin function fib2
	.p2align	2
_fib2:                                  ; @fib2
	.cfi_startproc
; %bb.0:                                ; %L6
	sub	sp, sp, #48
	.cfi_def_cfa_offset 48
	mov	w8, #1
	str	x0, [sp, #8]
	stp	x8, xzr, [sp, #24]
	ldr	x8, [sp, #8]
	cbz	x8, LBB1_2
LBB1_1:                                 ; %L8
                                        ; =>This Inner Loop Header: Depth=1
	ldp	x10, x9, [sp, #24]
	ldr	x8, [sp, #8]
	sub	x8, x8, #1
	add	x9, x9, x10
	stp	x9, x10, [sp, #24]
	stp	x8, x9, [sp, #8]
	ldr	x8, [sp, #8]
	cbnz	x8, LBB1_1
LBB1_2:                                 ; %L9
	ldr	x0, [sp, #32]
	str	x0, [sp, #40]
	add	sp, sp, #48
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:                                ; %L11
	sub	sp, sp, #96
	.cfi_def_cfa_offset 96
	stp	x20, x19, [sp, #64]             ; 16-byte Folded Spill
	stp	x29, x30, [sp, #80]             ; 16-byte Folded Spill
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	mov	w0, #16
	bl	_malloc
Lloh0:
	adrp	x19, l_.read@PAGE
Lloh1:
	add	x19, x19, l_.read@PAGEOFF
	str	x0, [sp, #48]
	add	x20, sp, #24
	mov	x0, x19
	str	x20, [sp]
	bl	_scanf
	ldr	x8, [sp, #48]
	mov	x0, x19
	ldr	x9, [sp, #24]
	str	x9, [x8]
	str	x20, [sp]
	bl	_scanf
	ldr	x8, [sp, #48]
	ldr	x9, [sp, #24]
	ldr	x0, [x8]
	str	x9, [x8, #8]
	bl	_fib1
	ldr	x8, [sp, #48]
	str	x0, [sp, #40]
	ldr	x8, [x8, #8]
	mov	x0, x8
	bl	_fib2
	ldr	x8, [sp, #48]
	str	x0, [sp, #32]
	mov	x0, x8
	bl	_free
	ldp	x9, x8, [sp, #32]
Lloh2:
	adrp	x0, l_.fstr2@PAGE
Lloh3:
	add	x0, x0, l_.fstr2@PAGEOFF
	stp	x8, x9, [sp]
	bl	_printf
	ldp	x29, x30, [sp, #80]             ; 16-byte Folded Reload
	mov	x0, xzr
	str	xzr, [sp, #56]
	ldp	x20, x19, [sp, #64]             ; 16-byte Folded Reload
	add	sp, sp, #96
	ret
	.loh AdrpAdd	Lloh2, Lloh3
	.loh AdrpAdd	Lloh0, Lloh1
	.cfi_endproc
                                        ; -- End function
	.comm	_.nilnums,8,3                   ; @.nilnums
	.section	__TEXT,__cstring,cstring_literals
l_.read:                                ; @.read
	.asciz	"%ld"

l_.fstr2:                               ; @.fstr2
	.asciz	"c=%ld\nd=%ld"

.subsections_via_symbols
