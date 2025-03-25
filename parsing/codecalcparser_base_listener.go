// Code generated from CodeCalcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CodeCalcParser
import "github.com/antlr4-go/antlr/v4"

// BaseCodeCalcParserListener is a complete listener for a parse tree produced by CodeCalcParser.
type BaseCodeCalcParserListener struct{}

var _ CodeCalcParserListener = &BaseCodeCalcParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCodeCalcParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCodeCalcParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCodeCalcParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCodeCalcParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModule is called when production module is entered.
func (s *BaseCodeCalcParserListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseCodeCalcParserListener) ExitModule(ctx *ModuleContext) {}

// EnterStatementsAppend is called when production StatementsAppend is entered.
func (s *BaseCodeCalcParserListener) EnterStatementsAppend(ctx *StatementsAppendContext) {}

// ExitStatementsAppend is called when production StatementsAppend is exited.
func (s *BaseCodeCalcParserListener) ExitStatementsAppend(ctx *StatementsAppendContext) {}

// EnterStatementsInitial is called when production StatementsInitial is entered.
func (s *BaseCodeCalcParserListener) EnterStatementsInitial(ctx *StatementsInitialContext) {}

// ExitStatementsInitial is called when production StatementsInitial is exited.
func (s *BaseCodeCalcParserListener) ExitStatementsInitial(ctx *StatementsInitialContext) {}

// EnterAssignStatement is called when production AssignStatement is entered.
func (s *BaseCodeCalcParserListener) EnterAssignStatement(ctx *AssignStatementContext) {}

// ExitAssignStatement is called when production AssignStatement is exited.
func (s *BaseCodeCalcParserListener) ExitAssignStatement(ctx *AssignStatementContext) {}

// EnterExpressionStatement is called when production ExpressionStatement is entered.
func (s *BaseCodeCalcParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production ExpressionStatement is exited.
func (s *BaseCodeCalcParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterBinaryExpression is called when production BinaryExpression is entered.
func (s *BaseCodeCalcParserListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production BinaryExpression is exited.
func (s *BaseCodeCalcParserListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseCodeCalcParserListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseCodeCalcParserListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterUnaryExpression is called when production UnaryExpression is entered.
func (s *BaseCodeCalcParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production UnaryExpression is exited.
func (s *BaseCodeCalcParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseCodeCalcParserListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseCodeCalcParserListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}
