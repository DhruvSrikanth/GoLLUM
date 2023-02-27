source_filename = "example5"
target triple = "arm64-apple-darwin22.2.0"

%struct.simple = type {i64}
%struct.foo = type {i64, i64, struct.simple}

@.nilsimple = common global %struct.simple* null
@.nilfoo = common global %struct.foo* null
@globalfoo = common global %struct.foo* null
@unusedGlobal = common global %struct.foo* null

define i64 @tailrecursive(i64 %num)
{
L0:
	%_retval = alloca i64
	%unused = alloca %struct.foo*
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	br label %L1

L1:
	%r0 = load i64, i64* %P_num
	store i64 0, i64* %r1
	%r2 = icmp sle i64 %r0, %r1
	br i1 %r2, label %L2, label %L3

L2:
	store i64 0, i64* %_retval
	%r3 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L4

L3:
	br label %L4

L4:
	%r4 = load %struct.foo*, %struct.foo** %unused
	%r5 = call i8* @malloc(i32 24)
	%r6 = bitcast i8* %r5 to %struct.foo*
	store %struct.foo* %r6, %struct.foo** %r4
	%r7 = load %struct.foo*, %struct.foo** @unusedGlobal
	%r8 = load %struct.foo*, %struct.foo** %unused
	store %struct.foo* %r8, %struct.foo** %r7
	%r9 = load i64, i64* %P_num
	store i64 1, i64* %r10
	%r11 = sub i64 %r9, %r10
	%r12 = call i64 @tailrecursive(i64 %r11)
	br label %L5

L5:
	store i64 0, i64* %_retval
	%r13 = load i64, i64* %r12
	ret i64 %r13


}

define i64 @add(i64 %x, i64 %y)
{
L6:
	%_retval = alloca i64
	%P_x = alloca i64
	store i64 %x, i64* %P_x
	%P_y = alloca i64
	store i64 %y, i64* %P_y
	%r14 = load i64, i64* %P_x
	%r15 = load i64, i64* %P_y
	%r16 = add i64 %r14, %r15
	store i64 %r16, i64* %_retval
	%r17 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L7

L7:
	%r18 = load i64, i64* %_retval
	ret i64 %r18


}

define i64 @domath(i64 %num)
{
L8:
	%_retval = alloca i64
	%math1 = alloca %struct.foo*
	%math2 = alloca %struct.foo*
	%tmp = alloca i64
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	%r19 = load %struct.foo*, %struct.foo** %math1
	%r20 = call i8* @malloc(i32 24)
	%r21 = bitcast i8* %r20 to %struct.foo*
	store %struct.foo* %r21, %struct.foo** %r19
	%r22 = load %struct.foo*, %struct.foo** %math1
	%r23 = getelementptr %struct.foo*, %struct.foo** %r22, i32 0, i32 2
	%r24 = call i8* @malloc(i32 8)
	%r25 = bitcast i8* %r24 to %struct.simple*
	store %struct.simple* %r25, %struct.simple** %r23
	%r26 = load %struct.foo*, %struct.foo** %math2
	%r27 = call i8* @malloc(i32 24)
	%r28 = bitcast i8* %r27 to %struct.foo*
	store %struct.foo* %r28, %struct.foo** %r26
	%r29 = load %struct.foo*, %struct.foo** %math2
	%r30 = getelementptr %struct.foo*, %struct.foo** %r29, i32 0, i32 2
	%r31 = call i8* @malloc(i32 8)
	%r32 = bitcast i8* %r31 to %struct.simple*
	store %struct.simple* %r32, %struct.simple** %r30
	%r33 = load %struct.foo*, %struct.foo** %math1
	%r34 = getelementptr %struct.foo*, %struct.foo** %r33, i32 0, i32 0
	%r35 = load i64, i64* %P_num
	store i64 %r35, i64* %r34
	%r36 = load %struct.foo*, %struct.foo** %math2
	%r37 = getelementptr %struct.foo*, %struct.foo** %r36, i32 0, i32 0
	store i64 3, i64* %r38
	store i64 %r38, i64* %r37
	%r39 = load %struct.foo*, %struct.foo** %math1
	%r40 = getelementptr %struct.foo*, %struct.foo** %r39, i32 0, i32 2
	%r41 = getelementptr %struct.simple*, %struct.simple** %r40, i32 0, i32 0
	%r42 = load %struct.foo*, %struct.foo** %math1
	%r43 = getelementptr %struct.foo*, %struct.foo** %r42, i32 0, i32 0
	store i64 %r43, i64* %r41
	%r44 = load %struct.foo*, %struct.foo** %math2
	%r45 = getelementptr %struct.foo*, %struct.foo** %r44, i32 0, i32 2
	%r46 = getelementptr %struct.simple*, %struct.simple** %r45, i32 0, i32 0
	%r47 = load %struct.foo*, %struct.foo** %math2
	%r48 = getelementptr %struct.foo*, %struct.foo** %r47, i32 0, i32 0
	store i64 %r48, i64* %r46
	br label %L9

L9:
	%r49 = load i64, i64* %P_num
	store i64 0, i64* %r50
	%r51 = icmp sgt i64 %r49, %r50
	br i1 %r51, label %L10, label %L11

L10:
	%r52 = load i64, i64* %tmp
	%r53 = load %struct.foo*, %struct.foo** %math1
	%r54 = getelementptr %struct.foo*, %struct.foo** %r53, i32 0, i32 0
	%r55 = load %struct.foo*, %struct.foo** %math2
	%r56 = getelementptr %struct.foo*, %struct.foo** %r55, i32 0, i32 0
	%r57 = mul i64 %r54, %r56
	store i64 %r57, i64* %r52
	%r58 = load i64, i64* %tmp
	%r59 = load i64, i64* %tmp
	%r60 = load %struct.foo*, %struct.foo** %math1
	%r61 = getelementptr %struct.foo*, %struct.foo** %r60, i32 0, i32 2
	%r62 = getelementptr %struct.simple*, %struct.simple** %r61, i32 0, i32 0
	%r63 = mul i64 %r59, %r62
	%r64 = load %struct.foo*, %struct.foo** %math2
	%r65 = getelementptr %struct.foo*, %struct.foo** %r64, i32 0, i32 0
	%r66 = sdiv i64 %r63, %r65
	store i64 %r66, i64* %r58
	%r67 = load i64, i64* %tmp
	%r68 = load %struct.foo*, %struct.foo** %math2
	%r69 = getelementptr %struct.foo*, %struct.foo** %r68, i32 0, i32 2
	%r70 = getelementptr %struct.simple*, %struct.simple** %r69, i32 0, i32 0
	%r71 = load %struct.foo*, %struct.foo** %math1
	%r72 = getelementptr %struct.foo*, %struct.foo** %r71, i32 0, i32 0
	%r73 = call i64 @add(i64 %r70, i64 %r72)
	store i64 %r73, i64* %r67
	%r74 = load i64, i64* %tmp
	%r75 = load %struct.foo*, %struct.foo** %math2
	%r76 = getelementptr %struct.foo*, %struct.foo** %r75, i32 0, i32 0
	%r77 = load %struct.foo*, %struct.foo** %math1
	%r78 = getelementptr %struct.foo*, %struct.foo** %r77, i32 0, i32 0
	%r79 = sub i64 %r76, %r78
	store i64 %r79, i64* %r74
	%r80 = load i64, i64* %P_num
	%r81 = load i64, i64* %P_num
	store i64 1, i64* %r82
	%r83 = sub i64 %r81, %r82
	store i64 %r83, i64* %r80
	br label %L9

L11:
	%r84 = load %struct.foo*, %struct.foo** %math1
	%r85 = bitcast %struct.foo* %r84 to i8*
	call void @free(i8* %r85)
	%r87 = load %struct.foo*, %struct.foo** %math2
	%r88 = bitcast %struct.foo* %r87 to i8*
	call void @free(i8* %r88)
	br label %L12

L12:
	store i64 0, i64* %_retval
	%r90 = load i64, i64* %r89
	ret i64 %r90


}

define i64 @objinstantiation(i64 %num)
{
L13:
	%_retval = alloca i64
	%tmp = alloca %struct.foo*
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	br label %L14

L14:
	%r91 = load i64, i64* %P_num
	store i64 0, i64* %r92
	%r93 = icmp sgt i64 %r91, %r92
	br i1 %r93, label %L15, label %L16

L15:
	%r94 = load %struct.foo*, %struct.foo** %tmp
	%r95 = call i8* @malloc(i32 24)
	%r96 = bitcast i8* %r95 to %struct.foo*
	store %struct.foo* %r96, %struct.foo** %r94
	%r97 = load %struct.foo*, %struct.foo** %tmp
	%r98 = bitcast %struct.foo* %r97 to i8*
	call void @free(i8* %r98)
	%r100 = load i64, i64* %P_num
	%r101 = load i64, i64* %P_num
	store i64 1, i64* %r102
	%r103 = sub i64 %r101, %r102
	store i64 %r103, i64* %r100
	br label %L14

L16:
	br label %L17

L17:
	store i64 0, i64* %_retval
	%r104 = load i64, i64* %r103
	ret i64 %r104


}

define i64 @ackermann(i64 %m, i64 %n)
{
L18:
	%_retval = alloca i64
	%P_m = alloca i64
	store i64 %m, i64* %P_m
	%P_n = alloca i64
	store i64 %n, i64* %P_n
	br label %L19

L19:
	%r105 = load i64, i64* %P_m
	store i64 0, i64* %r106
	%r107 = icmp eq i64 %r105, %r106
	br i1 %r107, label %L20, label %L21

L20:
	%r108 = load i64, i64* %P_n
	store i64 1, i64* %r109
	%r110 = add i64 %r108, %r109
	store i64 %r110, i64* %_retval
	%r111 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L22

L21:
	br label %L22

L22:
	br label %L23

L23:
	%r112 = load i64, i64* %P_n
	store i64 0, i64* %r113
	%r114 = icmp eq i64 %r112, %r113
	br i1 %r114, label %L24, label %L25

L24:
	%r115 = load i64, i64* %P_m
	store i64 1, i64* %r116
	%r117 = sub i64 %r115, %r116
	store i64 1, i64* %r118
	%r119 = call i64 @ackermann(i64 %r117, i64 %r118)
	store i64 %r119, i64* %_retval
	%r120 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L26

L25:
	%r121 = load i64, i64* %P_m
	store i64 1, i64* %r122
	%r123 = sub i64 %r121, %r122
	%r124 = load i64, i64* %P_m
	%r125 = load i64, i64* %P_n
	store i64 1, i64* %r126
	%r127 = sub i64 %r125, %r126
	%r128 = call i64 @ackermann(i64 %r124, i64 %r127)
	%r129 = call i64 @ackermann(i64 %r123, i64 %r128)
	store i64 %r129, i64* %_retval
	%r130 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L26

L26:
	br label %L27

L27:
	%r131 = load i64, i64* %_retval
	ret i64 %r131


}

define i64 @main()
{
L28:
	%_retval = alloca i64
	%a = alloca i64
	%b = alloca i64
	%c = alloca i64
	%d = alloca i64
	%e = alloca i64
	%temp = alloca i64
	%r132 = load i64, i64* %a
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r132)
	%r133 = load i64, i64* %b
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r133)
	%r134 = load i64, i64* %c
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r134)
	%r135 = load i64, i64* %d
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r135)
	%r136 = load i64, i64* %e
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r136)
	%r137 = load i64, i64* %a
	%r138 = call i64 @tailrecursive(i64 %r137)
	%r139 = load i64, i64* %b
	%r140 = call i64 @domath(i64 %r139)
	%r141 = load i64, i64* %c
	%r142 = call i64 @objinstantiation(i64 %r141)
	%r143 = load i64, i64* %temp
	%r144 = load i64, i64* %d
	%r145 = load i64, i64* %e
	%r146 = call i64 @ackermann(i64 %r144, i64 %r145)
	store i64 %r146, i64* %r143
	%r147 = load i64, i64* %a
	%r148 = load i64, i64* %r147
	%r149 = load i64, i64* %b
	%r150 = load i64, i64* %r149
	%r151 = load i64, i64* %c
	%r152 = load i64, i64* %r151
	%r153 = load i64, i64* %temp
	%r154 = load i64, i64* %r153
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([26 x i8], [26 x i8]* @.fstr2, i32 0, i32 0), i64 %r148, i64 %r150, i64 %r152, i64 %r154)
	br label %L29

L29:
	store i64 0, i64* %_retval
	%r155 = load i64, i64* %r154
	ret i64 %r155


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [26 x i8] c"a=%ld\nb=%ld\n%c=%ld,temp=%ld\00", align 1

