source_filename = "example2"
target triple = "arm64-apple-darwin22.2.0"

%struct.Node = type {i64, %struct.Node*, %struct.Node*}

@.nilNode = common global %struct.Node* null
@root = common global %struct.Node* null
@x = common global i64 0
@y = common global i64 0

define i64 @compare(i64 %cur, i64 %neuw)
{
L0:
	%compare_retval = alloca i64
	%P_cur = alloca i64
	store i64 %cur, i64* %P_cur
	%P_neuw = alloca i64
	store i64 %neuw, i64* %P_neuw
	br label %L1

L1:
	%r0 = load i64, i64* %P_cur
	%r1 = load i64, i64* %P_neuw
	%r2 = icmp slt i64 %r0, %r1
	br i1 %r2, label %L2, label %L3

L2:
	store i64 1, i64* %compare_retval
	%r3 = load i64, i64* %compare_retval
	ret i64 %r3
	br label %L4

L3:
	br label %L4

L4:
	%r4 = load i64, i64* %P_cur
	%r5 = load i64, i64* %P_neuw
	%r6 = icmp sgt i64 %r4, %r5
	br i1 %r6, label %L5, label %L6

L5:
	%r7 = mul i64 -1, 1
	store i64 %r7, i64* %compare_retval
	%r8 = load i64, i64* %compare_retval
	ret i64 %r8
	br label %L7

L6:
	store i64 0, i64* %compare_retval
	%r9 = load i64, i64* %compare_retval
	ret i64 %r9
	br label %L7

L7:
	br label %L4

L8:
	br label %L9

L9:
	%r10 = load i64, i64* %compare_retval
	ret i64 %r10


}

define i64 @addNode(i64 %numAdd, %struct.Node* %curr)
{
L10:
	%addNode_retval = alloca i64
	%compVal = alloca i64
	%newNode = alloca %struct.Node*
	%P_numAdd = alloca i64
	store i64 %numAdd, i64* %P_numAdd
	%P_curr = alloca %struct.Node*
	store %struct.Node* %curr, %struct.Node** %P_curr
	br label %L11

L11:
	%r11 = load %struct.Node*, %struct.Node** %P_curr
	%r12 = load %struct.Node*, %struct.Node** @.nilNode
	%r13 = icmp eq i64 %r11, %r12
	br i1 %r13, label %L12, label %L13

L12:
	%r14 = call i8* @malloc(i32 24)
	%r15 = bitcast i8* %r14 to %struct.Node*
	store %struct.Node* %r15, %struct.Node** %newNode
	%r16 = load %struct.Node*, %struct.Node** %newNode
	%r17 = getelementptr %struct.Node, %struct.Node* %r16, i32 0, i32 0
	%r18 = load i64, i64* %P_numAdd
	store i64 %r18, i64* %r17
	%r19 = load %struct.Node*, %struct.Node** %newNode
	store %struct.Node* %r19, %struct.Node** @root
	br label %L14

L13:
	%r20 = load %struct.Node*, %struct.Node** %P_curr
	%r21 = getelementptr %struct.Node, %struct.Node* %r20, i32 0, i32 0
	%r22 = load i64, i64* %r21
	%r23 = load i64, i64* %P_numAdd
	%r24 = call i64 @compare(i64 %r22, i64 %r23)
	store i64 %r24, i64* %compVal
	br label %L14

L14:
	%r25 = load i64, i64* %compVal
	%r26 = mul i64 -1, 1
	%r27 = icmp eq i64 %r25, %r26
	br i1 %r27, label %L15, label %L16

L15:
	br label %L16

L16:
	%r28 = load %struct.Node*, %struct.Node** %P_curr
	%r29 = getelementptr %struct.Node, %struct.Node* %r28, i32 0, i32 1
	%r30 = load %struct.Node*, %struct.Node** @.nilNode
	%r31 = icmp eq i64 %r29, %r30
	br i1 %r31, label %L17, label %L18

L17:
	%r32 = call i8* @malloc(i32 24)
	%r33 = bitcast i8* %r32 to %struct.Node*
	store %struct.Node* %r33, %struct.Node** %newNode
	%r34 = load %struct.Node*, %struct.Node** %newNode
	%r35 = getelementptr %struct.Node, %struct.Node* %r34, i32 0, i32 0
	%r36 = load i64, i64* %P_numAdd
	store i64 %r36, i64* %r35
	%r37 = load %struct.Node*, %struct.Node** %P_curr
	%r38 = getelementptr %struct.Node, %struct.Node* %r37, i32 0, i32 1
	%r39 = load %struct.Node*, %struct.Node** %newNode
	store %struct.Node* %r39, %struct.Node** %r38
	br label %L19

L18:
	%r40 = load i64, i64* %P_numAdd
	%r41 = load %struct.Node*, %struct.Node** %P_curr
	%r42 = getelementptr %struct.Node, %struct.Node* %r41, i32 0, i32 1
	%r43 = call i64 @addNode(i64 %r40, %struct.Node %r42)
	br label %L19

L19:
	br label %L17

L20:
	br label %L21

L21:
	%r44 = load i64, i64* %compVal
	%r45 = icmp eq i64 %r44, 1
	br i1 %r45, label %L22, label %L23

L22:
	br label %L23

L23:
	%r46 = load %struct.Node*, %struct.Node** %P_curr
	%r47 = getelementptr %struct.Node, %struct.Node* %r46, i32 0, i32 2
	%r48 = load %struct.Node*, %struct.Node** @.nilNode
	%r49 = icmp eq i64 %r47, %r48
	br i1 %r49, label %L24, label %L25

L24:
	%r50 = call i8* @malloc(i32 24)
	%r51 = bitcast i8* %r50 to %struct.Node*
	store %struct.Node* %r51, %struct.Node** %newNode
	%r52 = load %struct.Node*, %struct.Node** %newNode
	%r53 = getelementptr %struct.Node, %struct.Node* %r52, i32 0, i32 0
	%r54 = load i64, i64* %P_numAdd
	store i64 %r54, i64* %r53
	%r55 = load %struct.Node*, %struct.Node** %P_curr
	%r56 = getelementptr %struct.Node, %struct.Node* %r55, i32 0, i32 2
	%r57 = load %struct.Node*, %struct.Node** %newNode
	store %struct.Node* %r57, %struct.Node** %r56
	br label %L26

L25:
	%r58 = load i64, i64* %P_numAdd
	%r59 = load %struct.Node*, %struct.Node** %P_curr
	%r60 = getelementptr %struct.Node, %struct.Node* %r59, i32 0, i32 2
	%r61 = call i64 @addNode(i64 %r58, %struct.Node %r60)
	br label %L26

L26:
	br label %L24

L27:
	br label %L24

L28:
	br label %L17

L29:
	br label %L14

L30:
	br label %L31

L31:
	store i64 0, i64* %addNode_retval
	%r62 = load i64, i64* %addNode_retval
	ret i64 %r62


}

define i64 @printDepthTree(%struct.Node* %curr)
{
L32:
	%printDepthTree_retval = alloca i64
	%temp = alloca i64
	%P_curr = alloca %struct.Node*
	store %struct.Node* %curr, %struct.Node** %P_curr
	br label %L33

L33:
	%r63 = load %struct.Node*, %struct.Node** %P_curr
	%r64 = load %struct.Node*, %struct.Node** @.nilNode
	%r65 = icmp ne i64 %r63, %r64
	br i1 %r65, label %L34, label %L35

L34:
	br label %L35

L35:
	%r66 = load %struct.Node*, %struct.Node** %P_curr
	%r67 = getelementptr %struct.Node, %struct.Node* %r66, i32 0, i32 1
	%r68 = load %struct.Node*, %struct.Node** @.nilNode
	%r69 = icmp ne i64 %r67, %r68
	br i1 %r69, label %L36, label %L37

L36:
	%r70 = load %struct.Node*, %struct.Node** %P_curr
	%r71 = getelementptr %struct.Node, %struct.Node* %r70, i32 0, i32 1
	%r72 = call i64 @printDepthTree(%struct.Node %r71)
	br label %L38

L37:
	br label %L38

L38:
	%r73 = load %struct.Node*, %struct.Node** %P_curr
	%r74 = getelementptr %struct.Node, %struct.Node* %r73, i32 0, i32 0
	store i64 %r74, i64* %temp
	%r75 = load i64, i64* %temp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr1, i32 0, i32 0), i64 %r75)
	br label %L39

L39:
	%r76 = load %struct.Node*, %struct.Node** %P_curr
	%r77 = getelementptr %struct.Node, %struct.Node* %r76, i32 0, i32 2
	%r78 = load %struct.Node*, %struct.Node** @.nilNode
	%r79 = icmp ne i64 %r77, %r78
	br i1 %r79, label %L40, label %L41

L40:
	%r80 = load %struct.Node*, %struct.Node** %P_curr
	%r81 = getelementptr %struct.Node, %struct.Node* %r80, i32 0, i32 2
	%r82 = call i64 @printDepthTree(%struct.Node %r81)
	br label %L42

L41:
	br label %L42

L42:
	br label %L36

L43:
	br label %L36

L44:
	br label %L45

L45:
	store i64 0, i64* %printDepthTree_retval
	%r83 = load i64, i64* %printDepthTree_retval
	ret i64 %r83


}

define i64 @deleteLeavesTree(%struct.Node* %curr)
{
L46:
	%deleteLeavesTree_retval = alloca i64
	%P_curr = alloca %struct.Node*
	store %struct.Node* %curr, %struct.Node** %P_curr
	br label %L47

L47:
	%r84 = load %struct.Node*, %struct.Node** %P_curr
	%r85 = load %struct.Node*, %struct.Node** @.nilNode
	%r86 = icmp ne i64 %r84, %r85
	br i1 %r86, label %L48, label %L49

L48:
	br label %L49

L49:
	%r87 = load %struct.Node*, %struct.Node** %P_curr
	%r88 = getelementptr %struct.Node, %struct.Node* %r87, i32 0, i32 1
	%r89 = load %struct.Node*, %struct.Node** @.nilNode
	%r90 = icmp ne i64 %r88, %r89
	br i1 %r90, label %L50, label %L51

L50:
	%r91 = load %struct.Node*, %struct.Node** %P_curr
	%r92 = getelementptr %struct.Node, %struct.Node* %r91, i32 0, i32 1
	%r93 = call i64 @deleteLeavesTree(%struct.Node %r92)
	br label %L52

L51:
	br label %L52

L52:
	br label %L53

L53:
	%r94 = load %struct.Node*, %struct.Node** %P_curr
	%r95 = getelementptr %struct.Node, %struct.Node* %r94, i32 0, i32 2
	%r96 = load %struct.Node*, %struct.Node** @.nilNode
	%r97 = icmp ne i64 %r95, %r96
	br i1 %r97, label %L54, label %L55

L54:
	%r98 = load %struct.Node*, %struct.Node** %P_curr
	%r99 = getelementptr %struct.Node, %struct.Node* %r98, i32 0, i32 2
	%r100 = call i64 @deleteLeavesTree(%struct.Node %r99)
	br label %L56

L55:
	br label %L56

L56:
	%r101 = load %struct.Node*, %struct.Node** %P_curr
	%r102 = bitcast %struct.Node* %r101 to i8*
	call void @free(i8* %r102)
	br label %L50

L57:
	br label %L50

L58:
	br label %L59

L59:
	store i64 0, i64* %deleteLeavesTree_retval
	%r104 = load i64, i64* %deleteLeavesTree_retval
	ret i64 %r104


}

define i64 @main()
{
L60:
	%main_retval = alloca i64
	%input = alloca i64
	%temp = alloca i64
	%r105 = load %struct.Node*, %struct.Node** @.nilNode
	store %struct.Node* %r105, %struct.Node** @root
	store i64 0, i64* %input
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %input)
	br label %L61

L61:
	%r106 = load i64, i64* %input
	%r107 = icmp ne i64 %r106, 0
	br i1 %r107, label %L62, label %L63

L62:
	%r108 = load i64, i64* %input
	%r109 = load %struct.Node*, %struct.Node** @root
	%r110 = call i64 @addNode(i64 %r108, %struct.Node %r109)
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %input)
	br label %L61

L63:
	%r111 = load %struct.Node*, %struct.Node** @root
	%r112 = call i64 @printDepthTree(%struct.Node %r111)
	%r113 = load %struct.Node*, %struct.Node** @root
	%r114 = call i64 @deleteLeavesTree(%struct.Node %r113)
	br label %L64

L64:
	store i64 0, i64* %main_retval
	%r115 = load i64, i64* %main_retval
	ret i64 %r115


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

