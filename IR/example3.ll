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
	%r9 = icmp ne i64 %r7, %r8
	br i1 %r9, label %L7, label %L8

L7:
	br label %L8

L8:
	%r10 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r11 = getelementptr %struct.ListNode, %struct.ListNode* %r10, i32 0, i32 1
	%r12 = load i64, i64* %P_num
	%r13 = icmp eq i64 %r11, %r12
	br i1 %r13, label %L9, label %L10

L9:
	%r14 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r15 = getelementptr %struct.ListNode, %struct.ListNode* %r14, i32 0, i32 2
	store i64 %r15, i64* %find_retval
	%r16 = load i64, i64* %find_retval
	ret i64 %r16
	br label %L11

L10:
	%r17 = load i64, i64* %P_num
	%r18 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r19 = getelementptr %struct.ListNode, %struct.ListNode* %r18, i32 0, i32 0
	%r20 = load i64, i64* %r19
	%r21 = call i64 @find(i64 %r17, %struct.ListNode %r20)
	store i64 %r21, i64* %find_retval
	%r22 = load i64, i64* %find_retval
	ret i64 %r22
	br label %L11

L11:
	br label %L9

L12:
	br label %L9

L13:
	%r23 = mul i64 -1, 1
	store i64 %r23, i64* %find_retval
	%r24 = load i64, i64* %find_retval
	ret i64 %r24
	br label %L14

L14:
	%r25 = load i64, i64* %find_retval
	ret i64 %r25


}

define %struct.ListNode* @add(i64 %lookup, i64 %value, %struct.ListNode* %list)
{
L15:
	%add_retval = alloca %struct.ListNode*
	%P_lookup = alloca i64
	store i64 %lookup, i64* %P_lookup
	%P_value = alloca i64
	store i64 %value, i64* %P_value
	%P_list = alloca %struct.ListNode*
	store %struct.ListNode* %list, %struct.ListNode** %P_list
	br label %L16

L16:
	%r26 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r27 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	%r28 = icmp eq i64 %r26, %r27
	br i1 %r28, label %L17, label %L18

L17:
	%r29 = call i8* @malloc(i32 24)
	%r30 = bitcast i8* %r29 to %struct.ListNode*
	store %struct.ListNode* %r30, %struct.ListNode** %P_list
	%r31 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r32 = getelementptr %struct.ListNode, %struct.ListNode* %r31, i32 0, i32 1
	%r33 = load i64, i64* %P_lookup
	store i64 %r33, i64* %r32
	%r34 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r35 = getelementptr %struct.ListNode, %struct.ListNode* %r34, i32 0, i32 2
	%r36 = load i64, i64* %P_value
	store i64 %r36, i64* %r35
	%r37 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r38 = getelementptr %struct.ListNode, %struct.ListNode* %r37, i32 0, i32 0
	%r39 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	store %struct.ListNode* %r39, %struct.ListNode** %r38
	br label %L19

L18:
	%r40 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r41 = getelementptr %struct.ListNode, %struct.ListNode* %r40, i32 0, i32 0
	%r42 = load i64, i64* %P_lookup
	%r43 = load i64, i64* %P_value
	%r44 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r45 = getelementptr %struct.ListNode, %struct.ListNode* %r44, i32 0, i32 0
	%r46 = load i64, i64* %r45
	%r47 = call %struct.ListNode* @add(i64 %r42, i64 %r43, %struct.ListNode %r46)
	store %struct.ListNode* %r47, %struct.ListNode** %r41
	br label %L19

L19:
	%r48 = load %struct.ListNode*, %struct.ListNode** %P_list
	store %struct.ListNode* %r48, %struct.ListNode** %add_retval
	%r49 = load %struct.ListNode*, %struct.ListNode** %add_retval
	ret %struct.ListNode* %r49
	br label %L20

L20:
	%r50 = load %struct.ListNode*, %struct.ListNode** %add_retval
	ret %struct.ListNode* %r50


}

define i64 @factorial(i64 %num, %struct.ListNode* %list)
{
L21:
	%factorial_retval = alloca i64
	%lookup = alloca i64
	%fact = alloca i64
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	%P_list = alloca %struct.ListNode*
	store %struct.ListNode* %list, %struct.ListNode** %P_list
	%r51 = call i64 @timesone(i64 %r100)
	br label %L22

L22:
	%r52 = load i64, i64* %P_num
	%r53 = icmp eq i64 %r52, 1
	br i1 %r53, label %L23, label %L24

L23:
	store i64 1, i64* %factorial_retval
	%r54 = load i64, i64* %factorial_retval
	ret i64 %r54
	br label %L25

L24:
	%r55 = load i64, i64* %P_num
	%r56 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r57 = call i64 @find(i64 %r55, %struct.ListNode %r56)
	store i64 %r57, i64* %lookup
	br label %L25

L25:
	%r58 = load i64, i64* %lookup
	%r59 = mul i64 -1, 1
	%r60 = icmp ne i64 %r58, %r59
	br i1 %r60, label %L26, label %L27

L26:
	%r61 = load i64, i64* %lookup
	store i64 %r61, i64* %factorial_retval
	%r62 = load i64, i64* %factorial_retval
	ret i64 %r62
	br label %L28

L27:
	br label %L28

L28:
	%r63 = load i64, i64* %P_num
	%r64 = load i64, i64* %P_num
	%r65 = sub i64 %r64, 1
	%r66 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r67 = call i64 @factorial(i64 %r65, %struct.ListNode %r66)
	%r68 = mul i64 %r63, %r67
	store i64 %r68, i64* %fact
	br label %L29

L29:
	%r69 = load i64, i64* %fact
	%r70 = sdiv i64 %r69, 3
	%r71 = icmp eq i64 %r70, 0
	br i1 %r71, label %L30, label %L31

L30:
	%r72 = load i64, i64* %P_num
	%r73 = load i64, i64* %fact
	%r74 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r75 = call %struct.ListNode* @add(i64 %r72, i64 %r73, %struct.ListNode %r74)
	br label %L32

L31:
	br label %L32

L32:
	%r76 = load i64, i64* %fact
	store i64 %r76, i64* %factorial_retval
	%r77 = load i64, i64* %factorial_retval
	ret i64 %r77
	br label %L25

L33:
	br label %L34

L34:
	%r78 = load i64, i64* %factorial_retval
	ret i64 %r78


}

define i64 @maxfactorial(%struct.ListNode* %list, i64 %max)
{
L35:
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
	br label %L36

L36:
	%r83 = load %struct.Row*, %struct.Row** %row
	%r84 = load %struct.Row*, %struct.Row** @.nilRow
	%r85 = icmp ne i64 %r83, %r84
	br i1 %r85, label %L37, label %L38

L37:
	%r86 = load %struct.Row*, %struct.Row** %row
	%r87 = getelementptr %struct.Row, %struct.Row* %r86, i32 0, i32 1
	store %struct.Cell* %r87, %struct.Cell** %cell
	%r88 = load %struct.Row*, %struct.Row** %row
	%r89 = getelementptr %struct.Row, %struct.Row* %r88, i32 0, i32 0
	store %struct.Row* %r89, %struct.Row** %row
	br label %L38

L38:
	%r90 = load %struct.Cell*, %struct.Cell** %cell
	%r91 = load %struct.Cell*, %struct.Cell** @.nilCell
	%r92 = icmp ne i64 %r90, %r91
	br i1 %r92, label %L39, label %L40

L39:
	%r93 = load %struct.Cell*, %struct.Cell** %cell
	%r94 = getelementptr %struct.Cell, %struct.Cell* %r93, i32 0, i32 1
	%r95 = load i64, i64* %r94
	%r96 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r97 = call i64 @factorial(i64 %r95, %struct.ListNode %r96)
	store i64 %r97, i64* %fact
	%r98 = load %struct.Cell*, %struct.Cell** %cell
	%r99 = getelementptr %struct.Cell, %struct.Cell* %r98, i32 0, i32 0
	store %struct.Cell* %r99, %struct.Cell** %cell
	br label %L40

L40:
	%r100 = load i64, i64* %fact
	%r101 = load i64, i64* %P_max
	%r102 = icmp sgt i64 %r100, %r101
	br i1 %r102, label %L41, label %L42

L41:
	%r103 = load i64, i64* %fact
	store i64 %r103, i64* %P_max
	br label %L43

L42:
	br label %L43

L43:
	br label %L38

L44:
	br label %L36

L45:
	%r104 = load i64, i64* %P_max
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([18 x i8], [18 x i8]* @.fstr1, i32 0, i32 0), i64 %r104)
	br label %L46

L46:
	store i64 0, i64* %maxfactorial_retval
	%r105 = load i64, i64* %maxfactorial_retval
	ret i64 %r105


}

define i64 @newvalue(i64 %height, i64 %width)
{
L47:
	%newvalue_retval = alloca i64
	%position = alloca i64
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r106 = call i64 @timesone(i64 %r0)
	%r107 = load i64, i64* %P_height
	%r108 = load i64, i64* %P_width
	%r109 = mul i64 %r107, %r108
	store i64 %r109, i64* %position
	%r110 = load i64, i64* @maxrange
	%r111 = load i64, i64* %position
	%r112 = sdiv i64 %r110, %r111
	%r113 = load i64, i64* %P_height
	%r114 = add i64 %r112, %r113
	store i64 %r114, i64* %newvalue_retval
	%r115 = load i64, i64* %newvalue_retval
	ret i64 %r115
	br label %L48

L48:
	%r116 = load i64, i64* %newvalue_retval
	ret i64 %r116


}

define %struct.Cell* @newcell(%struct.Cell* %cell, i64 %height, i64 %width)
{
L49:
	%newcell_retval = alloca %struct.Cell*
	%P_cell = alloca %struct.Cell*
	store %struct.Cell* %cell, %struct.Cell** %P_cell
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r117 = load %struct.Cell*, %struct.Cell** %P_cell
	%r118 = getelementptr %struct.Cell, %struct.Cell* %r117, i32 0, i32 1
	%r119 = load i64, i64* %P_height
	%r120 = load i64, i64* %P_width
	%r121 = call i64 @newvalue(i64 %r119, i64 %r120)
	store i64 %r121, i64* %r118
	br label %L50

L50:
	%r122 = load i64, i64* %P_width
	%r123 = icmp sgt i64 %r122, 1
	br i1 %r123, label %L51, label %L52

L51:
	%r124 = load %struct.Cell*, %struct.Cell** %P_cell
	%r125 = getelementptr %struct.Cell, %struct.Cell* %r124, i32 0, i32 0
	%r126 = call i8* @malloc(i32 16)
	%r127 = bitcast i8* %r126 to %struct.Cell*
	%r128 = load i64, i64* %P_height
	%r129 = load i64, i64* %P_width
	%r130 = sub i64 %r129, 1
	%r131 = call %struct.Cell* @newcell(%struct.Cell %r127, i64 %r128, i64 %r130)
	store %struct.Cell* %r131, %struct.Cell** %r125
	br label %L53

L52:
	%r132 = load %struct.Cell*, %struct.Cell** %P_cell
	%r133 = getelementptr %struct.Cell, %struct.Cell* %r132, i32 0, i32 0
	%r134 = load %struct.Cell*, %struct.Cell** @.nilCell
	store %struct.Cell* %r134, %struct.Cell** %r133
	br label %L53

L53:
	%r135 = load %struct.Cell*, %struct.Cell** %P_cell
	store %struct.Cell* %r135, %struct.Cell** %newcell_retval
	%r136 = load %struct.Cell*, %struct.Cell** %newcell_retval
	ret %struct.Cell* %r136
	br label %L54

L54:
	%r137 = load %struct.Cell*, %struct.Cell** %newcell_retval
	ret %struct.Cell* %r137


}

define %struct.Row* @newrow(%struct.Row* %row, i64 %height, i64 %width)
{
L55:
	%newrow_retval = alloca %struct.Row*
	%P_row = alloca %struct.Row*
	store %struct.Row* %row, %struct.Row** %P_row
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r138 = load %struct.Row*, %struct.Row** %P_row
	%r139 = getelementptr %struct.Row, %struct.Row* %r138, i32 0, i32 1
	%r140 = call i8* @malloc(i32 16)
	%r141 = bitcast i8* %r140 to %struct.Cell*
	%r142 = load i64, i64* %P_height
	%r143 = load i64, i64* %P_width
	%r144 = call %struct.Cell* @newcell(%struct.Cell %r141, i64 %r142, i64 %r143)
	store %struct.Cell* %r144, %struct.Cell** %r139
	br label %L56

L56:
	%r145 = load i64, i64* %P_height
	%r146 = icmp sgt i64 %r145, 1
	br i1 %r146, label %L57, label %L58

L57:
	%r147 = load %struct.Row*, %struct.Row** %P_row
	%r148 = getelementptr %struct.Row, %struct.Row* %r147, i32 0, i32 0
	%r149 = call i8* @malloc(i32 16)
	%r150 = bitcast i8* %r149 to %struct.Row*
	%r151 = load i64, i64* %P_height
	%r152 = sub i64 %r151, 1
	%r153 = load i64, i64* %P_width
	%r154 = call %struct.Row* @newrow(%struct.Row %r150, i64 %r152, i64 %r153)
	store %struct.Row* %r154, %struct.Row** %r148
	br label %L59

L58:
	%r155 = load %struct.Row*, %struct.Row** %P_row
	%r156 = getelementptr %struct.Row, %struct.Row* %r155, i32 0, i32 0
	%r157 = load %struct.Row*, %struct.Row** @.nilRow
	store %struct.Row* %r157, %struct.Row** %r156
	br label %L59

L59:
	%r158 = load %struct.Row*, %struct.Row** %P_row
	store %struct.Row* %r158, %struct.Row** %newrow_retval
	%r159 = load %struct.Row*, %struct.Row** %newrow_retval
	ret %struct.Row* %r159
	br label %L60

L60:
	%r160 = load %struct.Row*, %struct.Row** %newrow_retval
	ret %struct.Row* %r160


}

define i64 @newmatrix(i64 %height, i64 %width)
{
L61:
	%newmatrix_retval = alloca i64
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r161 = call i8* @malloc(i32 16)
	%r162 = bitcast i8* %r161 to %struct.Row*
	%r163 = load i64, i64* %P_height
	%r164 = load i64, i64* %P_width
	%r165 = call %struct.Row* @newrow(%struct.Row %r162, i64 %r163, i64 %r164)
	store %struct.Row* %r165, %struct.Row** @matrix
	br label %L62

L62:
	store i64 0, i64* %newmatrix_retval
	%r166 = load i64, i64* %newmatrix_retval
	ret i64 %r166


}

define i64 @getmatrixsize(i64 %matrixsize)
{
L63:
	%getmatrixsize_retval = alloca i64
	%P_matrixsize = alloca i64
	store i64 %matrixsize, i64* %P_matrixsize
	br label %L64

L64:
	%r167 = load i64, i64* %P_matrixsize
	%r168 = icmp sle i64 %r167, 0
	br i1 %r168, label %L65, label %L66

L65:
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %P_matrixsize)
	%r169 = load i64, i64* %P_matrixsize
	%r170 = call i64 @getmatrixsize(i64 %r169)
	br label %L67

L66:
	%r171 = load i64, i64* %P_matrixsize
	store i64 %r171, i64* %getmatrixsize_retval
	%r172 = load i64, i64* %getmatrixsize_retval
	ret i64 %r172
	br label %L67

L67:
	%r173 = load i64, i64* %P_matrixsize
	store i64 %r173, i64* %getmatrixsize_retval
	%r174 = load i64, i64* %getmatrixsize_retval
	ret i64 %r174
	br label %L68

L68:
	%r175 = load i64, i64* %getmatrixsize_retval
	ret i64 %r175


}

define i64 @getmaxrange(i64 %maxrange)
{
L69:
	%getmaxrange_retval = alloca i64
	%P_maxrange = alloca i64
	store i64 %maxrange, i64* %P_maxrange
	br label %L70

L70:
	%r176 = load i64, i64* @maxrange
	%r177 = icmp sle i64 %r176, 1
	br i1 %r177, label %L71, label %L72

L71:
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* @maxrange)
	%r178 = load i64, i64* @maxrange
	%r179 = call i64 @getmaxrange(i64 %r178)
	br label %L73

L72:
	%r180 = load i64, i64* @maxrange
	store i64 %r180, i64* %getmaxrange_retval
	%r181 = load i64, i64* %getmaxrange_retval
	ret i64 %r181
	br label %L73

L73:
	%r182 = load i64, i64* @maxrange
	store i64 %r182, i64* %getmaxrange_retval
	%r183 = load i64, i64* %getmaxrange_retval
	ret i64 %r183
	br label %L74

L74:
	%r184 = load i64, i64* %getmaxrange_retval
	ret i64 %r184


}

define i64 @main()
{
L75:
	%main_retval = alloca i64
	%height = alloca i64
	%width = alloca i64
	store i64 0, i64* %height
	store i64 0, i64* %width
	store i64 0, i64* @maxrange
	%r185 = load i64, i64* %height
	%r186 = call i64 @getmatrixsize(i64 %r185)
	store i64 %r186, i64* %height
	%r187 = load i64, i64* %height
	store i64 %r187, i64* %width
	%r188 = load i64, i64* @maxrange
	%r189 = call i64 @getmaxrange(i64 %r188)
	store i64 %r189, i64* @maxrange
	%r190 = load i64, i64* %height
	%r191 = load i64, i64* %width
	%r192 = call i64 @newmatrix(i64 %r190, i64 %r191)
	%r193 = call i8* @malloc(i32 24)
	%r194 = bitcast i8* %r193 to %struct.ListNode*
	%r195 = call i64 @maxfactorial(%struct.ListNode %r194, i64 %r0)
	br label %L76

L76:
	store i64 0, i64* %main_retval
	%r196 = load i64, i64* %main_retval
	ret i64 %r196


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [18 x i8] c"Max Factorial=%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

