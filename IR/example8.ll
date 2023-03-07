source_filename = "example8"
target triple = "x86_64-linux-gnu"



define i64 @main()
{
L0:
	%main_retval = alloca i64
	%a = alloca i64
	%b = alloca i64
	%c = alloca i64
	store i64 129, i64* %a
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %b)
	%r0 = load i64, i64* %a
	%r1 = load i64, i64* %b
	%r2 = add i64 %r0, %r1
	store i64 %r2, i64* %c
	%r3 = load i64, i64* %a
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.fstr2, i32 0, i32 0), i64 %r3)
	%r4 = load i64, i64* %b
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.fstr3, i32 0, i32 0), i64 %r4)
	%r5 = load i64, i64* %c
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.fstr4, i32 0, i32 0), i64 %r5)
	br label %L1

L1:
	store i64 0, i64* %main_retval
	%r6 = load i64, i64* %main_retval
	ret i64 %r6


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [5 x i8] c"%ld\0A\00", align 1
@.fstr3 = private unnamed_addr constant [5 x i8] c"%ld\0A\00", align 1
@.fstr4 = private unnamed_addr constant [5 x i8] c"%ld\0A\00", align 1

