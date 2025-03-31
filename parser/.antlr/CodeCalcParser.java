// Generated from /home/excalibur/PersonalProjects/codecalc/parser/CodeCalcParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class CodeCalcParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Terminator=1, L_BRACE=2, R_BRACE=3, L_C_BRACE=4, R_C_BRACE=5, L_S_BRACE=6, 
		R_S_BRACE=7, ASSIGN=8, OP_ADD=9, OP_SUB=10, OP_MUL=11, OP_DIV=12, OP_MOD=13, 
		OP_NOT=14, OP_GRT=15, OP_LST=16, OP_GTE=17, OP_LTE=18, OP_EQU=19, OP_NEQ=20, 
		OP_OR=21, OP_AND=22, IF=23, ELSE=24, WHILE=25, CONTINUE=26, BREAK=27, 
		ARRAY=28, Number=29, Identifier=30, WhiteSpace=31, CommentMultiLine=32, 
		CommentSingleLine=33;
	public static final int
		RULE_module = 0, RULE_statements = 1, RULE_statement = 2, RULE_while_statement = 3, 
		RULE_if_statement = 4, RULE_expression = 5;
	private static String[] makeRuleNames() {
		return new String[] {
			"module", "statements", "statement", "while_statement", "if_statement", 
			"expression"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'='", "'+'", 
			"'-'", "'*'", "'/'", "'%'", "'!'", "'>'", "'<'", "'>='", "'<='", "'=='", 
			"'!='", "'|'", "'&'", "'if'", "'else'", "'while'", "'continue'", "'break'", 
			"'array'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "Terminator", "L_BRACE", "R_BRACE", "L_C_BRACE", "R_C_BRACE", "L_S_BRACE", 
			"R_S_BRACE", "ASSIGN", "OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "OP_MOD", 
			"OP_NOT", "OP_GRT", "OP_LST", "OP_GTE", "OP_LTE", "OP_EQU", "OP_NEQ", 
			"OP_OR", "OP_AND", "IF", "ELSE", "WHILE", "CONTINUE", "BREAK", "ARRAY", 
			"Number", "Identifier", "WhiteSpace", "CommentMultiLine", "CommentSingleLine"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "CodeCalcParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public CodeCalcParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ModuleContext extends ParserRuleContext {
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public ModuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_module; }
	}

	public final ModuleContext module() throws RecognitionException {
		ModuleContext _localctx = new ModuleContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_module);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(12);
			statements(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementsContext extends ParserRuleContext {
		public StatementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statements; }
	 
		public StatementsContext() { }
		public void copyFrom(StatementsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementsAppendContext extends StatementsContext {
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public StatementsAppendContext(StatementsContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementsInitialContext extends StatementsContext {
		public StatementsInitialContext(StatementsContext ctx) { copyFrom(ctx); }
	}

	public final StatementsContext statements() throws RecognitionException {
		return statements(0);
	}

	private StatementsContext statements(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		StatementsContext _localctx = new StatementsContext(_ctx, _parentState);
		StatementsContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_statements, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new StatementsInitialContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			}
			_ctx.stop = _input.LT(-1);
			setState(19);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new StatementsAppendContext(new StatementsContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_statements);
					setState(15);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(16);
					statement();
					}
					} 
				}
				setState(21);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	 
		public StatementContext() { }
		public void copyFrom(StatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStatementContext extends StatementContext {
		public TerminalNode BREAK() { return getToken(CodeCalcParser.BREAK, 0); }
		public TerminalNode Terminator() { return getToken(CodeCalcParser.Terminator, 0); }
		public BreakStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignIndexStatementContext extends StatementContext {
		public Token identifier;
		public TerminalNode L_S_BRACE() { return getToken(CodeCalcParser.L_S_BRACE, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode R_S_BRACE() { return getToken(CodeCalcParser.R_S_BRACE, 0); }
		public TerminalNode ASSIGN() { return getToken(CodeCalcParser.ASSIGN, 0); }
		public TerminalNode Terminator() { return getToken(CodeCalcParser.Terminator, 0); }
		public TerminalNode Identifier() { return getToken(CodeCalcParser.Identifier, 0); }
		public AssignIndexStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionStatementContext extends StatementContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CodeCalcParser.Terminator, 0); }
		public ExpressionStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayStatementContext extends StatementContext {
		public Token identifier;
		public TerminalNode ARRAY() { return getToken(CodeCalcParser.ARRAY, 0); }
		public TerminalNode L_S_BRACE() { return getToken(CodeCalcParser.L_S_BRACE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode R_S_BRACE() { return getToken(CodeCalcParser.R_S_BRACE, 0); }
		public TerminalNode Terminator() { return getToken(CodeCalcParser.Terminator, 0); }
		public TerminalNode Identifier() { return getToken(CodeCalcParser.Identifier, 0); }
		public ArrayStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PassThroughStatementContext extends StatementContext {
		public While_statementContext while_statement() {
			return getRuleContext(While_statementContext.class,0);
		}
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public PassThroughStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignStatementContext extends StatementContext {
		public Token identifier;
		public TerminalNode ASSIGN() { return getToken(CodeCalcParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CodeCalcParser.Terminator, 0); }
		public TerminalNode Identifier() { return getToken(CodeCalcParser.Identifier, 0); }
		public AssignStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStatementContext extends StatementContext {
		public TerminalNode CONTINUE() { return getToken(CodeCalcParser.CONTINUE, 0); }
		public TerminalNode Terminator() { return getToken(CodeCalcParser.Terminator, 0); }
		public ContinueStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BlankStatementContext extends StatementContext {
		public TerminalNode Terminator() { return getToken(CodeCalcParser.Terminator, 0); }
		public BlankStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_statement);
		try {
			setState(52);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				_localctx = new PassThroughStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(22);
				while_statement();
				}
				break;
			case 2:
				_localctx = new PassThroughStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(23);
				if_statement();
				}
				break;
			case 3:
				_localctx = new BreakStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(24);
				match(BREAK);
				setState(25);
				match(Terminator);
				}
				break;
			case 4:
				_localctx = new ContinueStatementContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(26);
				match(CONTINUE);
				setState(27);
				match(Terminator);
				}
				break;
			case 5:
				_localctx = new AssignStatementContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(28);
				((AssignStatementContext)_localctx).identifier = match(Identifier);
				setState(29);
				match(ASSIGN);
				setState(30);
				expression(0);
				setState(31);
				match(Terminator);
				}
				break;
			case 6:
				_localctx = new AssignIndexStatementContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(33);
				((AssignIndexStatementContext)_localctx).identifier = match(Identifier);
				setState(34);
				match(L_S_BRACE);
				setState(35);
				expression(0);
				setState(36);
				match(R_S_BRACE);
				setState(37);
				match(ASSIGN);
				setState(38);
				expression(0);
				setState(39);
				match(Terminator);
				}
				break;
			case 7:
				_localctx = new ArrayStatementContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(41);
				match(ARRAY);
				setState(42);
				((ArrayStatementContext)_localctx).identifier = match(Identifier);
				setState(43);
				match(L_S_BRACE);
				setState(44);
				expression(0);
				setState(45);
				match(R_S_BRACE);
				setState(46);
				match(Terminator);
				}
				break;
			case 8:
				_localctx = new ExpressionStatementContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(48);
				expression(0);
				setState(49);
				match(Terminator);
				}
				break;
			case 9:
				_localctx = new BlankStatementContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(51);
				match(Terminator);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class While_statementContext extends ParserRuleContext {
		public While_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_statement; }
	 
		public While_statementContext() { }
		public void copyFrom(While_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class WhileStatementContext extends While_statementContext {
		public TerminalNode WHILE() { return getToken(CodeCalcParser.WHILE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode L_C_BRACE() { return getToken(CodeCalcParser.L_C_BRACE, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode R_C_BRACE() { return getToken(CodeCalcParser.R_C_BRACE, 0); }
		public WhileStatementContext(While_statementContext ctx) { copyFrom(ctx); }
	}

	public final While_statementContext while_statement() throws RecognitionException {
		While_statementContext _localctx = new While_statementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_while_statement);
		try {
			_localctx = new WhileStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			match(WHILE);
			setState(55);
			expression(0);
			setState(56);
			match(L_C_BRACE);
			setState(57);
			statements(0);
			setState(58);
			match(R_C_BRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_statementContext extends ParserRuleContext {
		public If_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_statement; }
	 
		public If_statementContext() { }
		public void copyFrom(If_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends If_statementContext {
		public TerminalNode IF() { return getToken(CodeCalcParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode L_C_BRACE() { return getToken(CodeCalcParser.L_C_BRACE, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode R_C_BRACE() { return getToken(CodeCalcParser.R_C_BRACE, 0); }
		public IfStatementContext(If_statementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfElseIfStatementContext extends If_statementContext {
		public TerminalNode IF() { return getToken(CodeCalcParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode L_C_BRACE() { return getToken(CodeCalcParser.L_C_BRACE, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode R_C_BRACE() { return getToken(CodeCalcParser.R_C_BRACE, 0); }
		public TerminalNode ELSE() { return getToken(CodeCalcParser.ELSE, 0); }
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public IfElseIfStatementContext(If_statementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfElseStatementContext extends If_statementContext {
		public TerminalNode IF() { return getToken(CodeCalcParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<TerminalNode> L_C_BRACE() { return getTokens(CodeCalcParser.L_C_BRACE); }
		public TerminalNode L_C_BRACE(int i) {
			return getToken(CodeCalcParser.L_C_BRACE, i);
		}
		public List<StatementsContext> statements() {
			return getRuleContexts(StatementsContext.class);
		}
		public StatementsContext statements(int i) {
			return getRuleContext(StatementsContext.class,i);
		}
		public List<TerminalNode> R_C_BRACE() { return getTokens(CodeCalcParser.R_C_BRACE); }
		public TerminalNode R_C_BRACE(int i) {
			return getToken(CodeCalcParser.R_C_BRACE, i);
		}
		public TerminalNode ELSE() { return getToken(CodeCalcParser.ELSE, 0); }
		public IfElseStatementContext(If_statementContext ctx) { copyFrom(ctx); }
	}

	public final If_statementContext if_statement() throws RecognitionException {
		If_statementContext _localctx = new If_statementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_if_statement);
		try {
			setState(84);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				_localctx = new IfStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(60);
				match(IF);
				setState(61);
				expression(0);
				setState(62);
				match(L_C_BRACE);
				setState(63);
				statements(0);
				setState(64);
				match(R_C_BRACE);
				}
				break;
			case 2:
				_localctx = new IfElseIfStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(66);
				match(IF);
				setState(67);
				expression(0);
				setState(68);
				match(L_C_BRACE);
				setState(69);
				statements(0);
				setState(70);
				match(R_C_BRACE);
				setState(71);
				match(ELSE);
				setState(72);
				if_statement();
				}
				break;
			case 3:
				_localctx = new IfElseStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(74);
				match(IF);
				setState(75);
				expression(0);
				setState(76);
				match(L_C_BRACE);
				setState(77);
				statements(0);
				setState(78);
				match(R_C_BRACE);
				setState(79);
				match(ELSE);
				setState(80);
				match(L_C_BRACE);
				setState(81);
				statements(0);
				setState(82);
				match(R_C_BRACE);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	 
		public ExpressionContext() { }
		public void copyFrom(ExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BinaryExpressionContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OP_MUL() { return getToken(CodeCalcParser.OP_MUL, 0); }
		public TerminalNode OP_DIV() { return getToken(CodeCalcParser.OP_DIV, 0); }
		public TerminalNode OP_MOD() { return getToken(CodeCalcParser.OP_MOD, 0); }
		public TerminalNode OP_ADD() { return getToken(CodeCalcParser.OP_ADD, 0); }
		public TerminalNode OP_SUB() { return getToken(CodeCalcParser.OP_SUB, 0); }
		public TerminalNode OP_GRT() { return getToken(CodeCalcParser.OP_GRT, 0); }
		public TerminalNode OP_LST() { return getToken(CodeCalcParser.OP_LST, 0); }
		public TerminalNode OP_GTE() { return getToken(CodeCalcParser.OP_GTE, 0); }
		public TerminalNode OP_LTE() { return getToken(CodeCalcParser.OP_LTE, 0); }
		public TerminalNode OP_EQU() { return getToken(CodeCalcParser.OP_EQU, 0); }
		public TerminalNode OP_NEQ() { return getToken(CodeCalcParser.OP_NEQ, 0); }
		public TerminalNode OP_AND() { return getToken(CodeCalcParser.OP_AND, 0); }
		public TerminalNode OP_OR() { return getToken(CodeCalcParser.OP_OR, 0); }
		public BinaryExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralExpressionContext extends ExpressionContext {
		public Token number;
		public TerminalNode Number() { return getToken(CodeCalcParser.Number, 0); }
		public LiteralExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UnaryExpressionContext extends ExpressionContext {
		public Token op;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode R_BRACE() { return getToken(CodeCalcParser.R_BRACE, 0); }
		public TerminalNode L_BRACE() { return getToken(CodeCalcParser.L_BRACE, 0); }
		public TerminalNode OP_NOT() { return getToken(CodeCalcParser.OP_NOT, 0); }
		public TerminalNode OP_ADD() { return getToken(CodeCalcParser.OP_ADD, 0); }
		public TerminalNode OP_SUB() { return getToken(CodeCalcParser.OP_SUB, 0); }
		public UnaryExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AccessExpressionContext extends ExpressionContext {
		public Token identifier;
		public TerminalNode L_S_BRACE() { return getToken(CodeCalcParser.L_S_BRACE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode R_S_BRACE() { return getToken(CodeCalcParser.R_S_BRACE, 0); }
		public TerminalNode Identifier() { return getToken(CodeCalcParser.Identifier, 0); }
		public AccessExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierExpressionContext extends ExpressionContext {
		public Token identifier;
		public TerminalNode Identifier() { return getToken(CodeCalcParser.Identifier, 0); }
		public IdentifierExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				_localctx = new UnaryExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(87);
				((UnaryExpressionContext)_localctx).op = match(L_BRACE);
				setState(88);
				expression(0);
				setState(89);
				match(R_BRACE);
				}
				break;
			case 2:
				{
				_localctx = new LiteralExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(91);
				((LiteralExpressionContext)_localctx).number = match(Number);
				}
				break;
			case 3:
				{
				_localctx = new AccessExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(92);
				((AccessExpressionContext)_localctx).identifier = match(Identifier);
				setState(93);
				match(L_S_BRACE);
				setState(94);
				expression(0);
				setState(95);
				match(R_S_BRACE);
				}
				break;
			case 4:
				{
				_localctx = new IdentifierExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(97);
				((IdentifierExpressionContext)_localctx).identifier = match(Identifier);
				}
				break;
			case 5:
				{
				_localctx = new UnaryExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(98);
				((UnaryExpressionContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 17920L) != 0)) ) {
					((UnaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(99);
				expression(6);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(119);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(117);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(102);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(103);
						((BinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 14336L) != 0)) ) {
							((BinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(104);
						expression(6);
						}
						break;
					case 2:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(105);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(106);
						((BinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_ADD || _la==OP_SUB) ) {
							((BinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(107);
						expression(5);
						}
						break;
					case 3:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(108);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(109);
						((BinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 491520L) != 0)) ) {
							((BinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(110);
						expression(4);
						}
						break;
					case 4:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(111);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(112);
						((BinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_EQU || _la==OP_NEQ) ) {
							((BinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(113);
						expression(3);
						}
						break;
					case 5:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(114);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(115);
						((BinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_OR || _la==OP_AND) ) {
							((BinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(116);
						expression(2);
						}
						break;
					}
					} 
				}
				setState(121);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return statements_sempred((StatementsContext)_localctx, predIndex);
		case 5:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean statements_sempred(StatementsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 5);
		case 2:
			return precpred(_ctx, 4);
		case 3:
			return precpred(_ctx, 3);
		case 4:
			return precpred(_ctx, 2);
		case 5:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001!{\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002\u0002"+
		"\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002\u0005"+
		"\u0007\u0005\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0005\u0001\u0012\b\u0001\n\u0001\f\u0001\u0015\t\u0001\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u00025\b"+
		"\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0003\u0004U\b\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005e\b"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005v\b\u0005\n\u0005"+
		"\f\u0005y\t\u0005\u0001\u0005\u0000\u0002\u0002\n\u0006\u0000\u0002\u0004"+
		"\u0006\b\n\u0000\u0006\u0002\u0000\t\n\u000e\u000e\u0001\u0000\u000b\r"+
		"\u0001\u0000\t\n\u0001\u0000\u000f\u0012\u0001\u0000\u0013\u0014\u0001"+
		"\u0000\u0015\u0016\u0088\u0000\f\u0001\u0000\u0000\u0000\u0002\u000e\u0001"+
		"\u0000\u0000\u0000\u00044\u0001\u0000\u0000\u0000\u00066\u0001\u0000\u0000"+
		"\u0000\bT\u0001\u0000\u0000\u0000\nd\u0001\u0000\u0000\u0000\f\r\u0003"+
		"\u0002\u0001\u0000\r\u0001\u0001\u0000\u0000\u0000\u000e\u0013\u0006\u0001"+
		"\uffff\uffff\u0000\u000f\u0010\n\u0001\u0000\u0000\u0010\u0012\u0003\u0004"+
		"\u0002\u0000\u0011\u000f\u0001\u0000\u0000\u0000\u0012\u0015\u0001\u0000"+
		"\u0000\u0000\u0013\u0011\u0001\u0000\u0000\u0000\u0013\u0014\u0001\u0000"+
		"\u0000\u0000\u0014\u0003\u0001\u0000\u0000\u0000\u0015\u0013\u0001\u0000"+
		"\u0000\u0000\u00165\u0003\u0006\u0003\u0000\u00175\u0003\b\u0004\u0000"+
		"\u0018\u0019\u0005\u001b\u0000\u0000\u00195\u0005\u0001\u0000\u0000\u001a"+
		"\u001b\u0005\u001a\u0000\u0000\u001b5\u0005\u0001\u0000\u0000\u001c\u001d"+
		"\u0005\u001e\u0000\u0000\u001d\u001e\u0005\b\u0000\u0000\u001e\u001f\u0003"+
		"\n\u0005\u0000\u001f \u0005\u0001\u0000\u0000 5\u0001\u0000\u0000\u0000"+
		"!\"\u0005\u001e\u0000\u0000\"#\u0005\u0006\u0000\u0000#$\u0003\n\u0005"+
		"\u0000$%\u0005\u0007\u0000\u0000%&\u0005\b\u0000\u0000&\'\u0003\n\u0005"+
		"\u0000\'(\u0005\u0001\u0000\u0000(5\u0001\u0000\u0000\u0000)*\u0005\u001c"+
		"\u0000\u0000*+\u0005\u001e\u0000\u0000+,\u0005\u0006\u0000\u0000,-\u0003"+
		"\n\u0005\u0000-.\u0005\u0007\u0000\u0000./\u0005\u0001\u0000\u0000/5\u0001"+
		"\u0000\u0000\u000001\u0003\n\u0005\u000012\u0005\u0001\u0000\u000025\u0001"+
		"\u0000\u0000\u000035\u0005\u0001\u0000\u00004\u0016\u0001\u0000\u0000"+
		"\u00004\u0017\u0001\u0000\u0000\u00004\u0018\u0001\u0000\u0000\u00004"+
		"\u001a\u0001\u0000\u0000\u00004\u001c\u0001\u0000\u0000\u00004!\u0001"+
		"\u0000\u0000\u00004)\u0001\u0000\u0000\u000040\u0001\u0000\u0000\u0000"+
		"43\u0001\u0000\u0000\u00005\u0005\u0001\u0000\u0000\u000067\u0005\u0019"+
		"\u0000\u000078\u0003\n\u0005\u000089\u0005\u0004\u0000\u00009:\u0003\u0002"+
		"\u0001\u0000:;\u0005\u0005\u0000\u0000;\u0007\u0001\u0000\u0000\u0000"+
		"<=\u0005\u0017\u0000\u0000=>\u0003\n\u0005\u0000>?\u0005\u0004\u0000\u0000"+
		"?@\u0003\u0002\u0001\u0000@A\u0005\u0005\u0000\u0000AU\u0001\u0000\u0000"+
		"\u0000BC\u0005\u0017\u0000\u0000CD\u0003\n\u0005\u0000DE\u0005\u0004\u0000"+
		"\u0000EF\u0003\u0002\u0001\u0000FG\u0005\u0005\u0000\u0000GH\u0005\u0018"+
		"\u0000\u0000HI\u0003\b\u0004\u0000IU\u0001\u0000\u0000\u0000JK\u0005\u0017"+
		"\u0000\u0000KL\u0003\n\u0005\u0000LM\u0005\u0004\u0000\u0000MN\u0003\u0002"+
		"\u0001\u0000NO\u0005\u0005\u0000\u0000OP\u0005\u0018\u0000\u0000PQ\u0005"+
		"\u0004\u0000\u0000QR\u0003\u0002\u0001\u0000RS\u0005\u0005\u0000\u0000"+
		"SU\u0001\u0000\u0000\u0000T<\u0001\u0000\u0000\u0000TB\u0001\u0000\u0000"+
		"\u0000TJ\u0001\u0000\u0000\u0000U\t\u0001\u0000\u0000\u0000VW\u0006\u0005"+
		"\uffff\uffff\u0000WX\u0005\u0002\u0000\u0000XY\u0003\n\u0005\u0000YZ\u0005"+
		"\u0003\u0000\u0000Ze\u0001\u0000\u0000\u0000[e\u0005\u001d\u0000\u0000"+
		"\\]\u0005\u001e\u0000\u0000]^\u0005\u0006\u0000\u0000^_\u0003\n\u0005"+
		"\u0000_`\u0005\u0007\u0000\u0000`e\u0001\u0000\u0000\u0000ae\u0005\u001e"+
		"\u0000\u0000bc\u0007\u0000\u0000\u0000ce\u0003\n\u0005\u0006dV\u0001\u0000"+
		"\u0000\u0000d[\u0001\u0000\u0000\u0000d\\\u0001\u0000\u0000\u0000da\u0001"+
		"\u0000\u0000\u0000db\u0001\u0000\u0000\u0000ew\u0001\u0000\u0000\u0000"+
		"fg\n\u0005\u0000\u0000gh\u0007\u0001\u0000\u0000hv\u0003\n\u0005\u0006"+
		"ij\n\u0004\u0000\u0000jk\u0007\u0002\u0000\u0000kv\u0003\n\u0005\u0005"+
		"lm\n\u0003\u0000\u0000mn\u0007\u0003\u0000\u0000nv\u0003\n\u0005\u0004"+
		"op\n\u0002\u0000\u0000pq\u0007\u0004\u0000\u0000qv\u0003\n\u0005\u0003"+
		"rs\n\u0001\u0000\u0000st\u0007\u0005\u0000\u0000tv\u0003\n\u0005\u0002"+
		"uf\u0001\u0000\u0000\u0000ui\u0001\u0000\u0000\u0000ul\u0001\u0000\u0000"+
		"\u0000uo\u0001\u0000\u0000\u0000ur\u0001\u0000\u0000\u0000vy\u0001\u0000"+
		"\u0000\u0000wu\u0001\u0000\u0000\u0000wx\u0001\u0000\u0000\u0000x\u000b"+
		"\u0001\u0000\u0000\u0000yw\u0001\u0000\u0000\u0000\u0006\u00134Tduw";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}