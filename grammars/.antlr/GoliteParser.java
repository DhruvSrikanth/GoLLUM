// Generated from /Users/dhruvsrikanth/Work/Projects/MPCS/Quarter_5/Compilers/Project/gollum/grammars/GoliteParser.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GoliteParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		ASSIGN=1, AND=2, OR=3, NOT=4, LT=5, GT=6, LE=7, GE=8, EQ=9, NE=10, PLUS=11, 
		MINUS=12, MULT=13, DIV=14, LPAREN=15, RPAREN=16, LBRACE=17, RBRACE=18, 
		SEMICOLON=19, COMMA=20, VAR=21, TYPE=22, NEW=23, DELETE=24, STRUCT=25, 
		DOT=26, FUNC=27, RET=28, IF=29, ELSE=30, FOR=31, SCAN=32, PRINTF=33, INT_LIT=34, 
		STRING_LIT=35, BOOL_LIT=36, NIL_LIT=37, INT=38, BOOL=39, IDENT=40, WS=41, 
		COMMENT=42;
	public static final int
		RULE_program = 0, RULE_types = 1, RULE_typeDeclaration = 2, RULE_fields = 3, 
		RULE_morefields = 4, RULE_decl = 5, RULE_type = 6, RULE_declarations = 7, 
		RULE_declaration = 8, RULE_ids = 9, RULE_otherids = 10, RULE_functions = 11, 
		RULE_function = 12, RULE_parameters = 13, RULE_parametersPrime = 14, RULE_parametersPPrime = 15, 
		RULE_returnType = 16, RULE_statements = 17, RULE_statement = 18, RULE_read = 19, 
		RULE_block = 20, RULE_delete = 21, RULE_assignment = 22, RULE_print = 23, 
		RULE_conditional = 24, RULE_loop = 25, RULE_returnRule = 26, RULE_invocation = 27, 
		RULE_arguments = 28, RULE_lValue = 29, RULE_expression = 30, RULE_boolTerm = 31, 
		RULE_equalTerm = 32, RULE_relationTerm = 33, RULE_simpleTerm = 34, RULE_term = 35, 
		RULE_unaryTerm = 36, RULE_selectorTerm = 37, RULE_selectorTermPrime = 38, 
		RULE_subfactor = 39, RULE_functioncall = 40, RULE_allocation = 41, RULE_factor = 42;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "types", "typeDeclaration", "fields", "morefields", "decl", 
			"type", "declarations", "declaration", "ids", "otherids", "functions", 
			"function", "parameters", "parametersPrime", "parametersPPrime", "returnType", 
			"statements", "statement", "read", "block", "delete", "assignment", "print", 
			"conditional", "loop", "returnRule", "invocation", "arguments", "lValue", 
			"expression", "boolTerm", "equalTerm", "relationTerm", "simpleTerm", 
			"term", "unaryTerm", "selectorTerm", "selectorTermPrime", "subfactor", 
			"functioncall", "allocation", "factor"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'&&'", "'||'", "'!'", "'<'", "'>'", "'<='", "'>='", "'=='", 
			"'!='", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'{'", "'}'", "';'", 
			"','", "'var'", "'type'", "'new'", "'delete'", "'struct'", "'.'", "'func'", 
			"'return'", "'if'", "'else'", "'for'", "'scan'", "'printf'", null, null, 
			null, "'nil'", "'int'", "'bool'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ASSIGN", "AND", "OR", "NOT", "LT", "GT", "LE", "GE", "EQ", "NE", 
			"PLUS", "MINUS", "MULT", "DIV", "LPAREN", "RPAREN", "LBRACE", "RBRACE", 
			"SEMICOLON", "COMMA", "VAR", "TYPE", "NEW", "DELETE", "STRUCT", "DOT", 
			"FUNC", "RET", "IF", "ELSE", "FOR", "SCAN", "PRINTF", "INT_LIT", "STRING_LIT", 
			"BOOL_LIT", "NIL_LIT", "INT", "BOOL", "IDENT", "WS", "COMMENT"
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
	public String getGrammarFileName() { return "GoliteParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GoliteParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public TypesContext ty;
		public DeclarationsContext dcls;
		public FunctionsContext fns;
		public TerminalNode EOF() { return getToken(GoliteParser.EOF, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public DeclarationsContext declarations() {
			return getRuleContext(DeclarationsContext.class,0);
		}
		public FunctionsContext functions() {
			return getRuleContext(FunctionsContext.class,0);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			((ProgramContext)_localctx).ty = types();
			setState(87);
			((ProgramContext)_localctx).dcls = declarations();
			setState(88);
			((ProgramContext)_localctx).fns = functions();
			setState(89);
			match(EOF);
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

	public static class TypesContext extends ParserRuleContext {
		public List<TypeDeclarationContext> typeDeclaration() {
			return getRuleContexts(TypeDeclarationContext.class);
		}
		public TypeDeclarationContext typeDeclaration(int i) {
			return getRuleContext(TypeDeclarationContext.class,i);
		}
		public TypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types; }
	}

	public final TypesContext types() throws RecognitionException {
		TypesContext _localctx = new TypesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_types);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TYPE) {
				{
				{
				setState(91);
				typeDeclaration();
				}
				}
				setState(96);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class TypeDeclarationContext extends ParserRuleContext {
		public Token id;
		public FieldsContext flds;
		public TerminalNode TYPE() { return getToken(GoliteParser.TYPE, 0); }
		public TerminalNode STRUCT() { return getToken(GoliteParser.STRUCT, 0); }
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public FieldsContext fields() {
			return getRuleContext(FieldsContext.class,0);
		}
		public TypeDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDeclaration; }
	}

	public final TypeDeclarationContext typeDeclaration() throws RecognitionException {
		TypeDeclarationContext _localctx = new TypeDeclarationContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_typeDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(97);
			match(TYPE);
			setState(98);
			((TypeDeclarationContext)_localctx).id = match(IDENT);
			setState(99);
			match(STRUCT);
			setState(100);
			match(LBRACE);
			setState(101);
			((TypeDeclarationContext)_localctx).flds = fields();
			setState(102);
			match(RBRACE);
			setState(103);
			match(SEMICOLON);
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

	public static class FieldsContext extends ParserRuleContext {
		public DeclContext dcl;
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public List<MorefieldsContext> morefields() {
			return getRuleContexts(MorefieldsContext.class);
		}
		public MorefieldsContext morefields(int i) {
			return getRuleContext(MorefieldsContext.class,i);
		}
		public FieldsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fields; }
	}

	public final FieldsContext fields() throws RecognitionException {
		FieldsContext _localctx = new FieldsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_fields);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			((FieldsContext)_localctx).dcl = decl();
			setState(106);
			match(SEMICOLON);
			setState(110);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENT) {
				{
				{
				setState(107);
				morefields();
				}
				}
				setState(112);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class MorefieldsContext extends ParserRuleContext {
		public DeclContext dcl;
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public MorefieldsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_morefields; }
	}

	public final MorefieldsContext morefields() throws RecognitionException {
		MorefieldsContext _localctx = new MorefieldsContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_morefields);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(113);
			((MorefieldsContext)_localctx).dcl = decl();
			setState(114);
			match(SEMICOLON);
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

	public static class DeclContext extends ParserRuleContext {
		public Token id;
		public TypeContext ty;
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public DeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decl; }
	}

	public final DeclContext decl() throws RecognitionException {
		DeclContext _localctx = new DeclContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_decl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			((DeclContext)_localctx).id = match(IDENT);
			setState(117);
			((DeclContext)_localctx).ty = type();
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

	public static class TypeContext extends ParserRuleContext {
		public Token id;
		public TerminalNode INT() { return getToken(GoliteParser.INT, 0); }
		public TerminalNode BOOL() { return getToken(GoliteParser.BOOL, 0); }
		public TerminalNode MULT() { return getToken(GoliteParser.MULT, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_type);
		try {
			setState(123);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(119);
				match(INT);
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 2);
				{
				setState(120);
				match(BOOL);
				}
				break;
			case MULT:
				enterOuterAlt(_localctx, 3);
				{
				{
				setState(121);
				match(MULT);
				setState(122);
				((TypeContext)_localctx).id = match(IDENT);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	public static class DeclarationsContext extends ParserRuleContext {
		public List<DeclarationContext> declaration() {
			return getRuleContexts(DeclarationContext.class);
		}
		public DeclarationContext declaration(int i) {
			return getRuleContext(DeclarationContext.class,i);
		}
		public DeclarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarations; }
	}

	public final DeclarationsContext declarations() throws RecognitionException {
		DeclarationsContext _localctx = new DeclarationsContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(125);
				declaration();
				}
				}
				setState(130);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class DeclarationContext extends ParserRuleContext {
		public IdsContext idx;
		public TypeContext ty;
		public TerminalNode VAR() { return getToken(GoliteParser.VAR, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public IdsContext ids() {
			return getRuleContext(IdsContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(131);
			match(VAR);
			setState(132);
			((DeclarationContext)_localctx).idx = ids();
			setState(133);
			((DeclarationContext)_localctx).ty = type();
			setState(134);
			match(SEMICOLON);
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

	public static class IdsContext extends ParserRuleContext {
		public Token id;
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public List<OtheridsContext> otherids() {
			return getRuleContexts(OtheridsContext.class);
		}
		public OtheridsContext otherids(int i) {
			return getRuleContext(OtheridsContext.class,i);
		}
		public IdsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ids; }
	}

	public final IdsContext ids() throws RecognitionException {
		IdsContext _localctx = new IdsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_ids);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			((IdsContext)_localctx).id = match(IDENT);
			setState(140);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(137);
				otherids();
				}
				}
				setState(142);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class OtheridsContext extends ParserRuleContext {
		public Token id;
		public TerminalNode COMMA() { return getToken(GoliteParser.COMMA, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public OtheridsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_otherids; }
	}

	public final OtheridsContext otherids() throws RecognitionException {
		OtheridsContext _localctx = new OtheridsContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_otherids);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			match(COMMA);
			setState(144);
			((OtheridsContext)_localctx).id = match(IDENT);
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

	public static class FunctionsContext extends ParserRuleContext {
		public List<FunctionContext> function() {
			return getRuleContexts(FunctionContext.class);
		}
		public FunctionContext function(int i) {
			return getRuleContext(FunctionContext.class,i);
		}
		public FunctionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functions; }
	}

	public final FunctionsContext functions() throws RecognitionException {
		FunctionsContext _localctx = new FunctionsContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNC) {
				{
				{
				setState(146);
				function();
				}
				}
				setState(151);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class FunctionContext extends ParserRuleContext {
		public Token id;
		public ParametersContext params;
		public ReturnTypeContext rty;
		public DeclarationsContext dcls;
		public StatementsContext stmts;
		public TerminalNode FUNC() { return getToken(GoliteParser.FUNC, 0); }
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public DeclarationsContext declarations() {
			return getRuleContext(DeclarationsContext.class,0);
		}
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public ReturnTypeContext returnType() {
			return getRuleContext(ReturnTypeContext.class,0);
		}
		public FunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function; }
	}

	public final FunctionContext function() throws RecognitionException {
		FunctionContext _localctx = new FunctionContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_function);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			match(FUNC);
			setState(153);
			((FunctionContext)_localctx).id = match(IDENT);
			setState(154);
			((FunctionContext)_localctx).params = parameters();
			setState(156);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MULT) | (1L << INT) | (1L << BOOL))) != 0)) {
				{
				setState(155);
				((FunctionContext)_localctx).rty = returnType();
				}
			}

			setState(158);
			match(LBRACE);
			setState(159);
			((FunctionContext)_localctx).dcls = declarations();
			setState(160);
			((FunctionContext)_localctx).stmts = statements();
			setState(161);
			match(RBRACE);
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

	public static class ParametersContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public ParametersPrimeContext parametersPrime() {
			return getRuleContext(ParametersPrimeContext.class,0);
		}
		public ParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters; }
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(163);
			match(LPAREN);
			setState(165);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENT) {
				{
				setState(164);
				parametersPrime();
				}
			}

			setState(167);
			match(RPAREN);
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

	public static class ParametersPrimeContext extends ParserRuleContext {
		public DeclContext dcl;
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public List<ParametersPPrimeContext> parametersPPrime() {
			return getRuleContexts(ParametersPPrimeContext.class);
		}
		public ParametersPPrimeContext parametersPPrime(int i) {
			return getRuleContext(ParametersPPrimeContext.class,i);
		}
		public ParametersPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametersPrime; }
	}

	public final ParametersPrimeContext parametersPrime() throws RecognitionException {
		ParametersPrimeContext _localctx = new ParametersPrimeContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_parametersPrime);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			((ParametersPrimeContext)_localctx).dcl = decl();
			setState(173);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(170);
				parametersPPrime();
				}
				}
				setState(175);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class ParametersPPrimeContext extends ParserRuleContext {
		public DeclContext dcl;
		public TerminalNode COMMA() { return getToken(GoliteParser.COMMA, 0); }
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public ParametersPPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametersPPrime; }
	}

	public final ParametersPPrimeContext parametersPPrime() throws RecognitionException {
		ParametersPPrimeContext _localctx = new ParametersPPrimeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_parametersPPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(176);
			match(COMMA);
			setState(177);
			((ParametersPPrimeContext)_localctx).dcl = decl();
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

	public static class ReturnTypeContext extends ParserRuleContext {
		public TypeContext ty;
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public ReturnTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnType; }
	}

	public final ReturnTypeContext returnType() throws RecognitionException {
		ReturnTypeContext _localctx = new ReturnTypeContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_returnType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			((ReturnTypeContext)_localctx).ty = type();
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

	public static class StatementsContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public StatementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statements; }
	}

	public final StatementsContext statements() throws RecognitionException {
		StatementsContext _localctx = new StatementsContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_statements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LBRACE) | (1L << DELETE) | (1L << RET) | (1L << IF) | (1L << FOR) | (1L << SCAN) | (1L << PRINTF) | (1L << IDENT))) != 0)) {
				{
				{
				setState(181);
				statement();
				}
				}
				setState(186);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class StatementContext extends ParserRuleContext {
		public BlockContext bl;
		public AssignmentContext asmt;
		public PrintContext prnt;
		public DeleteContext del;
		public ReadContext rd;
		public ConditionalContext cond;
		public LoopContext lp;
		public ReturnRuleContext ret;
		public InvocationContext invoke;
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public PrintContext print() {
			return getRuleContext(PrintContext.class,0);
		}
		public DeleteContext delete() {
			return getRuleContext(DeleteContext.class,0);
		}
		public ReadContext read() {
			return getRuleContext(ReadContext.class,0);
		}
		public ConditionalContext conditional() {
			return getRuleContext(ConditionalContext.class,0);
		}
		public LoopContext loop() {
			return getRuleContext(LoopContext.class,0);
		}
		public ReturnRuleContext returnRule() {
			return getRuleContext(ReturnRuleContext.class,0);
		}
		public InvocationContext invocation() {
			return getRuleContext(InvocationContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_statement);
		try {
			setState(196);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(187);
				((StatementContext)_localctx).bl = block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(188);
				((StatementContext)_localctx).asmt = assignment();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(189);
				((StatementContext)_localctx).prnt = print();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(190);
				((StatementContext)_localctx).del = delete();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(191);
				((StatementContext)_localctx).rd = read();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(192);
				((StatementContext)_localctx).cond = conditional();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(193);
				((StatementContext)_localctx).lp = loop();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(194);
				((StatementContext)_localctx).ret = returnRule();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(195);
				((StatementContext)_localctx).invoke = invocation();
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

	public static class ReadContext extends ParserRuleContext {
		public TerminalNode SCAN() { return getToken(GoliteParser.SCAN, 0); }
		public LValueContext lValue() {
			return getRuleContext(LValueContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public ReadContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_read; }
	}

	public final ReadContext read() throws RecognitionException {
		ReadContext _localctx = new ReadContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_read);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(198);
			match(SCAN);
			setState(199);
			lValue();
			setState(200);
			match(SEMICOLON);
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

	public static class BlockContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
			match(LBRACE);
			setState(203);
			statements();
			setState(204);
			match(RBRACE);
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

	public static class DeleteContext extends ParserRuleContext {
		public TerminalNode DELETE() { return getToken(GoliteParser.DELETE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public DeleteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_delete; }
	}

	public final DeleteContext delete() throws RecognitionException {
		DeleteContext _localctx = new DeleteContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_delete);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(206);
			match(DELETE);
			setState(207);
			expression();
			setState(208);
			match(SEMICOLON);
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

	public static class AssignmentContext extends ParserRuleContext {
		public LValueContext lValue() {
			return getRuleContext(LValueContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(GoliteParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(210);
			lValue();
			setState(211);
			match(ASSIGN);
			setState(212);
			expression();
			setState(213);
			match(SEMICOLON);
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

	public static class PrintContext extends ParserRuleContext {
		public TerminalNode PRINTF() { return getToken(GoliteParser.PRINTF, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoliteParser.STRING_LIT, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public PrintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print; }
	}

	public final PrintContext print() throws RecognitionException {
		PrintContext _localctx = new PrintContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_print);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(215);
			match(PRINTF);
			setState(216);
			match(LPAREN);
			setState(217);
			match(STRING_LIT);
			setState(222);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(218);
				match(COMMA);
				setState(219);
				expression();
				}
				}
				setState(224);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(225);
			match(RPAREN);
			setState(226);
			match(SEMICOLON);
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

	public static class ConditionalContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(GoliteParser.IF, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(GoliteParser.ELSE, 0); }
		public ConditionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional; }
	}

	public final ConditionalContext conditional() throws RecognitionException {
		ConditionalContext _localctx = new ConditionalContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			match(IF);
			setState(229);
			match(LPAREN);
			setState(230);
			expression();
			setState(231);
			match(RPAREN);
			setState(232);
			block();
			setState(235);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(233);
				match(ELSE);
				setState(234);
				block();
				}
			}

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

	public static class LoopContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(GoliteParser.FOR, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public LoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loop; }
	}

	public final LoopContext loop() throws RecognitionException {
		LoopContext _localctx = new LoopContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_loop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(237);
			match(FOR);
			setState(238);
			match(LPAREN);
			setState(239);
			expression();
			setState(240);
			match(RPAREN);
			setState(241);
			block();
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

	public static class ReturnRuleContext extends ParserRuleContext {
		public TerminalNode RET() { return getToken(GoliteParser.RET, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ReturnRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnRule; }
	}

	public final ReturnRuleContext returnRule() throws RecognitionException {
		ReturnRuleContext _localctx = new ReturnRuleContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_returnRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			match(RET);
			setState(245);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NOT) | (1L << MINUS) | (1L << LPAREN) | (1L << NEW) | (1L << INT_LIT) | (1L << BOOL_LIT) | (1L << NIL_LIT) | (1L << IDENT))) != 0)) {
				{
				setState(244);
				expression();
				}
			}

			setState(247);
			match(SEMICOLON);
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

	public static class InvocationContext extends ParserRuleContext {
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public InvocationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_invocation; }
	}

	public final InvocationContext invocation() throws RecognitionException {
		InvocationContext _localctx = new InvocationContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_invocation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(249);
			match(IDENT);
			setState(250);
			arguments();
			setState(251);
			match(SEMICOLON);
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

	public static class ArgumentsContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(253);
			match(LPAREN);
			setState(262);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NOT) | (1L << MINUS) | (1L << LPAREN) | (1L << NEW) | (1L << INT_LIT) | (1L << BOOL_LIT) | (1L << NIL_LIT) | (1L << IDENT))) != 0)) {
				{
				setState(254);
				expression();
				setState(259);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(255);
					match(COMMA);
					setState(256);
					expression();
					}
					}
					setState(261);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(264);
			match(RPAREN);
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

	public static class LValueContext extends ParserRuleContext {
		public List<TerminalNode> IDENT() { return getTokens(GoliteParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(GoliteParser.IDENT, i);
		}
		public List<TerminalNode> DOT() { return getTokens(GoliteParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(GoliteParser.DOT, i);
		}
		public LValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lValue; }
	}

	public final LValueContext lValue() throws RecognitionException {
		LValueContext _localctx = new LValueContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_lValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(266);
			match(IDENT);
			setState(271);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(267);
				match(DOT);
				setState(268);
				match(IDENT);
				}
				}
				setState(273);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class ExpressionContext extends ParserRuleContext {
		public List<BoolTermContext> boolTerm() {
			return getRuleContexts(BoolTermContext.class);
		}
		public BoolTermContext boolTerm(int i) {
			return getRuleContext(BoolTermContext.class,i);
		}
		public List<TerminalNode> OR() { return getTokens(GoliteParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(GoliteParser.OR, i);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(274);
			boolTerm();
			setState(279);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(275);
				match(OR);
				setState(276);
				boolTerm();
				}
				}
				setState(281);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class BoolTermContext extends ParserRuleContext {
		public List<EqualTermContext> equalTerm() {
			return getRuleContexts(EqualTermContext.class);
		}
		public EqualTermContext equalTerm(int i) {
			return getRuleContext(EqualTermContext.class,i);
		}
		public List<TerminalNode> AND() { return getTokens(GoliteParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(GoliteParser.AND, i);
		}
		public BoolTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolTerm; }
	}

	public final BoolTermContext boolTerm() throws RecognitionException {
		BoolTermContext _localctx = new BoolTermContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_boolTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(282);
			equalTerm();
			setState(287);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(283);
				match(AND);
				setState(284);
				equalTerm();
				}
				}
				setState(289);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class EqualTermContext extends ParserRuleContext {
		public List<RelationTermContext> relationTerm() {
			return getRuleContexts(RelationTermContext.class);
		}
		public RelationTermContext relationTerm(int i) {
			return getRuleContext(RelationTermContext.class,i);
		}
		public List<TerminalNode> EQ() { return getTokens(GoliteParser.EQ); }
		public TerminalNode EQ(int i) {
			return getToken(GoliteParser.EQ, i);
		}
		public List<TerminalNode> NE() { return getTokens(GoliteParser.NE); }
		public TerminalNode NE(int i) {
			return getToken(GoliteParser.NE, i);
		}
		public EqualTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_equalTerm; }
	}

	public final EqualTermContext equalTerm() throws RecognitionException {
		EqualTermContext _localctx = new EqualTermContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_equalTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(290);
			relationTerm();
			setState(295);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EQ || _la==NE) {
				{
				{
				setState(291);
				_la = _input.LA(1);
				if ( !(_la==EQ || _la==NE) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(292);
				relationTerm();
				}
				}
				setState(297);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class RelationTermContext extends ParserRuleContext {
		public List<SimpleTermContext> simpleTerm() {
			return getRuleContexts(SimpleTermContext.class);
		}
		public SimpleTermContext simpleTerm(int i) {
			return getRuleContext(SimpleTermContext.class,i);
		}
		public List<TerminalNode> GE() { return getTokens(GoliteParser.GE); }
		public TerminalNode GE(int i) {
			return getToken(GoliteParser.GE, i);
		}
		public List<TerminalNode> LT() { return getTokens(GoliteParser.LT); }
		public TerminalNode LT(int i) {
			return getToken(GoliteParser.LT, i);
		}
		public List<TerminalNode> LE() { return getTokens(GoliteParser.LE); }
		public TerminalNode LE(int i) {
			return getToken(GoliteParser.LE, i);
		}
		public List<TerminalNode> GT() { return getTokens(GoliteParser.GT); }
		public TerminalNode GT(int i) {
			return getToken(GoliteParser.GT, i);
		}
		public RelationTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationTerm; }
	}

	public final RelationTermContext relationTerm() throws RecognitionException {
		RelationTermContext _localctx = new RelationTermContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_relationTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(298);
			simpleTerm();
			setState(303);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << LE) | (1L << GE))) != 0)) {
				{
				{
				setState(299);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << LE) | (1L << GE))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(300);
				simpleTerm();
				}
				}
				setState(305);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class SimpleTermContext extends ParserRuleContext {
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public List<TerminalNode> PLUS() { return getTokens(GoliteParser.PLUS); }
		public TerminalNode PLUS(int i) {
			return getToken(GoliteParser.PLUS, i);
		}
		public List<TerminalNode> MINUS() { return getTokens(GoliteParser.MINUS); }
		public TerminalNode MINUS(int i) {
			return getToken(GoliteParser.MINUS, i);
		}
		public SimpleTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleTerm; }
	}

	public final SimpleTermContext simpleTerm() throws RecognitionException {
		SimpleTermContext _localctx = new SimpleTermContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_simpleTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(306);
			term();
			setState(311);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PLUS || _la==MINUS) {
				{
				{
				setState(307);
				_la = _input.LA(1);
				if ( !(_la==PLUS || _la==MINUS) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(308);
				term();
				}
				}
				setState(313);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class TermContext extends ParserRuleContext {
		public List<UnaryTermContext> unaryTerm() {
			return getRuleContexts(UnaryTermContext.class);
		}
		public UnaryTermContext unaryTerm(int i) {
			return getRuleContext(UnaryTermContext.class,i);
		}
		public List<TerminalNode> MULT() { return getTokens(GoliteParser.MULT); }
		public TerminalNode MULT(int i) {
			return getToken(GoliteParser.MULT, i);
		}
		public List<TerminalNode> DIV() { return getTokens(GoliteParser.DIV); }
		public TerminalNode DIV(int i) {
			return getToken(GoliteParser.DIV, i);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(314);
			unaryTerm();
			setState(319);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MULT || _la==DIV) {
				{
				{
				setState(315);
				_la = _input.LA(1);
				if ( !(_la==MULT || _la==DIV) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(316);
				unaryTerm();
				}
				}
				setState(321);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class UnaryTermContext extends ParserRuleContext {
		public TerminalNode NOT() { return getToken(GoliteParser.NOT, 0); }
		public SelectorTermContext selectorTerm() {
			return getRuleContext(SelectorTermContext.class,0);
		}
		public TerminalNode MINUS() { return getToken(GoliteParser.MINUS, 0); }
		public UnaryTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryTerm; }
	}

	public final UnaryTermContext unaryTerm() throws RecognitionException {
		UnaryTermContext _localctx = new UnaryTermContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_unaryTerm);
		try {
			setState(327);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				enterOuterAlt(_localctx, 1);
				{
				setState(322);
				match(NOT);
				setState(323);
				selectorTerm();
				}
				break;
			case MINUS:
				enterOuterAlt(_localctx, 2);
				{
				setState(324);
				match(MINUS);
				setState(325);
				selectorTerm();
				}
				break;
			case LPAREN:
			case NEW:
			case INT_LIT:
			case BOOL_LIT:
			case NIL_LIT:
			case IDENT:
				enterOuterAlt(_localctx, 3);
				{
				setState(326);
				selectorTerm();
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	public static class SelectorTermContext extends ParserRuleContext {
		public FactorContext factor() {
			return getRuleContext(FactorContext.class,0);
		}
		public List<SelectorTermPrimeContext> selectorTermPrime() {
			return getRuleContexts(SelectorTermPrimeContext.class);
		}
		public SelectorTermPrimeContext selectorTermPrime(int i) {
			return getRuleContext(SelectorTermPrimeContext.class,i);
		}
		public SelectorTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectorTerm; }
	}

	public final SelectorTermContext selectorTerm() throws RecognitionException {
		SelectorTermContext _localctx = new SelectorTermContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_selectorTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(329);
			factor();
			setState(333);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(330);
				selectorTermPrime();
				}
				}
				setState(335);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class SelectorTermPrimeContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(GoliteParser.DOT, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public SelectorTermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectorTermPrime; }
	}

	public final SelectorTermPrimeContext selectorTermPrime() throws RecognitionException {
		SelectorTermPrimeContext _localctx = new SelectorTermPrimeContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_selectorTermPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(336);
			match(DOT);
			setState(337);
			match(IDENT);
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

	public static class SubfactorContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public SubfactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subfactor; }
	}

	public final SubfactorContext subfactor() throws RecognitionException {
		SubfactorContext _localctx = new SubfactorContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_subfactor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(339);
			match(LPAREN);
			setState(340);
			expression();
			setState(341);
			match(RPAREN);
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

	public static class FunctioncallContext extends ParserRuleContext {
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public FunctioncallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functioncall; }
	}

	public final FunctioncallContext functioncall() throws RecognitionException {
		FunctioncallContext _localctx = new FunctioncallContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_functioncall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(343);
			match(IDENT);
			setState(345);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LPAREN) {
				{
				setState(344);
				arguments();
				}
			}

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

	public static class AllocationContext extends ParserRuleContext {
		public Token key;
		public TerminalNode NEW() { return getToken(GoliteParser.NEW, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public AllocationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_allocation; }
	}

	public final AllocationContext allocation() throws RecognitionException {
		AllocationContext _localctx = new AllocationContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_allocation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(347);
			match(NEW);
			setState(348);
			((AllocationContext)_localctx).key = match(IDENT);
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

	public static class FactorContext extends ParserRuleContext {
		public SubfactorContext subfactor() {
			return getRuleContext(SubfactorContext.class,0);
		}
		public FunctioncallContext functioncall() {
			return getRuleContext(FunctioncallContext.class,0);
		}
		public TerminalNode INT_LIT() { return getToken(GoliteParser.INT_LIT, 0); }
		public AllocationContext allocation() {
			return getRuleContext(AllocationContext.class,0);
		}
		public TerminalNode BOOL_LIT() { return getToken(GoliteParser.BOOL_LIT, 0); }
		public TerminalNode NIL_LIT() { return getToken(GoliteParser.NIL_LIT, 0); }
		public FactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor; }
	}

	public final FactorContext factor() throws RecognitionException {
		FactorContext _localctx = new FactorContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_factor);
		try {
			setState(356);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				enterOuterAlt(_localctx, 1);
				{
				setState(350);
				subfactor();
				}
				break;
			case IDENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(351);
				functioncall();
				}
				break;
			case INT_LIT:
				enterOuterAlt(_localctx, 3);
				{
				setState(352);
				match(INT_LIT);
				}
				break;
			case NEW:
				enterOuterAlt(_localctx, 4);
				{
				setState(353);
				allocation();
				}
				break;
			case BOOL_LIT:
				enterOuterAlt(_localctx, 5);
				{
				setState(354);
				match(BOOL_LIT);
				}
				break;
			case NIL_LIT:
				enterOuterAlt(_localctx, 6);
				{
				setState(355);
				match(NIL_LIT);
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3,\u0169\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\3\2\3\2\3\2\3\2\3\2\3\3\7\3_\n\3\f\3\16\3b\13\3\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\5\3\5\3\5\7\5o\n\5\f\5\16\5r\13\5\3\6\3\6\3\6\3\7\3\7"+
		"\3\7\3\b\3\b\3\b\3\b\5\b~\n\b\3\t\7\t\u0081\n\t\f\t\16\t\u0084\13\t\3"+
		"\n\3\n\3\n\3\n\3\n\3\13\3\13\7\13\u008d\n\13\f\13\16\13\u0090\13\13\3"+
		"\f\3\f\3\f\3\r\7\r\u0096\n\r\f\r\16\r\u0099\13\r\3\16\3\16\3\16\3\16\5"+
		"\16\u009f\n\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\5\17\u00a8\n\17\3\17"+
		"\3\17\3\20\3\20\7\20\u00ae\n\20\f\20\16\20\u00b1\13\20\3\21\3\21\3\21"+
		"\3\22\3\22\3\23\7\23\u00b9\n\23\f\23\16\23\u00bc\13\23\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u00c7\n\24\3\25\3\25\3\25\3\25\3\26"+
		"\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31"+
		"\3\31\3\31\3\31\7\31\u00df\n\31\f\31\16\31\u00e2\13\31\3\31\3\31\3\31"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u00ee\n\32\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\34\3\34\5\34\u00f8\n\34\3\34\3\34\3\35\3\35\3\35\3\35\3\36"+
		"\3\36\3\36\3\36\7\36\u0104\n\36\f\36\16\36\u0107\13\36\5\36\u0109\n\36"+
		"\3\36\3\36\3\37\3\37\3\37\7\37\u0110\n\37\f\37\16\37\u0113\13\37\3 \3"+
		" \3 \7 \u0118\n \f \16 \u011b\13 \3!\3!\3!\7!\u0120\n!\f!\16!\u0123\13"+
		"!\3\"\3\"\3\"\7\"\u0128\n\"\f\"\16\"\u012b\13\"\3#\3#\3#\7#\u0130\n#\f"+
		"#\16#\u0133\13#\3$\3$\3$\7$\u0138\n$\f$\16$\u013b\13$\3%\3%\3%\7%\u0140"+
		"\n%\f%\16%\u0143\13%\3&\3&\3&\3&\3&\5&\u014a\n&\3\'\3\'\7\'\u014e\n\'"+
		"\f\'\16\'\u0151\13\'\3(\3(\3(\3)\3)\3)\3)\3*\3*\5*\u015c\n*\3+\3+\3+\3"+
		",\3,\3,\3,\3,\3,\5,\u0167\n,\3,\2\2-\2\4\6\b\n\f\16\20\22\24\26\30\32"+
		"\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTV\2\6\3\2\13\f\3\2\7\n\3\2\r"+
		"\16\3\2\17\20\2\u0165\2X\3\2\2\2\4`\3\2\2\2\6c\3\2\2\2\bk\3\2\2\2\ns\3"+
		"\2\2\2\fv\3\2\2\2\16}\3\2\2\2\20\u0082\3\2\2\2\22\u0085\3\2\2\2\24\u008a"+
		"\3\2\2\2\26\u0091\3\2\2\2\30\u0097\3\2\2\2\32\u009a\3\2\2\2\34\u00a5\3"+
		"\2\2\2\36\u00ab\3\2\2\2 \u00b2\3\2\2\2\"\u00b5\3\2\2\2$\u00ba\3\2\2\2"+
		"&\u00c6\3\2\2\2(\u00c8\3\2\2\2*\u00cc\3\2\2\2,\u00d0\3\2\2\2.\u00d4\3"+
		"\2\2\2\60\u00d9\3\2\2\2\62\u00e6\3\2\2\2\64\u00ef\3\2\2\2\66\u00f5\3\2"+
		"\2\28\u00fb\3\2\2\2:\u00ff\3\2\2\2<\u010c\3\2\2\2>\u0114\3\2\2\2@\u011c"+
		"\3\2\2\2B\u0124\3\2\2\2D\u012c\3\2\2\2F\u0134\3\2\2\2H\u013c\3\2\2\2J"+
		"\u0149\3\2\2\2L\u014b\3\2\2\2N\u0152\3\2\2\2P\u0155\3\2\2\2R\u0159\3\2"+
		"\2\2T\u015d\3\2\2\2V\u0166\3\2\2\2XY\5\4\3\2YZ\5\20\t\2Z[\5\30\r\2[\\"+
		"\7\2\2\3\\\3\3\2\2\2]_\5\6\4\2^]\3\2\2\2_b\3\2\2\2`^\3\2\2\2`a\3\2\2\2"+
		"a\5\3\2\2\2b`\3\2\2\2cd\7\30\2\2de\7*\2\2ef\7\33\2\2fg\7\23\2\2gh\5\b"+
		"\5\2hi\7\24\2\2ij\7\25\2\2j\7\3\2\2\2kl\5\f\7\2lp\7\25\2\2mo\5\n\6\2n"+
		"m\3\2\2\2or\3\2\2\2pn\3\2\2\2pq\3\2\2\2q\t\3\2\2\2rp\3\2\2\2st\5\f\7\2"+
		"tu\7\25\2\2u\13\3\2\2\2vw\7*\2\2wx\5\16\b\2x\r\3\2\2\2y~\7(\2\2z~\7)\2"+
		"\2{|\7\17\2\2|~\7*\2\2}y\3\2\2\2}z\3\2\2\2}{\3\2\2\2~\17\3\2\2\2\177\u0081"+
		"\5\22\n\2\u0080\177\3\2\2\2\u0081\u0084\3\2\2\2\u0082\u0080\3\2\2\2\u0082"+
		"\u0083\3\2\2\2\u0083\21\3\2\2\2\u0084\u0082\3\2\2\2\u0085\u0086\7\27\2"+
		"\2\u0086\u0087\5\24\13\2\u0087\u0088\5\16\b\2\u0088\u0089\7\25\2\2\u0089"+
		"\23\3\2\2\2\u008a\u008e\7*\2\2\u008b\u008d\5\26\f\2\u008c\u008b\3\2\2"+
		"\2\u008d\u0090\3\2\2\2\u008e\u008c\3\2\2\2\u008e\u008f\3\2\2\2\u008f\25"+
		"\3\2\2\2\u0090\u008e\3\2\2\2\u0091\u0092\7\26\2\2\u0092\u0093\7*\2\2\u0093"+
		"\27\3\2\2\2\u0094\u0096\5\32\16\2\u0095\u0094\3\2\2\2\u0096\u0099\3\2"+
		"\2\2\u0097\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098\31\3\2\2\2\u0099\u0097"+
		"\3\2\2\2\u009a\u009b\7\35\2\2\u009b\u009c\7*\2\2\u009c\u009e\5\34\17\2"+
		"\u009d\u009f\5\"\22\2\u009e\u009d\3\2\2\2\u009e\u009f\3\2\2\2\u009f\u00a0"+
		"\3\2\2\2\u00a0\u00a1\7\23\2\2\u00a1\u00a2\5\20\t\2\u00a2\u00a3\5$\23\2"+
		"\u00a3\u00a4\7\24\2\2\u00a4\33\3\2\2\2\u00a5\u00a7\7\21\2\2\u00a6\u00a8"+
		"\5\36\20\2\u00a7\u00a6\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8\u00a9\3\2\2\2"+
		"\u00a9\u00aa\7\22\2\2\u00aa\35\3\2\2\2\u00ab\u00af\5\f\7\2\u00ac\u00ae"+
		"\5 \21\2\u00ad\u00ac\3\2\2\2\u00ae\u00b1\3\2\2\2\u00af\u00ad\3\2\2\2\u00af"+
		"\u00b0\3\2\2\2\u00b0\37\3\2\2\2\u00b1\u00af\3\2\2\2\u00b2\u00b3\7\26\2"+
		"\2\u00b3\u00b4\5\f\7\2\u00b4!\3\2\2\2\u00b5\u00b6\5\16\b\2\u00b6#\3\2"+
		"\2\2\u00b7\u00b9\5&\24\2\u00b8\u00b7\3\2\2\2\u00b9\u00bc\3\2\2\2\u00ba"+
		"\u00b8\3\2\2\2\u00ba\u00bb\3\2\2\2\u00bb%\3\2\2\2\u00bc\u00ba\3\2\2\2"+
		"\u00bd\u00c7\5*\26\2\u00be\u00c7\5.\30\2\u00bf\u00c7\5\60\31\2\u00c0\u00c7"+
		"\5,\27\2\u00c1\u00c7\5(\25\2\u00c2\u00c7\5\62\32\2\u00c3\u00c7\5\64\33"+
		"\2\u00c4\u00c7\5\66\34\2\u00c5\u00c7\58\35\2\u00c6\u00bd\3\2\2\2\u00c6"+
		"\u00be\3\2\2\2\u00c6\u00bf\3\2\2\2\u00c6\u00c0\3\2\2\2\u00c6\u00c1\3\2"+
		"\2\2\u00c6\u00c2\3\2\2\2\u00c6\u00c3\3\2\2\2\u00c6\u00c4\3\2\2\2\u00c6"+
		"\u00c5\3\2\2\2\u00c7\'\3\2\2\2\u00c8\u00c9\7\"\2\2\u00c9\u00ca\5<\37\2"+
		"\u00ca\u00cb\7\25\2\2\u00cb)\3\2\2\2\u00cc\u00cd\7\23\2\2\u00cd\u00ce"+
		"\5$\23\2\u00ce\u00cf\7\24\2\2\u00cf+\3\2\2\2\u00d0\u00d1\7\32\2\2\u00d1"+
		"\u00d2\5> \2\u00d2\u00d3\7\25\2\2\u00d3-\3\2\2\2\u00d4\u00d5\5<\37\2\u00d5"+
		"\u00d6\7\3\2\2\u00d6\u00d7\5> \2\u00d7\u00d8\7\25\2\2\u00d8/\3\2\2\2\u00d9"+
		"\u00da\7#\2\2\u00da\u00db\7\21\2\2\u00db\u00e0\7%\2\2\u00dc\u00dd\7\26"+
		"\2\2\u00dd\u00df\5> \2\u00de\u00dc\3\2\2\2\u00df\u00e2\3\2\2\2\u00e0\u00de"+
		"\3\2\2\2\u00e0\u00e1\3\2\2\2\u00e1\u00e3\3\2\2\2\u00e2\u00e0\3\2\2\2\u00e3"+
		"\u00e4\7\22\2\2\u00e4\u00e5\7\25\2\2\u00e5\61\3\2\2\2\u00e6\u00e7\7\37"+
		"\2\2\u00e7\u00e8\7\21\2\2\u00e8\u00e9\5> \2\u00e9\u00ea\7\22\2\2\u00ea"+
		"\u00ed\5*\26\2\u00eb\u00ec\7 \2\2\u00ec\u00ee\5*\26\2\u00ed\u00eb\3\2"+
		"\2\2\u00ed\u00ee\3\2\2\2\u00ee\63\3\2\2\2\u00ef\u00f0\7!\2\2\u00f0\u00f1"+
		"\7\21\2\2\u00f1\u00f2\5> \2\u00f2\u00f3\7\22\2\2\u00f3\u00f4\5*\26\2\u00f4"+
		"\65\3\2\2\2\u00f5\u00f7\7\36\2\2\u00f6\u00f8\5> \2\u00f7\u00f6\3\2\2\2"+
		"\u00f7\u00f8\3\2\2\2\u00f8\u00f9\3\2\2\2\u00f9\u00fa\7\25\2\2\u00fa\67"+
		"\3\2\2\2\u00fb\u00fc\7*\2\2\u00fc\u00fd\5:\36\2\u00fd\u00fe\7\25\2\2\u00fe"+
		"9\3\2\2\2\u00ff\u0108\7\21\2\2\u0100\u0105\5> \2\u0101\u0102\7\26\2\2"+
		"\u0102\u0104\5> \2\u0103\u0101\3\2\2\2\u0104\u0107\3\2\2\2\u0105\u0103"+
		"\3\2\2\2\u0105\u0106\3\2\2\2\u0106\u0109\3\2\2\2\u0107\u0105\3\2\2\2\u0108"+
		"\u0100\3\2\2\2\u0108\u0109\3\2\2\2\u0109\u010a\3\2\2\2\u010a\u010b\7\22"+
		"\2\2\u010b;\3\2\2\2\u010c\u0111\7*\2\2\u010d\u010e\7\34\2\2\u010e\u0110"+
		"\7*\2\2\u010f\u010d\3\2\2\2\u0110\u0113\3\2\2\2\u0111\u010f\3\2\2\2\u0111"+
		"\u0112\3\2\2\2\u0112=\3\2\2\2\u0113\u0111\3\2\2\2\u0114\u0119\5@!\2\u0115"+
		"\u0116\7\5\2\2\u0116\u0118\5@!\2\u0117\u0115\3\2\2\2\u0118\u011b\3\2\2"+
		"\2\u0119\u0117\3\2\2\2\u0119\u011a\3\2\2\2\u011a?\3\2\2\2\u011b\u0119"+
		"\3\2\2\2\u011c\u0121\5B\"\2\u011d\u011e\7\4\2\2\u011e\u0120\5B\"\2\u011f"+
		"\u011d\3\2\2\2\u0120\u0123\3\2\2\2\u0121\u011f\3\2\2\2\u0121\u0122\3\2"+
		"\2\2\u0122A\3\2\2\2\u0123\u0121\3\2\2\2\u0124\u0129\5D#\2\u0125\u0126"+
		"\t\2\2\2\u0126\u0128\5D#\2\u0127\u0125\3\2\2\2\u0128\u012b\3\2\2\2\u0129"+
		"\u0127\3\2\2\2\u0129\u012a\3\2\2\2\u012aC\3\2\2\2\u012b\u0129\3\2\2\2"+
		"\u012c\u0131\5F$\2\u012d\u012e\t\3\2\2\u012e\u0130\5F$\2\u012f\u012d\3"+
		"\2\2\2\u0130\u0133\3\2\2\2\u0131\u012f\3\2\2\2\u0131\u0132\3\2\2\2\u0132"+
		"E\3\2\2\2\u0133\u0131\3\2\2\2\u0134\u0139\5H%\2\u0135\u0136\t\4\2\2\u0136"+
		"\u0138\5H%\2\u0137\u0135\3\2\2\2\u0138\u013b\3\2\2\2\u0139\u0137\3\2\2"+
		"\2\u0139\u013a\3\2\2\2\u013aG\3\2\2\2\u013b\u0139\3\2\2\2\u013c\u0141"+
		"\5J&\2\u013d\u013e\t\5\2\2\u013e\u0140\5J&\2\u013f\u013d\3\2\2\2\u0140"+
		"\u0143\3\2\2\2\u0141\u013f\3\2\2\2\u0141\u0142\3\2\2\2\u0142I\3\2\2\2"+
		"\u0143\u0141\3\2\2\2\u0144\u0145\7\6\2\2\u0145\u014a\5L\'\2\u0146\u0147"+
		"\7\16\2\2\u0147\u014a\5L\'\2\u0148\u014a\5L\'\2\u0149\u0144\3\2\2\2\u0149"+
		"\u0146\3\2\2\2\u0149\u0148\3\2\2\2\u014aK\3\2\2\2\u014b\u014f\5V,\2\u014c"+
		"\u014e\5N(\2\u014d\u014c\3\2\2\2\u014e\u0151\3\2\2\2\u014f\u014d\3\2\2"+
		"\2\u014f\u0150\3\2\2\2\u0150M\3\2\2\2\u0151\u014f\3\2\2\2\u0152\u0153"+
		"\7\34\2\2\u0153\u0154\7*\2\2\u0154O\3\2\2\2\u0155\u0156\7\21\2\2\u0156"+
		"\u0157\5> \2\u0157\u0158\7\22\2\2\u0158Q\3\2\2\2\u0159\u015b\7*\2\2\u015a"+
		"\u015c\5:\36\2\u015b\u015a\3\2\2\2\u015b\u015c\3\2\2\2\u015cS\3\2\2\2"+
		"\u015d\u015e\7\31\2\2\u015e\u015f\7*\2\2\u015fU\3\2\2\2\u0160\u0167\5"+
		"P)\2\u0161\u0167\5R*\2\u0162\u0167\7$\2\2\u0163\u0167\5T+\2\u0164\u0167"+
		"\7&\2\2\u0165\u0167\7\'\2\2\u0166\u0160\3\2\2\2\u0166\u0161\3\2\2\2\u0166"+
		"\u0162\3\2\2\2\u0166\u0163\3\2\2\2\u0166\u0164\3\2\2\2\u0166\u0165\3\2"+
		"\2\2\u0167W\3\2\2\2\35`p}\u0082\u008e\u0097\u009e\u00a7\u00af\u00ba\u00c6"+
		"\u00e0\u00ed\u00f7\u0105\u0108\u0111\u0119\u0121\u0129\u0131\u0139\u0141"+
		"\u0149\u014f\u015b\u0166";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}