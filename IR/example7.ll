source_filename = "example7"
target triple = "arm64-apple-darwin22.2.0"



define i64 @fact(i64 %x)
{
L0:
	%fact_retval = alloca i64
	%P_x = alloca i64
	store i64 %x, i64* %P_x
	br label %L1

L1:
	%r0 = load i64, i64* %P_x
	%r1 = icmp sle i64 %r0, 1
	br i1 %r1, label %L2, label %L3

L2:
	store i64 1, i64* %fact_retval
	%r2 = load i64, i64* %fact_retval
	ret i64 %r2

L3:
	%r3 = load i64, i64* %P_x
	%r4 = sub i64 %r3, 1
	%r5 = call i64 @fact(i64 %r4)
	%r6 = load i64, i64* %P_x
	%r7 = mul i64 %r6, %r5
	store i64 %r7, i64* %fact_retval
	%r8 = load i64, i64* %fact_retval
	ret i64 %r8

L4:
	br label %L5

L5:
	%r9 = load i64, i64* %fact_retval
	ret i64 %r9


}

define i64 @main()
{
L6:
	%main_retval = alloca i64
	%stop = alloca i64
	%factor = alloca i64
	%toStop = alloca i64
	%temp = alloca i64
	store i64 0, i64* %stop
	store i64 0, i64* %factor
	br label %L7

L7:
	%r10 = load i64, i64* %stop
	%r11 = xor i64 1, %r10
	%r12 = icmp eq i64 %r11, 1
	br i1 %r12, label %L8, label %L9

L8:
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %factor)
	%r13 = load i64, i64* %factor
	%r14 = call i64 @fact(i64 %r13)
	store i64 %r14, i64* %temp
	%r15 = load i64, i64* %temp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr2, i32 0, i32 0), i64 %r15)
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %toStop)
	br label %L9

L9:
	%r16 = load i64, i64* %toStop
	%r17 = icmp eq i64 %r16, 0
	br i1 %r17, label %L10, label %L11

L10:
	store i64 1, i64* %stop
	br label %L12

L11:
	br label %L12

L12:
	br label %L7

L13:
	br label %L14

L14:
	store i64 0, i64* %main_retval
	%r18 = load i64, i64* %main_retval
	ret i64 %r18


}


declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

