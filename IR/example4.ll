source_filename = "example4"
target triple = "arm64-apple-darwin22.2.0"

%struct.Node = type {i64, %struct.Node*}

@.nilNode = common global %struct.Node* null
@head = common global %struct.Node* null
@tail = common global %struct.Node* null

define i64 @Add(i64 %num)
{
L0:
	%Add_retval = alloca i64
	%newList = alloca %struct.Node*
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	%r0 = call i8* @malloc(i32 16)
	%r1 = bitcast i8* %r0 to %struct.Node*
	store %struct.Node* %r1, %struct.Node** %newList
	%r2 = load %struct.Node*, %struct.Node** %newList
	%r3 = getelementptr %struct.Node, %struct.Node* %r2, i32 0, i32 0
	%r4 = load i64, i64* %P_num
	store i64 %r4, i64* %r3
	%r5 = load %struct.Node*, %struct.Node** %newList
	%r6 = getelementptr %struct.Node, %struct.Node* %r5, i32 0, i32 1
	%r7 = load %struct.Node*, %struct.Node** @.nilNode
	store %struct.Node* %r7, %struct.Node** %r6
	br label %L1

L1:
	%r8 = load %struct.Node*, %struct.Node** @head
	%r9 = load i64, i64* @.nilNode
	%r10 = icmp eq i64 %r8, %r9
	br i1 %r10, label %L2, label %L3

L2:
	%r11 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r11, %struct.Node** @head
	%r12 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r12, %struct.Node** @tail
	br label %L4

L3:
	%r13 = load %struct.Node*, %struct.Node** @tail
	%r14 = getelementptr %struct.Node, %struct.Node* %r13, i32 0, i32 1
	%r15 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r15, %struct.Node** %r14
	%r16 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r16, %struct.Node** @tail
	br label %L4

L4:
	br label %L5

L5:
	store i64 0, i64* %Add_retval
	%r17 = load i64, i64* %Add_retval
	ret i64 %r17


}

define i64 @PrintList(%struct.Node* %curr)
{
L6:
	%PrintList_retval = alloca i64
	%printValue = alloca i64
	%P_curr = alloca %struct.Node*
	store %struct.Node* %curr, %struct.Node** %P_curr
	br label %L7

L7:
	%r18 = load %struct.Node*, %struct.Node** %P_curr
	%r19 = load %struct.Node*, %struct.Node** @tail
	%r20 = icmp eq i64 %r18, %r19
	br i1 %r20, label %L8, label %L9

L8:
	%r21 = load %struct.Node*, %struct.Node** %P_curr
	%r22 = getelementptr %struct.Node, %struct.Node* %r21, i32 0, i32 0
	%r23 = load i64, i64* %r22
	store i64 %r23, i64* %printValue
	%r24 = load i64, i64* %printValue
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr1, i32 0, i32 0), i64 %r24)
	br label %L10

L9:
	%r25 = load %struct.Node*, %struct.Node** %P_curr
	%r26 = getelementptr %struct.Node, %struct.Node* %r25, i32 0, i32 0
	%r27 = load i64, i64* %r26
	store i64 %r27, i64* %printValue
	%r28 = load i64, i64* %printValue
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr2, i32 0, i32 0), i64 %r28)
	%r29 = load %struct.Node*, %struct.Node** %P_curr
	%r30 = getelementptr %struct.Node, %struct.Node* %r29, i32 0, i32 1
	%r31 = load %struct.Node*, %struct.Node** %r30
	%r32 = call i64 @PrintList(%struct.Node* %r31)
	br label %L10

L10:
	br label %L11

L11:
	store i64 0, i64* %PrintList_retval
	%r33 = load i64, i64* %PrintList_retval
	ret i64 %r33


}

define i64 @Del(%struct.Node* %curr, i64 %num)
{
L12:
	%Del_retval = alloca i64
	%temp = alloca %struct.Node*
	%P_curr = alloca %struct.Node*
	store %struct.Node* %curr, %struct.Node** %P_curr
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	br label %L13

L13:
	%r34 = load %struct.Node*, %struct.Node** %P_curr
	%r35 = load i64, i64* @.nilNode
	%r36 = icmp eq i64 %r34, %r35
	br i1 %r36, label %L14, label %L15

L14:
	br label %L16

L15:
	br label %L16

L16:
	%r37 = load %struct.Node*, %struct.Node** @head
	%r38 = getelementptr %struct.Node, %struct.Node* %r37, i32 0, i32 0
	%r39 = load i64, i64* %r38
	%r40 = load i64, i64* %P_num
	%r41 = icmp eq i64 %r39, %r40
	br i1 %r41, label %L17, label %L18

L17:
	%r42 = load %struct.Node*, %struct.Node** @head
	store %struct.Node* %r42, %struct.Node** %temp
	%r43 = load %struct.Node*, %struct.Node** @head
	%r44 = getelementptr %struct.Node, %struct.Node* %r43, i32 0, i32 1
	%r45 = load %struct.Node*, %struct.Node** %r44
	store %struct.Node* %r45, %struct.Node** @head
	%r46 = load %struct.Node*, %struct.Node** %temp
	%r47 = bitcast %struct.Node* %r46 to i8*
	call void @free(i8* %r47)
	br label %L19

L18:
	br label %L19

L19:
	%r49 = load %struct.Node*, %struct.Node** %P_curr
	%r50 = getelementptr %struct.Node, %struct.Node* %r49, i32 0, i32 1
	%r51 = load i64, i64* %r50
	%r52 = load %struct.Node*, %struct.Node** @tail
	%r53 = icmp eq i64 %r51, %r52
	br i1 %r53, label %L20, label %L21

L20:
	%r54 = load %struct.Node*, %struct.Node** @tail
	store %struct.Node* %r54, %struct.Node** %temp
	%r55 = load %struct.Node*, %struct.Node** %P_curr
	store %struct.Node* %r55, %struct.Node** @tail
	%r56 = load %struct.Node*, %struct.Node** @tail
	%r57 = getelementptr %struct.Node, %struct.Node* %r56, i32 0, i32 1
	%r58 = load %struct.Node*, %struct.Node** @.nilNode
	store %struct.Node* %r58, %struct.Node** %r57
	%r59 = load %struct.Node*, %struct.Node** %temp
	%r60 = bitcast %struct.Node* %r59 to i8*
	call void @free(i8* %r60)
	br label %L22

L21:
	br label %L22

L22:
	%r62 = load %struct.Node*, %struct.Node** %P_curr
	%r63 = getelementptr %struct.Node, %struct.Node* %r62, i32 0, i32 1
	%r64 = load %struct.Node*, %struct.Node** %r63
	%r65 = getelementptr %struct.Node, %struct.Node* %r64, i32 0, i32 0
	%r66 = load i64, i64* %r65
	%r67 = load i64, i64* %P_num
	%r68 = icmp eq i64 %r66, %r67
	br i1 %r68, label %L23, label %L24

L23:
	%r69 = load %struct.Node*, %struct.Node** %P_curr
	%r70 = getelementptr %struct.Node, %struct.Node* %r69, i32 0, i32 1
	%r71 = load %struct.Node*, %struct.Node** %r70
	store %struct.Node* %r71, %struct.Node** %temp
	%r72 = load %struct.Node*, %struct.Node** %P_curr
	%r73 = getelementptr %struct.Node, %struct.Node* %r72, i32 0, i32 1
	%r74 = load %struct.Node*, %struct.Node** %P_curr
	%r75 = getelementptr %struct.Node, %struct.Node* %r74, i32 0, i32 1
	%r76 = load %struct.Node*, %struct.Node** %r75
	%r77 = getelementptr %struct.Node, %struct.Node* %r76, i32 0, i32 1
	%r78 = load %struct.Node*, %struct.Node** %r77
	store %struct.Node* %r78, %struct.Node** %r73
	%r79 = load %struct.Node*, %struct.Node** %temp
	%r80 = bitcast %struct.Node* %r79 to i8*
	call void @free(i8* %r80)
	br label %L25

L24:
	%r82 = load %struct.Node*, %struct.Node** %P_curr
	%r83 = getelementptr %struct.Node, %struct.Node* %r82, i32 0, i32 1
	%r84 = load %struct.Node*, %struct.Node** %r83
	%r85 = load i64, i64* %P_num
	%r86 = call i64 @Del(%struct.Node* %r84, i64 %r85)
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
	store i64 0, i64* %Del_retval
	%r87 = load i64, i64* %Del_retval
	ret i64 %r87


}

define i64 @main()
{
L30:
	%main_retval = alloca i64
	%x = alloca i64
	%y = alloca i64
	%i = alloca i64
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %x)
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %y)
	%r88 = call i64 @Add(i64 1)
	%r89 = call i64 @Add(i64 10)
	%r90 = call i64 @Add(i64 3)
	%r91 = call i64 @Add(i64 4)
	%r92 = load i64, i64* %x
	%r93 = call i64 @Add(i64 %r92)
	%r94 = load %struct.Node*, %struct.Node** @head
	%r95 = call i64 @PrintList(%struct.Node* %r94)
	store i64 0, i64* %i
	br label %L31

L31:
	%r96 = load i64, i64* %i
	%r97 = icmp slt i64 %r96, 50000000
	br i1 %r97, label %L32, label %L33

L32:
	%r98 = load i64, i64* %i
	%r99 = call i64 @Add(i64 %r98)
	%r100 = load i64, i64* %i
	%r101 = add i64 %r100, 1
	store i64 %r101, i64* %i
	br label %L31

L33:
	store i64 0, i64* %i
	br label %L34

L34:
	%r102 = load i64, i64* %i
	%r103 = icmp slt i64 %r102, 50000000
	br i1 %r103, label %L35, label %L36

L35:
	%r104 = load %struct.Node*, %struct.Node** @head
	%r105 = load i64, i64* %i
	%r106 = call i64 @Del(%struct.Node* %r104, i64 %r105)
	%r107 = load i64, i64* %i
	%r108 = add i64 %r107, 1
	store i64 %r108, i64* %i
	br label %L34

L36:
	%r109 = load %struct.Node*, %struct.Node** @head
	%r110 = load i64, i64* %y
	%r111 = call i64 @Del(%struct.Node* %r109, i64 %r110)
	%r112 = load %struct.Node*, %struct.Node** @head
	%r113 = call i64 @PrintList(%struct.Node* %r112)
	br label %L37

L37:
	store i64 0, i64* %main_retval
	%r114 = load i64, i64* %main_retval
	ret i64 %r114


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

