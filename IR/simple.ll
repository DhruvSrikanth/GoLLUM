source_filename = "simple"
target triple = "arm64-apple-macosx13.0.0"
%struct.Point2D = type {i64, i64}
@.nilPoint2D = common global %struct.Point2D* null   
@globalInit = common global i64 0

define %struct.Point2D* @Init(i64 %initVal)
{
L2: 
    ; delcare space for the return value
    %_retval = alloca %struct.Point2D*
    ; Parameters allocation (initVal int)
    %P_initVal = alloca i64
    ; var newPt *Point2D; 
    %newpt = alloca %struct.Point2D*
    ; Store the values of the register parameters at their stack memory locations 
    ; P_initVal = initVal 
    store i64 %initVal, i64* %P_initVal 
    ; newPt = nil; 
    %r26 = load %struct.Point2D*, %struct.Point2D** @.nilPoint2D
    store %struct.Point2D* %r26, %struct.Point2D** %newpt
    br label %L3
L3: 
    %r27 = load i64, i64* %P_initVal
	%r28 = icmp sgt i64 %r27, 0
	br i1 %r28, label %L4, label %L5
    
L4: 
    ; newPt = new Point2D; 
    %r29 = call i8* @malloc(i32 16)
	%r30 = bitcast i8* %r29 to %struct.Point2D*
	store %struct.Point2D* %r30, %struct.Point2D** %newpt
    ; newPt.x = initVal;
    %r31 = load i64, i64* %P_initVal 
    %r32 = load %struct.Point2D*, %struct.Point2D** %newpt
	%r33 = getelementptr %struct.Point2D, %struct.Point2D* %r32, i32 0, i32 0
    store i64 %r31, i64* %r33 
    ; newPt.y = initVal;
    %r34 = load i64, i64* %P_initVal 
    %r35 = load %struct.Point2D*, %struct.Point2D** %newpt
	%r36 = getelementptr %struct.Point2D, %struct.Point2D* %r35, i32 0, i32 1
    store i64 %r34, i64* %r36 
    ; return newPt; 
    %r37 = load %struct.Point2D*, %struct.Point2D** %newpt
    store %struct.Point2D* %r37, %struct.Point2D** %_retval
    %r38 = load %struct.Point2D*, %struct.Point2D** %_retval
    ret %struct.Point2D* %r38
L5: 
    ; False branch. Nothing happens here so jump to the if-exit L5
    br label %L6
L6: 
    ; return newPt; 
    %r39 = load %struct.Point2D*, %struct.Point2D** %newpt
    store %struct.Point2D* %r39, %struct.Point2D** %_retval
    %r40 = load %struct.Point2D*, %struct.Point2D** %_retval
    ret %struct.Point2D* %r40

}

define i64 @main()
{
L0: 
    ; delcare space for the return value
    %_retval = alloca i64
    ; var a, b int;
    ; var pt1, pt2 *Point2D; 
    %a = alloca i64      
    %b = alloca i64      
    %pt1 = alloca %struct.Point2D*
    %pt2 = alloca %struct.Point2D*
    ; a = 5;
    store i64 5, i64* %a    
    ; b = (a + 7) * 3; 
    %r0 = load i64, i64* %a 
    %r1 = add i64 7, %r0 
    %r2 = mul i64 %r1, 3
    store i64 %r2, i64* %b
    ; pt1 = new Point2D; 
	%r3 = call i8* @malloc(i32 16)
	%r4 = bitcast i8* %r3 to %struct.Point2D*
	store %struct.Point2D* %r4, %struct.Point2D** %pt1
    ; pt1.x = a; 
    %r6 = load i64, i64* %a     
	%r7 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r8 = getelementptr %struct.Point2D, %struct.Point2D* %r7, i32 0, i32 0
	store i64 %r6, i64* %r8
    ; pt1.y = b; 
    %r9 = load i64, i64* %b
    %r10 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r11 = getelementptr %struct.Point2D, %struct.Point2D* %r10, i32 0, i32 1
	store i64 %r9, i64* %r11      
    ; scan globalInit; 
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* @globalInit)
    ; pt2 = Init(globalInit);
    %r12 = load i64, i64* @globalInit
    %r14 = call %struct.Point2D* @Init(i64 %r12)
    store %struct.Point2D* %r14, %struct.Point2D** %pt2
    ; printf("offset=%d\npt2.x=%d\npt2.y=%d\n",globalInit,pt2.x,pt2.y);
    %r15 = load i64, i64* @globalInit
    %r16 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r17 = getelementptr %struct.Point2D, %struct.Point2D* %r16, i32 0, i32 0
    %r18 = load i64, i64* %r17 
    %r19 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r20 = getelementptr %struct.Point2D, %struct.Point2D* %r19, i32 0, i32 1
    %r21 = load i64, i64* %r20 
    call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr1, i32 0, i32 0), i64 %r15, i64 %r18, i64 %r21)
    ; delete pt1; 
    %r22 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r23 = bitcast %struct.Point2D* %r22 to i8*
	call void @free(i8* %r23)
    ; delete pt2; 
    %r24 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r25 = bitcast %struct.Point2D* %r24 to i8*
	call void @free(i8* %r25)
    br label %L1 
L1: 
    store i64 0, i64* %_retval 
    %r41 = load i64, i64* %_retval 
    ret i64 %r41 
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
@.fstr1 = private unnamed_addr constant [32 x i8] c"offset=%ld\0Apt2.x=%ld\0Apt2.y=%ld\0A\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

