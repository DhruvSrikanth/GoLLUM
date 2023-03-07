	.arch armv8-a

	.comm .nilsimple,8,8
	.comm .nilfoo,8,8
	.comm globalfoo,8,8
	.comm unusedGlobal,8,8

	.text

	.type tailrecursive, %function
	.global tailrecursive
	.p2align 2
tailrecursive:
.L0:
	sub sp, sp, #112
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #112]
	b .L1

.L1:
	ldr x1, [sp, #112]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	mov x2, #0
	cmp x1, x2
	cset x3, le
	str x3, [sp, #40]
	ldr x1, [sp, #40]
	mov x2, #0
	cmp x1, x2
	b.eq .L3

.L2:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #48]
	ldr x1, [sp, #48]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #112
	ret

.L3:
	b .L4

.L4:
	mov x0, #24
	bl malloc
	str x0, [sp, #56]
	ldr x1, [sp, #56]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	adrp x2, unusedGlobal
	add x2, x2, :lo12:unusedGlobal
	str x1, [x2]
	ldr x1, [sp, #112]
	str x1, [sp, #80]
	ldr x1, [sp, #80]
	mov x2, #1
	sub x3, x1, x2
	str x3, [sp, #88]
	ldr x1, [sp, #88]
	mov x0, x1
	bl tailrecursive
	str x0, [sp, #96]
	b .L5

.L5:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #112
	ret

	.size tailrecursive, (.-tailrecursive)

	.type add, %function
	.global add
	.p2align 2
add:
.L6:
	sub sp, sp, #80
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #56]
	str x1, [sp, #56]
	ldr x2, [sp, #56]
	str x2, [sp, #24]
	ldr x2, [sp, #56]
	str x2, [sp, #32]
	ldr x2, [sp, #24]
	ldr x3, [sp, #32]
	add x4, x2, x3
	str x4, [sp, #40]
	ldr x2, [sp, #40]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #48]
	ldr x2, [sp, #48]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #80
	ret

	.size add, (.-add)

	.type domath, %function
	.global domath
	.p2align 2
domath:
.L7:
	sub sp, sp, #640
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #640]
	mov x0, #24
	bl malloc
	str x0, [sp, #48]
	ldr x1, [sp, #48]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	add x1, x1, #16
	str x1, [sp, #72]
	mov x0, #8
	bl malloc
	str x0, [sp, #80]
	ldr x1, [sp, #80]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	str x1, [sp, #72]
	mov x0, #24
	bl malloc
	str x0, [sp, #96]
	ldr x1, [sp, #96]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	str x1, [sp, #112]
	ldr x1, [sp, #112]
	add x1, x1, #16
	str x1, [sp, #120]
	mov x0, #8
	bl malloc
	str x0, [sp, #128]
	ldr x1, [sp, #128]
	str x1, [sp, #136]
	ldr x1, [sp, #136]
	str x1, [sp, #120]
	ldr x1, [sp, #24]
	str x1, [sp, #144]
	ldr x1, [sp, #144]
	add x1, x1, #0
	str x1, [sp, #152]
	ldr x1, [sp, #640]
	str x1, [sp, #160]
	ldr x1, [sp, #160]
	str x1, [sp, #152]
	ldr x1, [sp, #32]
	str x1, [sp, #168]
	ldr x1, [sp, #168]
	add x1, x1, #0
	str x1, [sp, #176]
	mov x1, #3
	str x1, [sp, #176]
	ldr x1, [sp, #24]
	str x1, [sp, #184]
	ldr x1, [sp, #184]
	add x1, x1, #16
	str x1, [sp, #192]
	ldr x1, [sp, #192]
	str x1, [sp, #200]
	ldr x1, [sp, #200]
	add x1, x1, #0
	str x1, [sp, #208]
	ldr x1, [sp, #24]
	str x1, [sp, #216]
	ldr x1, [sp, #216]
	add x1, x1, #0
	str x1, [sp, #224]
	ldr x1, [sp, #224]
	str x1, [sp, #232]
	ldr x1, [sp, #232]
	str x1, [sp, #208]
	ldr x1, [sp, #32]
	str x1, [sp, #240]
	ldr x1, [sp, #240]
	add x1, x1, #16
	str x1, [sp, #248]
	ldr x1, [sp, #248]
	str x1, [sp, #256]
	ldr x1, [sp, #256]
	add x1, x1, #0
	str x1, [sp, #264]
	ldr x1, [sp, #32]
	str x1, [sp, #272]
	ldr x1, [sp, #272]
	add x1, x1, #0
	str x1, [sp, #280]
	ldr x1, [sp, #280]
	str x1, [sp, #288]
	ldr x1, [sp, #288]
	str x1, [sp, #264]
	b .L8

.L8:
	ldr x1, [sp, #640]
	str x1, [sp, #296]
	ldr x1, [sp, #296]
	mov x2, #0
	cmp x1, x2
	cset x3, gt
	str x3, [sp, #304]
	ldr x1, [sp, #304]
	mov x2, #0
	cmp x1, x2
	b.eq .L10

.L9:
	ldr x1, [sp, #24]
	str x1, [sp, #312]
	ldr x1, [sp, #312]
	add x1, x1, #0
	str x1, [sp, #320]
	ldr x1, [sp, #32]
	str x1, [sp, #328]
	ldr x1, [sp, #328]
	add x1, x1, #0
	str x1, [sp, #336]
	ldr x1, [sp, #320]
	str x1, [sp, #344]
	ldr x1, [sp, #336]
	str x1, [sp, #352]
	ldr x1, [sp, #344]
	ldr x2, [sp, #352]
	mul x3, x1, x2
	str x3, [sp, #360]
	ldr x1, [sp, #360]
	str x1, [sp, #40]
	ldr x1, [sp, #24]
	str x1, [sp, #368]
	ldr x1, [sp, #368]
	add x1, x1, #16
	str x1, [sp, #376]
	ldr x1, [sp, #376]
	str x1, [sp, #384]
	ldr x1, [sp, #384]
	add x1, x1, #0
	str x1, [sp, #392]
	ldr x1, [sp, #40]
	str x1, [sp, #400]
	ldr x1, [sp, #392]
	str x1, [sp, #408]
	ldr x1, [sp, #400]
	ldr x2, [sp, #408]
	mul x3, x1, x2
	str x3, [sp, #416]
	ldr x1, [sp, #32]
	str x1, [sp, #424]
	ldr x1, [sp, #424]
	add x1, x1, #0
	str x1, [sp, #432]
	ldr x1, [sp, #432]
	str x1, [sp, #440]
	ldr x1, [sp, #416]
	ldr x2, [sp, #440]
	sdiv x3, x1, x2
	str x3, [sp, #448]
	ldr x1, [sp, #448]
	str x1, [sp, #40]
	ldr x1, [sp, #32]
	str x1, [sp, #456]
	ldr x1, [sp, #456]
	add x1, x1, #16
	str x1, [sp, #464]
	ldr x1, [sp, #464]
	str x1, [sp, #472]
	ldr x1, [sp, #472]
	add x1, x1, #0
	str x1, [sp, #480]
	ldr x1, [sp, #480]
	str x1, [sp, #488]
	ldr x1, [sp, #24]
	str x1, [sp, #496]
	ldr x1, [sp, #496]
	add x1, x1, #0
	str x1, [sp, #504]
	ldr x1, [sp, #504]
	str x1, [sp, #512]
	ldr x1, [sp, #488]
	ldr x2, [sp, #512]
	mov x0, x1
	mov x1, x2
	bl add
	str x0, [sp, #520]
	ldr x1, [sp, #520]
	str x1, [sp, #40]
	ldr x1, [sp, #32]
	str x1, [sp, #528]
	ldr x1, [sp, #528]
	add x1, x1, #0
	str x1, [sp, #536]
	ldr x1, [sp, #24]
	str x1, [sp, #544]
	ldr x1, [sp, #544]
	add x1, x1, #0
	str x1, [sp, #552]
	ldr x1, [sp, #536]
	str x1, [sp, #560]
	ldr x1, [sp, #552]
	str x1, [sp, #568]
	ldr x1, [sp, #560]
	ldr x2, [sp, #568]
	sub x3, x1, x2
	str x3, [sp, #576]
	ldr x1, [sp, #576]
	str x1, [sp, #40]
	ldr x1, [sp, #640]
	str x1, [sp, #584]
	ldr x1, [sp, #584]
	mov x2, #1
	sub x3, x1, x2
	str x3, [sp, #592]
	ldr x1, [sp, #592]
	str x1, [sp, #640]
	b .L8

.L10:
	ldr x1, [sp, #24]
	str x1, [sp, #600]
	ldr x1, [sp, #600]
	str x1, [sp, #608]
	ldr x1, [sp, #608]
	mov x0, x1
	bl free
	ldr x1, [sp, #32]
	str x1, [sp, #616]
	ldr x1, [sp, #616]
	str x1, [sp, #624]
	ldr x1, [sp, #624]
	mov x0, x1
	bl free
	b .L11

.L11:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #632]
	ldr x1, [sp, #632]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #640
	ret

	.size domath, (.-domath)

	.type objinstantiation, %function
	.global objinstantiation
	.p2align 2
objinstantiation:
.L12:
	sub sp, sp, #112
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #104]
	b .L13

.L13:
	ldr x1, [sp, #104]
	str x1, [sp, #32]
	ldr x1, [sp, #32]
	mov x2, #0
	cmp x1, x2
	cset x3, gt
	str x3, [sp, #40]
	ldr x1, [sp, #40]
	mov x2, #0
	cmp x1, x2
	b.eq .L15

.L14:
	mov x0, #24
	bl malloc
	str x0, [sp, #48]
	ldr x1, [sp, #48]
	str x1, [sp, #56]
	ldr x1, [sp, #56]
	str x1, [sp, #24]
	ldr x1, [sp, #24]
	str x1, [sp, #64]
	ldr x1, [sp, #64]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	mov x0, x1
	bl free
	ldr x1, [sp, #104]
	str x1, [sp, #80]
	ldr x1, [sp, #80]
	mov x2, #1
	sub x3, x1, x2
	str x3, [sp, #88]
	ldr x1, [sp, #88]
	str x1, [sp, #104]
	b .L13

.L15:
	b .L16

.L16:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #96]
	ldr x1, [sp, #96]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #112
	ret

	.size objinstantiation, (.-objinstantiation)

	.type ackermann, %function
	.global ackermann
	.p2align 2
ackermann:
.L17:
	sub sp, sp, #208
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	str x0, [sp, #184]
	str x1, [sp, #192]
	b .L18

.L18:
	ldr x2, [sp, #184]
	str x2, [sp, #24]
	ldr x2, [sp, #24]
	mov x3, #0
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #32]
	ldr x2, [sp, #32]
	mov x3, #0
	cmp x2, x3
	b.eq .L20

.L19:
	ldr x2, [sp, #192]
	str x2, [sp, #40]
	ldr x2, [sp, #40]
	mov x3, #1
	add x4, x2, x3
	str x4, [sp, #48]
	ldr x2, [sp, #48]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #56]
	ldr x2, [sp, #56]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #208
	ret

.L20:
	b .L21

.L21:
	b .L22

.L22:
	ldr x2, [sp, #192]
	str x2, [sp, #64]
	ldr x2, [sp, #64]
	mov x3, #0
	cmp x2, x3
	cset x4, eq
	str x4, [sp, #72]
	ldr x2, [sp, #72]
	mov x3, #0
	cmp x2, x3
	b.eq .L24

.L23:
	ldr x2, [sp, #184]
	str x2, [sp, #80]
	ldr x2, [sp, #80]
	mov x3, #1
	sub x4, x2, x3
	str x4, [sp, #88]
	ldr x2, [sp, #88]
	mov x3, #1
	mov x0, x2
	mov x1, x3
	bl ackermann
	str x0, [sp, #96]
	ldr x2, [sp, #96]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #104]
	ldr x2, [sp, #104]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #208
	ret

.L24:
	ldr x2, [sp, #184]
	str x2, [sp, #112]
	ldr x2, [sp, #112]
	mov x3, #1
	sub x4, x2, x3
	str x4, [sp, #120]
	ldr x2, [sp, #184]
	str x2, [sp, #128]
	ldr x2, [sp, #192]
	str x2, [sp, #136]
	ldr x2, [sp, #136]
	mov x3, #1
	sub x4, x2, x3
	str x4, [sp, #144]
	ldr x2, [sp, #128]
	ldr x3, [sp, #144]
	mov x0, x2
	mov x1, x3
	bl ackermann
	str x0, [sp, #152]
	ldr x2, [sp, #120]
	ldr x3, [sp, #152]
	mov x0, x2
	mov x1, x3
	bl ackermann
	str x0, [sp, #160]
	ldr x2, [sp, #160]
	str x2, [sp, #16]
	ldr x2, [sp, #16]
	str x2, [sp, #168]
	ldr x2, [sp, #168]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #208
	ret

.L25:
	b .L26

.L26:
	ldr x2, [sp, #16]
	str x2, [sp, #176]
	ldr x2, [sp, #176]
	mov x0, x2
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #208
	ret

	.size ackermann, (.-ackermann)

	.type main, %function
	.global main
	.p2align 2
main:
.L27:
	sub sp, sp, #176
	sub sp, sp, #16
	stp x29, x30, [sp]
	mov fp, sp
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #24
	bl scanf
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #32
	bl scanf
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #40
	bl scanf
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #48
	bl scanf
	adrp x1, .READ
	add x1, x1, :lo12:.READ
	mov x0, x1
	add x1, sp, #56
	bl scanf
	ldr x1, [sp, #24]
	str x1, [sp, #72]
	ldr x1, [sp, #72]
	mov x0, x1
	bl tailrecursive
	str x0, [sp, #80]
	ldr x1, [sp, #32]
	str x1, [sp, #88]
	ldr x1, [sp, #88]
	mov x0, x1
	bl domath
	str x0, [sp, #96]
	ldr x1, [sp, #40]
	str x1, [sp, #104]
	ldr x1, [sp, #104]
	mov x0, x1
	bl objinstantiation
	str x0, [sp, #112]
	ldr x1, [sp, #48]
	str x1, [sp, #120]
	ldr x1, [sp, #56]
	str x1, [sp, #128]
	ldr x1, [sp, #120]
	ldr x2, [sp, #128]
	mov x0, x1
	mov x1, x2
	bl ackermann
	str x0, [sp, #136]
	ldr x1, [sp, #136]
	str x1, [sp, #64]
	ldr x1, [sp, #24]
	str x1, [sp, #144]
	ldr x1, [sp, #32]
	str x1, [sp, #152]
	ldr x1, [sp, #40]
	str x1, [sp, #160]
	ldr x1, [sp, #64]
	str x1, [sp, #168]
	ldr x1, [sp, #144]
	ldr x2, [sp, #152]
	ldr x3, [sp, #160]
	ldr x4, [sp, #168]
	mov x1, x1
	mov x2, x2
	mov x3, x3
	mov x4, x4
	adrp x5, .FSTR2
	add x5, x5, :lo12:.FSTR2
	mov x0, x5
	bl printf
	b .L28

.L28:
	mov x1, #0
	str x1, [sp, #16]
	ldr x1, [sp, #16]
	str x1, [sp, #176]
	ldr x1, [sp, #176]
	mov x0, x1
	ldp x29, x30, [sp]
	add sp, sp, #16
	add sp, sp, #176
	ret

	.size main, (.-main)


.READ:
	.asciz	"%ld"
	.size	.READ, 4

.FSTR2:
	.asciz	"a=%ld\nb=%ld\nc=%ld\n,temp=%ld\n"
	.size	.FSTR2, 33

