source_filename = "example9"
target triple = "arm64-apple-darwin22.2.0"



define i64 @main()
{
L0:
	%_retval = alloca i64
	%a = alloca i64
	%r0 = load i64, i64* %a
	store i64 0, i64* %r1
	store i64 %r1, i64* %r0
	br label %L1

L1:
	%r2 = load i64, i64* %a
	store i64 10, i64* %r3
	%r4 = icmp slt i64 %r2, %r3
	br i1 %r4, label %L2, label %L3

L2:
	%r5 = load i64, i64* %a
	%r6 = load i64, i64* %a
	store i64 1, i64* %r7
	%r8 = add i64 %r6, %r7
	store i64 %r8, i64* %r5
	br label %L1

L3:
	%r9 = load i64, i64* %a
	%r10 = load i64, i64* %r9
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.fstr1, i32 0, i32 0), i64 %r10)
	br label %L4

L4:
	store i64 0, i64* %_retval
	%r11 = load i64, i64* %r10
	ret i64 %r11


}


declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)

@.fstr1 = private unnamed_addr constant [7 x i8] c"a=%ld\n\00", align 1

