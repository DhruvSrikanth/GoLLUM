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
	%r9 = load %struct.Node*, %struct.Node** @.nilNode
	%r10 = icmp eq %struct.Node* %r8, %r9
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
	%r20 = icmp eq %struct.Node* %r18, %r19
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
	%r35 = load %struct.Node*, %struct.Node** @.nilNode
	%r36 = icmp eq %struct.Node* %r34, %r35
	br i1 %r36, label %L14, label %L15

L14:
	br label %L28

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
	br label %L27

L18:
	br label %L19

L19:
	%r48 = load %struct.Node*, %struct.Node** %P_curr
	%r49 = getelementptr %struct.Node, %struct.Node* %r48, i32 0, i32 1
	%r50 = load %struct.Node*, %struct.Node** %r49
	%r51 = load %struct.Node*, %struct.Node** @tail
	%r52 = icmp eq %struct.Node* %r50, %r51
	br i1 %r52, label %L20, label %L21

L20:
	%r53 = load %struct.Node*, %struct.Node** @tail
	store %struct.Node* %r53, %struct.Node** %temp
	%r54 = load %struct.Node*, %struct.Node** %P_curr
	store %struct.Node* %r54, %struct.Node** @tail
	%r55 = load %struct.Node*, %struct.Node** @tail
	%r56 = getelementptr %struct.Node, %struct.Node* %r55, i32 0, i32 1
	%r57 = load %struct.Node*, %struct.Node** @.nilNode
	store %struct.Node* %r57, %struct.Node** %r56
	%r58 = load %struct.Node*, %struct.Node** %temp
	%r59 = bitcast %struct.Node* %r58 to i8*
	call void @free(i8* %r59)
	br label %L26

L21:
	br label %L22

L22:
	%r60 = load %struct.Node*, %struct.Node** %P_curr
	%r61 = getelementptr %struct.Node, %struct.Node* %r60, i32 0, i32 1
	%r62 = load %struct.Node*, %struct.Node** %r61
	%r63 = getelementptr %struct.Node, %struct.Node* %r62, i32 0, i32 0
	%r64 = load i64, i64* %r63
	%r65 = load i64, i64* %P_num
	%r66 = icmp eq i64 %r64, %r65
	br i1 %r66, label %L23, label %L24

L23:
	%r67 = load %struct.Node*, %struct.Node** %P_curr
	%r68 = getelementptr %struct.Node, %struct.Node* %r67, i32 0, i32 1
	%r69 = load %struct.Node*, %struct.Node** %r68
	store %struct.Node* %r69, %struct.Node** %temp
	%r70 = load %struct.Node*, %struct.Node** %P_curr
	%r71 = getelementptr %struct.Node, %struct.Node* %r70, i32 0, i32 1
	%r72 = load %struct.Node*, %struct.Node** %P_curr
	%r73 = getelementptr %struct.Node, %struct.Node* %r72, i32 0, i32 1
	%r74 = load %struct.Node*, %struct.Node** %r73
	%r75 = getelementptr %struct.Node, %struct.Node* %r74, i32 0, i32 1
	%r76 = load %struct.Node*, %struct.Node** %r75
	store %struct.Node* %r76, %struct.Node** %r71
	%r77 = load %struct.Node*, %struct.Node** %temp
	%r78 = bitcast %struct.Node* %r77 to i8*
	call void @free(i8* %r78)
	br label %L25

L24:
	%r79 = load %struct.Node*, %struct.Node** %P_curr
	%r80 = getelementptr %struct.Node, %struct.Node* %r79, i32 0, i32 1
	%r81 = load %struct.Node*, %struct.Node** %r80
	%r82 = load i64, i64* %P_num
	%r83 = call i64 @Del(%struct.Node* %r81, i64 %r82)
	br label %L25

L25:
	br label %L26

L26:
	br label %L27

L27:
	br label %L28

L28:
	br label %L29

L29:
	store i64 0, i64* %Del_retval
	%r84 = load i64, i64* %Del_retval
	ret i64 %r84


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
	%r85 = call i64 @Add(i64 1)
	%r86 = call i64 @Add(i64 10)
	%r87 = call i64 @Add(i64 3)
	%r88 = call i64 @Add(i64 4)
	%r89 = load i64, i64* %x
	%r90 = call i64 @Add(i64 %r89)
	%r91 = load %struct.Node*, %struct.Node** @head
	%r92 = call i64 @PrintList(%struct.Node* %r91)
	store i64 0, i64* %i
	br label %L31

L31:
	%r93 = load i64, i64* %i
	%r94 = icmp slt i64 %r93, 50000000
	br i1 %r94, label %L32, label %L33

L32:
	%r95 = load i64, i64* %i
	%r96 = call i64 @Add(i64 %r95)
	%r97 = load i64, i64* %i
	%r98 = add i64 %r97, 1
	store i64 %r98, i64* %i
	br label %L31

L33:
	store i64 0, i64* %i
	br label %L34

L34:
	%r99 = load i64, i64* %i
	%r100 = icmp slt i64 %r99, 50000000
	br i1 %r100, label %L35, label %L36

L35:
	%r101 = load %struct.Node*, %struct.Node** @head
	%r102 = load i64, i64* %i
	%r103 = call i64 @Del(%struct.Node* %r101, i64 %r102)
	%r104 = load i64, i64* %i
	%r105 = add i64 %r104, 1
	store i64 %r105, i64* %i
	br label %L34

L36:
	%r106 = load %struct.Node*, %struct.Node** @head
	%r107 = load i64, i64* %y
	%r108 = call i64 @Del(%struct.Node* %r106, i64 %r107)
	%r109 = load %struct.Node*, %struct.Node** @head
	%r110 = call i64 @PrintList(%struct.Node* %r109)
	br label %L37

L37:
	store i64 0, i64* %main_retval
	%r111 = load i64, i64* %main_retval
	ret i64 %r111


}


declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
declare i8* @malloc(i32)

@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

