// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar CodeCalcLexer;

Terminator: ';';

ASSIGN: '=';

OP_ADD : '+';
OP_SUB : '-';

Number     : [1-9][0-9]*;
Identifier : [\p{L}_] [\p{L}\p{N}_]*;

// Parse out
WhiteSpace: [\p{White_Space}]+ -> channel(HIDDEN);

CommentMultiLine  : '`' .*? '`'                   -> channel(HIDDEN);
CommentSingleLine : '#' (~['"\r\n\\])* '\r'? '\n' -> channel(HIDDEN);