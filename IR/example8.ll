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

L3:
	br label %L4

L4:
	%r13 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r13, %struct.Point2D** %Init_retval
	%r14 = load %struct.Point2D*, %struct.Point2D** %Init_retval
	ret %struct.Point2D* %r14


}

define i64 @main()
{
L5:
	%main_retval = alloca i64
	%a = alloca i64
	%b = alloca i64
	%pt1 = alloca %struct.Point2D*
	%pt2 = alloca %struct.Point2D*
	store i64 5, i64* %a
	%r15 = load i64, i64* %a
	%r16 = add i64 %r15, 7
	%r17 = mul i64 %r16, 3
	store i64 %r17, i64* %b
	%r18 = call i8* @malloc(i32 16)
	%r19 = bitcast i8* %r18 to %struct.Point2D*
	store %struct.Point2D* %r19, %struct.Point2D** %pt1
	%r20 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r21 = getelementptr %struct.Point2D, %struct.Point2D* %r20, i32 0, i32 0
	%r22 = load i64, i64* %a
	store i64 %r22, i64* %r21
	%r23 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r24 = getelementptr %struct.Point2D, %struct.Point2D* %r23, i32 0, i32 1
	%r25 = load i64, i64* %b
	store i64 %r25, i64* %r24
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* @globalInit)
	%r26 = load i64, i64* @globalInit
	%r27 = call %struct.Point2D* @Init(i64 %r26)
	store %struct.Point2D* %r27, %struct.Point2D** %pt2
	%r28 = load i64, i64* @globalInit
	%r29 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r30 = getelementptr %struct.Point2D, %struct.Point2D* %r29, i32 0, i32 0
	%r31 = load i64, i64* %r30
	%r32 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r33 = getelementptr %struct.Point2D, %struct.Point2D* %r32, i32 0, i32 1
	%r34 = load i64, i64* %r33
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr2, i32 0, i32 0), i64 %r28, i64 %r31, i64 %r34)
	%r35 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r36 = bitcast %struct.Point2D* %r35 to i8*
	call void @free(i8* %r36)
	%r38 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r39 = bitcast %struct.Point2D* %r38 to i8*
	call void @free(i8* %r39)
	br label %L6

L6:
	store i64 0, i64* %main_retval
	%r41 = load i64, i64* %main_retval
	ret i64 %r41


}


declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [32 x i8] c"offset=%ld\0Apt2.x=%ld\0Apt2.y=%ld\0A\00", align 1

