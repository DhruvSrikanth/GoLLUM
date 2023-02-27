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
	%_retval = alloca i64
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
	store i64 1, i64* %r3
	store i64 %r3, i64* %_retval
	%r4 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L4

L3:
	br label %L4

L4:
	%r5 = load i64, i64* %P_cur
	%r6 = load i64, i64* %P_neuw
	%r7 = icmp sgt i64 %r5, %r6
	br i1 %r7, label %L5, label %L6

L5:
	store i64 -1, i64* %r8
	store i64 1, i64* %r9
	%r10 = mul i64 %r8, %r9
	store i64 %r10, i64* %_retval
	%r11 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L7

L6:
	store i64 0, i64* %r12
	store i64 %r12, i64* %_retval
	%r13 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L7

L7:
	br label %L4

L8:
	br label %L9

L9:
	%r14 = load i64, i64* %_retval
	ret i64 %r14


}

define i64 @addNode(i64 %numAdd, %struct.Node* %curr)
{
L10:
	%_retval = alloca i64
	%compVal = alloca i64
	%newNode = alloca %struct.Node*
	%P_numAdd = alloca i64
	store i64 %numAdd, i64* %P_numAdd
	%P_curr = alloca %struct.Node*
	store %struct.Node %curr, %struct.Node* %P_curr
	br label %L11

L11:
	%r15 = load %struct.Node*, %struct.Node** %P_curr
	store i64 null, i64* %r16
	%r17 = icmp eq i64 %r15, %r16
	br i1 %r17, label %L12, label %L13

L12:
	%r18 = load %struct.Node*, %struct.Node** %newNode
	%r19 = call i8* @malloc(i32 24)
	%r20 = bitcast i8* %r19 to %struct.Node*
	store %struct.Node* %r20, %struct.Node** %r18
	%r21 = load %struct.Node*, %struct.Node** %newNode
	%r22 = getelementptr %struct.Node*, %struct.Node** %r21, i32 0, i32 0
	%r23 = load i64, i64* %P_numAdd
	store i64 %r23, i64* %r22
	%r24 = load %struct.Node*, %struct.Node** @root
	%r25 = load %struct.Node*, %struct.Node** %newNode
	store %struct.Node* %r25, %struct.Node** %r24
	br label %L14

L13:
	%r26 = load i64, i64* %compVal
	%r27 = load %struct.Node*, %struct.Node** %P_curr
	%r28 = getelementptr %struct.Node*, %struct.Node** %r27, i32 0, i32 0
	%r29 = load i64, i64* %P_numAdd
	%r30 = call i64 @compare(i64 %r28, i64 %r29)
	store i64 %r30, i64* %r26
	br label %L14

L14:
	%r31 = load i64, i64* %compVal
	store i64 -1, i64* %r32
	store i64 1, i64* %r33
	%r34 = mul i64 %r32, %r33
	%r35 = icmp eq i64 %r31, %r34
	br i1 %r35, label %L15, label %L16

L15:
	br label %L16

L16:
	%r36 = load %struct.Node*, %struct.Node** %P_curr
	%r37 = getelementptr %struct.Node*, %struct.Node** %r36, i32 0, i32 1
	store i64 null, i64* %r38
	%r39 = icmp eq i64 %r37, %r38
	br i1 %r39, label %L17, label %L18

L17:
	%r40 = load %struct.Node*, %struct.Node** %newNode
	%r41 = call i8* @malloc(i32 24)
	%r42 = bitcast i8* %r41 to %struct.Node*
	store %struct.Node* %r42, %struct.Node** %r40
	%r43 = load %struct.Node*, %struct.Node** %newNode
	%r44 = getelementptr %struct.Node*, %struct.Node** %r43, i32 0, i32 0
	%r45 = load i64, i64* %P_numAdd
	store i64 %r45, i64* %r44
	%r46 = load %struct.Node*, %struct.Node** %P_curr
	%r47 = getelementptr %struct.Node*, %struct.Node** %r46, i32 0, i32 1
	%r48 = load %struct.Node*, %struct.Node** %newNode
	store %struct.Node* %r48, %struct.Node** %r47
	br label %L19

L18:
	%r49 = load i64, i64* %P_numAdd
	%r50 = load %struct.Node*, %struct.Node** %P_curr
	%r51 = getelementptr %struct.Node*, %struct.Node** %r50, i32 0, i32 1
	%r52 = call i64 @addNode(i64 %r49, %struct.Node %r51)
	br label %L19

L19:
	br label %L17

L20:
	br label %L21

L21:
	%r53 = load i64, i64* %compVal
	store i64 1, i64* %r54
	%r55 = icmp eq i64 %r53, %r54
	br i1 %r55, label %L22, label %L23

L22:
	br label %L23

L23:
	%r56 = load %struct.Node*, %struct.Node** %P_curr
	%r57 = getelementptr %struct.Node*, %struct.Node** %r56, i32 0, i32 2
	store i64 null, i64* %r58
	%r59 = icmp eq i64 %r57, %r58
	br i1 %r59, label %L24, label %L25

L24:
	%r60 = load %struct.Node*, %struct.Node** %newNode
	%r61 = call i8* @malloc(i32 24)
	%r62 = bitcast i8* %r61 to %struct.Node*
	store %struct.Node* %r62, %struct.Node** %r60
	%r63 = load %struct.Node*, %struct.Node** %newNode
	%r64 = getelementptr %struct.Node*, %struct.Node** %r63, i32 0, i32 0
	%r65 = load i64, i64* %P_numAdd
	store i64 %r65, i64* %r64
	%r66 = load %struct.Node*, %struct.Node** %P_curr
	%r67 = getelementptr %struct.Node*, %struct.Node** %r66, i32 0, i32 2
	%r68 = load %struct.Node*, %struct.Node** %newNode
	store %struct.Node* %r68, %struct.Node** %r67
	br label %L26

L25:
	%r69 = load i64, i64* %P_numAdd
	%r70 = load %struct.Node*, %struct.Node** %P_curr
	%r71 = getelementptr %struct.Node*, %struct.Node** %r70, i32 0, i32 2
	%r72 = call i64 @addNode(i64 %r69, %struct.Node %r71)
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
	store i64 0, i64* %_retval
	%r73 = load i64, i64* %r72
	ret i64 %r73


}

define i64 @printDepthTree(%struct.Node* %curr)
{
L32:
	%_retval = alloca i64
	%temp = alloca i64
	%P_curr = alloca %struct.Node*
	store %struct.Node %curr, %struct.Node* %P_curr
	br label %L33

L33:
	%r74 = load %struct.Node*, %struct.Node** %P_curr
	store i64 null, i64* %r75
	%r76 = icmp ne i64 %r74, %r75
	br i1 %r76, label %L34, label %L35

L34:
	br label %L35

L35:
	%r77 = load %struct.Node*, %struct.Node** %P_curr
	%r78 = getelementptr %struct.Node*, %struct.Node** %r77, i32 0, i32 1
	store i64 null, i64* %r79
	%r80 = icmp ne i64 %r78, %r79
	br i1 %r80, label %L36, label %L37

L36:
	%r81 = load %struct.Node*, %struct.Node** %P_curr
	%r82 = getelementptr %struct.Node*, %struct.Node** %r81, i32 0, i32 1
	%r83 = call i64 @printDepthTree(%struct.Node %r82)
	br label %L38

L37:
	br label %L38

L38:
	%r84 = load i64, i64* %temp
	%r85 = load %struct.Node*, %struct.Node** %P_curr
	%r86 = getelementptr %struct.Node*, %struct.Node** %r85, i32 0, i32 0
	store i64 %r86, i64* %r84
	%r87 = load i64, i64* %temp
	%r88 = load i64, i64* %r87
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.fstr1, i32 0, i32 0), i64 %r88)
	br label %L39

L39:
	%r89 = load %struct.Node*, %struct.Node** %P_curr
	%r90 = getelementptr %struct.Node*, %struct.Node** %r89, i32 0, i32 2
	store i64 null, i64* %r91
	%r92 = icmp ne i64 %r90, %r91
	br i1 %r92, label %L40, label %L41

L40:
	%r93 = load %struct.Node*, %struct.Node** %P_curr
	%r94 = getelementptr %struct.Node*, %struct.Node** %r93, i32 0, i32 2
	%r95 = call i64 @printDepthTree(%struct.Node %r94)
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
	store i64 0, i64* %_retval
	%r96 = load i64, i64* %r95
	ret i64 %r96


}

define i64 @deleteLeavesTree(%struct.Node* %curr)
{
L46:
	%_retval = alloca i64
	%P_curr = alloca %struct.Node*
	store %struct.Node %curr, %struct.Node* %P_curr
	br label %L47

L47:
	%r97 = load %struct.Node*, %struct.Node** %P_curr
	store i64 null, i64* %r98
	%r99 = icmp ne i64 %r97, %r98
	br i1 %r99, label %L48, label %L49

L48:
	br label %L49

L49:
	%r100 = load %struct.Node*, %struct.Node** %P_curr
	%r101 = getelementptr %struct.Node*, %struct.Node** %r100, i32 0, i32 1
	store i64 null, i64* %r102
	%r103 = icmp ne i64 %r101, %r102
	br i1 %r103, label %L50, label %L51

L50:
	%r104 = load %struct.Node*, %struct.Node** %P_curr
	%r105 = getelementptr %struct.Node*, %struct.Node** %r104, i32 0, i32 1
	%r106 = call i64 @deleteLeavesTree(%struct.Node %r105)
	br label %L52

L51:
	br label %L52

L52:
	br label %L53

L53:
	%r107 = load %struct.Node*, %struct.Node** %P_curr
	%r108 = getelementptr %struct.Node*, %struct.Node** %r107, i32 0, i32 2
	store i64 null, i64* %r109
	%r110 = icmp ne i64 %r108, %r109
	br i1 %r110, label %L54, label %L55

L54:
	%r111 = load %struct.Node*, %struct.Node** %P_curr
	%r112 = getelementptr %struct.Node*, %struct.Node** %r111, i32 0, i32 2
	%r113 = call i64 @deleteLeavesTree(%struct.Node %r112)
	br label %L56

L55:
	br label %L56

L56:
	%r114 = load %struct.Node*, %struct.Node** %P_curr
	%r115 = bitcast %struct.Node* %r114 to i8*
	call void @free(i8* %r115)
	br label %L50

L57:
	br label %L50

L58:
	br label %L59

L59:
	store i64 0, i64* %_retval
	%r117 = load i64, i64* %r116
	ret i64 %r117


}

define i64 @main()
{
L60:
	%_retval = alloca i64
	%input = alloca i64
	%temp = alloca i64
	%r118 = load %struct.Node*, %struct.Node** @root
	%r119 = load %struct.Node*, %struct.Node** @.nilNode
	store %struct.Node* %r119, %struct.Node** %r118
	%r120 = load i64, i64* %input
	store i64 0, i64* %r121
	store i64 %r121, i64* %r120
	%r122 = load i64, i64* %input
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r122)
	br label %L61

L61:
	%r123 = load i64, i64* %input
	store i64 0, i64* %r124
	%r125 = icmp ne i64 %r123, %r124
	br i1 %r125, label %L62, label %L63

L62:
	%r126 = load i64, i64* %input
	%r127 = load %struct.Node*, %struct.Node** @root
	%r128 = call i64 @addNode(i64 %r126, %struct.Node %r127)
	%r129 = load i64, i64* %input
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r129)
	br label %L61

L63:
	%r130 = load %struct.Node*, %struct.Node** @root
	%r131 = call i64 @printDepthTree(%struct.Node %r130)
	%r132 = load %struct.Node*, %struct.Node** @root
	%r133 = call i64 @deleteLeavesTree(%struct.Node %r132)
	br label %L64

L64:
	store i64 0, i64* %_retval
	%r134 = load i64, i64* %r133
	ret i64 %r134


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [3 x i8] c"%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

