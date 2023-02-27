source_filename = "example8"
target triple = "arm64-apple-darwin22.2.0"

%struct.Point2D = type {i64, i64}

@.nilPoint2D = common global %struct.Point2D* null
@globalInit = common global i64 0

define %struct.Point2D* @Init(i64 %initVal)
{
L0:
	%_retval = alloca %struct.Point2D*
	%newPt = alloca %struct.Point2D*
	%P_initVal = alloca i64
	store i64 %initVal, i64* %P_initVal
	%r0 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r1 = load %struct.Point2D*, %struct.Point2D** @.nilPoint2D
	store %struct.Point2D* %r1, %struct.Point2D** %r0

L1:

L2:
	%r2 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r3 = call i8* @malloc(i32 16)
	%r4 = bitcast i8* %r3 to %struct.Point2D*
	store %struct.Point2D* %r4, %struct.Point2D** %r2
	%r5 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r6 = getelementptr %struct.Point2D*, %struct.Point2D** %r5, i32 0, i32 0
	%r7 = load i64, i64* %P_initVal
	store i64 %r7, i64* %r6
	%r8 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r9 = getelementptr %struct.Point2D*, %struct.Point2D** %r8, i32 0, i32 1
	%r10 = load i64, i64* %P_initVal
	store i64 %r10, i64* %r9
	%r11 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r11, %struct.Point2D** %_retval
	%r12 = load %struct.Point2D*, %struct.Point2D** %_retval
	ret %struct.Point2D* %_retval

L3:

L4:
	%r13 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r13, %struct.Point2D** %_retval
	%r14 = load %struct.Point2D*, %struct.Point2D** %_retval
	ret %struct.Point2D* %_retval
	br label %L5

L5:
	%r15 = load %struct.Point2D, %struct.Point2D* %_retval
	ret %struct.Point2D %r15


}

define i64 @main()
{
L6:
	%_retval = alloca i64
	%a = alloca i64
	%b = alloca i64
	%pt1 = alloca %struct.Point2D*
	%pt2 = alloca %struct.Point2D*
	%r16 = load i64, i64* %a
	store i64 5, i64* %r17
	store i64 %r17, i64* %r16
	%r18 = load i64, i64* %b
	%r19 = load i64, i64* %a
	store i64 7, i64* %r20
	%r21 = add i64 %r19, %r20
	store i64 3, i64* %r22
	%r23 = mul i64 %r21, %r22
	store i64 %r23, i64* %r18
	%r24 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r25 = call i8* @malloc(i32 16)
	%r26 = bitcast i8* %r25 to %struct.Point2D*
	store %struct.Point2D* %r26, %struct.Point2D** %r24
	%r27 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r28 = getelementptr %struct.Point2D*, %struct.Point2D** %r27, i32 0, i32 0
	%r29 = load i64, i64* %a
	store i64 %r29, i64* %r28
	%r30 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r31 = getelementptr %struct.Point2D*, %struct.Point2D** %r30, i32 0, i32 1
	%r32 = load i64, i64* %b
	store i64 %r32, i64* %r31
	%r33 = load i64, i64* @globalInit
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r33)
	%r34 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r35 = load i64, i64* @globalInit
	%r36 = call %struct.Point2D* @Init(i64 %r35)
	store %struct.Point2D* %r36, %struct.Point2D** %r34
	%r37 = load i64, i64* @globalInit
	%r38 = load i64, i64* r37
	%r39 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r40 = getelementptr %struct.Point2D*, %struct.Point2D** %r39, i32 0, i32 0
	%r41 = load i64, i64* r40
	%r42 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r43 = getelementptr %struct.Point2D*, %struct.Point2D** %r42, i32 0, i32 1
	%r44 = load i64, i64* r43
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr2, i32 0, i32 0), i64 %r38, i64 %r41, i64 %r44)
	%r45 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r46 = bitcast %struct.Point2D* %r45 to i8*
	call void @free(i8* %r46)
	%r48 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r49 = bitcast %struct.Point2D* %r48 to i8*
	call void @free(i8* %r49)
	br label %L7

L7:
	store i64 0, i64* %_retval
	%r51 = load i64, i64* r50
	ret i64 %r51


}


declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [32 x i8] c"offset=%ld\npt2.x=%ld\npt2.y=%ld\n\00", align 1

