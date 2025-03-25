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

	// EnterAssignStatement is called when entering the AssignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterExpressionStatement is called when entering the ExpressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterBinaryExpression is called when entering the BinaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterLiteralExpression is called when entering the LiteralExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterUnaryExpression is called when entering the UnaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitStatementsAppend is called when exiting the StatementsAppend production.
	ExitStatementsAppend(c *StatementsAppendContext)

	// ExitStatementsInitial is called when exiting the StatementsInitial production.
	ExitStatementsInitial(c *StatementsInitialContext)

	// ExitAssignStatement is called when exiting the AssignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitExpressionStatement is called when exiting the ExpressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitBinaryExpression is called when exiting the BinaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitLiteralExpression is called when exiting the LiteralExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitUnaryExpression is called when exiting the UnaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)
}
