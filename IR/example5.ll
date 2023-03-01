source_filename = "example5"
target triple = "arm64-apple-darwin22.2.0"

%struct.simple = type {i64}
%struct.foo = type {i64, i64, %struct.simple*}

@.nilsimple = common global %struct.simple* null
@.nilfoo = common global %struct.foo* null
@globalfoo = common global %struct.foo* null
@unusedGlobal = common global %struct.foo* null

define i64 @tailrecursive(i64 %num)
{
L0:
	%tailrecursive_retval = alloca i64
	%unused = alloca %struct.foo*
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	br label %L1

L1:
	%r0 = load i64, i64* %P_num
	%r1 = icmp sle i64 %r0, 0
	br i1 %r1, label %L2, label %L3

L2:
	store i64 0, i64* %tailrecursive_retval
	%r2 = load i64, i64* %tailrecursive_retval
	ret i64 %r2
	br label %L4

L3:
	br label %L4

L4:
	%r3 = call i8* @malloc(i32 24)
	%r4 = bitcast i8* %r3 to %struct.foo*
	store %struct.foo* %r4, %struct.foo** %unused
	%r5 = load %struct.foo*, %struct.foo** %unused
	store %struct.foo* %r5, %struct.foo** @unusedGlobal
	%r6 = load i64, i64* %P_num
	%r7 = sub i64 %r6, 1
	%r8 = call i64 @tailrecursive(i64 %r7)
	br label %L5

L5:
	store i64 0, i64* %tailrecursive_retval
	%r9 = load i64, i64* %tailrecursive_retval
	ret i64 %r9


}

define i64 @add(i64 %x, i64 %y)
{
L6:
	%add_retval = alloca i64
	%P_x = alloca i64
	store i64 %x, i64* %P_x
	%P_y = alloca i64
	store i64 %y, i64* %P_y
	%r10 = load i64, i64* %P_x
	%r11 = load i64, i64* %P_y
	%r12 = add i64 %r10, %r11
	store i64 %r12, i64* %add_retval
	%r13 = load i64, i64* %add_retval
	ret i64 %r13
	br label %L7

L7:
	%r14 = load i64, i64* %add_retval
	ret i64 %r14


}

define i64 @domath(i64 %num)
{
L8:
	%domath_retval = alloca i64
	%math1 = alloca %struct.foo*
	%math2 = alloca %struct.foo*
	%tmp = alloca i64
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	%r15 = call i8* @malloc(i32 24)
	%r16 = bitcast i8* %r15 to %struct.foo*
	store %struct.foo* %r16, %struct.foo** %math1
	%r17 = load %struct.foo*, %struct.foo** %math1
	%r18 = getelementptr %struct.foo, %struct.foo* %r17, i32 0, i32 2
	%r19 = call i8* @malloc(i32 8)
	%r20 = bitcast i8* %r19 to %struct.simple*
	store %struct.simple* %r20, %struct.simple** %r18
	%r21 = call i8* @malloc(i32 24)
	%r22 = bitcast i8* %r21 to %struct.foo*
	store %struct.foo* %r22, %struct.foo** %math2
	%r23 = load %struct.foo*, %struct.foo** %math2
	%r24 = getelementptr %struct.foo, %struct.foo* %r23, i32 0, i32 2
	%r25 = call i8* @malloc(i32 8)
	%r26 = bitcast i8* %r25 to %struct.simple*
	store %struct.simple* %r26, %struct.simple** %r24
	%r27 = load %struct.foo*, %struct.foo** %math1
	%r28 = getelementptr %struct.foo, %struct.foo* %r27, i32 0, i32 0
	%r29 = load i64, i64* %P_num
	store i64 %r29, i64* %r28
	%r30 = load %struct.foo*, %struct.foo** %math2
	%r31 = getelementptr %struct.foo, %struct.foo* %r30, i32 0, i32 0
	store i64 3, i64* %r31
	%r32 = load %struct.foo*, %struct.foo** %math1
	%r33 = getelementptr %struct.foo, %struct.foo* %r32, i32 0, i32 2
	%r34 = load %struct.simple*, %struct.simple** %r33
	%r35 = getelementptr %struct.simple, %struct.simple* %r34, i32 0, i32 0
	%r36 = load %struct.foo*, %struct.foo** %math1
	%r37 = getelementptr %struct.foo, %struct.foo* %r36, i32 0, i32 0
	store i64 %r37, i64* %r34
	%r38 = load %struct.foo*, %struct.foo** %math2
	%r39 = getelementptr %struct.foo, %struct.foo* %r38, i32 0, i32 2
	%r40 = load %struct.simple*, %struct.simple** %r39
	%r41 = getelementptr %struct.simple, %struct.simple* %r40, i32 0, i32 0
	%r42 = load %struct.foo*, %struct.foo** %math2
	%r43 = getelementptr %struct.foo, %struct.foo* %r42, i32 0, i32 0
	store i64 %r43, i64* %r40
	br label %L9

L9:
	%r44 = load i64, i64* %P_num
	%r45 = icmp sgt i64 %r44, 0
	br i1 %r45, label %L10, label %L11

L10:
	%r46 = load %struct.foo*, %struct.foo** %math1
	%r47 = getelementptr %struct.foo, %struct.foo* %r46, i32 0, i32 0
	%r48 = load %struct.foo*, %struct.foo** %math2
	%r49 = getelementptr %struct.foo, %struct.foo* %r48, i32 0, i32 0
	%r50 = mul i64 %r47, %r49
	store i64 %r50, i64* %tmp
	%r51 = load i64, i64* %tmp
	%r52 = load %struct.foo*, %struct.foo** %math1
	%r53 = getelementptr %struct.foo, %struct.foo* %r52, i32 0, i32 2
	%r54 = load %struct.simple*, %struct.simple** %r53
	%r55 = getelementptr %struct.simple, %struct.simple* %r54, i32 0, i32 0
	%r56 = mul i64 %r51, %r55
	%r57 = load %struct.foo*, %struct.foo** %math2
	%r58 = getelementptr %struct.foo, %struct.foo* %r57, i32 0, i32 0
	%r59 = sdiv i64 %r56, %r58
	store i64 %r59, i64* %tmp
	%r60 = load %struct.foo*, %struct.foo** %math2
	%r61 = getelementptr %struct.foo, %struct.foo* %r60, i32 0, i32 2
	%r62 = load %struct.simple*, %struct.simple** %r61
	%r63 = getelementptr %struct.simple, %struct.simple* %r62, i32 0, i32 0
	%r64 = load i64, i64* %r63
	%r65 = load %struct.foo*, %struct.foo** %math1
	%r66 = getelementptr %struct.foo, %struct.foo* %r65, i32 0, i32 0
	%r67 = load i64, i64* %r66
	%r68 = call i64 @add(i64 %r64, i64 %r67)
	store i64 %r68, i64* %tmp
	%r69 = load %struct.foo*, %struct.foo** %math2
	%r70 = getelementptr %struct.foo, %struct.foo* %r69, i32 0, i32 0
	%r71 = load %struct.foo*, %struct.foo** %math1
	%r72 = getelementptr %struct.foo, %struct.foo* %r71, i32 0, i32 0
	%r73 = sub i64 %r70, %r72
	store i64 %r73, i64* %tmp
	%r74 = load i64, i64* %P_num
	%r75 = sub i64 %r74, 1
	store i64 %r75, i64* %P_num
	br label %L9

L11:
	%r76 = load %struct.foo*, %struct.foo** %math1
	%r77 = bitcast %struct.foo* %r76 to i8*
	call void @free(i8* %r77)
	%r79 = load %struct.foo*, %struct.foo** %math2
	%r80 = bitcast %struct.foo* %r79 to i8*
	call void @free(i8* %r80)
	br label %L12

L12:
	store i64 0, i64* %domath_retval
	%r82 = load i64, i64* %domath_retval
	ret i64 %r82


}

define i64 @objinstantiation(i64 %num)
{
L13:
	%objinstantiation_retval = alloca i64
	%tmp = alloca %struct.foo*
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	br label %L14

L14:
	%r83 = load i64, i64* %P_num
	%r84 = icmp sgt i64 %r83, 0
	br i1 %r84, label %L15, label %L16

L15:
	%r85 = call i8* @malloc(i32 24)
	%r86 = bitcast i8* %r85 to %struct.foo*
	store %struct.foo* %r86, %struct.foo** %tmp
	%r87 = load %struct.foo*, %struct.foo** %tmp
	%r88 = bitcast %struct.foo* %r87 to i8*
	call void @free(i8* %r88)
	%r90 = load i64, i64* %P_num
	%r91 = sub i64 %r90, 1
	store i64 %r91, i64* %P_num
	br label %L14

L16:
	br label %L17

L17:
	store i64 0, i64* %objinstantiation_retval
	%r92 = load i64, i64* %objinstantiation_retval
	ret i64 %r92


}

define i64 @ackermann(i64 %m, i64 %n)
{
L18:
	%ackermann_retval = alloca i64
	%P_m = alloca i64
	store i64 %m, i64* %P_m
	%P_n = alloca i64
	store i64 %n, i64* %P_n
	br label %L19

L19:
	%r93 = load i64, i64* %P_m
	%r94 = icmp eq i64 %r93, 0
	br i1 %r94, label %L20, label %L21

L20:
	%r95 = load i64, i64* %P_n
	%r96 = add i64 %r95, 1
	store i64 %r96, i64* %ackermann_retval
	%r97 = load i64, i64* %ackermann_retval
	ret i64 %r97
	br label %L22

L21:
	br label %L22

L22:
	br label %L23

L23:
	%r98 = load i64, i64* %P_n
	%r99 = icmp eq i64 %r98, 0
	br i1 %r99, label %L24, label %L25

L24:
	%r100 = load i64, i64* %P_m
	%r101 = sub i64 %r100, 1
	store i64 1, i64* %r102
	%r103 = call i64 @ackermann(i64 %r101, i64 %r102)
	store i64 %r103, i64* %ackermann_retval
	%r104 = load i64, i64* %ackermann_retval
	ret i64 %r104
	br label %L26

L25:
	%r105 = load i64, i64* %P_m
	%r106 = sub i64 %r105, 1
	%r107 = load i64, i64* %P_m
	%r108 = load i64, i64* %P_n
	%r109 = sub i64 %r108, 1
	%r110 = call i64 @ackermann(i64 %r107, i64 %r109)
	%r111 = call i64 @ackermann(i64 %r106, i64 %r110)
	store i64 %r111, i64* %ackermann_retval
	%r112 = load i64, i64* %ackermann_retval
	ret i64 %r112
	br label %L26

L26:
	br label %L27

L27:
	%r113 = load i64, i64* %ackermann_retval
	ret i64 %r113


}

define i64 @main()
{
L28:
	%main_retval = alloca i64
	%a = alloca i64
	%b = alloca i64
	%c = alloca i64
	%d = alloca i64
	%e = alloca i64
	%temp = alloca i64
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %a)
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %b)
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %c)
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %d)
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %e)
	%r114 = load i64, i64* %a
	%r115 = call i64 @tailrecursive(i64 %r114)
	%r116 = load i64, i64* %b
	%r117 = call i64 @domath(i64 %r116)
	%r118 = load i64, i64* %c
	%r119 = call i64 @objinstantiation(i64 %r118)
	%r120 = load i64, i64* %d
	%r121 = load i64, i64* %e
	%r122 = call i64 @ackermann(i64 %r120, i64 %r121)
	store i64 %r122, i64* %temp
	%r123 = load i64, i64* %a
	%r124 = load i64, i64* %b
	%r125 = load i64, i64* %c
	%r126 = load i64, i64* %temp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([28 x i8], [28 x i8]* @.fstr2, i32 0, i32 0), i64 %r123, i64 %r124, i64 %r125, i64 %r126)
	br label %L29

L29:
	store i64 0, i64* %main_retval
	%r127 = load i64, i64* %main_retval
	ret i64 %r127


}


declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [28 x i8] c"a=%ld\0Ab=%ld\0A%c=%ld,temp=%ld\00", align 1

