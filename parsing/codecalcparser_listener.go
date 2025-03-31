// Code generated from CodeCalcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CodeCalcParser
import "github.com/antlr4-go/antlr/v4"

// CodeCalcParserListener is a complete listener for a parse tree produced by CodeCalcParser.
type CodeCalcParserListener interface {
	antlr.ParseTreeListener

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterStatementsAppend is called when entering the StatementsAppend production.
	EnterStatementsAppend(c *StatementsAppendContext)

	// EnterStatementsInitial is called when entering the StatementsInitial production.
	EnterStatementsInitial(c *StatementsInitialContext)

	// EnterPassThroughStatement is called when entering the PassThroughStatement production.
	EnterPassThroughStatement(c *PassThroughStatementContext)

	// EnterBreakStatement is called when entering the BreakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the ContinueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterAssignStatement is called when entering the AssignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterAssignIndexStatement is called when entering the AssignIndexStatement production.
	EnterAssignIndexStatement(c *AssignIndexStatementContext)

	// EnterArrayStatement is called when entering the ArrayStatement production.
	EnterArrayStatement(c *ArrayStatementContext)

	// EnterExpressionStatement is called when entering the ExpressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterBlankStatement is called when entering the BlankStatement production.
	EnterBlankStatement(c *BlankStatementContext)

	// EnterWhileStatement is called when entering the WhileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterIfStatement is called when entering the IfStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIfElseIfStatement is called when entering the IfElseIfStatement production.
	EnterIfElseIfStatement(c *IfElseIfStatementContext)

	// EnterIfElseStatement is called when entering the IfElseStatement production.
	EnterIfElseStatement(c *IfElseStatementContext)

	// EnterBinaryExpression is called when entering the BinaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterLiteralExpression is called when entering the LiteralExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterUnaryExpression is called when entering the UnaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterAccessExpression is called when entering the AccessExpression production.
	EnterAccessExpression(c *AccessExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitStatementsAppend is called when exiting the StatementsAppend production.
	ExitStatementsAppend(c *StatementsAppendContext)

	// ExitStatementsInitial is called when exiting the StatementsInitial production.
	ExitStatementsInitial(c *StatementsInitialContext)

	// ExitPassThroughStatement is called when exiting the PassThroughStatement production.
	ExitPassThroughStatement(c *PassThroughStatementContext)

	// ExitBreakStatement is called when exiting the BreakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the ContinueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitAssignStatement is called when exiting the AssignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitAssignIndexStatement is called when exiting the AssignIndexStatement production.
	ExitAssignIndexStatement(c *AssignIndexStatementContext)

	// ExitArrayStatement is called when exiting the ArrayStatement production.
	ExitArrayStatement(c *ArrayStatementContext)

	// ExitExpressionStatement is called when exiting the ExpressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitBlankStatement is called when exiting the BlankStatement production.
	ExitBlankStatement(c *BlankStatementContext)

	// ExitWhileStatement is called when exiting the WhileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitIfStatement is called when exiting the IfStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIfElseIfStatement is called when exiting the IfElseIfStatement production.
	ExitIfElseIfStatement(c *IfElseIfStatementContext)

	// ExitIfElseStatement is called when exiting the IfElseStatement production.
	ExitIfElseStatement(c *IfElseStatementContext)

	// ExitBinaryExpression is called when exiting the BinaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitLiteralExpression is called when exiting the LiteralExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitUnaryExpression is called when exiting the UnaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitAccessExpression is called when exiting the AccessExpression production.
	ExitAccessExpression(c *AccessExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)
}
