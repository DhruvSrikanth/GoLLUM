parser grammar GoliteParser; 

options {
    tokenVocab = GoliteLexer;
}

// Production Rules:
/*
Program = Types Declarations Functions 'eof'                                                       ;
Types = {TypeDeclaration}                                                                          ;
TypeDeclaration = 'type' 'id' 'struct' '{' Fields '}' ';'                                          ;
Fields = Decl ';' {Decl ';'}                                                                       ;
Decl = 'id' Type                                                                                   ;
Type = 'int' | 'bool' | '*' 'id'                                                                   ;
Declarations = {Declaration}                                                                       ;
Declaration = 'var' Ids Type ';'                                                                   ;
Ids = 'id' {',' 'id'}                                                                              ;
Functions = {Function}                                                                             ;
Function = 'func' 'id' Parameters [ReturnType] '{' Declarations Statements '}'                     ;
Parameters = '(' [ Decl {',' Decl}] ')'                                                            ;
ReturnType = type                                                                                  ;
Statements = {Statement}                                                                           ;
Statement = Block | Assignment | Print | Delete | Conditional | Loop | Return | Read | Invocation  ;
Read =  'scan' expression ';'                                                                      ;
Block = '{' Statements '}'                                                                         ;
Delete = 'delete' Expression ';'                                                                   ;
Assignment = LValue '=' (Expression | 'scan') ';'                                                  ;
Print = 'printf' '(' 'string' { ',' Expression} ')'  ';'                                           ;
Conditional = 'if' '(' Expression ')' Block ['else' Block]                                         ;
Loop = 'for' '(' Expression ')' Block                                                              ;
Return = 'return' [Expression] ';'                                                                 ;
Invocation = 'id' Arguments ';'                                                                    ;
Arguments = '(' [Expression {',' Expression}] ')'                                                  ;
LValue = 'id' {'.' id}                                                                             ;
Expression = BoolTerm {'||' BoolTerm}                                                              ;
BoolTerm = EqualTerm {'&&' EqualTerm}                                                              ;
EqualTerm =  RelationTerm {('=='| '!=') RelationTerm}                                              ;
RelationTerm = SimpleTerm {('>'| '<' | '<=' | '>=') SimpleTerm}                                    ;
SimpleTerm = Term {('+'| '-') Term}                                                                ;
Term = UnaryTerm {('*'| '/') UnaryTerm}                                                            ;
UnaryTerm = '!' SelectorTerm | '-' SelectorTerm | SelectorTerm                                     ;
SelectorTerm = Factor {'.' 'id'}                                                                   ;
Factor = '(' Expression ')' | 'id' [Arguments] | 'number' | 'new' 'id' | 'true' | 'false' | 'nil'  ;
*/

program: types declarations functions EOF;

types: typeDeclaration*;

typeDeclaration: TYPE IDENT STRUCT LBRACE fields RBRACE SEMICOLON;

fields: decl SEMICOLON (decl SEMICOLON)*;

decl: IDENT type;

type: INT | BOOL | (MULT IDENT);

declarations: declaration*;

declaration: VAR ids type SEMICOLON;

ids: IDENT (COMMA IDENT)*;

functions: function*;

function: FUNC IDENT parameters returnType? LBRACE declarations statements RBRACE;

parameters: LPAREN (decl (COMMA decl)*)? RPAREN;

returnType: type;

statements: statement*;

statement: block | assignment | print | delete | conditional | loop | returnRule | invocation;

read: SCAN expression SEMICOLON;

block: LBRACE statements RBRACE;

delete: DELETE expression SEMICOLON;

assignment: lValue ASSIGN (expression | SCAN) SEMICOLON;

print: PRINTF LPAREN STRING_LIT (COMMA expression)* RPAREN SEMICOLON;

conditional: IF LPAREN expression RPAREN block (ELSE block)?;

loop: FOR LPAREN expression RPAREN block;

returnRule: RET expression? SEMICOLON;

invocation: IDENT arguments SEMICOLON;

arguments: LPAREN (expression (COMMA expression)*)? RPAREN;

lValue: IDENT (DOT IDENT)*;

expression: boolTerm (OR boolTerm)*;

boolTerm: equalTerm (AND equalTerm)*;

equalTerm: relationTerm ((EQ | NE) relationTerm)*;

relationTerm: simpleTerm ((GT | LT | LE | GT) simpleTerm)*;

simpleTerm: term ((PLUS | MINUS) term)*;

term: unaryTerm ((MULT | DIV) unaryTerm)*;

unaryTerm: NOT selectorTerm | MINUS selectorTerm | selectorTerm; // Not sure what this line does especially the minus and is the ! a not?

selectorTerm: factor (DOT IDENT)*;

factor: LPAREN expression RPAREN | IDENT arguments? | INT_LIT | NEW IDENT | BOOL_LIT | NIL_LIT; // Can i nest the true | false as BOOL_LIT?

