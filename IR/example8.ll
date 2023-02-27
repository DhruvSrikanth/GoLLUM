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
	br label %L1

L1:
	%r2 = load i64, i64* %P_initVal
	store i64 0, i64* %r3
	%r4 = icmp sgt i64 %r2, %r3
	br i1 %r4, label %L2, label %L3

L2:
	%r5 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r6 = call i8* @malloc(i32 16)
	%r7 = bitcast i8* %r6 to %struct.Point2D*
	store %struct.Point2D* %r7, %struct.Point2D** %r5
	%r8 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r9 = getelementptr %struct.Point2D*, %struct.Point2D** %r8, i32 0, i32 0
	%r10 = load i64, i64* %P_initVal
	store i64 %r10, i64* %r9
	%r11 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r12 = getelementptr %struct.Point2D*, %struct.Point2D** %r11, i32 0, i32 1
	%r13 = load i64, i64* %P_initVal
	store i64 %r13, i64* %r12
	%r14 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r14, %struct.Point2D** %_retval
	%r15 = load %struct.Point2D*, %struct.Point2D** %_retval
	ret %struct.Point2D* %_retval
	br label %L4

L3:
	br label %L4

L4:
	%r16 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r16, %struct.Point2D** %_retval
	%r17 = load %struct.Point2D*, %struct.Point2D** %_retval
	ret %struct.Point2D* %_retval
	br label %L5

L5:
	%r18 = load %struct.Point2D, %struct.Point2D* %_retval
	ret %struct.Point2D %r18


}

define i64 @main()
{
L6:
	%_retval = alloca i64
	%a = alloca i64
	%b = alloca i64
	%pt1 = alloca %struct.Point2D*
	%pt2 = alloca %struct.Point2D*
	%r19 = load i64, i64* %a
	store i64 5, i64* %r20
	store i64 %r20, i64* %r19
	%r21 = load i64, i64* %b
	%r22 = load i64, i64* %a
	store i64 7, i64* %r23
	%r24 = add i64 %r22, %r23
	store i64 3, i64* %r25
	%r26 = mul i64 %r24, %r25
	store i64 %r26, i64* %r21
	%r27 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r28 = call i8* @malloc(i32 16)
	%r29 = bitcast i8* %r28 to %struct.Point2D*
	store %struct.Point2D* %r29, %struct.Point2D** %r27
	%r30 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r31 = getelementptr %struct.Point2D*, %struct.Point2D** %r30, i32 0, i32 0
	%r32 = load i64, i64* %a
	store i64 %r32, i64* %r31
	%r33 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r34 = getelementptr %struct.Point2D*, %struct.Point2D** %r33, i32 0, i32 1
	%r35 = load i64, i64* %b
	store i64 %r35, i64* %r34
	%r36 = load i64, i64* @globalInit
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r36)
	%r37 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r38 = load i64, i64* @globalInit
	%r39 = call %struct.Point2D* @Init(i64 %r38)
	store %struct.Point2D* %r39, %struct.Point2D** %r37
	%r40 = load i64, i64* @globalInit
	%r41 = load i64, i64* r40
	%r42 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r43 = getelementptr %struct.Point2D*, %struct.Point2D** %r42, i32 0, i32 0
	%r44 = load i64, i64* r43
	%r45 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r46 = getelementptr %struct.Point2D*, %struct.Point2D** %r45, i32 0, i32 1
	%r47 = load i64, i64* r46
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr2, i32 0, i32 0), i64 %r41, i64 %r44, i64 %r47)
	%r48 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r49 = bitcast %struct.Point2D* %r48 to i8*
	call void @free(i8* %r49)
	%r51 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r52 = bitcast %struct.Point2D* %r51 to i8*
	call void @free(i8* %r52)
	br label %L7

L7:
	store i64 0, i64* %_retval
	%r54 = load i64, i64* r53
	ret i64 %r54


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [32 x i8] c"offset=%ld\npt2.x=%ld\npt2.y=%ld\n\00", align 1

