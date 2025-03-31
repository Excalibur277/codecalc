// Code generated from CodeCalcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CodeCalcParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CodeCalcParser.
type CodeCalcParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CodeCalcParser#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#StatementsAppend.
	VisitStatementsAppend(ctx *StatementsAppendContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#StatementsInitial.
	VisitStatementsInitial(ctx *StatementsInitialContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#PassThroughStatement.
	VisitPassThroughStatement(ctx *PassThroughStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#BreakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#ContinueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#AssignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#AssignIndexStatement.
	VisitAssignIndexStatement(ctx *AssignIndexStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#ArrayStatement.
	VisitArrayStatement(ctx *ArrayStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#ExpressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#BlankStatement.
	VisitBlankStatement(ctx *BlankStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#WhileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#IfStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#IfElseIfStatement.
	VisitIfElseIfStatement(ctx *IfElseIfStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#IfElseStatement.
	VisitIfElseStatement(ctx *IfElseStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#BinaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#UnaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#AccessExpression.
	VisitAccessExpression(ctx *AccessExpressionContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}
}
