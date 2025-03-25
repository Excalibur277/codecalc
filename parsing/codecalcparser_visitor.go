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

	// Visit a parse tree produced by CodeCalcParser#AssignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#ExpressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#BinaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#UnaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by CodeCalcParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}
}
