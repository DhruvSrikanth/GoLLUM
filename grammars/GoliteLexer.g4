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

// Does this precedence look okay?


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

// STRING_LIT: .*; // How to declare a string literal? this is incorrect because it matches the empty string

BOOL_LIT: 'true' | 'false';

NIL_LIT: 'nil';

// TYPES
INT: 'int';

STRING: 'string'; // Is this a type?

BOOL: 'bool';

PTR: '*'IDENT; // Is this how you define a pointer type? Would this have higher precedence than the '*' and identifier operators?

// Identifiers
IDENT: [a-zA-Z][a-zA-Z0-9]*;

// Skip
// whitespace, newline, carriage return, tab
WS: [ \t\r\n]+ -> skip;

// Comments
COMMENT: '//'.*?'\t'*'\r'*'\n' -> skip; // Is this the right way to declare a comment? It is suggesting that I had a ? for a non-greedy match, but I don't know what that means.
