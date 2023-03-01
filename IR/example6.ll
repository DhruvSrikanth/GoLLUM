source_filename = "example6"
target triple = "arm64-apple-darwin22.2.0"



define i64 @isqrt(i64 %a)
{
L0:
	%isqrt_retval = alloca i64
	%square = alloca i64
	%delta = alloca i64
	%P_a = alloca i64
	store i64 %a, i64* %P_a
	store i64 1, i64* %square
	store i64 3, i64* %delta
	br label %L1

L1:
	%r0 = load i64, i64* %square
	%r1 = load i64, i64* %P_a
	%r2 = icmp sle i64 %r0, %r1
	br i1 %r2, label %L2, label %L3

L2:
	%r3 = load i64, i64* %square
	%r4 = load i64, i64* %delta
	%r5 = add i64 %r3, %r4
	store i64 %r5, i64* %square
	%r6 = load i64, i64* %delta
	%r7 = add i64 %r6, 2
	store i64 %r7, i64* %delta
	br label %L1

L3:
	%r8 = load i64, i64* %delta
	%r9 = sdiv i64 %r8, 2
	%r10 = sub i64 %r9, 1
	store i64 %r10, i64* %isqrt_retval
	%r11 = load i64, i64* %isqrt_retval
	ret i64 %r11
	br label %L4

L4:
	%r12 = load i64, i64* %isqrt_retval
	ret i64 %r12


}

define i64 @prime(i64 %a)
{
L5:
	%prime_retval = alloca i64
	%max = alloca i64
	%divisor = alloca i64
	%remainder = alloca i64
	%P_a = alloca i64
	store i64 %a, i64* %P_a
	br label %L6

L6:
	%r13 = load i64, i64* %P_a
	%r14 = icmp slt i64 %r13, 2
	br i1 %r14, label %L7, label %L8

L7:
	store i64 0, i64* %prime_retval
	%r15 = load i64, i64* %prime_retval
	ret i64 %r15
	br label %L9

L8:
	%r16 = load i64, i64* %P_a
	%r17 = call i64 @isqrt(i64 %r16)
	store i64 %r17, i64* %max
	store i64 2, i64* %divisor
	br label %L9

L9:
	%r18 = load i64, i64* %divisor
	%r19 = load i64, i64* %max
	%r20 = icmp sle i64 %r18, %r19
	br i1 %r20, label %L10, label %L11

L10:
	%r21 = load i64, i64* %P_a
	%r22 = load i64, i64* %P_a
	%r23 = load i64, i64* %divisor
	%r24 = sdiv i64 %r22, %r23
	%r25 = load i64, i64* %divisor
	%r26 = mul i64 %r24, %r25
	%r27 = sub i64 %r21, %r26
	store i64 %r27, i64* %remainder
	br label %L11

L11:
	%r28 = load i64, i64* %remainder
	%r29 = icmp eq i64 %r28, 0
	br i1 %r29, label %L12, label %L13

L12:
	store i64 0, i64* %prime_retval
	%r30 = load i64, i64* %prime_retval
	ret i64 %r30
	br label %L14

L13:
	br label %L14

L14:
	%r31 = load i64, i64* %divisor
	%r32 = add i64 %r31, 1
	store i64 %r32, i64* %divisor
	br label %L9

L15:
	store i64 1, i64* %prime_retval
	%r33 = load i64, i64* %prime_retval
	ret i64 %r33
	br label %L9

L16:
	br label %L17

L17:
	%r34 = load i64, i64* %prime_retval
	ret i64 %r34


}

define i64 @main()
{
L18:
	%main_retval = alloca i64
	%limit = alloca i64
	%a = alloca i64
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %limit)
	store i64 0, i64* %a
	br label %L19

L19:
	%r35 = load i64, i64* %a
	%r36 = load i64, i64* %limit
	%r37 = icmp sle i64 %r35, %r36
	br i1 %r37, label %L20, label %L21

L20:
	br label %L21

L21:
	%r38 = load i64, i64* %a
	%r39 = call i64 @prime(i64 %r38)
	%r40 = icmp eq i64 %r39, 1
	br i1 %r40, label %L22, label %L23

L22:
	%r41 = load i64, i64* %a
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr2, i32 0, i32 0), i64 %r41)
	br label %L24

L23:
	br label %L24

L24:
	%r42 = load i64, i64* %a
	%r43 = add i64 %r42, 1
	store i64 %r43, i64* %a
	br label %L19

L25:
	br label %L26

L26:
	store i64 0, i64* %main_retval
	%r44 = load i64, i64* %main_retval
	ret i64 %r44


}


declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

