source_filename = "example1"
target triple = "arm64-apple-darwin22.2.0"

%struct.nums = type {i64, i64}

@.nilnums = common global %struct.nums* null

define i64 @fib1(i64 %f1)
{
L0:
	%fib1_retval = alloca i64
	%P_f1 = alloca i64
	store i64 %f1, i64* %P_f1
	br label %L1

L1:
	%r0 = load i64, i64* %P_f1
	%r1 = icmp slt i64 %r0, 2
	br i1 %r1, label %L2, label %L3

L2:
	%r2 = load i64, i64* %P_f1
	store i64 %r2, i64* %fib1_retval
	%r3 = load i64, i64* %fib1_retval
	ret i64 %r3

L3:
	%r4 = load i64, i64* %P_f1
	%r5 = sub i64 %r4, 1
	%r6 = call i64 @fib1(i64 %r5)
	%r7 = load i64, i64* %P_f1
	%r8 = sub i64 %r7, 2
	%r9 = call i64 @fib1(i64 %r8)
	%r10 = add i64 %r6, %r9
	store i64 %r10, i64* %fib1_retval
	%r11 = load i64, i64* %fib1_retval
	ret i64 %r11

L4:
	br label %L5

L5:
	%r12 = load i64, i64* %fib1_retval
	ret i64 %r12


}

define i64 @fib2(i64 %f2)
{
L6:
	%fib2_retval = alloca i64
	%fir = alloca i64
	%sec = alloca i64
	%temp = alloca i64
	%P_f2 = alloca i64
	store i64 %f2, i64* %P_f2
	store i64 0, i64* %fir
	store i64 1, i64* %sec
	br label %L7

L7:
	%r13 = load i64, i64* %P_f2
	%r14 = icmp ne i64 %r13, 0
	br i1 %r14, label %L8, label %L9

L8:
	%r15 = load i64, i64* %P_f2
	%r16 = sub i64 %r15, 1
	store i64 %r16, i64* %P_f2
	%r17 = load i64, i64* %fir
	%r18 = load i64, i64* %sec
	%r19 = add i64 %r17, %r18
	store i64 %r19, i64* %temp
	%r20 = load i64, i64* %sec
	store i64 %r20, i64* %fir
	%r21 = load i64, i64* %temp
	store i64 %r21, i64* %sec
	br label %L7

L9:
	%r22 = load i64, i64* %fir
	store i64 %r22, i64* %fib2_retval
	%r23 = load i64, i64* %fib2_retval
	ret i64 %r23


}

define i64 @main()
{
L10:
	%main_retval = alloca i64
	%x = alloca %struct.nums*
	%c = alloca i64
	%d = alloca i64
	%temp = alloca i64
	%r24 = call i8* @malloc(i32 16)
	%r25 = bitcast i8* %r24 to %struct.nums*
	store %struct.nums* %r25, %struct.nums** %x
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %temp)
	%r26 = load %struct.nums*, %struct.nums** %x
	%r27 = getelementptr %struct.nums, %struct.nums* %r26, i32 0, i32 0
	%r28 = load i64, i64* %temp
	store i64 %r28, i64* %r27
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %temp)
	%r29 = load %struct.nums*, %struct.nums** %x
	%r30 = getelementptr %struct.nums, %struct.nums* %r29, i32 0, i32 1
	%r31 = load i64, i64* %temp
	store i64 %r31, i64* %r30
	%r32 = load %struct.nums*, %struct.nums** %x
	%r33 = getelementptr %struct.nums, %struct.nums* %r32, i32 0, i32 0
	%r34 = load i64, i64* %r33
	%r35 = call i64 @fib1(i64 %r34)
	store i64 %r35, i64* %c
	%r36 = load %struct.nums*, %struct.nums** %x
	%r37 = getelementptr %struct.nums, %struct.nums* %r36, i32 0, i32 1
	%r38 = load i64, i64* %r37
	%r39 = call i64 @fib2(i64 %r38)
	store i64 %r39, i64* %d
	%r40 = load %struct.nums*, %struct.nums** %x
	%r41 = bitcast %struct.nums* %r40 to i8*
	call void @free(i8* %r41)
	%r43 = load i64, i64* %c
	%r44 = load i64, i64* %d
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([12 x i8], [12 x i8]* @.fstr2, i32 0, i32 0), i64 %r43, i64 %r44)
	br label %L11

L11:
	store i64 0, i64* %main_retval
	%r45 = load i64, i64* %main_retval
	ret i64 %r45


}


declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)
declare void @free(i8*)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [12 x i8] c"c=%ld\0Ad=%ld\00", align 1

