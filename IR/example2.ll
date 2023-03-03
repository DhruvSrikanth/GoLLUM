source_filename = "example2"
target triple = "arm64-apple-darwin22.2.0"

%struct.Node = type {i64, %struct.Node*, %struct.Node*}

@.nilNode = common global %struct.Node* null
@root = common global %struct.Node* null

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

L6:
	store i64 0, i64* %compare_retval
	%r9 = load i64, i64* %compare_retval
	ret i64 %r9

L7:
	br label %L8

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
	%r13 = icmp eq %struct.Node* %r11, %r12
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
	br label %L30

L13:
	%r20 = load %struct.Node*, %struct.Node** %P_curr
	%r21 = getelementptr %struct.Node, %struct.Node* %r20, i32 0, i32 0
	%r22 = load i64, i64* %r21
	%r23 = load i64, i64* %P_numAdd
	%r24 = call i64 @compare(i64 %r22, i64 %r23)
	store i64 %r24, i64* %compVal
	br label %L14

L14:
	%r25 = mul i64 -1, 1
	%r26 = load i64, i64* %compVal
	%r27 = icmp eq i64 %r26, %r25
	br i1 %r27, label %L15, label %L20

L15:
	br label %L16

L16:
	%r28 = load %struct.Node*, %struct.Node** %P_curr
	%r29 = getelementptr %struct.Node, %struct.Node* %r28, i32 0, i32 1
	%r30 = load %struct.Node*, %struct.Node** %r29
	%r31 = load %struct.Node*, %struct.Node** @.nilNode
	%r32 = icmp eq %struct.Node* %r30, %r31
	br i1 %r32, label %L17, label %L18

L17:
	%r33 = call i8* @malloc(i32 24)
	%r34 = bitcast i8* %r33 to %struct.Node*
	store %struct.Node* %r34, %struct.Node** %newNode
	%r35 = load %struct.Node*, %struct.Node** %newNode
	%r36 = getelementptr %struct.Node, %struct.Node* %r35, i32 0, i32 0
	%r37 = load i64, i64* %P_numAdd
	store i64 %r37, i64* %r36
	%r38 = load %struct.Node*, %struct.Node** %P_curr
	%r39 = getelementptr %struct.Node, %struct.Node* %r38, i32 0, i32 1
	%r40 = load %struct.Node*, %struct.Node** %newNode
	store %struct.Node* %r40, %struct.Node** %r39
	br label %L19

L18:
	%r41 = load i64, i64* %P_numAdd
	%r42 = load %struct.Node*, %struct.Node** %P_curr
	%r43 = getelementptr %struct.Node, %struct.Node* %r42, i32 0, i32 1
	%r44 = load %struct.Node*, %struct.Node** %r43
	%r45 = call i64 @addNode(i64 %r41, %struct.Node* %r44)
	br label %L19

L19:
	br label %L29

L20:
	br label %L21

L21:
	%r46 = load i64, i64* %compVal
	%r47 = icmp eq i64 %r46, 1
	br i1 %r47, label %L22, label %L27

L22:
	br label %L23

L23:
	%r48 = load %struct.Node*, %struct.Node** %P_curr
	%r49 = getelementptr %struct.Node, %struct.Node* %r48, i32 0, i32 2
	%r50 = load %struct.Node*, %struct.Node** %r49
	%r51 = load %struct.Node*, %struct.Node** @.nilNode
	%r52 = icmp eq %struct.Node* %r50, %r51
	br i1 %r52, label %L24, label %L25

L24:
	%r53 = call i8* @malloc(i32 24)
	%r54 = bitcast i8* %r53 to %struct.Node*
	store %struct.Node* %r54, %struct.Node** %newNode
	%r55 = load %struct.Node*, %struct.Node** %newNode
	%r56 = getelementptr %struct.Node, %struct.Node* %r55, i32 0, i32 0
	%r57 = load i64, i64* %P_numAdd
	store i64 %r57, i64* %r56
	%r58 = load %struct.Node*, %struct.Node** %P_curr
	%r59 = getelementptr %struct.Node, %struct.Node* %r58, i32 0, i32 2
	%r60 = load %struct.Node*, %struct.Node** %newNode
	store %struct.Node* %r60, %struct.Node** %r59
	br label %L26

L25:
	%r61 = load i64, i64* %P_numAdd
	%r62 = load %struct.Node*, %struct.Node** %P_curr
	%r63 = getelementptr %struct.Node, %struct.Node* %r62, i32 0, i32 2
	%r64 = load %struct.Node*, %struct.Node** %r63
	%r65 = call i64 @addNode(i64 %r61, %struct.Node* %r64)
	br label %L26

L26:
	br label %L28

L27:
	br label %L28

L28:
	br label %L29

L29:
	br label %L30

L30:
	br label %L31

L31:
	store i64 0, i64* %addNode_retval
	%r66 = load i64, i64* %addNode_retval
	ret i64 %r66


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
	%r67 = load %struct.Node*, %struct.Node** %P_curr
	%r68 = load %struct.Node*, %struct.Node** @.nilNode
	%r69 = icmp ne %struct.Node* %r67, %r68
	br i1 %r69, label %L34, label %L43

L34:
	br label %L35

L35:
	%r70 = load %struct.Node*, %struct.Node** %P_curr
	%r71 = getelementptr %struct.Node, %struct.Node* %r70, i32 0, i32 1
	%r72 = load %struct.Node*, %struct.Node** %r71
	%r73 = load %struct.Node*, %struct.Node** @.nilNode
	%r74 = icmp ne %struct.Node* %r72, %r73
	br i1 %r74, label %L36, label %L37

L36:
	%r75 = load %struct.Node*, %struct.Node** %P_curr
	%r76 = getelementptr %struct.Node, %struct.Node* %r75, i32 0, i32 1
	%r77 = load %struct.Node*, %struct.Node** %r76
	%r78 = call i64 @printDepthTree(%struct.Node* %r77)
	br label %L38

L37:
	br label %L38

L38:
	%r79 = load %struct.Node*, %struct.Node** %P_curr
	%r80 = getelementptr %struct.Node, %struct.Node* %r79, i32 0, i32 0
	%r81 = load i64, i64* %r80
	store i64 %r81, i64* %temp
	%r82 = load i64, i64* %temp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr1, i32 0, i32 0), i64 %r82)
	br label %L39

L39:
	%r83 = load %struct.Node*, %struct.Node** %P_curr
	%r84 = getelementptr %struct.Node, %struct.Node* %r83, i32 0, i32 2
	%r85 = load %struct.Node*, %struct.Node** %r84
	%r86 = load %struct.Node*, %struct.Node** @.nilNode
	%r87 = icmp ne %struct.Node* %r85, %r86
	br i1 %r87, label %L40, label %L41

L40:
	%r88 = load %struct.Node*, %struct.Node** %P_curr
	%r89 = getelementptr %struct.Node, %struct.Node* %r88, i32 0, i32 2
	%r90 = load %struct.Node*, %struct.Node** %r89
	%r91 = call i64 @printDepthTree(%struct.Node* %r90)
	br label %L42

L41:
	br label %L42

L42:
	br label %L44

L43:
	br label %L44

L44:
	br label %L45

L45:
	store i64 0, i64* %printDepthTree_retval
	%r92 = load i64, i64* %printDepthTree_retval
	ret i64 %r92


}

define i64 @deleteLeavesTree(%struct.Node* %curr)
{
L46:
	%deleteLeavesTree_retval = alloca i64
	%P_curr = alloca %struct.Node*
	store %struct.Node* %curr, %struct.Node** %P_curr
	br label %L47

L47:
	%r93 = load %struct.Node*, %struct.Node** %P_curr
	%r94 = load %struct.Node*, %struct.Node** @.nilNode
	%r95 = icmp ne %struct.Node* %r93, %r94
	br i1 %r95, label %L48, label %L57

L48:
	br label %L49

L49:
	%r96 = load %struct.Node*, %struct.Node** %P_curr
	%r97 = getelementptr %struct.Node, %struct.Node* %r96, i32 0, i32 1
	%r98 = load %struct.Node*, %struct.Node** %r97
	%r99 = load %struct.Node*, %struct.Node** @.nilNode
	%r100 = icmp ne %struct.Node* %r98, %r99
	br i1 %r100, label %L50, label %L51

L50:
	%r101 = load %struct.Node*, %struct.Node** %P_curr
	%r102 = getelementptr %struct.Node, %struct.Node* %r101, i32 0, i32 1
	%r103 = load %struct.Node*, %struct.Node** %r102
	%r104 = call i64 @deleteLeavesTree(%struct.Node* %r103)
	br label %L52

L51:
	br label %L52

L52:
	br label %L53

L53:
	%r105 = load %struct.Node*, %struct.Node** %P_curr
	%r106 = getelementptr %struct.Node, %struct.Node* %r105, i32 0, i32 2
	%r107 = load %struct.Node*, %struct.Node** %r106
	%r108 = load %struct.Node*, %struct.Node** @.nilNode
	%r109 = icmp ne %struct.Node* %r107, %r108
	br i1 %r109, label %L54, label %L55

L54:
	%r110 = load %struct.Node*, %struct.Node** %P_curr
	%r111 = getelementptr %struct.Node, %struct.Node* %r110, i32 0, i32 2
	%r112 = load %struct.Node*, %struct.Node** %r111
	%r113 = call i64 @deleteLeavesTree(%struct.Node* %r112)
	br label %L56

L55:
	br label %L56

L56:
	%r114 = load %struct.Node*, %struct.Node** %P_curr
	%r115 = bitcast %struct.Node* %r114 to i8*
	call void @free(i8* %r115)
	br label %L58

L57:
	br label %L58

L58:
	br label %L59

L59:
	store i64 0, i64* %deleteLeavesTree_retval
	%r117 = load i64, i64* %deleteLeavesTree_retval
	ret i64 %r117


}

define i64 @main()
{
L60:
	%main_retval = alloca i64
	%input = alloca i64
	%temp = alloca i64
	%r118 = load %struct.Node*, %struct.Node** @.nilNode
	store %struct.Node* %r118, %struct.Node** @root
	store i64 0, i64* %input
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %input)
	br label %L61

L61:
	%r119 = load i64, i64* %input
	%r120 = icmp ne i64 %r119, 0
	br i1 %r120, label %L62, label %L63

L62:
	%r121 = load i64, i64* %input
	%r122 = load %struct.Node*, %struct.Node** @root
	%r123 = call i64 @addNode(i64 %r121, %struct.Node* %r122)
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %input)
	br label %L61

L63:
	%r124 = load %struct.Node*, %struct.Node** @root
	%r125 = call i64 @printDepthTree(%struct.Node* %r124)
	%r126 = load %struct.Node*, %struct.Node** @root
	%r127 = call i64 @deleteLeavesTree(%struct.Node* %r126)
	br label %L64

L64:
	store i64 0, i64* %main_retval
	%r128 = load i64, i64* %main_retval
	ret i64 %r128


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

