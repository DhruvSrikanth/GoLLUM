source_filename = "example8"
target triple = "arm64-apple-darwin22.2.0"



define i64 @something(i64 %x)
{
L0:
	%something_retval = alloca i64
	%P_x = alloca i64
	store i64 %x, i64* %P_x
	%r0 = load i64, i64* %P_x
	%r1 = add i64 %r0, 1
	store i64 %r1, i64* %something_retval
	%r2 = load i64, i64* %something_retval
	ret i64 %r2
	br label %L1

L1:
	%r3 = load i64, i64* %something_retval
	ret i64 %r3


}

define i64 @main()
{
L2:
	%main_retval = alloca i64
	%z = alloca i64
	store i64 1, i64* %r4
	%r5 = call i64 @something(i64 1)
	store i64 %r5, i64* %z
	%r6 = load i64, i64* %z
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.fstr1, i32 0, i32 0), i64 %r6)
	br label %L3

L3:
	store i64 0, i64* %main_retval
	%r7 = load i64, i64* %main_retval
	ret i64 %r7


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [5 x i8] c"%ld\0A\00", align 1

