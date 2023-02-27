source_filename = "example1"
target triple = "arm64-apple-darwin22.2.0"

%struct.nums = type {i64, i64}

@.nilnums = common global %struct.nums* null

define i64 @fib1(i64 %f1)
{
L0:
	%_retval = alloca i64
	%P_f1 = alloca i64
	store i64 %f1, i64* %P_f1
	br label %L1

L1:
	%r0 = load i64, i64* %P_f1
	store i64 2, i64* %r1
	%r2 = icmp slt i64 %r0, %r1
	br i1 %r2, label %L2, label %L3

L2:
	%r3 = load i64, i64* %P_f1
	store i64 %r3, i64* %_retval
	%r4 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L4

L3:
	%r5 = load i64, i64* %P_f1
	store i64 1, i64* %r6
	%r7 = sub i64 %r5, %r6
	%r8 = call i64 @fib1(i64 %r7)
	%r9 = load i64, i64* %P_f1
	store i64 2, i64* %r10
	%r11 = sub i64 %r9, %r10
	%r12 = call i64 @fib1(i64 %r11)
	%r13 = add i64 %r8, %r12
	store i64 %r13, i64* %_retval
	%r14 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L4

L4:
	br label %L5

L5:
	%r15 = load i64, i64* %_retval
	ret i64 %r15


}

define i64 @fib2(i64 %f2)
{
L6:
	%_retval = alloca i64
	%fir = alloca i64
	%sec = alloca i64
	%temp = alloca i64
	%P_f2 = alloca i64
	store i64 %f2, i64* %P_f2
	%r16 = load i64, i64* %fir
	store i64 0, i64* %r17
	store i64 %r17, i64* %r16
	%r18 = load i64, i64* %sec
	store i64 1, i64* %r19
	store i64 %r19, i64* %r18
	br label %L7

L7:
	%r20 = load i64, i64* %P_f2
	store i64 0, i64* %r21
	%r22 = icmp ne i64 %r20, %r21
	br i1 %r22, label %L8, label %L9

L8:
	%r23 = load i64, i64* %P_f2
	%r24 = load i64, i64* %P_f2
	store i64 1, i64* %r25
	%r26 = sub i64 %r24, %r25
	store i64 %r26, i64* %r23
	%r27 = load i64, i64* %temp
	%r28 = load i64, i64* %fir
	%r29 = load i64, i64* %sec
	%r30 = add i64 %r28, %r29
	store i64 %r30, i64* %r27
	%r31 = load i64, i64* %fir
	%r32 = load i64, i64* %sec
	store i64 %r32, i64* %r31
	%r33 = load i64, i64* %sec
	%r34 = load i64, i64* %temp
	store i64 %r34, i64* %r33
	br label %L7

L9:
	%r35 = load i64, i64* %fir
	store i64 %r35, i64* %_retval
	%r36 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L10

L10:
	%r37 = load i64, i64* %_retval
	ret i64 %r37


}

define i64 @main()
{
L11:
	%_retval = alloca i64
	%x = alloca %struct.nums*
	%c = alloca i64
	%d = alloca i64
	%temp = alloca i64
	%r38 = load %struct.nums*, %struct.nums** %x
	%r39 = call i8* @malloc(i32 16)
	%r40 = bitcast i8* %r39 to %struct.nums*
	store %struct.nums* %r40, %struct.nums** %r38
	%r41 = load i64, i64* %temp
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r41)
	%r42 = load %struct.nums*, %struct.nums** %x
	%r43 = getelementptr %struct.nums*, %struct.nums** %r42, i32 0, i32 0
	%r44 = load i64, i64* %temp
	store i64 %r44, i64* %r43
	%r45 = load i64, i64* %temp
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r45)
	%r46 = load %struct.nums*, %struct.nums** %x
	%r47 = getelementptr %struct.nums*, %struct.nums** %r46, i32 0, i32 1
	%r48 = load i64, i64* %temp
	store i64 %r48, i64* %r47
	%r49 = load i64, i64* %c
	%r50 = load %struct.nums*, %struct.nums** %x
	%r51 = getelementptr %struct.nums*, %struct.nums** %r50, i32 0, i32 0
	%r52 = call i64 @fib1(i64 %r51)
	store i64 %r52, i64* %r49
	%r53 = load i64, i64* %d
	%r54 = load %struct.nums*, %struct.nums** %x
	%r55 = getelementptr %struct.nums*, %struct.nums** %r54, i32 0, i32 1
	%r56 = call i64 @fib2(i64 %r55)
	store i64 %r56, i64* %r53
	%r57 = load %struct.nums*, %struct.nums** %x
	%r58 = bitcast %struct.nums* %r57 to i8*
	call void @free(i8* %r58)
	%r60 = load i64, i64* %c
	%r61 = load i64, i64* %r60
	%r62 = load i64, i64* %d
	%r63 = load i64, i64* %r62
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.fstr2, i32 0, i32 0), i64 %r61, i64 %r63)
	br label %L12

L12:
	store i64 0, i64* %_retval
	%r64 = load i64, i64* %r63
	ret i64 %r64


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [11 x i8] c"c=%ld\nd=%ld\00", align 1

