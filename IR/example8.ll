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
	ret %struct.Point2D*%_retval

L3:

L4:
	%r13 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r13, %struct.Point2D** %_retval
	%r14 = load %struct.Point2D*, %struct.Point2D** %_retval
	ret %struct.Point2D*%_retval

L5:


}

define i64 @main()
{
L6:
	%_retval = alloca i64
	%a = alloca i64
	%b = alloca i64
	%pt1 = alloca %struct.Point2D*
	%pt2 = alloca %struct.Point2D*
	%r15 = load i64, i64* %a
	store i64 5, i64* %r16
	store i64 %r16, i64* %r15
	%r17 = load i64, i64* %b
	%r18 = load i64, i64* %a
	store i64 7, i64* %r19
	%r20 = add i64 %r18, %r19
	store i64 3, i64* %r21
	%r22 = mul i64 %r20, %r21
	store i64 %r22, i64* %r17
	%r23 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r24 = call i8* @malloc(i32 16)
	%r25 = bitcast i8* %r24 to %struct.Point2D*
	store %struct.Point2D* %r25, %struct.Point2D** %r23
	%r26 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r27 = getelementptr %struct.Point2D*, %struct.Point2D** %r26, i32 0, i32 0
	%r28 = load i64, i64* %a
	store i64 %r28, i64* %r27
	%r29 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r30 = getelementptr %struct.Point2D*, %struct.Point2D** %r29, i32 0, i32 1
	%r31 = load i64, i64* %b
	store i64 %r31, i64* %r30
	%r32 = load i64, i64* @globalInit
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r32)
	%r33 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r34 = load i64, i64* @globalInit
	%r35 = call %struct.Point2D* @Init(i64 %r34)
	store %struct.Point2D* %r35, %struct.Point2D** %r33
	%r36 = load i64, i64* @globalInit
	%r37 = load i64, i64* r36
	%r38 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r39 = getelementptr %struct.Point2D*, %struct.Point2D** %r38, i32 0, i32 0
	%r40 = load i64, i64* r39
	%r41 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r42 = getelementptr %struct.Point2D*, %struct.Point2D** %r41, i32 0, i32 1
	%r43 = load i64, i64* r42
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr2, i32 0, i32 0), i64 %r37, i64 %r40, i64 %r43)
	%r44 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r45 = bitcast %struct.Point2D* %r44 to i8*
	call void @free(i8* %r45)
	%r47 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r48 = bitcast %struct.Point2D* %r47 to i8*
	call void @free(i8* %r48)

L7:


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [32 x i8] c"offset=%ld\npt2.x=%ld\npt2.y=%ld\n\00", align 1

