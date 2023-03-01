source_filename = "example8"
target triple = "arm64-apple-darwin22.2.0"

%struct.Point2D = type {i64, i64}

@.nilPoint2D = common global %struct.Point2D* null
@globalInit = common global i64 0

define %struct.Point2D* @Init(i64 %initVal)
{
L0:
	%Init_retval = alloca %struct.Point2D*
	%newPt = alloca %struct.Point2D*
	%P_initVal = alloca i64
	store i64 %initVal, i64* %P_initVal
	%r0 = load %struct.Point2D*, %struct.Point2D** @.nilPoint2D
	store %struct.Point2D* %r0, %struct.Point2D** %newPt
	br label %L1

L1:
	%r1 = load i64, i64* %P_initVal
	%r2 = icmp sgt i64 %r1, 0
	br i1 %r2, label %L2, label %L3

L2:
	%r3 = call i8* @malloc(i32 16)
	%r4 = bitcast i8* %r3 to %struct.Point2D*
	store %struct.Point2D* %r4, %struct.Point2D** %newPt
	%r5 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r6 = getelementptr %struct.Point2D, %struct.Point2D* %r5, i32 0, i32 0
	%r7 = load i64, i64* %P_initVal
	store i64 %r7, i64* %r6
	%r8 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r9 = getelementptr %struct.Point2D, %struct.Point2D* %r8, i32 0, i32 1
	%r10 = load i64, i64* %P_initVal
	store i64 %r10, i64* %r9
	%r11 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r11, %struct.Point2D** %Init_retval
	%r12 = load %struct.Point2D*, %struct.Point2D** %Init_retval
	ret %struct.Point2D* %r12
	br label %L4

L3:
	br label %L4

L4:
	%r13 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r13, %struct.Point2D** %Init_retval
	%r14 = load %struct.Point2D*, %struct.Point2D** %Init_retval
	ret %struct.Point2D* %r14
	br label %L5

L5:
	%r15 = load %struct.Point2D*, %struct.Point2D** %Init_retval
	ret %struct.Point2D* %r15


}

define i64 @main()
{
L6:
	%main_retval = alloca i64
	%a = alloca i64
	%b = alloca i64
	%pt1 = alloca %struct.Point2D*
	%pt2 = alloca %struct.Point2D*
	store i64 5, i64* %a
	%r16 = load i64, i64* %a
	%r17 = add i64 %r16, 7
	%r18 = mul i64 %r17, 3
	store i64 %r18, i64* %b
	%r19 = call i8* @malloc(i32 16)
	%r20 = bitcast i8* %r19 to %struct.Point2D*
	store %struct.Point2D* %r20, %struct.Point2D** %pt1
	%r21 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r22 = getelementptr %struct.Point2D, %struct.Point2D* %r21, i32 0, i32 0
	%r23 = load i64, i64* %a
	store i64 %r23, i64* %r22
	%r24 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r25 = getelementptr %struct.Point2D, %struct.Point2D* %r24, i32 0, i32 1
	%r26 = load i64, i64* %b
	store i64 %r26, i64* %r25
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* @globalInit)
	%r27 = load i64, i64* @globalInit
	%r28 = call %struct.Point2D* @Init(i64 %r27)
	store %struct.Point2D* %r28, %struct.Point2D** %pt2
	%r29 = load i64, i64* @globalInit
	%r30 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r31 = getelementptr %struct.Point2D, %struct.Point2D* %r30, i32 0, i32 0
	%r32 = load i64, i64* %r31
	%r33 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r34 = getelementptr %struct.Point2D, %struct.Point2D* %r33, i32 0, i32 1
	%r35 = load i64, i64* %r34
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr2, i32 0, i32 0), i64 %r29, i64 %r32, i64 %r35)
	%r36 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r37 = bitcast %struct.Point2D* %r36 to i8*
	call void @free(i8* %r37)
	%r39 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r40 = bitcast %struct.Point2D* %r39 to i8*
	call void @free(i8* %r40)
	br label %L7

L7:
	store i64 0, i64* %main_retval
	%r42 = load i64, i64* %main_retval
	ret i64 %r42


}


declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [32 x i8] c"offset=%ld\0Apt2.x=%ld\0Apt2.y=%ld\0A\00", align 1

