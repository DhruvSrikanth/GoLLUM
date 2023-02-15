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
Read =  'scan' LValue ';'                                                                          ;
Block = '{' Statements '}'                                                                         ;
Delete = 'delete' Expression ';'                                                                   ;
Assignment = LValue '=' Expression ';'                                                             ;
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
RelationTerm = SimpleTerm {('>' | '<' | '<=' | '>=') SimpleTerm}                                   ;
SimpleTerm = Term {('+'| '-') Term}                                                                ;
Term = UnaryTerm {('*'| '/') UnaryTerm}                                                            ;
UnaryTerm = '!' SelectorTerm | '-' SelectorTerm | SelectorTerm                                     ;
SelectorTerm = Factor {'.' 'id'}                                                                   ;
Factor = '(' Expression ')' | 'id' [Arguments] | 'number' | 'new' 'id' | 'true' | 'false' | 'nil'  ;
*/

program: ty=types dcls=declarations fns=functions EOF;

types: typeDeclaration*;

typeDeclaration: TYPE id=IDENT STRUCT LBRACE flds=fields RBRACE SEMICOLON;

fields: dcl=decl SEMICOLON morefields*;

morefields: dcl=decl SEMICOLON; // Split fields rule

decl: id=IDENT ty=type;

type: INT | BOOL | (MULT id=IDENT);

declarations: declaration*;

declaration: VAR idx=ids ty=type SEMICOLON;

ids: id=IDENT otherids*;

otherids: COMMA id=IDENT; // Split from ids grammar rule

functions: function*;

function: FUNC id=IDENT params=parameters rty=returnType? LBRACE dcls=declarations stmts=statements RBRACE;

parameters: LPAREN parametersPrime? RPAREN;

parametersPrime: dcl=decl parametersPPrime*; // Split from parameters grammar rule

parametersPPrime: COMMA dcl=decl; // Split from parameters grammar rule

returnType: ty=type;

statements: statement*;

statement: bl=block | asmt=assignment | prnt=print | del=delete | rd=read | cond=conditional | lp=loop | ret=returnRule | invoke=invocation;

read: SCAN lval=lValue SEMICOLON;

block: LBRACE stmts=statements RBRACE;

delete: DELETE expr=expression SEMICOLON;

assignment: lval=lValue ASSIGN expr=expression SEMICOLON;

print: PRINTF LPAREN str=STRING_LIT printPrime* RPAREN SEMICOLON;

printPrime: COMMA expr=expression; // Split from print grammar rule

conditional: IF LPAREN expr=expression RPAREN bl=block then=conditionalPrime?;

conditionalPrime: ELSE bl=block; // Split from conditional grammar rule

loop: FOR LPAREN expr=expression RPAREN bl=block;

returnRule: RET expr=expression? SEMICOLON;

invocation: id=IDENT args=arguments SEMICOLON;

arguments: LPAREN args=argumentsPrime? RPAREN; // Split from arguments grammar rule

argumentsPrime: expr=expression argumentsPrimePrime*;

argumentsPrimePrime: COMMA expr=expression; // Split from arguments grammar rule

lValue: id=IDENT lValuePrime*;

lValuePrime: DOT id=IDENT; // Split from lValue grammar rule

expression: bt=boolTerm expressionPrime*;

expressionPrime: op=OR bt=boolTerm; // Split from expression grammar rule

boolTerm: eq=equalTerm boolTermPrime*;

boolTermPrime: op=AND eq=equalTerm; // Split from boolTerm grammar rule

equalTerm: rt=relationTerm equalTermPrime*;

equalTermPrime: op=(EQ | NE) rt=relationTerm; // Split from equalTerm grammar rule

relationTerm: st=simpleTerm relationTermPrime*;

relationTermPrime: op=(GE | LT | LE | GT) st=simpleTerm; // Split from relationTerm grammar rule

simpleTerm: t=term simpleTermPrime*;

simpleTermPrime: op=(PLUS | MINUS) t=term; // Split from simpleTerm grammar rule

term: ut=unaryTerm termPrime*;

termPrime: op=(MULT | DIV) ut=unaryTerm; // Split from term grammar rule

unaryTerm: unaryTermBool | unaryTermInt | selectorTerm;

unaryTermBool: op=NOT st=selectorTerm; // Split from unaryTerm grammar rule

unaryTermInt: op=MINUS st=selectorTerm; // Split from unaryTerm grammar rule

selectorTerm: f=factor selectorTermPrime*;

selectorTermPrime: DOT id=IDENT; // Split from selectorTerm grammer

factor: subfactor | variableInvocation | INT_LIT | allocation | BOOL_LIT | NIL_LIT;

subfactor: LPAREN expr=expression RPAREN; // Split from factor grammer

variableInvocation: id=IDENT args=arguments?; // Split from factor grammer, not sure if this is a function call

allocation: NEW key=IDENT; // Split from factor grammer