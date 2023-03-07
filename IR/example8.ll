source_filename = "example8"
target triple = "x86_64-linux-gnu"

%struct.nums = type {i64, i64}

@.nilnums = common global %struct.nums* null

define i64 @main()
{
L0:
	%main_retval = alloca i64
	%pair = alloca %struct.nums*
	%a = alloca i64
	%b = alloca i64
	%r0 = call i8* @malloc(i32 16)
	%r1 = bitcast i8* %r0 to %struct.nums*
	store %struct.nums* %r1, %struct.nums** %pair
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %a)
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %b)
	%r2 = load %struct.nums*, %struct.nums** %pair
	%r3 = getelementptr %struct.nums, %struct.nums* %r2, i32 0, i32 0
	%r4 = load i64, i64* %a
	store i64 %r4, i64* %r3
	%r5 = load %struct.nums*, %struct.nums** %pair
	%r6 = getelementptr %struct.nums, %struct.nums* %r5, i32 0, i32 1
	%r7 = load i64, i64* %b
	store i64 %r7, i64* %r6
	%r8 = load %struct.nums*, %struct.nums** %pair
	%r9 = bitcast %struct.nums* %r8 to i8*
	call void @free(i8* %r9)
	br label %L1

L1:
	store i64 0, i64* %main_retval
	%r10 = load i64, i64* %main_retval
	ret i64 %r10


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

