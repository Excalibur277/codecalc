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
		"", "';'", "'('", "')'", "'{'", "'}'", "'='", "'+'", "'-'", "'*'", "'/'",
		"'!'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='", "'|'", "'&'", "'if'",
		"'else'", "'while'", "'continue'", "'break'",
	}
	staticData.SymbolicNames = []string{
		"", "Terminator", "L_BRACE", "R_BRACE", "L_C_BRACE", "R_C_BRACE", "ASSIGN",
		"OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "OP_NOT", "OP_GRT", "OP_LST",
		"OP_GTE", "OP_LTE", "OP_EQU", "OP_NEQ", "OP_OR", "OP_AND", "IF", "ELSE",
		"WHILE", "CONTINUE", "BREAK", "Number", "Identifier", "WhiteSpace",
		"CommentMultiLine", "CommentSingleLine",
	}
	staticData.RuleNames = []string{
		"module", "statements", "statement", "while_statement", "if_statement",
		"expression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 99, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 18, 8, 1, 10, 1, 12,
		1, 21, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 37, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 69, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 80, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 5, 5, 94, 8, 5, 10, 5, 12, 5, 97, 9, 5, 1, 5, 0, 2, 2, 10,
		6, 0, 2, 4, 6, 8, 10, 0, 5, 2, 0, 7, 8, 11, 11, 1, 0, 9, 10, 1, 0, 7, 8,
		1, 0, 12, 15, 1, 0, 18, 19, 107, 0, 12, 1, 0, 0, 0, 2, 14, 1, 0, 0, 0,
		4, 36, 1, 0, 0, 0, 6, 38, 1, 0, 0, 0, 8, 68, 1, 0, 0, 0, 10, 79, 1, 0,
		0, 0, 12, 13, 3, 2, 1, 0, 13, 1, 1, 0, 0, 0, 14, 19, 6, 1, -1, 0, 15, 16,
		10, 1, 0, 0, 16, 18, 3, 4, 2, 0, 17, 15, 1, 0, 0, 0, 18, 21, 1, 0, 0, 0,
		19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 3, 1, 0, 0, 0, 21, 19, 1, 0,
		0, 0, 22, 37, 3, 6, 3, 0, 23, 37, 3, 8, 4, 0, 24, 25, 5, 24, 0, 0, 25,
		37, 5, 1, 0, 0, 26, 27, 5, 23, 0, 0, 27, 37, 5, 1, 0, 0, 28, 29, 5, 26,
		0, 0, 29, 30, 5, 6, 0, 0, 30, 31, 3, 10, 5, 0, 31, 32, 5, 1, 0, 0, 32,
		37, 1, 0, 0, 0, 33, 34, 3, 10, 5, 0, 34, 35, 5, 1, 0, 0, 35, 37, 1, 0,
		0, 0, 36, 22, 1, 0, 0, 0, 36, 23, 1, 0, 0, 0, 36, 24, 1, 0, 0, 0, 36, 26,
		1, 0, 0, 0, 36, 28, 1, 0, 0, 0, 36, 33, 1, 0, 0, 0, 37, 5, 1, 0, 0, 0,
		38, 39, 5, 22, 0, 0, 39, 40, 3, 10, 5, 0, 40, 41, 5, 4, 0, 0, 41, 42, 3,
		2, 1, 0, 42, 43, 5, 5, 0, 0, 43, 7, 1, 0, 0, 0, 44, 45, 5, 20, 0, 0, 45,
		46, 3, 10, 5, 0, 46, 47, 5, 4, 0, 0, 47, 48, 3, 2, 1, 0, 48, 49, 5, 5,
		0, 0, 49, 69, 1, 0, 0, 0, 50, 51, 5, 20, 0, 0, 51, 52, 3, 10, 5, 0, 52,
		53, 5, 4, 0, 0, 53, 54, 3, 2, 1, 0, 54, 55, 5, 5, 0, 0, 55, 56, 5, 21,
		0, 0, 56, 57, 3, 8, 4, 0, 57, 69, 1, 0, 0, 0, 58, 59, 5, 20, 0, 0, 59,
		60, 3, 10, 5, 0, 60, 61, 5, 4, 0, 0, 61, 62, 3, 2, 1, 0, 62, 63, 5, 5,
		0, 0, 63, 64, 5, 21, 0, 0, 64, 65, 5, 4, 0, 0, 65, 66, 3, 2, 1, 0, 66,
		67, 5, 5, 0, 0, 67, 69, 1, 0, 0, 0, 68, 44, 1, 0, 0, 0, 68, 50, 1, 0, 0,
		0, 68, 58, 1, 0, 0, 0, 69, 9, 1, 0, 0, 0, 70, 71, 6, 5, -1, 0, 71, 72,
		5, 2, 0, 0, 72, 73, 3, 10, 5, 0, 73, 74, 5, 3, 0, 0, 74, 80, 1, 0, 0, 0,
		75, 80, 5, 25, 0, 0, 76, 80, 5, 26, 0, 0, 77, 78, 7, 0, 0, 0, 78, 80, 3,
		10, 5, 5, 79, 70, 1, 0, 0, 0, 79, 75, 1, 0, 0, 0, 79, 76, 1, 0, 0, 0, 79,
		77, 1, 0, 0, 0, 80, 95, 1, 0, 0, 0, 81, 82, 10, 4, 0, 0, 82, 83, 7, 1,
		0, 0, 83, 94, 3, 10, 5, 5, 84, 85, 10, 3, 0, 0, 85, 86, 7, 2, 0, 0, 86,
		94, 3, 10, 5, 4, 87, 88, 10, 2, 0, 0, 88, 89, 7, 3, 0, 0, 89, 94, 3, 10,
		5, 3, 90, 91, 10, 1, 0, 0, 91, 92, 7, 4, 0, 0, 92, 94, 3, 10, 5, 2, 93,
		81, 1, 0, 0, 0, 93, 84, 1, 0, 0, 0, 93, 87, 1, 0, 0, 0, 93, 90, 1, 0, 0,
		0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 11,
		1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 6, 19, 36, 68, 79, 93, 95,
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
	CodeCalcParserL_BRACE           = 2
	CodeCalcParserR_BRACE           = 3
	CodeCalcParserL_C_BRACE         = 4
	CodeCalcParserR_C_BRACE         = 5
	CodeCalcParserASSIGN            = 6
	CodeCalcParserOP_ADD            = 7
	CodeCalcParserOP_SUB            = 8
	CodeCalcParserOP_MUL            = 9
	CodeCalcParserOP_DIV            = 10
	CodeCalcParserOP_NOT            = 11
	CodeCalcParserOP_GRT            = 12
	CodeCalcParserOP_LST            = 13
	CodeCalcParserOP_GTE            = 14
	CodeCalcParserOP_LTE            = 15
	CodeCalcParserOP_EQU            = 16
	CodeCalcParserOP_NEQ            = 17
	CodeCalcParserOP_OR             = 18
	CodeCalcParserOP_AND            = 19
	CodeCalcParserIF                = 20
	CodeCalcParserELSE              = 21
	CodeCalcParserWHILE             = 22
	CodeCalcParserCONTINUE          = 23
	CodeCalcParserBREAK             = 24
	CodeCalcParserNumber            = 25
	CodeCalcParserIdentifier        = 26
	CodeCalcParserWhiteSpace        = 27
	CodeCalcParserCommentMultiLine  = 28
	CodeCalcParserCommentSingleLine = 29
)

// CodeCalcParser rules.
const (
	CodeCalcParserRULE_module          = 0
	CodeCalcParserRULE_statements      = 1
	CodeCalcParserRULE_statement       = 2
	CodeCalcParserRULE_while_statement = 3
	CodeCalcParserRULE_if_statement    = 4
	CodeCalcParserRULE_expression      = 5
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
		p.SetState(12)
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
	p.SetState(19)
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
			p.SetState(15)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(16)
				p.Statement()
			}

		}
		p.SetState(21)
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

type BreakStatementContext struct {
	StatementContext
}

func NewBreakStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatementContext {
	var p = new(BreakStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserBREAK, 0)
}

func (s *BreakStatementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserTerminator, 0)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
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

type PassThroughStatementContext struct {
	StatementContext
}

func NewPassThroughStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassThroughStatementContext {
	var p = new(PassThroughStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PassThroughStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassThroughStatementContext) While_statement() IWhile_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_statementContext)
}

func (s *PassThroughStatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *PassThroughStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterPassThroughStatement(s)
	}
}

func (s *PassThroughStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitPassThroughStatement(s)
	}
}

func (s *PassThroughStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitPassThroughStatement(s)

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

type ContinueStatementContext struct {
	StatementContext
}

func NewContinueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserCONTINUE, 0)
}

func (s *ContinueStatementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserTerminator, 0)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeCalcParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CodeCalcParserRULE_statement)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPassThroughStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.While_statement()
		}

	case 2:
		localctx = NewPassThroughStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.If_statement()
		}

	case 3:
		localctx = NewBreakStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Match(CodeCalcParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.Match(CodeCalcParserTerminator)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewContinueStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(26)
			p.Match(CodeCalcParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.Match(CodeCalcParserTerminator)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(28)

			var _m = p.Match(CodeCalcParserIdentifier)

			localctx.(*AssignStatementContext).identifier = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Match(CodeCalcParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.expression(0)
		}
		{
			p.SetState(31)
			p.Match(CodeCalcParserTerminator)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewExpressionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(33)
			p.expression(0)
		}
		{
			p.SetState(34)
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

// IWhile_statementContext is an interface to support dynamic dispatch.
type IWhile_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhile_statementContext differentiates from other interfaces.
	IsWhile_statementContext()
}

type While_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_statementContext() *While_statementContext {
	var p = new(While_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_while_statement
	return p
}

func InitEmptyWhile_statementContext(p *While_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_while_statement
}

func (*While_statementContext) IsWhile_statementContext() {}

func NewWhile_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_statementContext {
	var p = new(While_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeCalcParserRULE_while_statement

	return p
}

func (s *While_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *While_statementContext) CopyAll(ctx *While_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *While_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileStatementContext struct {
	While_statementContext
}

func NewWhileStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatementContext {
	var p = new(WhileStatementContext)

	InitEmptyWhile_statementContext(&p.While_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*While_statementContext))

	return p
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserWHILE, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
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

func (s *WhileStatementContext) L_C_BRACE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserL_C_BRACE, 0)
}

func (s *WhileStatementContext) Statements() IStatementsContext {
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

func (s *WhileStatementContext) R_C_BRACE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserR_C_BRACE, 0)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeCalcParser) While_statement() (localctx IWhile_statementContext) {
	localctx = NewWhile_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CodeCalcParserRULE_while_statement)
	localctx = NewWhileStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(CodeCalcParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.expression(0)
	}
	{
		p.SetState(40)
		p.Match(CodeCalcParserL_C_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.statements(0)
	}
	{
		p.SetState(42)
		p.Match(CodeCalcParserR_C_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CodeCalcParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeCalcParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) CopyAll(ctx *If_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfStatementContext struct {
	If_statementContext
}

func NewIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementContext {
	var p = new(IfStatementContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserIF, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
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

func (s *IfStatementContext) L_C_BRACE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserL_C_BRACE, 0)
}

func (s *IfStatementContext) Statements() IStatementsContext {
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

func (s *IfStatementContext) R_C_BRACE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserR_C_BRACE, 0)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseIfStatementContext struct {
	If_statementContext
}

func NewIfElseIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseIfStatementContext {
	var p = new(IfElseIfStatementContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *IfElseIfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseIfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserIF, 0)
}

func (s *IfElseIfStatementContext) Expression() IExpressionContext {
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

func (s *IfElseIfStatementContext) L_C_BRACE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserL_C_BRACE, 0)
}

func (s *IfElseIfStatementContext) Statements() IStatementsContext {
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

func (s *IfElseIfStatementContext) R_C_BRACE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserR_C_BRACE, 0)
}

func (s *IfElseIfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserELSE, 0)
}

func (s *IfElseIfStatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *IfElseIfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterIfElseIfStatement(s)
	}
}

func (s *IfElseIfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitIfElseIfStatement(s)
	}
}

func (s *IfElseIfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitIfElseIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseStatementContext struct {
	If_statementContext
}

func NewIfElseStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseStatementContext {
	var p = new(IfElseStatementContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *IfElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserIF, 0)
}

func (s *IfElseStatementContext) Expression() IExpressionContext {
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

func (s *IfElseStatementContext) AllL_C_BRACE() []antlr.TerminalNode {
	return s.GetTokens(CodeCalcParserL_C_BRACE)
}

func (s *IfElseStatementContext) L_C_BRACE(i int) antlr.TerminalNode {
	return s.GetToken(CodeCalcParserL_C_BRACE, i)
}

func (s *IfElseStatementContext) AllStatements() []IStatementsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementsContext); ok {
			len++
		}
	}

	tst := make([]IStatementsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementsContext); ok {
			tst[i] = t.(IStatementsContext)
			i++
		}
	}

	return tst
}

func (s *IfElseStatementContext) Statements(i int) IStatementsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
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

	return t.(IStatementsContext)
}

func (s *IfElseStatementContext) AllR_C_BRACE() []antlr.TerminalNode {
	return s.GetTokens(CodeCalcParserR_C_BRACE)
}

func (s *IfElseStatementContext) R_C_BRACE(i int) antlr.TerminalNode {
	return s.GetToken(CodeCalcParserR_C_BRACE, i)
}

func (s *IfElseStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserELSE, 0)
}

func (s *IfElseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.EnterIfElseStatement(s)
	}
}

func (s *IfElseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeCalcParserListener); ok {
		listenerT.ExitIfElseStatement(s)
	}
}

func (s *IfElseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeCalcParserVisitor:
		return t.VisitIfElseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeCalcParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CodeCalcParserRULE_if_statement)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Match(CodeCalcParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.expression(0)
		}
		{
			p.SetState(46)
			p.Match(CodeCalcParserL_C_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.statements(0)
		}
		{
			p.SetState(48)
			p.Match(CodeCalcParserR_C_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewIfElseIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Match(CodeCalcParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.expression(0)
		}
		{
			p.SetState(52)
			p.Match(CodeCalcParserL_C_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.statements(0)
		}
		{
			p.SetState(54)
			p.Match(CodeCalcParserR_C_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Match(CodeCalcParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.If_statement()
		}

	case 3:
		localctx = NewIfElseStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.Match(CodeCalcParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.expression(0)
		}
		{
			p.SetState(60)
			p.Match(CodeCalcParserL_C_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.statements(0)
		}
		{
			p.SetState(62)
			p.Match(CodeCalcParserR_C_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Match(CodeCalcParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(CodeCalcParserL_C_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.statements(0)
		}
		{
			p.SetState(66)
			p.Match(CodeCalcParserR_C_BRACE)
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

func (s *BinaryExpressionContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_AND, 0)
}

func (s *BinaryExpressionContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserOP_OR, 0)
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

func (s *UnaryExpressionContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserR_BRACE, 0)
}

func (s *UnaryExpressionContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(CodeCalcParserL_BRACE, 0)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, CodeCalcParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CodeCalcParserL_BRACE:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(71)

			var _m = p.Match(CodeCalcParserL_BRACE)

			localctx.(*UnaryExpressionContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.expression(0)
		}
		{
			p.SetState(73)
			p.Match(CodeCalcParserR_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CodeCalcParserNumber:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(75)

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
			p.SetState(76)

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
			p.SetState(77)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2432) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(78)
			p.expression(5)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(93)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CodeCalcParserRULE_expression)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(82)

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
					p.SetState(83)
					p.expression(5)
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CodeCalcParserRULE_expression)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(85)

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
					p.SetState(86)
					p.expression(4)
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CodeCalcParserRULE_expression)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(88)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&61440) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)
					p.expression(3)
				}

			case 4:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CodeCalcParserRULE_expression)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(91)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CodeCalcParserOP_OR || _la == CodeCalcParserOP_AND) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(92)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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

	case 5:
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
