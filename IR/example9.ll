source_filename = "example9"
target triple = "arm64-apple-darwin22.2.0"



define i64 @main()
{
L0:
	%main_retval = alloca i64
	%a = alloca i64
	store i64 0, i64* %a
	br label %L1

L1:
	%r0 = load i64, i64* %a
	%r1 = icmp slt i64 %r0, 10
	br i1 %r1, label %L2, label %L3

L2:
	%r2 = load i64, i64* %a
	%r3 = add i64 %r2, 1
	store i64 %r3, i64* %a
	br label %L1

L3:
	%r4 = load i64, i64* %a
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.fstr1, i32 0, i32 0), i64 %r4)
	br label %L4

L4:
	store i64 0, i64* %main_retval
	%r5 = load i64, i64* %main_retval
	ret i64 %r5


}


declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)

@.fstr1 = private unnamed_addr constant [7 x i8] c"a=%ld\0A\00", align 1

