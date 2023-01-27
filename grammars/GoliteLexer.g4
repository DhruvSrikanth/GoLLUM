lexer grammar GoliteLexer;


// Precedence levels:
// 1. assignment
// 2: ||, &&
// 3: ==, !=, <, <=, >, >=
// 4: +, -, *, /
// 5: delimiter
// 6: keyword
// 7: literals
// 8: types
// 9: identifiers
// 10: skip


// Assignment operator
ASSIGN: '=';

// Logical operators
AND: '&&';

OR: '||';

NOT: '!';

// Relational operators
LT: '<';

GT: '>';

LE: '<=';

GE: '>=';

EQ: '==';

NE: '!=';

// Operators
PLUS: '+';

MINUS: '-';

MULT: '*';

DIV: '/';

// Delimiters
LPAREN: '(';

RPAREN: ')';

LBRACE: '{';

RBRACE: '}';

SEMICOLON: ';';

COMMA: ',';

// Keywords
VAR: 'var';

TYPE: 'type';

NEW: 'new';

DELETE: 'delete';

STRUCT: 'struct';

DOT: '.';

FUNC: 'func';

RET: 'return';

IF: 'if';

ELSE: 'else';

FOR: 'for';

SCAN: 'scan';

PRINTF: 'printf';

// Literals
INT_LIT: [-]?[1-9][0-9]* | [-]?[0];

STRING_LIT: '"'.*?'"';

BOOL_LIT: 'true' | 'false';

NIL_LIT: 'nil';

// TYPES
INT: 'int';

BOOL: 'bool';

// Identifiers
IDENT: [a-zA-Z][a-zA-Z0-9]*;

// Skip
// whitespace, newline, carriage return, tab
WS: [ \t\r\n]+ -> skip;

// Comments (only single line comments)
COMMENT: '//'.*?'\n' -> skip;
