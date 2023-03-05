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


}

define i64 @domath(i64 %num)
{
L7:
	%domath_retval = alloca i64
	%math1 = alloca %struct.foo*
	%math2 = alloca %struct.foo*
	%tmp = alloca i64
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	%r14 = call i8* @malloc(i32 24)
	%r15 = bitcast i8* %r14 to %struct.foo*
	store %struct.foo* %r15, %struct.foo** %math1
	%r16 = load %struct.foo*, %struct.foo** %math1
	%r17 = getelementptr %struct.foo, %struct.foo* %r16, i32 0, i32 2
	%r18 = call i8* @malloc(i32 8)
	%r19 = bitcast i8* %r18 to %struct.simple*
	store %struct.simple* %r19, %struct.simple** %r17
	%r20 = call i8* @malloc(i32 24)
	%r21 = bitcast i8* %r20 to %struct.foo*
	store %struct.foo* %r21, %struct.foo** %math2
	%r22 = load %struct.foo*, %struct.foo** %math2
	%r23 = getelementptr %struct.foo, %struct.foo* %r22, i32 0, i32 2
	%r24 = call i8* @malloc(i32 8)
	%r25 = bitcast i8* %r24 to %struct.simple*
	store %struct.simple* %r25, %struct.simple** %r23
	%r26 = load %struct.foo*, %struct.foo** %math1
	%r27 = getelementptr %struct.foo, %struct.foo* %r26, i32 0, i32 0
	%r28 = load i64, i64* %P_num
	store i64 %r28, i64* %r27
	%r29 = load %struct.foo*, %struct.foo** %math2
	%r30 = getelementptr %struct.foo, %struct.foo* %r29, i32 0, i32 0
	store i64 3, i64* %r30
	%r31 = load %struct.foo*, %struct.foo** %math1
	%r32 = getelementptr %struct.foo, %struct.foo* %r31, i32 0, i32 2
	%r33 = load %struct.simple*, %struct.simple** %r32
	%r34 = getelementptr %struct.simple, %struct.simple* %r33, i32 0, i32 0
	%r35 = load %struct.foo*, %struct.foo** %math1
	%r36 = getelementptr %struct.foo, %struct.foo* %r35, i32 0, i32 0
	%r37 = load i64, i64* %r36
	store i64 %r37, i64* %r34
	%r38 = load %struct.foo*, %struct.foo** %math2
	%r39 = getelementptr %struct.foo, %struct.foo* %r38, i32 0, i32 2
	%r40 = load %struct.simple*, %struct.simple** %r39
	%r41 = getelementptr %struct.simple, %struct.simple* %r40, i32 0, i32 0
	%r42 = load %struct.foo*, %struct.foo** %math2
	%r43 = getelementptr %struct.foo, %struct.foo* %r42, i32 0, i32 0
	%r44 = load i64, i64* %r43
	store i64 %r44, i64* %r41
	br label %L8

L8:
	%r45 = load i64, i64* %P_num
	%r46 = icmp sgt i64 %r45, 0
	br i1 %r46, label %L9, label %L10

L9:
	%r47 = load %struct.foo*, %struct.foo** %math1
	%r48 = getelementptr %struct.foo, %struct.foo* %r47, i32 0, i32 0
	%r49 = load %struct.foo*, %struct.foo** %math2
	%r50 = getelementptr %struct.foo, %struct.foo* %r49, i32 0, i32 0
	%r51 = load i64, i64* %r48
	%r52 = load i64, i64* %r50
	%r53 = mul i64 %r51, %r52
	store i64 %r53, i64* %tmp
	%r54 = load %struct.foo*, %struct.foo** %math1
	%r55 = getelementptr %struct.foo, %struct.foo* %r54, i32 0, i32 2
	%r56 = load %struct.simple*, %struct.simple** %r55
	%r57 = getelementptr %struct.simple, %struct.simple* %r56, i32 0, i32 0
	%r58 = load i64, i64* %tmp
	%r59 = load i64, i64* %r57
	%r60 = mul i64 %r58, %r59
	%r61 = load %struct.foo*, %struct.foo** %math2
	%r62 = getelementptr %struct.foo, %struct.foo* %r61, i32 0, i32 0
	%r63 = load i64, i64* %r62
	%r64 = sdiv i64 %r60, %r63
	store i64 %r64, i64* %tmp
	%r65 = load %struct.foo*, %struct.foo** %math2
	%r66 = getelementptr %struct.foo, %struct.foo* %r65, i32 0, i32 2
	%r67 = load %struct.simple*, %struct.simple** %r66
	%r68 = getelementptr %struct.simple, %struct.simple* %r67, i32 0, i32 0
	%r69 = load i64, i64* %r68
	%r70 = load %struct.foo*, %struct.foo** %math1
	%r71 = getelementptr %struct.foo, %struct.foo* %r70, i32 0, i32 0
	%r72 = load i64, i64* %r71
	%r73 = call i64 @add(i64 %r69, i64 %r72)
	store i64 %r73, i64* %tmp
	%r74 = load %struct.foo*, %struct.foo** %math2
	%r75 = getelementptr %struct.foo, %struct.foo* %r74, i32 0, i32 0
	%r76 = load %struct.foo*, %struct.foo** %math1
	%r77 = getelementptr %struct.foo, %struct.foo* %r76, i32 0, i32 0
	%r78 = load i64, i64* %r75
	%r79 = load i64, i64* %r77
	%r80 = sub i64 %r78, %r79
	store i64 %r80, i64* %tmp
	%r81 = load i64, i64* %P_num
	%r82 = sub i64 %r81, 1
	store i64 %r82, i64* %P_num
	br label %L8

L10:
	%r83 = load %struct.foo*, %struct.foo** %math1
	%r84 = bitcast %struct.foo* %r83 to i8*
	call void @free(i8* %r84)
	%r85 = load %struct.foo*, %struct.foo** %math2
	%r86 = bitcast %struct.foo* %r85 to i8*
	call void @free(i8* %r86)
	br label %L11

L11:
	store i64 0, i64* %domath_retval
	%r87 = load i64, i64* %domath_retval
	ret i64 %r87


}

define i64 @objinstantiation(i64 %num)
{
L12:
	%objinstantiation_retval = alloca i64
	%tmp = alloca %struct.foo*
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	br label %L13

L13:
	%r88 = load i64, i64* %P_num
	%r89 = icmp sgt i64 %r88, 0
	br i1 %r89, label %L14, label %L15

L14:
	%r90 = call i8* @malloc(i32 24)
	%r91 = bitcast i8* %r90 to %struct.foo*
	store %struct.foo* %r91, %struct.foo** %tmp
	%r92 = load %struct.foo*, %struct.foo** %tmp
	%r93 = bitcast %struct.foo* %r92 to i8*
	call void @free(i8* %r93)
	%r94 = load i64, i64* %P_num
	%r95 = sub i64 %r94, 1
	store i64 %r95, i64* %P_num
	br label %L13

L15:
	br label %L16

L16:
	store i64 0, i64* %objinstantiation_retval
	%r96 = load i64, i64* %objinstantiation_retval
	ret i64 %r96


}

define i64 @ackermann(i64 %m, i64 %n)
{
L17:
	%ackermann_retval = alloca i64
	%P_m = alloca i64
	store i64 %m, i64* %P_m
	%P_n = alloca i64
	store i64 %n, i64* %P_n
	br label %L18

L18:
	%r97 = load i64, i64* %P_m
	%r98 = icmp eq i64 %r97, 0
	br i1 %r98, label %L19, label %L20

L19:
	%r99 = load i64, i64* %P_n
	%r100 = add i64 %r99, 1
	store i64 %r100, i64* %ackermann_retval
	%r101 = load i64, i64* %ackermann_retval
	ret i64 %r101

L20:
	br label %L21

L21:
	br label %L22

L22:
	%r102 = load i64, i64* %P_n
	%r103 = icmp eq i64 %r102, 0
	br i1 %r103, label %L23, label %L24

L23:
	%r104 = load i64, i64* %P_m
	%r105 = sub i64 %r104, 1
	%r106 = call i64 @ackermann(i64 %r105, i64 1)
	store i64 %r106, i64* %ackermann_retval
	%r107 = load i64, i64* %ackermann_retval
	ret i64 %r107

L24:
	%r108 = load i64, i64* %P_m
	%r109 = sub i64 %r108, 1
	%r110 = load i64, i64* %P_m
	%r111 = load i64, i64* %P_n
	%r112 = sub i64 %r111, 1
	%r113 = call i64 @ackermann(i64 %r110, i64 %r112)
	%r114 = call i64 @ackermann(i64 %r109, i64 %r113)
	store i64 %r114, i64* %ackermann_retval
	%r115 = load i64, i64* %ackermann_retval
	ret i64 %r115

L25:
	br label %L26

L26:
	%r116 = load i64, i64* %ackermann_retval
	ret i64 %r116


}

define i64 @main()
{
L27:
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
	%r117 = load i64, i64* %a
	%r118 = call i64 @tailrecursive(i64 %r117)
	%r119 = load i64, i64* %b
	%r120 = call i64 @domath(i64 %r119)
	%r121 = load i64, i64* %c
	%r122 = call i64 @objinstantiation(i64 %r121)
	%r123 = load i64, i64* %d
	%r124 = load i64, i64* %e
	%r125 = call i64 @ackermann(i64 %r123, i64 %r124)
	store i64 %r125, i64* %temp
	%r126 = load i64, i64* %a
	%r127 = load i64, i64* %b
	%r128 = load i64, i64* %c
	%r129 = load i64, i64* %temp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([29 x i8], [29 x i8]* @.fstr2, i32 0, i32 0), i64 %r126, i64 %r127, i64 %r128, i64 %r129)
	br label %L28

L28:
	store i64 0, i64* %main_retval
	%r130 = load i64, i64* %main_retval
	ret i64 %r130


}


declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)
declare void @free(i8*)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [29 x i8] c"a=%ld\0Ab=%ld\0Ac=%ld\0A,temp=%ld\0A\00", align 1

