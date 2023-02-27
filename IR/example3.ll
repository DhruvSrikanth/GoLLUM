source_filename = "example3"
target triple = "arm64-apple-darwin22.2.0"

%struct.Cell = type {%struct.Cell*, i64}
%struct.Row = type {%struct.Row*, struct.Cell}
%struct.ListNode = type {%struct.ListNode*, i64, i64}

@.nilCell = common global %struct.Cell* null
@.nilRow = common global %struct.Row* null
@.nilListNode = common global %struct.ListNode* null
@matrix = common global %struct.Row* null
@maxrange = common global i64 0

define i64 @timesone(i64 %iter)
{
L0:
	%_retval = alloca i64
	%P_iter = alloca i64
	store i64 %iter, i64* %P_iter
	br label %L1

L1:
	%r0 = load i64, i64* %P_iter
	store i64 0, i64* %r1
	%r2 = icmp sgt i64 %r0, %r1
	br i1 %r2, label %L2, label %L3

L2:
	%r3 = load i64, i64* @maxrange
	%r4 = load i64, i64* @maxrange
	store i64 1, i64* %r5
	%r6 = mul i64 %r4, %r5
	store i64 %r6, i64* %r3
	%r7 = load i64, i64* %P_iter
	%r8 = load i64, i64* %P_iter
	store i64 1, i64* %r9
	%r10 = sub i64 %r8, %r9
	store i64 %r10, i64* %r7
	br label %L1

L3:
	br label %L4

L4:
	store i64 0, i64* %_retval
	%r11 = load i64, i64* %r10
	ret i64 %r11


}

define i64 @find(i64 %num, %struct.ListNode* %list)
{
L5:
	%_retval = alloca i64
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	%P_list = alloca %struct.ListNode*
	store %struct.ListNode %list, %struct.ListNode* %P_list
	br label %L6

L6:
	%r12 = load %struct.ListNode*, %struct.ListNode** %P_list
	store i64 null, i64* %r13
	%r14 = icmp ne i64 %r12, %r13
	br i1 %r14, label %L7, label %L8

L7:
	br label %L8

L8:
	%r15 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r16 = getelementptr %struct.ListNode*, %struct.ListNode** %r15, i32 0, i32 1
	%r17 = load i64, i64* %P_num
	%r18 = icmp eq i64 %r16, %r17
	br i1 %r18, label %L9, label %L10

L9:
	%r19 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r20 = getelementptr %struct.ListNode*, %struct.ListNode** %r19, i32 0, i32 2
	store i64 %r20, i64* %_retval
	%r21 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L11

L10:
	%r22 = load i64, i64* %P_num
	%r23 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r24 = getelementptr %struct.ListNode*, %struct.ListNode** %r23, i32 0, i32 0
	%r25 = call i64 @find(i64 %r22, %struct.ListNode %r24)
	store i64 %r25, i64* %_retval
	%r26 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L11

L11:
	br label %L9

L12:
	br label %L9

L13:
	store i64 -1, i64* %r27
	store i64 1, i64* %r28
	%r29 = mul i64 %r27, %r28
	store i64 %r29, i64* %_retval
	%r30 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L14

L14:
	%r31 = load i64, i64* %_retval
	ret i64 %r31


}

define %struct.ListNode* @add(i64 %lookup, i64 %value, %struct.ListNode* %list)
{
L15:
	%_retval = alloca %struct.ListNode*
	%P_lookup = alloca i64
	store i64 %lookup, i64* %P_lookup
	%P_value = alloca i64
	store i64 %value, i64* %P_value
	%P_list = alloca %struct.ListNode*
	store %struct.ListNode %list, %struct.ListNode* %P_list
	br label %L16

L16:
	%r32 = load %struct.ListNode*, %struct.ListNode** %P_list
	store i64 null, i64* %r33
	%r34 = icmp eq i64 %r32, %r33
	br i1 %r34, label %L17, label %L18

L17:
	%r35 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r36 = call i8* @malloc(i32 24)
	%r37 = bitcast i8* %r36 to %struct.ListNode*
	store %struct.ListNode* %r37, %struct.ListNode** %r35
	%r38 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r39 = getelementptr %struct.ListNode*, %struct.ListNode** %r38, i32 0, i32 1
	%r40 = load i64, i64* %P_lookup
	store i64 %r40, i64* %r39
	%r41 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r42 = getelementptr %struct.ListNode*, %struct.ListNode** %r41, i32 0, i32 2
	%r43 = load i64, i64* %P_value
	store i64 %r43, i64* %r42
	%r44 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r45 = getelementptr %struct.ListNode*, %struct.ListNode** %r44, i32 0, i32 0
	%r46 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	store %struct.ListNode* %r46, %struct.ListNode** %r45
	br label %L19

L18:
	%r47 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r48 = getelementptr %struct.ListNode*, %struct.ListNode** %r47, i32 0, i32 0
	%r49 = load i64, i64* %P_lookup
	%r50 = load i64, i64* %P_value
	%r51 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r52 = getelementptr %struct.ListNode*, %struct.ListNode** %r51, i32 0, i32 0
	%r53 = call %struct.ListNode* @add(i64 %r49, i64 %r50, %struct.ListNode %r52)
	store %struct.ListNode* %r53, %struct.ListNode** %r48
	br label %L19

L19:
	%r54 = load %struct.ListNode*, %struct.ListNode** %P_list
	store %struct.ListNode* %r54, %struct.ListNode** %_retval
	%r55 = load %struct.ListNode*, %struct.ListNode** %_retval
	ret %struct.ListNode* %_retval
	br label %L20

L20:
	%r56 = load %struct.ListNode, %struct.ListNode* %_retval
	ret %struct.ListNode %r56


}

define i64 @factorial(i64 %num, %struct.ListNode* %list)
{
L21:
	%_retval = alloca i64
	%lookup = alloca i64
	%fact = alloca i64
	%P_num = alloca i64
	store i64 %num, i64* %P_num
	%P_list = alloca %struct.ListNode*
	store %struct.ListNode %list, %struct.ListNode* %P_list
	store i64 100, i64* %r57
	%r58 = call i64 @timesone(i64 %r57)
	br label %L22

L22:
	%r59 = load i64, i64* %P_num
	store i64 1, i64* %r60
	%r61 = icmp eq i64 %r59, %r60
	br i1 %r61, label %L23, label %L24

L23:
	store i64 1, i64* %r62
	store i64 %r62, i64* %_retval
	%r63 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L25

L24:
	%r64 = load i64, i64* %lookup
	%r65 = load i64, i64* %P_num
	%r66 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r67 = call i64 @find(i64 %r65, %struct.ListNode %r66)
	store i64 %r67, i64* %r64
	br label %L25

L25:
	%r68 = load i64, i64* %lookup
	store i64 -1, i64* %r69
	store i64 1, i64* %r70
	%r71 = mul i64 %r69, %r70
	%r72 = icmp ne i64 %r68, %r71
	br i1 %r72, label %L26, label %L27

L26:
	%r73 = load i64, i64* %lookup
	store i64 %r73, i64* %_retval
	%r74 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L28

L27:
	br label %L28

L28:
	%r75 = load i64, i64* %fact
	%r76 = load i64, i64* %P_num
	%r77 = load i64, i64* %P_num
	store i64 1, i64* %r78
	%r79 = sub i64 %r77, %r78
	%r80 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r81 = call i64 @factorial(i64 %r79, %struct.ListNode %r80)
	%r82 = mul i64 %r76, %r81
	store i64 %r82, i64* %r75
	br label %L29

L29:
	%r83 = load i64, i64* %fact
	store i64 3, i64* %r84
	%r85 = sdiv i64 %r83, %r84
	store i64 0, i64* %r86
	%r87 = icmp eq i64 %r85, %r86
	br i1 %r87, label %L30, label %L31

L30:
	%r88 = load i64, i64* %P_num
	%r89 = load i64, i64* %fact
	%r90 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r91 = call %struct.ListNode* @add(i64 %r88, i64 %r89, %struct.ListNode %r90)
	br label %L32

L31:
	br label %L32

L32:
	%r92 = load i64, i64* %fact
	store i64 %r92, i64* %_retval
	%r93 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L25

L33:
	br label %L34

L34:
	%r94 = load i64, i64* %_retval
	ret i64 %r94


}

define i64 @maxfactorial(%struct.ListNode* %list, i64 %max)
{
L35:
	%_retval = alloca i64
	%row = alloca %struct.Row*
	%cell = alloca %struct.Cell*
	%fact = alloca i64
	%P_list = alloca %struct.ListNode*
	store %struct.ListNode %list, %struct.ListNode* %P_list
	%P_max = alloca i64
	store i64 %max, i64* %P_max
	%r95 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r96 = getelementptr %struct.ListNode*, %struct.ListNode** %r95, i32 0, i32 0
	%r97 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	store %struct.ListNode* %r97, %struct.ListNode** %r96
	%r98 = load %struct.Row*, %struct.Row** %row
	%r99 = load %struct.Row*, %struct.Row** @matrix
	store %struct.Row* %r99, %struct.Row** %r98
	br label %L36

L36:
	%r100 = load %struct.Row*, %struct.Row** %row
	store i64 null, i64* %r101
	%r102 = icmp ne i64 %r100, %r101
	br i1 %r102, label %L37, label %L38

L37:
	%r103 = load %struct.Cell*, %struct.Cell** %cell
	%r104 = load %struct.Row*, %struct.Row** %row
	%r105 = getelementptr %struct.Row*, %struct.Row** %r104, i32 0, i32 1
	store %struct.Cell* %r105, %struct.Cell** %r103
	%r106 = load %struct.Row*, %struct.Row** %row
	%r107 = load %struct.Row*, %struct.Row** %row
	%r108 = getelementptr %struct.Row*, %struct.Row** %r107, i32 0, i32 0
	store %struct.Row* %r108, %struct.Row** %r106
	br label %L38

L38:
	%r109 = load %struct.Cell*, %struct.Cell** %cell
	store i64 null, i64* %r110
	%r111 = icmp ne i64 %r109, %r110
	br i1 %r111, label %L39, label %L40

L39:
	%r112 = load i64, i64* %fact
	%r113 = load %struct.Cell*, %struct.Cell** %cell
	%r114 = getelementptr %struct.Cell*, %struct.Cell** %r113, i32 0, i32 1
	%r115 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r116 = call i64 @factorial(i64 %r114, %struct.ListNode %r115)
	store i64 %r116, i64* %r112
	%r117 = load %struct.Cell*, %struct.Cell** %cell
	%r118 = load %struct.Cell*, %struct.Cell** %cell
	%r119 = getelementptr %struct.Cell*, %struct.Cell** %r118, i32 0, i32 0
	store %struct.Cell* %r119, %struct.Cell** %r117
	br label %L40

L40:
	%r120 = load i64, i64* %fact
	%r121 = load i64, i64* %P_max
	%r122 = icmp sgt i64 %r120, %r121
	br i1 %r122, label %L41, label %L42

L41:
	%r123 = load i64, i64* %P_max
	%r124 = load i64, i64* %fact
	store i64 %r124, i64* %r123
	br label %L43

L42:
	br label %L43

L43:
	br label %L38

L44:
	br label %L36

L45:
	%r125 = load i64, i64* %P_max
	%r126 = load i64, i64* %r125
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([17 x i8], [17 x i8]* @.fstr1, i32 0, i32 0), i64 %r126)
	br label %L46

L46:
	store i64 0, i64* %_retval
	%r127 = load i64, i64* %r126
	ret i64 %r127


}

define i64 @newvalue(i64 %height, i64 %width)
{
L47:
	%_retval = alloca i64
	%position = alloca i64
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	store i64 0, i64* %r128
	%r129 = call i64 @timesone(i64 %r128)
	%r130 = load i64, i64* %position
	%r131 = load i64, i64* %P_height
	%r132 = load i64, i64* %P_width
	%r133 = mul i64 %r131, %r132
	store i64 %r133, i64* %r130
	%r134 = load i64, i64* @maxrange
	%r135 = load i64, i64* %position
	%r136 = sdiv i64 %r134, %r135
	%r137 = load i64, i64* %P_height
	%r138 = add i64 %r136, %r137
	store i64 %r138, i64* %_retval
	%r139 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L48

L48:
	%r140 = load i64, i64* %_retval
	ret i64 %r140


}

define %struct.Cell* @newcell(%struct.Cell* %cell, i64 %height, i64 %width)
{
L49:
	%_retval = alloca %struct.Cell*
	%P_cell = alloca %struct.Cell*
	store %struct.Cell %cell, %struct.Cell* %P_cell
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r141 = load %struct.Cell*, %struct.Cell** %P_cell
	%r142 = getelementptr %struct.Cell*, %struct.Cell** %r141, i32 0, i32 1
	%r143 = load i64, i64* %P_height
	%r144 = load i64, i64* %P_width
	%r145 = call i64 @newvalue(i64 %r143, i64 %r144)
	store i64 %r145, i64* %r142
	br label %L50

L50:
	%r146 = load i64, i64* %P_width
	store i64 1, i64* %r147
	%r148 = icmp sgt i64 %r146, %r147
	br i1 %r148, label %L51, label %L52

L51:
	%r149 = load %struct.Cell*, %struct.Cell** %P_cell
	%r150 = getelementptr %struct.Cell*, %struct.Cell** %r149, i32 0, i32 0
	%r151 = call i8* @malloc(i32 16)
	%r152 = bitcast i8* %r151 to %struct.Cell*
	%r153 = load i64, i64* %P_height
	%r154 = load i64, i64* %P_width
	store i64 1, i64* %r155
	%r156 = sub i64 %r154, %r155
	%r157 = call %struct.Cell* @newcell(%struct.Cell %r152, i64 %r153, i64 %r156)
	store %struct.Cell* %r157, %struct.Cell** %r150
	br label %L53

L52:
	%r158 = load %struct.Cell*, %struct.Cell** %P_cell
	%r159 = getelementptr %struct.Cell*, %struct.Cell** %r158, i32 0, i32 0
	%r160 = load %struct.Cell*, %struct.Cell** @.nilCell
	store %struct.Cell* %r160, %struct.Cell** %r159
	br label %L53

L53:
	%r161 = load %struct.Cell*, %struct.Cell** %P_cell
	store %struct.Cell* %r161, %struct.Cell** %_retval
	%r162 = load %struct.Cell*, %struct.Cell** %_retval
	ret %struct.Cell* %_retval
	br label %L54

L54:
	%r163 = load %struct.Cell, %struct.Cell* %_retval
	ret %struct.Cell %r163


}

define %struct.Row* @newrow(%struct.Row* %row, i64 %height, i64 %width)
{
L55:
	%_retval = alloca %struct.Row*
	%P_row = alloca %struct.Row*
	store %struct.Row %row, %struct.Row* %P_row
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r164 = load %struct.Row*, %struct.Row** %P_row
	%r165 = getelementptr %struct.Row*, %struct.Row** %r164, i32 0, i32 1
	%r166 = call i8* @malloc(i32 16)
	%r167 = bitcast i8* %r166 to %struct.Cell*
	%r168 = load i64, i64* %P_height
	%r169 = load i64, i64* %P_width
	%r170 = call %struct.Cell* @newcell(%struct.Cell %r167, i64 %r168, i64 %r169)
	store %struct.Cell* %r170, %struct.Cell** %r165
	br label %L56

L56:
	%r171 = load i64, i64* %P_height
	store i64 1, i64* %r172
	%r173 = icmp sgt i64 %r171, %r172
	br i1 %r173, label %L57, label %L58

L57:
	%r174 = load %struct.Row*, %struct.Row** %P_row
	%r175 = getelementptr %struct.Row*, %struct.Row** %r174, i32 0, i32 0
	%r176 = call i8* @malloc(i32 16)
	%r177 = bitcast i8* %r176 to %struct.Row*
	%r178 = load i64, i64* %P_height
	store i64 1, i64* %r179
	%r180 = sub i64 %r178, %r179
	%r181 = load i64, i64* %P_width
	%r182 = call %struct.Row* @newrow(%struct.Row %r177, i64 %r180, i64 %r181)
	store %struct.Row* %r182, %struct.Row** %r175
	br label %L59

L58:
	%r183 = load %struct.Row*, %struct.Row** %P_row
	%r184 = getelementptr %struct.Row*, %struct.Row** %r183, i32 0, i32 0
	%r185 = load %struct.Row*, %struct.Row** @.nilRow
	store %struct.Row* %r185, %struct.Row** %r184
	br label %L59

L59:
	%r186 = load %struct.Row*, %struct.Row** %P_row
	store %struct.Row* %r186, %struct.Row** %_retval
	%r187 = load %struct.Row*, %struct.Row** %_retval
	ret %struct.Row* %_retval
	br label %L60

L60:
	%r188 = load %struct.Row, %struct.Row* %_retval
	ret %struct.Row %r188


}

define i64 @newmatrix(i64 %height, i64 %width)
{
L61:
	%_retval = alloca i64
	%P_height = alloca i64
	store i64 %height, i64* %P_height
	%P_width = alloca i64
	store i64 %width, i64* %P_width
	%r189 = load %struct.Row*, %struct.Row** @matrix
	%r190 = call i8* @malloc(i32 16)
	%r191 = bitcast i8* %r190 to %struct.Row*
	%r192 = load i64, i64* %P_height
	%r193 = load i64, i64* %P_width
	%r194 = call %struct.Row* @newrow(%struct.Row %r191, i64 %r192, i64 %r193)
	store %struct.Row* %r194, %struct.Row** %r189
	br label %L62

L62:
	store i64 0, i64* %_retval
	%r195 = load i64, i64* %r194
	ret i64 %r195


}

define i64 @getmatrixsize(i64 %matrixsize)
{
L63:
	%_retval = alloca i64
	%P_matrixsize = alloca i64
	store i64 %matrixsize, i64* %P_matrixsize
	br label %L64

L64:
	%r196 = load i64, i64* %P_matrixsize
	store i64 0, i64* %r197
	%r198 = icmp sle i64 %r196, %r197
	br i1 %r198, label %L65, label %L66

L65:
	%r199 = load i64, i64* %P_matrixsize
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r199)
	%r200 = load i64, i64* %P_matrixsize
	%r201 = call i64 @getmatrixsize(i64 %r200)
	br label %L67

L66:
	%r202 = load i64, i64* %P_matrixsize
	store i64 %r202, i64* %_retval
	%r203 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L67

L67:
	%r204 = load i64, i64* %P_matrixsize
	store i64 %r204, i64* %_retval
	%r205 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L68

L68:
	%r206 = load i64, i64* %_retval
	ret i64 %r206


}

define i64 @getmaxrange(i64 %maxrange)
{
L69:
	%_retval = alloca i64
	%P_maxrange = alloca i64
	store i64 %maxrange, i64* %P_maxrange
	br label %L70

L70:
	%r207 = load i64, i64* @maxrange
	store i64 1, i64* %r208
	%r209 = icmp sle i64 %r207, %r208
	br i1 %r209, label %L71, label %L72

L71:
	%r210 = load i64, i64* @maxrange
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %r210)
	%r211 = load i64, i64* @maxrange
	%r212 = call i64 @getmaxrange(i64 %r211)
	br label %L73

L72:
	%r213 = load i64, i64* @maxrange
	store i64 %r213, i64* %_retval
	%r214 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L73

L73:
	%r215 = load i64, i64* @maxrange
	store i64 %r215, i64* %_retval
	%r216 = load i64, i64* %_retval
	ret i64 %_retval
	br label %L74

L74:
	%r217 = load i64, i64* %_retval
	ret i64 %r217


}

define i64 @main()
{
L75:
	%_retval = alloca i64
	%height = alloca i64
	%width = alloca i64
	%r218 = load i64, i64* %height
	store i64 0, i64* %r219
	store i64 %r219, i64* %r218
	%r220 = load i64, i64* %width
	store i64 0, i64* %r221
	store i64 %r221, i64* %r220
	%r222 = load i64, i64* @maxrange
	store i64 0, i64* %r223
	store i64 %r223, i64* %r222
	%r224 = load i64, i64* %height
	%r225 = load i64, i64* %height
	%r226 = call i64 @getmatrixsize(i64 %r225)
	store i64 %r226, i64* %r224
	%r227 = load i64, i64* %width
	%r228 = load i64, i64* %height
	store i64 %r228, i64* %r227
	%r229 = load i64, i64* @maxrange
	%r230 = load i64, i64* @maxrange
	%r231 = call i64 @getmaxrange(i64 %r230)
	store i64 %r231, i64* %r229
	%r232 = load i64, i64* %height
	%r233 = load i64, i64* %width
	%r234 = call i64 @newmatrix(i64 %r232, i64 %r233)
	%r235 = call i8* @malloc(i32 24)
	%r236 = bitcast i8* %r235 to %struct.ListNode*
	store i64 0, i64* %r237
	%r238 = call i64 @maxfactorial(%struct.ListNode %r236, i64 %r237)
	br label %L76

L76:
	store i64 0, i64* %_retval
	%r239 = load i64, i64* %r238
	ret i64 %r239


}


declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [17 x i8] c"Max Factorial=%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1

