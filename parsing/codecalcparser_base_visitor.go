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

func (v *BaseCodeCalcParserVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeCalcParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
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

func (v *BaseCodeCalcParserVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
