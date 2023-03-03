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


}

define i64 @prime(i64 %a)
{
L4:
	%prime_retval = alloca i64
	%max = alloca i64
	%divisor = alloca i64
	%remainder = alloca i64
	%P_a = alloca i64
	store i64 %a, i64* %P_a
	br label %L5

L5:
	%r12 = load i64, i64* %P_a
	%r13 = icmp slt i64 %r12, 2
	br i1 %r13, label %L6, label %L7

L6:
	store i64 0, i64* %prime_retval
	%r14 = load i64, i64* %prime_retval
	ret i64 %r14

L7:
	%r15 = load i64, i64* %P_a
	%r16 = call i64 @isqrt(i64 %r15)
	store i64 %r16, i64* %max
	store i64 2, i64* %divisor
	br label %L8

L8:
	%r17 = load i64, i64* %divisor
	%r18 = load i64, i64* %max
	%r19 = icmp sle i64 %r17, %r18
	br i1 %r19, label %L9, label %L10

L9:
	%r20 = load i64, i64* %P_a
	%r21 = load i64, i64* %divisor
	%r22 = sdiv i64 %r20, %r21
	%r23 = load i64, i64* %divisor
	%r24 = mul i64 %r22, %r23
	%r25 = load i64, i64* %P_a
	%r26 = sub i64 %r25, %r24
	store i64 %r26, i64* %remainder
	br label %L10

L10:
	%r27 = load i64, i64* %remainder
	%r28 = icmp eq i64 %r27, 0
	br i1 %r28, label %L11, label %L12

L11:
	store i64 0, i64* %prime_retval
	%r29 = load i64, i64* %prime_retval
	ret i64 %r29

L12:
	br label %L13

L13:
	%r30 = load i64, i64* %divisor
	%r31 = add i64 %r30, 1
	store i64 %r31, i64* %divisor
	br label %L8

L14:
	store i64 1, i64* %prime_retval
	%r32 = load i64, i64* %prime_retval
	ret i64 %r32

L15:
	br label %L16

L16:
	%r33 = load i64, i64* %prime_retval
	ret i64 %r33


}

define i64 @main()
{
L17:
	%main_retval = alloca i64
	%limit = alloca i64
	%a = alloca i64
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %limit)
	store i64 0, i64* %a
	br label %L18

L18:
	%r34 = load i64, i64* %a
	%r35 = load i64, i64* %limit
	%r36 = icmp sle i64 %r34, %r35
	br i1 %r36, label %L19, label %L20

L19:
	br label %L20

L20:
	%r37 = load i64, i64* %a
	%r38 = call i64 @prime(i64 %r37)
	%r39 = icmp eq i64 %r38, 1
	br i1 %r39, label %L21, label %L22

L21:
	%r40 = load i64, i64* %a
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr2, i32 0, i32 0), i64 %r40)
	br label %L23

L22:
	br label %L23

L23:
	%r41 = load i64, i64* %a
	%r42 = add i64 %r41, 1
	store i64 %r42, i64* %a
	br label %L18

L24:
	br label %L25

L25:
	store i64 0, i64* %main_retval
	%r43 = load i64, i64* %main_retval
	ret i64 %r43


}


declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

