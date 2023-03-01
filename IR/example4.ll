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
	store i64 %r22, i64* %printValue
	%r23 = load i64, i64* %printValue
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr1, i32 0, i32 0), i64 %r23)
	br label %L10

L9:
	%r24 = load %struct.Node*, %struct.Node** %P_curr
	%r25 = getelementptr %struct.Node, %struct.Node* %r24, i32 0, i32 0
	store i64 %r25, i64* %printValue
	%r26 = load i64, i64* %printValue
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr2, i32 0, i32 0), i64 %r26)
	%r27 = load %struct.Node*, %struct.Node** %P_curr
	%r28 = getelementptr %struct.Node, %struct.Node* %r27, i32 0, i32 1
	%r29 = call i64 @PrintList(%struct.Node %r28)
	br label %L10

L10:
	br label %L11

L11:
	store i64 0, i64* %PrintList_retval
	%r30 = load i64, i64* %PrintList_retval
	ret i64 %r30


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
	%r31 = load %struct.Node*, %struct.Node** %P_curr
	%r32 = load %struct.Node*, %struct.Node** @.nilNode
	%r33 = icmp eq i64 %r31, %r32
	br i1 %r33, label %L14, label %L15

L14:
	br label %L16

L15:
	br label %L16

L16:
	%r34 = load %struct.Node*, %struct.Node** @head
	%r35 = getelementptr %struct.Node, %struct.Node* %r34, i32 0, i32 0
	%r36 = load i64, i64* %P_num
	%r37 = icmp eq i64 %r35, %r36
	br i1 %r37, label %L17, label %L18

L17:
	%r38 = load %struct.Node*, %struct.Node** @head
	store %struct.Node* %r38, %struct.Node** %temp
	%r39 = load %struct.Node*, %struct.Node** @head
	%r40 = getelementptr %struct.Node, %struct.Node* %r39, i32 0, i32 1
	store %struct.Node* %r40, %struct.Node** @head
	%r41 = load %struct.Node*, %struct.Node** %temp
	%r42 = bitcast %struct.Node* %r41 to i8*
	call void @free(i8* %r42)
	br label %L19

L18:
	br label %L19

L19:
	%r44 = load %struct.Node*, %struct.Node** %P_curr
	%r45 = getelementptr %struct.Node, %struct.Node* %r44, i32 0, i32 1
	%r46 = load %struct.Node*, %struct.Node** @tail
	%r47 = icmp eq i64 %r45, %r46
	br i1 %r47, label %L20, label %L21

L20:
	%r48 = load %struct.Node*, %struct.Node** @tail
	store %struct.Node* %r48, %struct.Node** %temp
	%r49 = load %struct.Node*, %struct.Node** %P_curr
	store %struct.Node* %r49, %struct.Node** @tail
	%r50 = load %struct.Node*, %struct.Node** @tail
	%r51 = getelementptr %struct.Node, %struct.Node* %r50, i32 0, i32 1
	%r52 = load %struct.Node*, %struct.Node** @.nilNode
	store %struct.Node* %r52, %struct.Node** %r51
	%r53 = load %struct.Node*, %struct.Node** %temp
	%r54 = bitcast %struct.Node* %r53 to i8*
	call void @free(i8* %r54)
	br label %L22

L21:
	br label %L22

L22:
	%r56 = load %struct.Node*, %struct.Node** %P_curr
	%r57 = getelementptr %struct.Node, %struct.Node* %r56, i32 0, i32 1
	%r58 = load %struct.Node*, %struct.Node** %r57
	%r59 = getelementptr %struct.Node, %struct.Node* %r58, i32 0, i32 0
	%r60 = load i64, i64* %P_num
	%r61 = icmp eq i64 %r59, %r60
	br i1 %r61, label %L23, label %L24

L23:
	%r62 = load %struct.Node*, %struct.Node** %P_curr
	%r63 = getelementptr %struct.Node, %struct.Node* %r62, i32 0, i32 1
	store %struct.Node* %r63, %struct.Node** %temp
	%r64 = load %struct.Node*, %struct.Node** %P_curr
	%r65 = getelementptr %struct.Node, %struct.Node* %r64, i32 0, i32 1
	%r66 = load %struct.Node*, %struct.Node** %P_curr
	%r67 = getelementptr %struct.Node, %struct.Node* %r66, i32 0, i32 1
	%r68 = load %struct.Node*, %struct.Node** %r67
	%r69 = getelementptr %struct.Node, %struct.Node* %r68, i32 0, i32 1
	store %struct.Node* %r69, %struct.Node** %r65
	%r70 = load %struct.Node*, %struct.Node** %temp
	%r71 = bitcast %struct.Node* %r70 to i8*
	call void @free(i8* %r71)
	br label %L25

L24:
	%r73 = load %struct.Node*, %struct.Node** %P_curr
	%r74 = getelementptr %struct.Node, %struct.Node* %r73, i32 0, i32 1
	%r75 = load i64, i64* %P_num
	%r76 = call i64 @Del(%struct.Node %r74, i64 %r75)
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
	%r77 = load i64, i64* %Del_retval
	ret i64 %r77


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
	%r78 = call i64 @Add(i64 %r1)
	%r79 = call i64 @Add(i64 %r10)
	%r80 = call i64 @Add(i64 %r3)
	%r81 = call i64 @Add(i64 %r4)
	%r82 = load i64, i64* %x
	%r83 = call i64 @Add(i64 %r82)
	%r84 = load %struct.Node*, %struct.Node** @head
	%r85 = call i64 @PrintList(%struct.Node %r84)
	store i64 0, i64* %i
	br label %L31

L31:
	%r86 = load i64, i64* %i
	%r87 = icmp slt i64 %r86, 50000000
	br i1 %r87, label %L32, label %L33

L32:
	%r88 = load i64, i64* %i
	%r89 = call i64 @Add(i64 %r88)
	%r90 = load i64, i64* %i
	%r91 = add i64 %r90, 1
	store i64 %r91, i64* %i
	br label %L31

L33:
	store i64 0, i64* %i
	br label %L34

L34:
	%r92 = load i64, i64* %i
	%r93 = icmp slt i64 %r92, 50000000
	br i1 %r93, label %L35, label %L36

L35:
	%r94 = load %struct.Node*, %struct.Node** @head
	%r95 = load i64, i64* %i
	%r96 = call i64 @Del(%struct.Node %r94, i64 %r95)
	%r97 = load i64, i64* %i
	%r98 = add i64 %r97, 1
	store i64 %r98, i64* %i
	br label %L34

L36:
	%r99 = load %struct.Node*, %struct.Node** @head
	%r100 = load i64, i64* %y
	%r101 = call i64 @Del(%struct.Node %r99, i64 %r100)
	%r102 = load %struct.Node*, %struct.Node** @head
	%r103 = call i64 @PrintList(%struct.Node %r102)
	br label %L37

L37:
	store i64 0, i64* %main_retval
	%r104 = load i64, i64* %main_retval
	ret i64 %r104


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

