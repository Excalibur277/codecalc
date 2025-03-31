// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar CodeCalcLexer;

Terminator : ';';
L_BRACE    : '(';
R_BRACE    : ')';
L_C_BRACE  : '{';
R_C_BRACE  : '}';
L_S_BRACE  : '[';
R_S_BRACE  : ']';

ASSIGN: '=';

OP_ADD : '+';
OP_SUB : '-';
OP_MUL : '*';
OP_DIV : '/';
OP_MOD : '%';
OP_NOT : '!';

OP_GRT : '>';
OP_LST : '<';
OP_GTE : '>=';
OP_LTE : '<=';
OP_EQU : '==';
OP_NEQ : '!=';

OP_OR  : '|';
OP_AND : '&';

IF       : 'if';
ELSE     : 'else';
WHILE    : 'while';
CONTINUE : 'continue';
BREAK    : 'break';
ARRAY    : 'array';

Number     : [0-9][0-9]*;
Identifier : [\p{L}_] [\p{L}\p{N}_]*;

// Parse out
WhiteSpace: [\p{White_Space}]+ -> channel(HIDDEN);

CommentMultiLine  : '`' .*? '`'                             -> channel(HIDDEN);
CommentSingleLine : '#' (~['"\r\n\\])* (('\r'? '\n') | EOF) -> channel(HIDDEN);