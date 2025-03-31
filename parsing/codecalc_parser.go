// Code generated from CodeCalcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CodeCalcParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CodeCalcParser struct {
	*antlr.BaseParser
}

var CodeCalcParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func codecalcparserParserInit() {
	staticData := &CodeCalcParserParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'='", "'+'", "'-'", "'*'", "'/'", "'!'", "'>'", "'<'", "'>='",
		"'<='", "'=='", "'!='",
	}
	staticData.SymbolicNames = []string{
		"", "Terminator", "ASSIGN", "OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV",
		"OP_NOT", "OP_GRT", "OP_LST", "OP_GTE", "OP_LTE", "OP_EQU", "OP_NEQ",
		"Number", "Identifier", "WhiteSpace", "CommentMultiLine", "CommentSingleLine",
	}
	staticData.RuleNames = []string{
		"module", "statements", "statement", "expression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 53, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 5, 1, 14, 8, 1, 10, 1, 12, 1, 17, 9, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 27, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 34, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 48, 8, 3, 10, 3, 12, 3, 51, 9, 3, 1, 3, 0,
		2, 2, 6, 4, 0, 2, 4, 6, 0, 5, 2, 0, 3, 4, 7, 7, 1, 0, 5, 6, 1, 0, 3, 4,
		1, 0, 8, 11, 1, 0, 12, 13, 56, 0, 8, 1, 0, 0, 0, 2, 10, 1, 0, 0, 0, 4,
		26, 1, 0, 0, 0, 6, 33, 1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 1, 1, 0, 0, 0,
		10, 15, 6, 1, -1, 0, 11, 12, 10, 1, 0, 0, 12, 14, 3, 4, 2, 0, 13, 11, 1,
		0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16,
		3, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 19, 5, 15, 0, 0, 19, 20, 5, 2, 0,
		0, 20, 21, 3, 6, 3, 0, 21, 22, 5, 1, 0, 0, 22, 27, 1, 0, 0, 0, 23, 24,
		3, 6, 3, 0, 24, 25, 5, 1, 0, 0, 25, 27, 1, 0, 0, 0, 26, 18, 1, 0, 0, 0,
		26, 23, 1, 0, 0, 0, 27, 5, 1, 0, 0, 0, 28, 29, 6, 3, -1, 0, 29, 34, 5,
		14, 0, 0, 30, 34, 5, 15, 0, 0, 31, 32, 7, 0, 0, 0, 32, 34, 3, 6, 3, 5,
		33, 28, 1, 0, 0, 0, 33, 30, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 49, 1,
		0, 0, 0, 35, 36, 10, 4, 0, 0, 36, 37, 7, 1, 0, 0, 37, 48, 3, 6, 3, 5, 38,
		39, 10, 3, 0, 0, 39, 40, 7, 2, 0, 0, 40, 48, 3, 6, 3, 4, 41, 42, 10, 2,
		0, 0, 42, 43, 7, 3, 0, 0, 43, 48, 3, 6, 3, 3, 44, 45, 10, 1, 0, 0, 45,
		46, 7, 4, 0, 0, 46, 48, 3, 6, 3, 2, 47, 35, 1, 0, 0, 0, 47, 38, 1, 0, 0,
		0, 47, 41, 1, 0, 0, 0, 47, 44, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 7, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0,
		5, 15, 26, 33, 47, 49,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CodeCalcParserInit initializes any static state used to implement CodeCalcParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCodeCalcParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CodeCalcParserInit() {
	staticData := &CodeCalcParserParserStaticData
	staticData.once.Do(codecalcparserParserInit)
}

// NewCodeCalcParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCodeCalcParser(input antlr.TokenStream) *CodeCalcParser {
	CodeCalcParserInit()
	this := new(CodeCalcParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CodeCalcParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CodeCalcParser.g4"

	return this
}

// CodeCalcParser tokens.
const (
	CodeCalcParserEOF               = antlr.TokenEOF
	CodeCalcParserTerminator        = 1
	CodeCalcParserASSIGN            = 2
	CodeCalcParserOP_ADD            = 3
	CodeCalcParserOP_SUB            = 4
	CodeCalcParserOP_MUL            = 5
	CodeCalcParserOP_DIV            = 6
	CodeCalcParserOP_NOT            = 7
	CodeCalcParserOP_GRT            = 8
	CodeCalcParserOP_LST            = 9
	CodeCalcParserOP_GTE            = 10
	CodeCalcParserOP_LTE            = 11
	CodeCalcParserOP_EQU            = 12
	CodeCalcParserOP_NEQ            = 13
	CodeCalcParserNumber            = 14
	CodeCalcParserIdentifier        = 15
	CodeCalcParserWhiteSpace        = 16
	CodeCalcParserCommentMultiLine  = 17
	CodeCalcParserCommentSingleLine = 18
)

// CodeCalcParser rules.
const (
	CodeCalcParserRULE_module     = 0
	CodeCalcParserRULE_statements = 1
	CodeCalcParserRULE_statement  = 2
	CodeCalcParserRULE_expression = 3
)

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statements() IStatementsContext

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_module
	return p
}

func InitEmptyModuleContext(p *ModuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_module
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeCalcParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitModule(s)
	}
}

func (s *ModuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitModule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeCalcParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CodeCalcParserRULE_module)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.statements(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeCalcParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) CopyAll(ctx *StatementsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatementsAppendContext struct {
	StatementsContext
}

func NewStatementsAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementsAppendContext {
	var p = new(StatementsAppendContext)

	InitEmptyStatementsContext(&p.StatementsContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementsContext))

	return p
}

func (s *StatementsAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsAppendContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *StatementsAppendContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterStatementsAppend(s)
	}
}

func (s *StatementsAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitStatementsAppend(s)
	}
}

func (s *StatementsAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitStatementsAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementsInitialContext struct {
	StatementsContext
}

func NewStatementsInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementsInitialContext {
	var p = new(StatementsInitialContext)

	InitEmptyStatementsContext(&p.StatementsContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementsContext))

	return p
}

func (s *StatementsInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterStatementsInitial(s)
	}
}

func (s *StatementsInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitStatementsInitial(s)
	}
}

func (s *StatementsInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitStatementsInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeCalcParser) Statements() (localctx IStatementsContext) {
	return p.statements(0)
}

func (p *CodeCalcParser) statements(_p int) (localctx IStatementsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStatementsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CodeCalcParserRULE_statements, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewStatementsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewStatementsAppendContext(p, NewStatementsContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CodeCalcParserRULE_statements)
			p.SetState(11)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(12)
				p.Statement()
			}

		}
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeCalcParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionStatementContext struct {
	StatementContext
}

func NewExpressionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserTerminator, 0)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignStatementContext struct {
	StatementContext
	identifier antlr.Token
}

func NewAssignStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStatementContext {
	var p = new(AssignStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AssignStatementContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *AssignStatementContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *AssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserASSIGN, 0)
}

func (s *AssignStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignStatementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserTerminator, 0)
}

func (s *AssignStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserIdentifier, 0)
}

func (s *AssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterAssignStatement(s)
	}
}

func (s *AssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitAssignStatement(s)
	}
}

func (s *AssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeCalcParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CodeCalcParserRULE_statement)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)

			var _m = p.Match(CodeCalcParserIdentifier)

			localctx.(*AssignStatementContext).identifier = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.Match(CodeCalcParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(20)
			p.expression(0)
		}
		{
			p.SetState(21)
			p.Match(CodeCalcParserTerminator)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExpressionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.expression(0)
		}
		{
			p.SetState(24)
			p.Match(CodeCalcParserTerminator)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeCalcParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_MUL, 0)
}

func (s *BinaryExpressionContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_DIV, 0)
}

func (s *BinaryExpressionContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_ADD, 0)
}

func (s *BinaryExpressionContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_SUB, 0)
}

func (s *BinaryExpressionContext) OP_GRT() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_GRT, 0)
}

func (s *BinaryExpressionContext) OP_LST() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_LST, 0)
}

func (s *BinaryExpressionContext) OP_GTE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_GTE, 0)
}

func (s *BinaryExpressionContext) OP_LTE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_LTE, 0)
}

func (s *BinaryExpressionContext) OP_EQU() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_EQU, 0)
}

func (s *BinaryExpressionContext) OP_NEQ() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_NEQ, 0)
}

func (s *BinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	ExpressionContext
	number antlr.Token
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetNumber() antlr.Token { return s.number }

func (s *LiteralExpressionContext) SetNumber(v antlr.Token) { s.number = v }

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserNumber, 0)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_NOT, 0)
}

func (s *UnaryExpressionContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_ADD, 0)
}

func (s *UnaryExpressionContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_SUB, 0)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	ExpressionContext
	identifier antlr.Token
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *IdentifierExpressionContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeCalcParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CodeCalcParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, CodeCalcParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CodeCalcParserNumber:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(29)

			var _m = p.Match(CodeCalcParserNumber)

			localctx.(*LiteralExpressionContext).number = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CodeCalcParserIdentifier:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)

			var _m = p.Match(CodeCalcParserIdentifier)

			localctx.(*IdentifierExpressionContext).identifier = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CodeCalcParserOP_ADD, CodeCalcParserOP_SUB, CodeCalcParserOP_NOT:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&152) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(32)
			p.expression(5)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CodeCalcParserRULE_expression)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(36)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CodeCalcParserOP_MUL || _la == CodeCalcParserOP_DIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(37)
					p.expression(5)
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CodeCalcParserRULE_expression)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(39)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CodeCalcParserOP_ADD || _la == CodeCalcParserOP_SUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(40)
					p.expression(4)
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CodeCalcParserRULE_expression)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(42)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3840) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(43)
					p.expression(3)
				}

			case 4:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CodeCalcParserRULE_expression)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(45)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CodeCalcParserOP_EQU || _la == CodeCalcParserOP_NEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(46)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *CodeCalcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *StatementsContext = nil
		if localctx != nil {
			t = localctx.(*StatementsContext)
		}
		return p.Statements_Sempred(t, predIndex)

	case 3:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CodeCalcParser) Statements_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CodeCalcParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
