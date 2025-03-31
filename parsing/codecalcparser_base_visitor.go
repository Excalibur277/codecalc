// Code generated from CodeCalcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CodeCalcParser
import "github.com/antlr4-go/antlr/v4"

type BaseCodeCalcParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCodeCalcParserVisitor) VisitModule(ctx *ModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitStatementsAppend(ctx *StatementsAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitStatementsInitial(ctx *StatementsInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitPassThroughStatement(ctx *PassThroughStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitAssignIndexStatement(ctx *AssignIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitArrayStatement(ctx *ArrayStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitBlankStatement(ctx *BlankStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitIfElseIfStatement(ctx *IfElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitIfElseStatement(ctx *IfElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitAccessExpression(ctx *AccessExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
