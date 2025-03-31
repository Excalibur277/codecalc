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
    : while_statement                                                                     # PassThroughStatement
    | if_statement                                                                        # PassThroughStatement
    | BREAK Terminator                                                                    # BreakStatement
    | CONTINUE Terminator                                                                 # ContinueStatement
    | identifier = Identifier ASSIGN expression Terminator                                # AssignStatement
    | identifier = Identifier L_S_BRACE expression R_S_BRACE ASSIGN expression Terminator # AssignIndexStatement
    | ARRAY identifier = Identifier L_S_BRACE expression R_S_BRACE Terminator             # ArrayStatement
    | expression Terminator                                                               # ExpressionStatement
    | Terminator                                                                          # BlankStatement
    ;

while_statement
    : WHILE expression L_C_BRACE statements R_C_BRACE # WhileStatement
    ;

if_statement
    : IF expression L_C_BRACE statements R_C_BRACE                                     # IfStatement
    | IF expression L_C_BRACE statements R_C_BRACE ELSE if_statement                   # IfElseIfStatement
    | IF expression L_C_BRACE statements R_C_BRACE ELSE L_C_BRACE statements R_C_BRACE # IfElseStatement
    ;

expression
    : op = L_BRACE expression R_BRACE                                # UnaryExpression
    | number = Number                                                # LiteralExpression
    | identifier = Identifier L_S_BRACE expression R_S_BRACE         # AccessExpression
    | identifier = Identifier                                        # IdentifierExpression
    | op = (OP_NOT | OP_ADD | OP_SUB) expression                     # UnaryExpression
    | expression op = (OP_MUL | OP_DIV | OP_MOD) expression          # BinaryExpression
    | expression op = (OP_ADD | OP_SUB) expression                   # BinaryExpression
    | expression op = (OP_GRT | OP_LST | OP_GTE | OP_LTE) expression # BinaryExpression
    | expression op = (OP_EQU | OP_NEQ) expression                   # BinaryExpression
    | expression op = (OP_AND | OP_OR) expression                    # BinaryExpression
    ;