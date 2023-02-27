source_filename = "example4"
target triple = "arm64-apple-darwin22.2.0"

%struct.Node = type {i64, %struct.Node*}

@.nilNode = common global %struct.Node* null
@head = common global %struct.Node* null
@tail = common global %struct.Node* null

define i64 @Add(i64 %num)
{
L0:
	%_retval = alloca i64
	%newList = alloca %struct.Node*
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	%r0 = load %struct.Node*, %struct.Node** %newList
	%r1 = call i8* @malloc(i32 16)
	%r2 = bitcast i8* %r1 to %struct.Node*
	store %struct.Node* %r2, %struct.Node** %r0
	%r3 = load %struct.Node*, %struct.Node** %newList
	%r4 = getelementptr %struct.Node*, %struct.Node** %r3, i32 0, i32 0
	%r5 = load i64, i64* %P_num
	store i64 %r5, i64* %r4
	%r6 = load %struct.Node*, %struct.Node** %newList
	%r7 = getelementptr %struct.Node*, %struct.Node** %r6, i32 0, i32 1
	%r8 = load %struct.Node*, %struct.Node** @.nilNode
	store %struct.Node* %r8, %struct.Node** %r7
	br label %L1

L1:
	%r9 = load %struct.Node*, %struct.Node** @head
	store i64 null, i64* %r10
	%r11 = icmp eq i64 %r9, %r10
	br i1 %r11, label %L2, label %L3

L2:
	%r12 = load %struct.Node*, %struct.Node** @head
	%r13 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r13, %struct.Node** %r12
	%r14 = load %struct.Node*, %struct.Node** @tail
	%r15 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r15, %struct.Node** %r14
	br label %L4

L3:
	%r16 = load %struct.Node*, %struct.Node** @tail
	%r17 = getelementptr %struct.Node*, %struct.Node** %r16, i32 0, i32 1
	%r18 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r18, %struct.Node** %r17
	%r19 = load %struct.Node*, %struct.Node** @tail
	%r20 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r20, %struct.Node** %r19
	br label %L4

L4:
	br label %L5

L5:
	store i64 0, i64* %_retval
	%r21 = load i64, i64* %r20
	ret i64 %r21


}

define i64 @PrintList(%struct.Node* %curr)
{
L6:
	%_retval = alloca i64
	%printValue = alloca i64
	%P_curr = alloca %struct.Node*
	store %struct.Node %curr, %struct.Node* %P_curr
	br label %L7

L7:
	%r22 = load %struct.Node*, %struct.Node** %P_curr
	%r23 = load %struct.Node*, %struct.Node** @tail
	%r24 = icmp eq i64 %r22, %r23
	br i1 %r24, label %L8, label %L9

L8:
	%r25 = load i64, i64* %printValue
	%r26 = load %struct.Node*, %struct.Node** %P_curr
	%r27 = getelementptr %struct.Node*, %struct.Node** %r26, i32 0, i32 0
	store i64 %r27, i64* %r25
	%r28 = load i64, i64* %printValue
	%r29 = load i64, i64* %r28
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.fstr1, i32 0, i32 0), i64 %r29)
	br label %L10

L9:
	%r30 = load i64, i64* %printValue
	%r31 = load %struct.Node*, %struct.Node** %P_curr
	%r32 = getelementptr %struct.Node*, %struct.Node** %r31, i32 0, i32 0
	store i64 %r32, i64* %r30
	%r33 = load i64, i64* %printValue
	%r34 = load i64, i64* %r33
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.fstr2, i32 0, i32 0), i64 %r34)
	%r35 = load %struct.Node*, %struct.Node** %P_curr
	%r36 = getelementptr %struct.Node*, %struct.Node** %r35, i32 0, i32 1
	%r37 = call i64 @PrintList(%struct.Node %r36)
	br label %L10

L10:
	br label %L11

L11:
	store i64 0, i64* %_retval
	%r38 = load i64, i64* %r37
	ret i64 %r38


}

define i64 @Del(%struct.Node* %curr, i64 %num)
{
L12:
	%_retval = alloca i64
	%temp = alloca %struct.Node*
	%P_curr = alloca %struct.Node*
	store %struct.Node %curr, %struct.Node* %P_curr
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	br label %L13

L13:
	%r39 = load %struct.Node*, %struct.Node** %P_curr
	store i64 null, i64* %r40
	%r41 = icmp eq i64 %r39, %r40
	br i1 %r41, label %L14, label %L15

L14:
	br label %L16

L15:
	br label %L16

L16:
	%r42 = load %struct.Node*, %struct.Node** @head
	%r43 = getelementptr %struct.Node*, %struct.Node** %r42, i32 0, i32 0
	%r44 = load i64, i64* %P_num
	%r45 = icmp eq i64 %r43, %r44
	br i1 %r45, label %L17, label %L18

L17:
	%r46 = load %struct.Node*, %struct.Node** %temp
	%r47 = load %struct.Node*, %struct.Node** @head
	store %struct.Node* %r47, %struct.Node** %r46
	%r48 = load %struct.Node*, %struct.Node** @head
	%r49 = load %struct.Node*, %struct.Node** @head
	%r50 = getelementptr %struct.Node*, %struct.Node** %r49, i32 0, i32 1
	store %struct.Node* %r50, %struct.Node** %r48
	%r51 = load %struct.Node*, %struct.Node** %temp
	%r52 = bitcast %struct.Node* %r51 to i8*
	call void @free(i8* %r52)
	br label %L19

L18:
	br label %L19

L19:
	%r54 = load %struct.Node*, %struct.Node** %P_curr
	%r55 = getelementptr %struct.Node*, %struct.Node** %r54, i32 0, i32 1
	%r56 = load %struct.Node*, %struct.Node** @tail
	%r57 = icmp eq i64 %r55, %r56
	br i1 %r57, label %L20, label %L21

L20:
	%r58 = load %struct.Node*, %struct.Node** %temp
	%r59 = load %struct.Node*, %struct.Node** @tail
	store %struct.Node* %r59, %struct.Node** %r58
	%r60 = load %struct.Node*, %struct.Node** @tail
	%r61 = load %struct.Node*, %struct.Node** %P_curr
	store %struct.Node* %r61, %struct.Node** %r60
	%r62 = load %struct.Node*, %struct.Node** @tail
	%r63 = getelementptr %struct.Node*, %struct.Node** %r62, i32 0, i32 1
	%r64 = load %struct.Node*, %struct.Node** @.nilNode
	store %struct.Node* %r64, %struct.Node** %r63
	%r65 = load %struct.Node*, %struct.Node** %temp
	%r66 = bitcast %struct.Node* %r65 to i8*
	call void @free(i8* %r66)
	br label %L22

L21:
	br label %L22

L22:
	%r68 = load %struct.Node*, %struct.Node** %P_curr
	%r69 = getelementptr %struct.Node*, %struct.Node** %r68, i32 0, i32 1
	%r70 = getelementptr %struct.Node*, %struct.Node** %r69, i32 0, i32 0
	%r71 = load i64, i64* %P_num
	%r72 = icmp eq i64 %r70, %r71
	br i1 %r72, label %L23, label %L24

L23:
	%r73 = load %struct.Node*, %struct.Node** %temp
	%r74 = load %struct.Node*, %struct.Node** %P_curr
	%r75 = getelementptr %struct.Node*, %struct.Node** %r74, i32 0, i32 1
	store %struct.Node* %r75, %struct.Node** %r73
	%r76 = load %struct.Node*, %struct.Node** %P_curr
	%r77 = getelementptr %struct.Node*, %struct.Node** %r76, i32 0, i32 1
	%r78 = load %struct.Node*, %struct.Node** %P_curr
	%r79 = getelementptr %struct.Node*, %struct.Node** %r78, i32 0, i32 1
	%r80 = getelementptr %struct.Node*, %struct.Node** %r79, i32 0, i32 1
	store %struct.Node* %r80, %struct.Node** %r77
	%r81 = load %struct.Node*, %struct.Node** %temp
	%r82 = bitcast %struct.Node* %r81 to i8*
	call void @free(i8* %r82)
	br label %L25

L24:
	%r84 = load %struct.Node*, %struct.Node** %P_curr
	%r85 = getelementptr %struct.Node*, %struct.Node** %r84, i32 0, i32 1
	%r86 = load i64, i64* %P_num
	%r87 = call i64 @Del(%struct.Node %r85, i64 %r86)
	br label %L25

L25:
	br label %L22

L26:
	br label %L19

L27:
	br label %L16

L28:
	br label %L29

L29:
	store i64 0, i64* %_retval
	%r88 = load i64, i64* %r87
	ret i64 %r88


}

define i64 @main()
{
L30:
	%_retval = alloca i64
	%x = alloca i64
	%y = alloca i64
	%i = alloca i64
	%r89 = load i64, i64* %x
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r89)
	%r90 = load i64, i64* %y
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r90)
	store i64 1, i64* %r91
	%r92 = call i64 @Add(i64 %r91)
	store i64 10, i64* %r93
	%r94 = call i64 @Add(i64 %r93)
	store i64 3, i64* %r95
	%r96 = call i64 @Add(i64 %r95)
	store i64 4, i64* %r97
	%r98 = call i64 @Add(i64 %r97)
	%r99 = load i64, i64* %x
	%r100 = call i64 @Add(i64 %r99)
	%r101 = load %struct.Node*, %struct.Node** @head
	%r102 = call i64 @PrintList(%struct.Node %r101)
	%r103 = load i64, i64* %i
	store i64 0, i64* %r104
	store i64 %r104, i64* %r103
	br label %L31

L31:
	%r105 = load i64, i64* %i
	store i64 50000000, i64* %r106
	%r107 = icmp slt i64 %r105, %r106
	br i1 %r107, label %L32, label %L33

L32:
	%r108 = load i64, i64* %i
	%r109 = call i64 @Add(i64 %r108)
	%r110 = load i64, i64* %i
	%r111 = load i64, i64* %i
	store i64 1, i64* %r112
	%r113 = add i64 %r111, %r112
	store i64 %r113, i64* %r110
	br label %L31

L33:
	%r114 = load i64, i64* %i
	store i64 0, i64* %r115
	store i64 %r115, i64* %r114
	br label %L34

L34:
	%r116 = load i64, i64* %i
	store i64 50000000, i64* %r117
	%r118 = icmp slt i64 %r116, %r117
	br i1 %r118, label %L35, label %L36

L35:
	%r119 = load %struct.Node*, %struct.Node** @head
	%r120 = load i64, i64* %i
	%r121 = call i64 @Del(%struct.Node %r119, i64 %r120)
	%r122 = load i64, i64* %i
	%r123 = load i64, i64* %i
	store i64 1, i64* %r124
	%r125 = add i64 %r123, %r124
	store i64 %r125, i64* %r122
	br label %L34

L36:
	%r126 = load %struct.Node*, %struct.Node** @head
	%r127 = load i64, i64* %y
	%r128 = call i64 @Del(%struct.Node %r126, i64 %r127)
	%r129 = load %struct.Node*, %struct.Node** @head
	%r130 = call i64 @PrintList(%struct.Node %r129)
	br label %L37

L37:
	store i64 0, i64* %_retval
	%r131 = load i64, i64* %r130
	ret i64 %r131


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [3 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [3 x i8] c"%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

