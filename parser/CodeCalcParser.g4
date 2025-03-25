// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, singleLineOverrulesHangingColon true, alignColons hanging
parser grammar CodeCalcParser;

options {
    tokenVocab = CodeCalcLexer;
}

module
    : statements
    ;

statements
    :                      # StatementsInitial
    | statements statement # StatementsAppend
    ;

statement
    : identifier = Identifier ASSIGN expression Terminator # AssignStatement
    | expression Terminator                                # ExpressionStatement
    ;

expression
    : number = Number                              # LiteralExpression
    | identifier = Identifier                      # IdentifierExpression
    | op = (OP_ADD | OP_SUB) expression            # UnaryExpression
    | expression op = (OP_ADD | OP_SUB) expression # BinaryExpression
    ;