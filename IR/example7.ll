source_filename = "example7"
target triple = "arm64-apple-darwin22.2.0"



define i64 @fact(i64 %x)
{
L0:
	%_retval = alloca i64
	%P_x = alloca i64
	store i64 %x, i64* %P_x
	br label %L1

L1:
	%r0 = load i64, i64* %P_x
	store i64 1, i64* %r1
	%r2 = icmp sle i64 %r0, %r1
	br i1 %r2, label %L2, label %L3

L2:
	store i64 1, i64* %r3
	store i64 %r3, i64* %_retval
	%r4 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L4

L3:
	%r5 = load i64, i64* %P_x
	%r6 = load i64, i64* %P_x
	store i64 1, i64* %r7
	%r8 = sub i64 %r6, %r7
	%r9 = call i64 @fact(i64 %r8)
	%r10 = mul i64 %r5, %r9
	store i64 %r10, i64* %_retval
	%r11 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L4

L4:
	br label %L5

L5:
	%r12 = load i64, i64* %_retval
	ret i64 %r12


}

define i64 @main()
{
L6:
	%_retval = alloca i64
	%stop = alloca i64
	%factor = alloca i64
	%toStop = alloca i64
	%temp = alloca i64
	%r13 = load i64, i64* %stop
	store i64 0, i64* %r14
	store i64 %r14, i64* %r13
	%r15 = load i64, i64* %factor
	store i64 0, i64* %r16
	store i64 %r16, i64* %r15
	br label %L7

L7:
	store i64 1, i64* %r17
	%r18 = load i64, i64* %stop
	%r19 = xor i64 %r17, %r18
	br i1 %r19, label %L8, label %L9

L8:
	%r20 = load i64, i64* %factor
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r20)
	%r21 = load i64, i64* %temp
	%r22 = load i64, i64* %factor
	%r23 = call i64 @fact(i64 %r22)
	store i64 %r23, i64* %r21
	%r24 = load i64, i64* %temp
	%r25 = load i64, i64* %r24
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.fstr2, i32 0, i32 0), i64 %r25)
	%r26 = load i64, i64* %toStop
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r26)
	br label %L9

L9:
	%r27 = load i64, i64* %toStop
	store i64 0, i64* %r28
	%r29 = icmp eq i64 %r27, %r28
	br i1 %r29, label %L10, label %L11

L10:
	%r30 = load i64, i64* %stop
	store i64 1, i64* %r31
	store i64 %r31, i64* %r30
	br label %L12

L11:
	br label %L12

L12:
	br label %L7

L13:
	br label %L14

L14:
	store i64 0, i64* %_retval
	%r32 = load i64, i64* %r31
	ret i64 %r32


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [3 x i8] c"%ld\00", align 1

