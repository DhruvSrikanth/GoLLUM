source_filename = "example3"
target triple = "arm64-apple-darwin22.2.0"

%struct.Cell = type {%struct.Cell*, i64}
%struct.Row = type {%struct.Row*, %struct.Cell*}
%struct.ListNode = type {%struct.ListNode*, i64, i64}

@.nilCell = common global %struct.Cell* null
@.nilRow = common global %struct.Row* null
@.nilListNode = common global %struct.ListNode* null
@matrix = common global %struct.Row* null
@maxrange = common global i64 0

define i64 @timesone(i64 %iter)
{
L0:
	%timesone_retval = alloca i64
	%P_iter = alloca i64
	store i64 %iter, i64* %P_iter
	br label %L1

L1:
	%r0 = load i64, i64* %P_iter
	%r1 = icmp sgt i64 %r0, 0
	br i1 %r1, label %L2, label %L3

L2:
	%r2 = load i64, i64* @maxrange
	%r3 = mul i64 %r2, 1
	store i64 %r3, i64* @maxrange
	%r4 = load i64, i64* %P_iter
	%r5 = sub i64 %r4, 1
	store i64 %r5, i64* %P_iter
	br label %L1

L3:
	br label %L4

L4:
	store i64 0, i64* %timesone_retval
	%r6 = load i64, i64* %timesone_retval
	ret i64 %r6


}

define i64 @find(i64 %num, %struct.ListNode* %list)
{
L5:
	%find_retval = alloca i64
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	%P_list = alloca %struct.ListNode*
	store %struct.ListNode* %list, %struct.ListNode** %P_list
	br label %L6

L6:
	%r7 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r8 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	%r9 = icmp ne %struct.ListNode* %r7, %r8
	br i1 %r9, label %L7, label %L8

L7:
	br label %L8

L8:
	%r10 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r11 = getelementptr %struct.ListNode, %struct.ListNode* %r10, i32 0, i32 1
	%r12 = load i64, i64* %r11
	%r13 = load i64, i64* %P_num
	%r14 = icmp eq i64 %r12, %r13
	br i1 %r14, label %L9, label %L10

L9:
	%r15 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r16 = getelementptr %struct.ListNode, %struct.ListNode* %r15, i32 0, i32 2
	%r17 = load i64, i64* %r16
	store i64 %r17, i64* %find_retval
	%r18 = load i64, i64* %find_retval
	ret i64 %r18

L10:
	%r19 = load i64, i64* %P_num
	%r20 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r21 = getelementptr %struct.ListNode, %struct.ListNode* %r20, i32 0, i32 0
	%r22 = load %struct.ListNode*, %struct.ListNode** %r21
	%r23 = call i64 @find(i64 %r19, %struct.ListNode* %r22)
	store i64 %r23, i64* %find_retval
	%r24 = load i64, i64* %find_retval
	ret i64 %r24

L11:
	br label %L9

L12:
	br label %L9

L13:
	%r25 = mul i64 -1, 1
	store i64 %r25, i64* %find_retval
	%r26 = load i64, i64* %find_retval
	ret i64 %r26


}

define %struct.ListNode* @add(i64 %lookup, i64 %value, %struct.ListNode* %list)
{
L14:
	%add_retval = alloca %struct.ListNode*
	%P_lookup = alloca i64
	store i64 %lookup, i64* %P_lookup
	%P_value = alloca i64
	store i64 %value, i64* %P_value
	%P_list = alloca %struct.ListNode*
	store %struct.ListNode* %list, %struct.ListNode** %P_list
	br label %L15

L15:
	%r27 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r28 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	%r29 = icmp eq %struct.ListNode* %r27, %r28
	br i1 %r29, label %L16, label %L17

L16:
	%r30 = call i8* @malloc(i32 24)
	%r31 = bitcast i8* %r30 to %struct.ListNode*
	store %struct.ListNode* %r31, %struct.ListNode** %P_list
	%r32 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r33 = getelementptr %struct.ListNode, %struct.ListNode* %r32, i32 0, i32 1
	%r34 = load i64, i64* %P_lookup
	store i64 %r34, i64* %r33
	%r35 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r36 = getelementptr %struct.ListNode, %struct.ListNode* %r35, i32 0, i32 2
	%r37 = load i64, i64* %P_value
	store i64 %r37, i64* %r36
	%r38 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r39 = getelementptr %struct.ListNode, %struct.ListNode* %r38, i32 0, i32 0
	%r40 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	store %struct.ListNode* %r40, %struct.ListNode** %r39
	br label %L18

L17:
	%r41 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r42 = getelementptr %struct.ListNode, %struct.ListNode* %r41, i32 0, i32 0
	%r43 = load i64, i64* %P_lookup
	%r44 = load i64, i64* %P_value
	%r45 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r46 = getelementptr %struct.ListNode, %struct.ListNode* %r45, i32 0, i32 0
	%r47 = load %struct.ListNode*, %struct.ListNode** %r46
	%r48 = call %struct.ListNode* @add(i64 %r43, i64 %r44, %struct.ListNode* %r47)
	store %struct.ListNode* %r48, %struct.ListNode** %r42
	br label %L18

L18:
	%r49 = load %struct.ListNode*, %struct.ListNode** %P_list
	store %struct.ListNode* %r49, %struct.ListNode** %add_retval
	%r50 = load %struct.ListNode*, %struct.ListNode** %add_retval
	ret %struct.ListNode* %r50


}

define i64 @factorial(i64 %num, %struct.ListNode* %list)
{
L19:
	%factorial_retval = alloca i64
	%lookup = alloca i64
	%fact = alloca i64
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	%P_list = alloca %struct.ListNode*
	store %struct.ListNode* %list, %struct.ListNode** %P_list
	%r51 = call i64 @timesone(i64 100)
	br label %L20

L20:
	%r52 = load i64, i64* %P_num
	%r53 = icmp eq i64 %r52, 1
	br i1 %r53, label %L21, label %L22

L21:
	store i64 1, i64* %factorial_retval
	%r54 = load i64, i64* %factorial_retval
	ret i64 %r54

L22:
	%r55 = load i64, i64* %P_num
	%r56 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r57 = call i64 @find(i64 %r55, %struct.ListNode* %r56)
	store i64 %r57, i64* %lookup
	br label %L23

L23:
	%r58 = mul i64 -1, 1
	%r59 = load i64, i64* %lookup
	%r60 = icmp ne i64 %r59, %r58
	br i1 %r60, label %L24, label %L25

L24:
	%r61 = load i64, i64* %lookup
	store i64 %r61, i64* %factorial_retval
	%r62 = load i64, i64* %factorial_retval
	ret i64 %r62

L25:
	br label %L26

L26:
	%r63 = load i64, i64* %P_num
	%r64 = sub i64 %r63, 1
	%r65 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r66 = call i64 @factorial(i64 %r64, %struct.ListNode* %r65)
	%r67 = load i64, i64* %P_num
	%r68 = mul i64 %r67, %r66
	store i64 %r68, i64* %fact
	br label %L27

L27:
	%r69 = load i64, i64* %fact
	%r70 = sdiv i64 %r69, 3
	%r71 = icmp eq i64 %r70, 0
	br i1 %r71, label %L28, label %L29

L28:
	%r72 = load i64, i64* %P_num
	%r73 = load i64, i64* %fact
	%r74 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r75 = call %struct.ListNode* @add(i64 %r72, i64 %r73, %struct.ListNode* %r74)
	br label %L30

L29:
	br label %L30

L30:
	%r76 = load i64, i64* %fact
	store i64 %r76, i64* %factorial_retval
	%r77 = load i64, i64* %factorial_retval
	ret i64 %r77

L31:
	br label %L32

L32:
	%r78 = load i64, i64* %factorial_retval
	ret i64 %r78


}

define i64 @maxfactorial(%struct.ListNode* %list, i64 %max)
{
L33:
	%maxfactorial_retval = alloca i64
	%row = alloca %struct.Row*
	%cell = alloca %struct.Cell*
	%fact = alloca i64
	%P_list = alloca %struct.ListNode*
	store %struct.ListNode* %list, %struct.ListNode** %P_list
	%P_max = alloca i64
	store i64 %max, i64* %P_max
	%r79 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r80 = getelementptr %struct.ListNode, %struct.ListNode* %r79, i32 0, i32 0
	%r81 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	store %struct.ListNode* %r81, %struct.ListNode** %r80
	%r82 = load %struct.Row*, %struct.Row** @matrix
	store %struct.Row* %r82, %struct.Row** %row
	br label %L34

L34:
	%r83 = load %struct.Row*, %struct.Row** %row
	%r84 = load %struct.Row*, %struct.Row** @.nilRow
	%r85 = icmp ne %struct.Row* %r83, %r84
	br i1 %r85, label %L35, label %L43

L35:
	%r86 = load %struct.Row*, %struct.Row** %row
	%r87 = getelementptr %struct.Row, %struct.Row* %r86, i32 0, i32 1
	%r88 = load %struct.Cell*, %struct.Cell** %r87
	store %struct.Cell* %r88, %struct.Cell** %cell
	%r89 = load %struct.Row*, %struct.Row** %row
	%r90 = getelementptr %struct.Row, %struct.Row* %r89, i32 0, i32 0
	%r91 = load %struct.Row*, %struct.Row** %r90
	store %struct.Row* %r91, %struct.Row** %row
	br label %L36

L36:
	%r92 = load %struct.Cell*, %struct.Cell** %cell
	%r93 = load %struct.Cell*, %struct.Cell** @.nilCell
	%r94 = icmp ne %struct.Cell* %r92, %r93
	br i1 %r94, label %L37, label %L42

L37:
	%r95 = load %struct.Cell*, %struct.Cell** %cell
	%r96 = getelementptr %struct.Cell, %struct.Cell* %r95, i32 0, i32 1
	%r97 = load i64, i64* %r96
	%r98 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r99 = call i64 @factorial(i64 %r97, %struct.ListNode* %r98)
	store i64 %r99, i64* %fact
	%r100 = load %struct.Cell*, %struct.Cell** %cell
	%r101 = getelementptr %struct.Cell, %struct.Cell* %r100, i32 0, i32 0
	%r102 = load %struct.Cell*, %struct.Cell** %r101
	store %struct.Cell* %r102, %struct.Cell** %cell
	br label %L38

L38:
	%r103 = load i64, i64* %fact
	%r104 = load i64, i64* %P_max
	%r105 = icmp sgt i64 %r103, %r104
	br i1 %r105, label %L39, label %L40

L39:
	%r106 = load i64, i64* %fact
	store i64 %r106, i64* %P_max
	br label %L41

L40:
	br label %L41

L41:
	br label %L36

L42:
	br label %L34

L43:
	%r107 = load i64, i64* %P_max
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([18 x i8], [18 x i8]* @.fstr1, i32 0, i32 0), i64 %r107)
	br label %L44

L44:
	store i64 0, i64* %maxfactorial_retval
	%r108 = load i64, i64* %maxfactorial_retval
	ret i64 %r108


}

define i64 @newvalue(i64 %height, i64 %width)
{
L45:
	%newvalue_retval = alloca i64
	%position = alloca i64
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r109 = call i64 @timesone(i64 0)
	%r110 = load i64, i64* %P_height
	%r111 = load i64, i64* %P_width
	%r112 = mul i64 %r110, %r111
	store i64 %r112, i64* %position
	%r113 = load i64, i64* @maxrange
	%r114 = load i64, i64* %position
	%r115 = sdiv i64 %r113, %r114
	%r116 = load i64, i64* %P_height
	%r117 = add i64 %r115, %r116
	store i64 %r117, i64* %newvalue_retval
	%r118 = load i64, i64* %newvalue_retval
	ret i64 %r118


}

define %struct.Cell* @newcell(%struct.Cell* %cell, i64 %height, i64 %width)
{
L46:
	%newcell_retval = alloca %struct.Cell*
	%P_cell = alloca %struct.Cell*
	store %struct.Cell* %cell, %struct.Cell** %P_cell
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r119 = load %struct.Cell*, %struct.Cell** %P_cell
	%r120 = getelementptr %struct.Cell, %struct.Cell* %r119, i32 0, i32 1
	%r121 = load i64, i64* %P_height
	%r122 = load i64, i64* %P_width
	%r123 = call i64 @newvalue(i64 %r121, i64 %r122)
	store i64 %r123, i64* %r120
	br label %L47

L47:
	%r124 = load i64, i64* %P_width
	%r125 = icmp sgt i64 %r124, 1
	br i1 %r125, label %L48, label %L49

L48:
	%r126 = load %struct.Cell*, %struct.Cell** %P_cell
	%r127 = getelementptr %struct.Cell, %struct.Cell* %r126, i32 0, i32 0
	%r128 = call i8* @malloc(i32 16)
	%r129 = bitcast i8* %r128 to %struct.Cell*
	%r130 = load i64, i64* %P_height
	%r131 = load i64, i64* %P_width
	%r132 = sub i64 %r131, 1
	%r133 = call %struct.Cell* @newcell(%struct.Cell* %r129, i64 %r130, i64 %r132)
	store %struct.Cell* %r133, %struct.Cell** %r127
	br label %L50

L49:
	%r134 = load %struct.Cell*, %struct.Cell** %P_cell
	%r135 = getelementptr %struct.Cell, %struct.Cell* %r134, i32 0, i32 0
	%r136 = load %struct.Cell*, %struct.Cell** @.nilCell
	store %struct.Cell* %r136, %struct.Cell** %r135
	br label %L50

L50:
	%r137 = load %struct.Cell*, %struct.Cell** %P_cell
	store %struct.Cell* %r137, %struct.Cell** %newcell_retval
	%r138 = load %struct.Cell*, %struct.Cell** %newcell_retval
	ret %struct.Cell* %r138


}

define %struct.Row* @newrow(%struct.Row* %row, i64 %height, i64 %width)
{
L51:
	%newrow_retval = alloca %struct.Row*
	%P_row = alloca %struct.Row*
	store %struct.Row* %row, %struct.Row** %P_row
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r139 = load %struct.Row*, %struct.Row** %P_row
	%r140 = getelementptr %struct.Row, %struct.Row* %r139, i32 0, i32 1
	%r141 = call i8* @malloc(i32 16)
	%r142 = bitcast i8* %r141 to %struct.Cell*
	%r143 = load i64, i64* %P_height
	%r144 = load i64, i64* %P_width
	%r145 = call %struct.Cell* @newcell(%struct.Cell* %r142, i64 %r143, i64 %r144)
	store %struct.Cell* %r145, %struct.Cell** %r140
	br label %L52

L52:
	%r146 = load i64, i64* %P_height
	%r147 = icmp sgt i64 %r146, 1
	br i1 %r147, label %L53, label %L54

L53:
	%r148 = load %struct.Row*, %struct.Row** %P_row
	%r149 = getelementptr %struct.Row, %struct.Row* %r148, i32 0, i32 0
	%r150 = call i8* @malloc(i32 16)
	%r151 = bitcast i8* %r150 to %struct.Row*
	%r152 = load i64, i64* %P_height
	%r153 = sub i64 %r152, 1
	%r154 = load i64, i64* %P_width
	%r155 = call %struct.Row* @newrow(%struct.Row* %r151, i64 %r153, i64 %r154)
	store %struct.Row* %r155, %struct.Row** %r149
	br label %L55

L54:
	%r156 = load %struct.Row*, %struct.Row** %P_row
	%r157 = getelementptr %struct.Row, %struct.Row* %r156, i32 0, i32 0
	%r158 = load %struct.Row*, %struct.Row** @.nilRow
	store %struct.Row* %r158, %struct.Row** %r157
	br label %L55

L55:
	%r159 = load %struct.Row*, %struct.Row** %P_row
	store %struct.Row* %r159, %struct.Row** %newrow_retval
	%r160 = load %struct.Row*, %struct.Row** %newrow_retval
	ret %struct.Row* %r160


}

define i64 @newmatrix(i64 %height, i64 %width)
{
L56:
	%newmatrix_retval = alloca i64
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r161 = call i8* @malloc(i32 16)
	%r162 = bitcast i8* %r161 to %struct.Row*
	%r163 = load i64, i64* %P_height
	%r164 = load i64, i64* %P_width
	%r165 = call %struct.Row* @newrow(%struct.Row* %r162, i64 %r163, i64 %r164)
	store %struct.Row* %r165, %struct.Row** @matrix
	br label %L57

L57:
	store i64 0, i64* %newmatrix_retval
	%r166 = load i64, i64* %newmatrix_retval
	ret i64 %r166


}

define i64 @getmatrixsize(i64 %matrixsize)
{
L58:
	%getmatrixsize_retval = alloca i64
	%P_matrixsize = alloca i64
	store i64 %matrixsize, i64* %P_matrixsize
	br label %L59

L59:
	%r167 = load i64, i64* %P_matrixsize
	%r168 = icmp sle i64 %r167, 0
	br i1 %r168, label %L60, label %L61

L60:
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %P_matrixsize)
	%r169 = load i64, i64* %P_matrixsize
	%r170 = call i64 @getmatrixsize(i64 %r169)
	br label %L62

L61:
	%r171 = load i64, i64* %P_matrixsize
	store i64 %r171, i64* %getmatrixsize_retval
	%r172 = load i64, i64* %getmatrixsize_retval
	ret i64 %r172

L62:
	%r173 = load i64, i64* %P_matrixsize
	store i64 %r173, i64* %getmatrixsize_retval
	%r174 = load i64, i64* %getmatrixsize_retval
	ret i64 %r174


}

define i64 @getmaxrange(i64 %maxrange)
{
L63:
	%getmaxrange_retval = alloca i64
	%P_maxrange = alloca i64
	store i64 %maxrange, i64* %P_maxrange
	br label %L64

L64:
	%r175 = load i64, i64* @maxrange
	%r176 = icmp sle i64 %r175, 1
	br i1 %r176, label %L65, label %L66

L65:
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* @maxrange)
	%r177 = load i64, i64* @maxrange
	%r178 = call i64 @getmaxrange(i64 %r177)
	br label %L67

L66:
	%r179 = load i64, i64* @maxrange
	store i64 %r179, i64* %getmaxrange_retval
	%r180 = load i64, i64* %getmaxrange_retval
	ret i64 %r180

L67:
	%r181 = load i64, i64* @maxrange
	store i64 %r181, i64* %getmaxrange_retval
	%r182 = load i64, i64* %getmaxrange_retval
	ret i64 %r182


}

define i64 @main()
{
L68:
	%main_retval = alloca i64
	%height = alloca i64
	%width = alloca i64
	store i64 0, i64* %height
	store i64 0, i64* %width
	store i64 0, i64* @maxrange
	%r183 = load i64, i64* %height
	%r184 = call i64 @getmatrixsize(i64 %r183)
	store i64 %r184, i64* %height
	%r185 = load i64, i64* %height
	store i64 %r185, i64* %width
	%r186 = load i64, i64* @maxrange
	%r187 = call i64 @getmaxrange(i64 %r186)
	store i64 %r187, i64* @maxrange
	%r188 = load i64, i64* %height
	%r189 = load i64, i64* %width
	%r190 = call i64 @newmatrix(i64 %r188, i64 %r189)
	%r191 = call i8* @malloc(i32 24)
	%r192 = bitcast i8* %r191 to %struct.ListNode*
	%r193 = call i64 @maxfactorial(%struct.ListNode* %r192, i64 0)
	br label %L69

L69:
	store i64 0, i64* %main_retval
	%r194 = load i64, i64* %main_retval
	ret i64 %r194


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [18 x i8] c"Max Factorial=%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

