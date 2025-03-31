package listener

import (
	"fmt"

	"codecalc.com/m/v2/ast"
	"codecalc.com/m/v2/listener/listenerstack"
	"codecalc.com/m/v2/parsing"
	"github.com/antlr4-go/antlr/v4"
)

var _ parsing.CodeCalcParserListener = &CodeCalcListener{}

type CodeCalcListener struct {
	parsing.BaseCodeCalcParserListener
	stack listenerstack.ListenerStack
	queue listenerstack.ListenerQueue
}

func Dequeue[T any](l *CodeCalcListener) T {
	if d, ok := l.queue.Dequeue().(T); ok {
		return d
	} else {
		panic(fmt.Sprintf("popping invalid type, got: %T, expected %T", d, new(T)))
	}
}

func Push[T any](l *CodeCalcListener, d T) {
	//fmt.Printf("Pushing: %s\n", d)
	l.stack.Push(d)
}

func NewListener() *CodeCalcListener {
	return &CodeCalcListener{}
}

func (l *CodeCalcListener) Pop(amount int) {
	if len(l.queue) > 0 {
		panic(fmt.Sprintf("queue has %d values remaining", len(l.queue)))
	}
	l.queue = l.stack.Pop(amount)
}

func (l *CodeCalcListener) ExitModule(c *parsing.ModuleContext) {
	l.Pop(1)
	Push(l, ast.NewModule(Dequeue[[]ast.Statement](l)))
}

func (l *CodeCalcListener) ExitStatementsInitial(c *parsing.StatementsInitialContext) {
	Push(l, []ast.Statement{})
}

func (l *CodeCalcListener) ExitStatementsAppend(c *parsing.StatementsAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.Statement](l), Dequeue[ast.Statement](l)))
}

func (l *CodeCalcListener) ExitAssignStatement(c *parsing.AssignStatementContext) {
	l.Pop(1)
	Push(l, ast.NewAssignStatement(c.GetIdentifier(), Dequeue[ast.Expression](l)))
}

func (l *CodeCalcListener) ExitExpressionStatement(c *parsing.ExpressionStatementContext) {
	l.Pop(1)
	Push(l, ast.NewExpressionStatement(Dequeue[ast.Expression](l)))
}

func (l *CodeCalcListener) ExitBreakStatement(c *parsing.BreakStatementContext) {
	Push(l, ast.NewBreakStatement())
}

func (l *CodeCalcListener) ExitContinueStatement(c *parsing.ContinueStatementContext) {
	Push(l, ast.NewContinueStatement())
}

func (l *CodeCalcListener) ExitWhileStatement(c *parsing.WhileStatementContext) {
	l.Pop(2)
	Push(l, ast.NewWhileStatement(Dequeue[ast.Expression](l), Dequeue[[]ast.Statement](l)))
}

func (l *CodeCalcListener) ExitIfStatement(c *parsing.IfStatementContext) {
	l.Pop(2)
	Push(l, ast.NewIfStatement(Dequeue[ast.Expression](l), Dequeue[[]ast.Statement](l)))
}

func (l *CodeCalcListener) ExitIfElseIfStatement(c *parsing.IfElseIfStatementContext) {
	l.Pop(3)
	Push(l, ast.NewIfElseStatement(Dequeue[ast.Expression](l), Dequeue[[]ast.Statement](l), Dequeue[ast.Statement](l)))
}

func (l *CodeCalcListener) ExitIfElseStatement(c *parsing.IfElseStatementContext) {
	l.Pop(3)
	Push(l, ast.NewIfElseStatement(Dequeue[ast.Expression](l), Dequeue[[]ast.Statement](l), ast.NewElseStatement(Dequeue[[]ast.Statement](l))))
}

func (l *CodeCalcListener) ExitBinaryExpression(c *parsing.BinaryExpressionContext) {
	l.Pop(2)
	Push(l, ast.NewBinaryExpression(Dequeue[ast.Expression](l), c.GetOp(), Dequeue[ast.Expression](l)))
}

func (l *CodeCalcListener) ExitUnaryExpression(c *parsing.UnaryExpressionContext) {
	l.Pop(1)
	Push(l, ast.NewUnaryExpression(c.GetOp(), Dequeue[ast.Expression](l)))
}

func (l *CodeCalcListener) ExitLiteralExpression(c *parsing.LiteralExpressionContext) {
	Push(l, ast.NewLiteralExpression(c.GetNumber()))
}

func (l *CodeCalcListener) ExitIdentifierExpression(c *parsing.IdentifierExpressionContext) {
	Push(l, ast.NewIdentifierExpression(c.GetIdentifier()))
}

func (l *CodeCalcListener) ExitEveryRule(c antlr.ParserRuleContext) {
	//fmt.Printf("Exiting Rule: %s\n", parsing.CodeCalcParserParserStaticData.RuleNames[c.GetRuleIndex()])
}
