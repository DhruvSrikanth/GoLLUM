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
		RULE_printPrime = 24, RULE_conditional = 25, RULE_conditionalPrime = 26, 
		RULE_loop = 27, RULE_returnRule = 28, RULE_invocation = 29, RULE_arguments = 30, 
		RULE_argumentsPrime = 31, RULE_argumentsPrimePrime = 32, RULE_lValue = 33, 
		RULE_lValuePrime = 34, RULE_expression = 35, RULE_expressionPrime = 36, 
		RULE_boolTerm = 37, RULE_boolTermPrime = 38, RULE_equalTerm = 39, RULE_equalTermPrime = 40, 
		RULE_relationTerm = 41, RULE_relationTermPrime = 42, RULE_simpleTerm = 43, 
		RULE_simpleTermPrime = 44, RULE_term = 45, RULE_termPrime = 46, RULE_unaryTerm = 47, 
		RULE_unaryTermBool = 48, RULE_unaryTermInt = 49, RULE_selectorTerm = 50, 
		RULE_selectorTermPrime = 51, RULE_factor = 52, RULE_subfactor = 53, RULE_functioncall = 54, 
		RULE_allocation = 55;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "types", "typeDeclaration", "fields", "morefields", "decl", 
			"type", "declarations", "declaration", "ids", "otherids", "functions", 
			"function", "parameters", "parametersPrime", "parametersPPrime", "returnType", 
			"statements", "statement", "read", "block", "delete", "assignment", "print", 
			"printPrime", "conditional", "conditionalPrime", "loop", "returnRule", 
			"invocation", "arguments", "argumentsPrime", "argumentsPrimePrime", "lValue", 
			"lValuePrime", "expression", "expressionPrime", "boolTerm", "boolTermPrime", 
			"equalTerm", "equalTermPrime", "relationTerm", "relationTermPrime", "simpleTerm", 
			"simpleTermPrime", "term", "termPrime", "unaryTerm", "unaryTermBool", 
			"unaryTermInt", "selectorTerm", "selectorTermPrime", "factor", "subfactor", 
			"functioncall", "allocation"
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
			setState(112);
			((ProgramContext)_localctx).ty = types();
			setState(113);
			((ProgramContext)_localctx).dcls = declarations();
			setState(114);
			((ProgramContext)_localctx).fns = functions();
			setState(115);
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
			setState(120);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TYPE) {
				{
				{
				setState(117);
				typeDeclaration();
				}
				}
				setState(122);
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
			setState(123);
			match(TYPE);
			setState(124);
			((TypeDeclarationContext)_localctx).id = match(IDENT);
			setState(125);
			match(STRUCT);
			setState(126);
			match(LBRACE);
			setState(127);
			((TypeDeclarationContext)_localctx).flds = fields();
			setState(128);
			match(RBRACE);
			setState(129);
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
			setState(131);
			((FieldsContext)_localctx).dcl = decl();
			setState(132);
			match(SEMICOLON);
			setState(136);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENT) {
				{
				{
				setState(133);
				morefields();
				}
				}
				setState(138);
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
			setState(139);
			((MorefieldsContext)_localctx).dcl = decl();
			setState(140);
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
			setState(142);
			((DeclContext)_localctx).id = match(IDENT);
			setState(143);
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
			setState(149);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(145);
				match(INT);
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 2);
				{
				setState(146);
				match(BOOL);
				}
				break;
			case MULT:
				enterOuterAlt(_localctx, 3);
				{
				{
				setState(147);
				match(MULT);
				setState(148);
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
			setState(154);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(151);
				declaration();
				}
				}
				setState(156);
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
			setState(157);
			match(VAR);
			setState(158);
			((DeclarationContext)_localctx).idx = ids();
			setState(159);
			((DeclarationContext)_localctx).ty = type();
			setState(160);
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
			setState(162);
			((IdsContext)_localctx).id = match(IDENT);
			setState(166);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(163);
				otherids();
				}
				}
				setState(168);
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
			setState(169);
			match(COMMA);
			setState(170);
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
			setState(175);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNC) {
				{
				{
				setState(172);
				function();
				}
				}
				setState(177);
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
			setState(178);
			match(FUNC);
			setState(179);
			((FunctionContext)_localctx).id = match(IDENT);
			setState(180);
			((FunctionContext)_localctx).params = parameters();
			setState(182);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MULT) | (1L << INT) | (1L << BOOL))) != 0)) {
				{
				setState(181);
				((FunctionContext)_localctx).rty = returnType();
				}
			}

			setState(184);
			match(LBRACE);
			setState(185);
			((FunctionContext)_localctx).dcls = declarations();
			setState(186);
			((FunctionContext)_localctx).stmts = statements();
			setState(187);
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
			setState(189);
			match(LPAREN);
			setState(191);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENT) {
				{
				setState(190);
				parametersPrime();
				}
			}

			setState(193);
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
			setState(195);
			((ParametersPrimeContext)_localctx).dcl = decl();
			setState(199);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(196);
				parametersPPrime();
				}
				}
				setState(201);
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
			setState(202);
			match(COMMA);
			setState(203);
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
			setState(205);
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
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LBRACE) | (1L << DELETE) | (1L << RET) | (1L << IF) | (1L << FOR) | (1L << SCAN) | (1L << PRINTF) | (1L << IDENT))) != 0)) {
				{
				{
				setState(207);
				statement();
				}
				}
				setState(212);
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
			setState(222);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(213);
				((StatementContext)_localctx).bl = block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(214);
				((StatementContext)_localctx).asmt = assignment();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(215);
				((StatementContext)_localctx).prnt = print();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(216);
				((StatementContext)_localctx).del = delete();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(217);
				((StatementContext)_localctx).rd = read();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(218);
				((StatementContext)_localctx).cond = conditional();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(219);
				((StatementContext)_localctx).lp = loop();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(220);
				((StatementContext)_localctx).ret = returnRule();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(221);
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
		public LValueContext lval;
		public TerminalNode SCAN() { return getToken(GoliteParser.SCAN, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public LValueContext lValue() {
			return getRuleContext(LValueContext.class,0);
		}
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
			setState(224);
			match(SCAN);
			setState(225);
			((ReadContext)_localctx).lval = lValue();
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

	public static class BlockContext extends ParserRuleContext {
		public StatementsContext stmts;
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
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
			setState(228);
			match(LBRACE);
			setState(229);
			((BlockContext)_localctx).stmts = statements();
			setState(230);
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
		public ExpressionContext expr;
		public TerminalNode DELETE() { return getToken(GoliteParser.DELETE, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
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
			setState(232);
			match(DELETE);
			setState(233);
			((DeleteContext)_localctx).expr = expression();
			setState(234);
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
		public LValueContext lval;
		public ExpressionContext expr;
		public TerminalNode ASSIGN() { return getToken(GoliteParser.ASSIGN, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public LValueContext lValue() {
			return getRuleContext(LValueContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
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
			setState(236);
			((AssignmentContext)_localctx).lval = lValue();
			setState(237);
			match(ASSIGN);
			setState(238);
			((AssignmentContext)_localctx).expr = expression();
			setState(239);
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
		public Token str;
		public TerminalNode PRINTF() { return getToken(GoliteParser.PRINTF, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoliteParser.STRING_LIT, 0); }
		public List<PrintPrimeContext> printPrime() {
			return getRuleContexts(PrintPrimeContext.class);
		}
		public PrintPrimeContext printPrime(int i) {
			return getRuleContext(PrintPrimeContext.class,i);
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
			setState(241);
			match(PRINTF);
			setState(242);
			match(LPAREN);
			setState(243);
			((PrintContext)_localctx).str = match(STRING_LIT);
			setState(247);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(244);
				printPrime();
				}
				}
				setState(249);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(250);
			match(RPAREN);
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

	public static class PrintPrimeContext extends ParserRuleContext {
		public ExpressionContext expr;
		public TerminalNode COMMA() { return getToken(GoliteParser.COMMA, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PrintPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printPrime; }
	}

	public final PrintPrimeContext printPrime() throws RecognitionException {
		PrintPrimeContext _localctx = new PrintPrimeContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_printPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(253);
			match(COMMA);
			setState(254);
			((PrintPrimeContext)_localctx).expr = expression();
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
		public ExpressionContext expr;
		public BlockContext bl;
		public ConditionalPrimeContext then;
		public TerminalNode IF() { return getToken(GoliteParser.IF, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ConditionalPrimeContext conditionalPrime() {
			return getRuleContext(ConditionalPrimeContext.class,0);
		}
		public ConditionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional; }
	}

	public final ConditionalContext conditional() throws RecognitionException {
		ConditionalContext _localctx = new ConditionalContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(256);
			match(IF);
			setState(257);
			match(LPAREN);
			setState(258);
			((ConditionalContext)_localctx).expr = expression();
			setState(259);
			match(RPAREN);
			setState(260);
			((ConditionalContext)_localctx).bl = block();
			setState(262);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(261);
				((ConditionalContext)_localctx).then = conditionalPrime();
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

	public static class ConditionalPrimeContext extends ParserRuleContext {
		public BlockContext bl;
		public TerminalNode ELSE() { return getToken(GoliteParser.ELSE, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ConditionalPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionalPrime; }
	}

	public final ConditionalPrimeContext conditionalPrime() throws RecognitionException {
		ConditionalPrimeContext _localctx = new ConditionalPrimeContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_conditionalPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(264);
			match(ELSE);
			setState(265);
			((ConditionalPrimeContext)_localctx).bl = block();
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
		public ExpressionContext expr;
		public BlockContext bl;
		public TerminalNode FOR() { return getToken(GoliteParser.FOR, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
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
		enterRule(_localctx, 54, RULE_loop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(267);
			match(FOR);
			setState(268);
			match(LPAREN);
			setState(269);
			((LoopContext)_localctx).expr = expression();
			setState(270);
			match(RPAREN);
			setState(271);
			((LoopContext)_localctx).bl = block();
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
		public ExpressionContext expr;
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
		enterRule(_localctx, 56, RULE_returnRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(273);
			match(RET);
			setState(275);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NOT) | (1L << MINUS) | (1L << LPAREN) | (1L << NEW) | (1L << INT_LIT) | (1L << BOOL_LIT) | (1L << NIL_LIT) | (1L << IDENT))) != 0)) {
				{
				setState(274);
				((ReturnRuleContext)_localctx).expr = expression();
				}
			}

			setState(277);
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
		public Token id;
		public ArgumentsContext args;
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public InvocationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_invocation; }
	}

	public final InvocationContext invocation() throws RecognitionException {
		InvocationContext _localctx = new InvocationContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_invocation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(279);
			((InvocationContext)_localctx).id = match(IDENT);
			setState(280);
			((InvocationContext)_localctx).args = arguments();
			setState(281);
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
		public ArgumentsPrimeContext args;
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public ArgumentsPrimeContext argumentsPrime() {
			return getRuleContext(ArgumentsPrimeContext.class,0);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(283);
			match(LPAREN);
			setState(285);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NOT) | (1L << MINUS) | (1L << LPAREN) | (1L << NEW) | (1L << INT_LIT) | (1L << BOOL_LIT) | (1L << NIL_LIT) | (1L << IDENT))) != 0)) {
				{
				setState(284);
				((ArgumentsContext)_localctx).args = argumentsPrime();
				}
			}

			setState(287);
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

	public static class ArgumentsPrimeContext extends ParserRuleContext {
		public ExpressionContext expr;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<ArgumentsPrimePrimeContext> argumentsPrimePrime() {
			return getRuleContexts(ArgumentsPrimePrimeContext.class);
		}
		public ArgumentsPrimePrimeContext argumentsPrimePrime(int i) {
			return getRuleContext(ArgumentsPrimePrimeContext.class,i);
		}
		public ArgumentsPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argumentsPrime; }
	}

	public final ArgumentsPrimeContext argumentsPrime() throws RecognitionException {
		ArgumentsPrimeContext _localctx = new ArgumentsPrimeContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_argumentsPrime);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(289);
			((ArgumentsPrimeContext)_localctx).expr = expression();
			setState(293);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(290);
				argumentsPrimePrime();
				}
				}
				setState(295);
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

	public static class ArgumentsPrimePrimeContext extends ParserRuleContext {
		public ExpressionContext expr;
		public TerminalNode COMMA() { return getToken(GoliteParser.COMMA, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ArgumentsPrimePrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argumentsPrimePrime; }
	}

	public final ArgumentsPrimePrimeContext argumentsPrimePrime() throws RecognitionException {
		ArgumentsPrimePrimeContext _localctx = new ArgumentsPrimePrimeContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_argumentsPrimePrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(296);
			match(COMMA);
			setState(297);
			((ArgumentsPrimePrimeContext)_localctx).expr = expression();
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
		public Token id;
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public List<LValuePrimeContext> lValuePrime() {
			return getRuleContexts(LValuePrimeContext.class);
		}
		public LValuePrimeContext lValuePrime(int i) {
			return getRuleContext(LValuePrimeContext.class,i);
		}
		public LValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lValue; }
	}

	public final LValueContext lValue() throws RecognitionException {
		LValueContext _localctx = new LValueContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_lValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(299);
			((LValueContext)_localctx).id = match(IDENT);
			setState(303);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(300);
				lValuePrime();
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

	public static class LValuePrimeContext extends ParserRuleContext {
		public Token id;
		public TerminalNode DOT() { return getToken(GoliteParser.DOT, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public LValuePrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lValuePrime; }
	}

	public final LValuePrimeContext lValuePrime() throws RecognitionException {
		LValuePrimeContext _localctx = new LValuePrimeContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_lValuePrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(306);
			match(DOT);
			setState(307);
			((LValuePrimeContext)_localctx).id = match(IDENT);
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
		public BoolTermContext bt;
		public BoolTermContext boolTerm() {
			return getRuleContext(BoolTermContext.class,0);
		}
		public List<ExpressionPrimeContext> expressionPrime() {
			return getRuleContexts(ExpressionPrimeContext.class);
		}
		public ExpressionPrimeContext expressionPrime(int i) {
			return getRuleContext(ExpressionPrimeContext.class,i);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(309);
			((ExpressionContext)_localctx).bt = boolTerm();
			setState(313);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(310);
				expressionPrime();
				}
				}
				setState(315);
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

	public static class ExpressionPrimeContext extends ParserRuleContext {
		public Token op;
		public BoolTermContext bt;
		public TerminalNode OR() { return getToken(GoliteParser.OR, 0); }
		public BoolTermContext boolTerm() {
			return getRuleContext(BoolTermContext.class,0);
		}
		public ExpressionPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionPrime; }
	}

	public final ExpressionPrimeContext expressionPrime() throws RecognitionException {
		ExpressionPrimeContext _localctx = new ExpressionPrimeContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_expressionPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(316);
			((ExpressionPrimeContext)_localctx).op = match(OR);
			setState(317);
			((ExpressionPrimeContext)_localctx).bt = boolTerm();
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
		public EqualTermContext eq;
		public EqualTermContext equalTerm() {
			return getRuleContext(EqualTermContext.class,0);
		}
		public List<BoolTermPrimeContext> boolTermPrime() {
			return getRuleContexts(BoolTermPrimeContext.class);
		}
		public BoolTermPrimeContext boolTermPrime(int i) {
			return getRuleContext(BoolTermPrimeContext.class,i);
		}
		public BoolTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolTerm; }
	}

	public final BoolTermContext boolTerm() throws RecognitionException {
		BoolTermContext _localctx = new BoolTermContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_boolTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(319);
			((BoolTermContext)_localctx).eq = equalTerm();
			setState(323);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(320);
				boolTermPrime();
				}
				}
				setState(325);
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

	public static class BoolTermPrimeContext extends ParserRuleContext {
		public Token op;
		public EqualTermContext eq;
		public TerminalNode AND() { return getToken(GoliteParser.AND, 0); }
		public EqualTermContext equalTerm() {
			return getRuleContext(EqualTermContext.class,0);
		}
		public BoolTermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolTermPrime; }
	}

	public final BoolTermPrimeContext boolTermPrime() throws RecognitionException {
		BoolTermPrimeContext _localctx = new BoolTermPrimeContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_boolTermPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(326);
			((BoolTermPrimeContext)_localctx).op = match(AND);
			setState(327);
			((BoolTermPrimeContext)_localctx).eq = equalTerm();
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
		public RelationTermContext rt;
		public RelationTermContext relationTerm() {
			return getRuleContext(RelationTermContext.class,0);
		}
		public List<EqualTermPrimeContext> equalTermPrime() {
			return getRuleContexts(EqualTermPrimeContext.class);
		}
		public EqualTermPrimeContext equalTermPrime(int i) {
			return getRuleContext(EqualTermPrimeContext.class,i);
		}
		public EqualTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_equalTerm; }
	}

	public final EqualTermContext equalTerm() throws RecognitionException {
		EqualTermContext _localctx = new EqualTermContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_equalTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(329);
			((EqualTermContext)_localctx).rt = relationTerm();
			setState(333);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EQ || _la==NE) {
				{
				{
				setState(330);
				equalTermPrime();
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

	public static class EqualTermPrimeContext extends ParserRuleContext {
		public Token op;
		public RelationTermContext rt;
		public RelationTermContext relationTerm() {
			return getRuleContext(RelationTermContext.class,0);
		}
		public TerminalNode EQ() { return getToken(GoliteParser.EQ, 0); }
		public TerminalNode NE() { return getToken(GoliteParser.NE, 0); }
		public EqualTermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_equalTermPrime; }
	}

	public final EqualTermPrimeContext equalTermPrime() throws RecognitionException {
		EqualTermPrimeContext _localctx = new EqualTermPrimeContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_equalTermPrime);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(336);
			((EqualTermPrimeContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==EQ || _la==NE) ) {
				((EqualTermPrimeContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(337);
			((EqualTermPrimeContext)_localctx).rt = relationTerm();
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
		public SimpleTermContext st;
		public SimpleTermContext simpleTerm() {
			return getRuleContext(SimpleTermContext.class,0);
		}
		public List<RelationTermPrimeContext> relationTermPrime() {
			return getRuleContexts(RelationTermPrimeContext.class);
		}
		public RelationTermPrimeContext relationTermPrime(int i) {
			return getRuleContext(RelationTermPrimeContext.class,i);
		}
		public RelationTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationTerm; }
	}

	public final RelationTermContext relationTerm() throws RecognitionException {
		RelationTermContext _localctx = new RelationTermContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_relationTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(339);
			((RelationTermContext)_localctx).st = simpleTerm();
			setState(343);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << LE) | (1L << GE))) != 0)) {
				{
				{
				setState(340);
				relationTermPrime();
				}
				}
				setState(345);
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

	public static class RelationTermPrimeContext extends ParserRuleContext {
		public Token op;
		public SimpleTermContext st;
		public SimpleTermContext simpleTerm() {
			return getRuleContext(SimpleTermContext.class,0);
		}
		public TerminalNode GE() { return getToken(GoliteParser.GE, 0); }
		public TerminalNode LT() { return getToken(GoliteParser.LT, 0); }
		public TerminalNode LE() { return getToken(GoliteParser.LE, 0); }
		public TerminalNode GT() { return getToken(GoliteParser.GT, 0); }
		public RelationTermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationTermPrime; }
	}

	public final RelationTermPrimeContext relationTermPrime() throws RecognitionException {
		RelationTermPrimeContext _localctx = new RelationTermPrimeContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_relationTermPrime);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(346);
			((RelationTermPrimeContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << LE) | (1L << GE))) != 0)) ) {
				((RelationTermPrimeContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(347);
			((RelationTermPrimeContext)_localctx).st = simpleTerm();
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
		public TermContext t;
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public List<SimpleTermPrimeContext> simpleTermPrime() {
			return getRuleContexts(SimpleTermPrimeContext.class);
		}
		public SimpleTermPrimeContext simpleTermPrime(int i) {
			return getRuleContext(SimpleTermPrimeContext.class,i);
		}
		public SimpleTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleTerm; }
	}

	public final SimpleTermContext simpleTerm() throws RecognitionException {
		SimpleTermContext _localctx = new SimpleTermContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_simpleTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(349);
			((SimpleTermContext)_localctx).t = term();
			setState(353);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PLUS || _la==MINUS) {
				{
				{
				setState(350);
				simpleTermPrime();
				}
				}
				setState(355);
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

	public static class SimpleTermPrimeContext extends ParserRuleContext {
		public Token op;
		public TermContext t;
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TerminalNode PLUS() { return getToken(GoliteParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(GoliteParser.MINUS, 0); }
		public SimpleTermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleTermPrime; }
	}

	public final SimpleTermPrimeContext simpleTermPrime() throws RecognitionException {
		SimpleTermPrimeContext _localctx = new SimpleTermPrimeContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_simpleTermPrime);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(356);
			((SimpleTermPrimeContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==PLUS || _la==MINUS) ) {
				((SimpleTermPrimeContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(357);
			((SimpleTermPrimeContext)_localctx).t = term();
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
		public UnaryTermContext ut;
		public UnaryTermContext unaryTerm() {
			return getRuleContext(UnaryTermContext.class,0);
		}
		public List<TermPrimeContext> termPrime() {
			return getRuleContexts(TermPrimeContext.class);
		}
		public TermPrimeContext termPrime(int i) {
			return getRuleContext(TermPrimeContext.class,i);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(359);
			((TermContext)_localctx).ut = unaryTerm();
			setState(363);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MULT || _la==DIV) {
				{
				{
				setState(360);
				termPrime();
				}
				}
				setState(365);
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

	public static class TermPrimeContext extends ParserRuleContext {
		public Token op;
		public UnaryTermContext ut;
		public UnaryTermContext unaryTerm() {
			return getRuleContext(UnaryTermContext.class,0);
		}
		public TerminalNode MULT() { return getToken(GoliteParser.MULT, 0); }
		public TerminalNode DIV() { return getToken(GoliteParser.DIV, 0); }
		public TermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_termPrime; }
	}

	public final TermPrimeContext termPrime() throws RecognitionException {
		TermPrimeContext _localctx = new TermPrimeContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_termPrime);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(366);
			((TermPrimeContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==MULT || _la==DIV) ) {
				((TermPrimeContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(367);
			((TermPrimeContext)_localctx).ut = unaryTerm();
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
		public UnaryTermBoolContext unaryTermBool() {
			return getRuleContext(UnaryTermBoolContext.class,0);
		}
		public UnaryTermIntContext unaryTermInt() {
			return getRuleContext(UnaryTermIntContext.class,0);
		}
		public SelectorTermContext selectorTerm() {
			return getRuleContext(SelectorTermContext.class,0);
		}
		public UnaryTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryTerm; }
	}

	public final UnaryTermContext unaryTerm() throws RecognitionException {
		UnaryTermContext _localctx = new UnaryTermContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_unaryTerm);
		try {
			setState(372);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				enterOuterAlt(_localctx, 1);
				{
				setState(369);
				unaryTermBool();
				}
				break;
			case MINUS:
				enterOuterAlt(_localctx, 2);
				{
				setState(370);
				unaryTermInt();
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
				setState(371);
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

	public static class UnaryTermBoolContext extends ParserRuleContext {
		public Token op;
		public SelectorTermContext st;
		public TerminalNode NOT() { return getToken(GoliteParser.NOT, 0); }
		public SelectorTermContext selectorTerm() {
			return getRuleContext(SelectorTermContext.class,0);
		}
		public UnaryTermBoolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryTermBool; }
	}

	public final UnaryTermBoolContext unaryTermBool() throws RecognitionException {
		UnaryTermBoolContext _localctx = new UnaryTermBoolContext(_ctx, getState());
		enterRule(_localctx, 96, RULE_unaryTermBool);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(374);
			((UnaryTermBoolContext)_localctx).op = match(NOT);
			setState(375);
			((UnaryTermBoolContext)_localctx).st = selectorTerm();
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

	public static class UnaryTermIntContext extends ParserRuleContext {
		public Token op;
		public SelectorTermContext st;
		public TerminalNode MINUS() { return getToken(GoliteParser.MINUS, 0); }
		public SelectorTermContext selectorTerm() {
			return getRuleContext(SelectorTermContext.class,0);
		}
		public UnaryTermIntContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryTermInt; }
	}

	public final UnaryTermIntContext unaryTermInt() throws RecognitionException {
		UnaryTermIntContext _localctx = new UnaryTermIntContext(_ctx, getState());
		enterRule(_localctx, 98, RULE_unaryTermInt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(377);
			((UnaryTermIntContext)_localctx).op = match(MINUS);
			setState(378);
			((UnaryTermIntContext)_localctx).st = selectorTerm();
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
		public FactorContext f;
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
		enterRule(_localctx, 100, RULE_selectorTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(380);
			((SelectorTermContext)_localctx).f = factor();
			setState(384);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(381);
				selectorTermPrime();
				}
				}
				setState(386);
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
		public Token id;
		public TerminalNode DOT() { return getToken(GoliteParser.DOT, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public SelectorTermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectorTermPrime; }
	}

	public final SelectorTermPrimeContext selectorTermPrime() throws RecognitionException {
		SelectorTermPrimeContext _localctx = new SelectorTermPrimeContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_selectorTermPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(387);
			match(DOT);
			setState(388);
			((SelectorTermPrimeContext)_localctx).id = match(IDENT);
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
		enterRule(_localctx, 104, RULE_factor);
		try {
			setState(396);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				enterOuterAlt(_localctx, 1);
				{
				setState(390);
				subfactor();
				}
				break;
			case IDENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(391);
				functioncall();
				}
				break;
			case INT_LIT:
				enterOuterAlt(_localctx, 3);
				{
				setState(392);
				match(INT_LIT);
				}
				break;
			case NEW:
				enterOuterAlt(_localctx, 4);
				{
				setState(393);
				allocation();
				}
				break;
			case BOOL_LIT:
				enterOuterAlt(_localctx, 5);
				{
				setState(394);
				match(BOOL_LIT);
				}
				break;
			case NIL_LIT:
				enterOuterAlt(_localctx, 6);
				{
				setState(395);
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

	public static class SubfactorContext extends ParserRuleContext {
		public ExpressionContext expr;
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public SubfactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subfactor; }
	}

	public final SubfactorContext subfactor() throws RecognitionException {
		SubfactorContext _localctx = new SubfactorContext(_ctx, getState());
		enterRule(_localctx, 106, RULE_subfactor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(398);
			match(LPAREN);
			setState(399);
			((SubfactorContext)_localctx).expr = expression();
			setState(400);
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
		public Token id;
		public ArgumentsContext args;
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
		enterRule(_localctx, 108, RULE_functioncall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(402);
			((FunctioncallContext)_localctx).id = match(IDENT);
			setState(404);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LPAREN) {
				{
				setState(403);
				((FunctioncallContext)_localctx).args = arguments();
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
		enterRule(_localctx, 110, RULE_allocation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(406);
			match(NEW);
			setState(407);
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3,\u019c\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\3\2\3\2\3\2\3\2\3\2\3\3"+
		"\7\3y\n\3\f\3\16\3|\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\7"+
		"\5\u0089\n\5\f\5\16\5\u008c\13\5\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b\3"+
		"\b\5\b\u0098\n\b\3\t\7\t\u009b\n\t\f\t\16\t\u009e\13\t\3\n\3\n\3\n\3\n"+
		"\3\n\3\13\3\13\7\13\u00a7\n\13\f\13\16\13\u00aa\13\13\3\f\3\f\3\f\3\r"+
		"\7\r\u00b0\n\r\f\r\16\r\u00b3\13\r\3\16\3\16\3\16\3\16\5\16\u00b9\n\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\17\3\17\5\17\u00c2\n\17\3\17\3\17\3\20\3\20"+
		"\7\20\u00c8\n\20\f\20\16\20\u00cb\13\20\3\21\3\21\3\21\3\22\3\22\3\23"+
		"\7\23\u00d3\n\23\f\23\16\23\u00d6\13\23\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\5\24\u00e1\n\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26"+
		"\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\7\31"+
		"\u00f8\n\31\f\31\16\31\u00fb\13\31\3\31\3\31\3\31\3\32\3\32\3\32\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\5\33\u0109\n\33\3\34\3\34\3\34\3\35\3\35\3\35"+
		"\3\35\3\35\3\35\3\36\3\36\5\36\u0116\n\36\3\36\3\36\3\37\3\37\3\37\3\37"+
		"\3 \3 \5 \u0120\n \3 \3 \3!\3!\7!\u0126\n!\f!\16!\u0129\13!\3\"\3\"\3"+
		"\"\3#\3#\7#\u0130\n#\f#\16#\u0133\13#\3$\3$\3$\3%\3%\7%\u013a\n%\f%\16"+
		"%\u013d\13%\3&\3&\3&\3\'\3\'\7\'\u0144\n\'\f\'\16\'\u0147\13\'\3(\3(\3"+
		"(\3)\3)\7)\u014e\n)\f)\16)\u0151\13)\3*\3*\3*\3+\3+\7+\u0158\n+\f+\16"+
		"+\u015b\13+\3,\3,\3,\3-\3-\7-\u0162\n-\f-\16-\u0165\13-\3.\3.\3.\3/\3"+
		"/\7/\u016c\n/\f/\16/\u016f\13/\3\60\3\60\3\60\3\61\3\61\3\61\5\61\u0177"+
		"\n\61\3\62\3\62\3\62\3\63\3\63\3\63\3\64\3\64\7\64\u0181\n\64\f\64\16"+
		"\64\u0184\13\64\3\65\3\65\3\65\3\66\3\66\3\66\3\66\3\66\3\66\5\66\u018f"+
		"\n\66\3\67\3\67\3\67\3\67\38\38\58\u0197\n8\39\39\39\39\2\2:\2\4\6\b\n"+
		"\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\\"+
		"^`bdfhjlnp\2\6\3\2\13\f\3\2\7\n\3\2\r\16\3\2\17\20\2\u018b\2r\3\2\2\2"+
		"\4z\3\2\2\2\6}\3\2\2\2\b\u0085\3\2\2\2\n\u008d\3\2\2\2\f\u0090\3\2\2\2"+
		"\16\u0097\3\2\2\2\20\u009c\3\2\2\2\22\u009f\3\2\2\2\24\u00a4\3\2\2\2\26"+
		"\u00ab\3\2\2\2\30\u00b1\3\2\2\2\32\u00b4\3\2\2\2\34\u00bf\3\2\2\2\36\u00c5"+
		"\3\2\2\2 \u00cc\3\2\2\2\"\u00cf\3\2\2\2$\u00d4\3\2\2\2&\u00e0\3\2\2\2"+
		"(\u00e2\3\2\2\2*\u00e6\3\2\2\2,\u00ea\3\2\2\2.\u00ee\3\2\2\2\60\u00f3"+
		"\3\2\2\2\62\u00ff\3\2\2\2\64\u0102\3\2\2\2\66\u010a\3\2\2\28\u010d\3\2"+
		"\2\2:\u0113\3\2\2\2<\u0119\3\2\2\2>\u011d\3\2\2\2@\u0123\3\2\2\2B\u012a"+
		"\3\2\2\2D\u012d\3\2\2\2F\u0134\3\2\2\2H\u0137\3\2\2\2J\u013e\3\2\2\2L"+
		"\u0141\3\2\2\2N\u0148\3\2\2\2P\u014b\3\2\2\2R\u0152\3\2\2\2T\u0155\3\2"+
		"\2\2V\u015c\3\2\2\2X\u015f\3\2\2\2Z\u0166\3\2\2\2\\\u0169\3\2\2\2^\u0170"+
		"\3\2\2\2`\u0176\3\2\2\2b\u0178\3\2\2\2d\u017b\3\2\2\2f\u017e\3\2\2\2h"+
		"\u0185\3\2\2\2j\u018e\3\2\2\2l\u0190\3\2\2\2n\u0194\3\2\2\2p\u0198\3\2"+
		"\2\2rs\5\4\3\2st\5\20\t\2tu\5\30\r\2uv\7\2\2\3v\3\3\2\2\2wy\5\6\4\2xw"+
		"\3\2\2\2y|\3\2\2\2zx\3\2\2\2z{\3\2\2\2{\5\3\2\2\2|z\3\2\2\2}~\7\30\2\2"+
		"~\177\7*\2\2\177\u0080\7\33\2\2\u0080\u0081\7\23\2\2\u0081\u0082\5\b\5"+
		"\2\u0082\u0083\7\24\2\2\u0083\u0084\7\25\2\2\u0084\7\3\2\2\2\u0085\u0086"+
		"\5\f\7\2\u0086\u008a\7\25\2\2\u0087\u0089\5\n\6\2\u0088\u0087\3\2\2\2"+
		"\u0089\u008c\3\2\2\2\u008a\u0088\3\2\2\2\u008a\u008b\3\2\2\2\u008b\t\3"+
		"\2\2\2\u008c\u008a\3\2\2\2\u008d\u008e\5\f\7\2\u008e\u008f\7\25\2\2\u008f"+
		"\13\3\2\2\2\u0090\u0091\7*\2\2\u0091\u0092\5\16\b\2\u0092\r\3\2\2\2\u0093"+
		"\u0098\7(\2\2\u0094\u0098\7)\2\2\u0095\u0096\7\17\2\2\u0096\u0098\7*\2"+
		"\2\u0097\u0093\3\2\2\2\u0097\u0094\3\2\2\2\u0097\u0095\3\2\2\2\u0098\17"+
		"\3\2\2\2\u0099\u009b\5\22\n\2\u009a\u0099\3\2\2\2\u009b\u009e\3\2\2\2"+
		"\u009c\u009a\3\2\2\2\u009c\u009d\3\2\2\2\u009d\21\3\2\2\2\u009e\u009c"+
		"\3\2\2\2\u009f\u00a0\7\27\2\2\u00a0\u00a1\5\24\13\2\u00a1\u00a2\5\16\b"+
		"\2\u00a2\u00a3\7\25\2\2\u00a3\23\3\2\2\2\u00a4\u00a8\7*\2\2\u00a5\u00a7"+
		"\5\26\f\2\u00a6\u00a5\3\2\2\2\u00a7\u00aa\3\2\2\2\u00a8\u00a6\3\2\2\2"+
		"\u00a8\u00a9\3\2\2\2\u00a9\25\3\2\2\2\u00aa\u00a8\3\2\2\2\u00ab\u00ac"+
		"\7\26\2\2\u00ac\u00ad\7*\2\2\u00ad\27\3\2\2\2\u00ae\u00b0\5\32\16\2\u00af"+
		"\u00ae\3\2\2\2\u00b0\u00b3\3\2\2\2\u00b1\u00af\3\2\2\2\u00b1\u00b2\3\2"+
		"\2\2\u00b2\31\3\2\2\2\u00b3\u00b1\3\2\2\2\u00b4\u00b5\7\35\2\2\u00b5\u00b6"+
		"\7*\2\2\u00b6\u00b8\5\34\17\2\u00b7\u00b9\5\"\22\2\u00b8\u00b7\3\2\2\2"+
		"\u00b8\u00b9\3\2\2\2\u00b9\u00ba\3\2\2\2\u00ba\u00bb\7\23\2\2\u00bb\u00bc"+
		"\5\20\t\2\u00bc\u00bd\5$\23\2\u00bd\u00be\7\24\2\2\u00be\33\3\2\2\2\u00bf"+
		"\u00c1\7\21\2\2\u00c0\u00c2\5\36\20\2\u00c1\u00c0\3\2\2\2\u00c1\u00c2"+
		"\3\2\2\2\u00c2\u00c3\3\2\2\2\u00c3\u00c4\7\22\2\2\u00c4\35\3\2\2\2\u00c5"+
		"\u00c9\5\f\7\2\u00c6\u00c8\5 \21\2\u00c7\u00c6\3\2\2\2\u00c8\u00cb\3\2"+
		"\2\2\u00c9\u00c7\3\2\2\2\u00c9\u00ca\3\2\2\2\u00ca\37\3\2\2\2\u00cb\u00c9"+
		"\3\2\2\2\u00cc\u00cd\7\26\2\2\u00cd\u00ce\5\f\7\2\u00ce!\3\2\2\2\u00cf"+
		"\u00d0\5\16\b\2\u00d0#\3\2\2\2\u00d1\u00d3\5&\24\2\u00d2\u00d1\3\2\2\2"+
		"\u00d3\u00d6\3\2\2\2\u00d4\u00d2\3\2\2\2\u00d4\u00d5\3\2\2\2\u00d5%\3"+
		"\2\2\2\u00d6\u00d4\3\2\2\2\u00d7\u00e1\5*\26\2\u00d8\u00e1\5.\30\2\u00d9"+
		"\u00e1\5\60\31\2\u00da\u00e1\5,\27\2\u00db\u00e1\5(\25\2\u00dc\u00e1\5"+
		"\64\33\2\u00dd\u00e1\58\35\2\u00de\u00e1\5:\36\2\u00df\u00e1\5<\37\2\u00e0"+
		"\u00d7\3\2\2\2\u00e0\u00d8\3\2\2\2\u00e0\u00d9\3\2\2\2\u00e0\u00da\3\2"+
		"\2\2\u00e0\u00db\3\2\2\2\u00e0\u00dc\3\2\2\2\u00e0\u00dd\3\2\2\2\u00e0"+
		"\u00de\3\2\2\2\u00e0\u00df\3\2\2\2\u00e1\'\3\2\2\2\u00e2\u00e3\7\"\2\2"+
		"\u00e3\u00e4\5D#\2\u00e4\u00e5\7\25\2\2\u00e5)\3\2\2\2\u00e6\u00e7\7\23"+
		"\2\2\u00e7\u00e8\5$\23\2\u00e8\u00e9\7\24\2\2\u00e9+\3\2\2\2\u00ea\u00eb"+
		"\7\32\2\2\u00eb\u00ec\5H%\2\u00ec\u00ed\7\25\2\2\u00ed-\3\2\2\2\u00ee"+
		"\u00ef\5D#\2\u00ef\u00f0\7\3\2\2\u00f0\u00f1\5H%\2\u00f1\u00f2\7\25\2"+
		"\2\u00f2/\3\2\2\2\u00f3\u00f4\7#\2\2\u00f4\u00f5\7\21\2\2\u00f5\u00f9"+
		"\7%\2\2\u00f6\u00f8\5\62\32\2\u00f7\u00f6\3\2\2\2\u00f8\u00fb\3\2\2\2"+
		"\u00f9\u00f7\3\2\2\2\u00f9\u00fa\3\2\2\2\u00fa\u00fc\3\2\2\2\u00fb\u00f9"+
		"\3\2\2\2\u00fc\u00fd\7\22\2\2\u00fd\u00fe\7\25\2\2\u00fe\61\3\2\2\2\u00ff"+
		"\u0100\7\26\2\2\u0100\u0101\5H%\2\u0101\63\3\2\2\2\u0102\u0103\7\37\2"+
		"\2\u0103\u0104\7\21\2\2\u0104\u0105\5H%\2\u0105\u0106\7\22\2\2\u0106\u0108"+
		"\5*\26\2\u0107\u0109\5\66\34\2\u0108\u0107\3\2\2\2\u0108\u0109\3\2\2\2"+
		"\u0109\65\3\2\2\2\u010a\u010b\7 \2\2\u010b\u010c\5*\26\2\u010c\67\3\2"+
		"\2\2\u010d\u010e\7!\2\2\u010e\u010f\7\21\2\2\u010f\u0110\5H%\2\u0110\u0111"+
		"\7\22\2\2\u0111\u0112\5*\26\2\u01129\3\2\2\2\u0113\u0115\7\36\2\2\u0114"+
		"\u0116\5H%\2\u0115\u0114\3\2\2\2\u0115\u0116\3\2\2\2\u0116\u0117\3\2\2"+
		"\2\u0117\u0118\7\25\2\2\u0118;\3\2\2\2\u0119\u011a\7*\2\2\u011a\u011b"+
		"\5> \2\u011b\u011c\7\25\2\2\u011c=\3\2\2\2\u011d\u011f\7\21\2\2\u011e"+
		"\u0120\5@!\2\u011f\u011e\3\2\2\2\u011f\u0120\3\2\2\2\u0120\u0121\3\2\2"+
		"\2\u0121\u0122\7\22\2\2\u0122?\3\2\2\2\u0123\u0127\5H%\2\u0124\u0126\5"+
		"B\"\2\u0125\u0124\3\2\2\2\u0126\u0129\3\2\2\2\u0127\u0125\3\2\2\2\u0127"+
		"\u0128\3\2\2\2\u0128A\3\2\2\2\u0129\u0127\3\2\2\2\u012a\u012b\7\26\2\2"+
		"\u012b\u012c\5H%\2\u012cC\3\2\2\2\u012d\u0131\7*\2\2\u012e\u0130\5F$\2"+
		"\u012f\u012e\3\2\2\2\u0130\u0133\3\2\2\2\u0131\u012f\3\2\2\2\u0131\u0132"+
		"\3\2\2\2\u0132E\3\2\2\2\u0133\u0131\3\2\2\2\u0134\u0135\7\34\2\2\u0135"+
		"\u0136\7*\2\2\u0136G\3\2\2\2\u0137\u013b\5L\'\2\u0138\u013a\5J&\2\u0139"+
		"\u0138\3\2\2\2\u013a\u013d\3\2\2\2\u013b\u0139\3\2\2\2\u013b\u013c\3\2"+
		"\2\2\u013cI\3\2\2\2\u013d\u013b\3\2\2\2\u013e\u013f\7\5\2\2\u013f\u0140"+
		"\5L\'\2\u0140K\3\2\2\2\u0141\u0145\5P)\2\u0142\u0144\5N(\2\u0143\u0142"+
		"\3\2\2\2\u0144\u0147\3\2\2\2\u0145\u0143\3\2\2\2\u0145\u0146\3\2\2\2\u0146"+
		"M\3\2\2\2\u0147\u0145\3\2\2\2\u0148\u0149\7\4\2\2\u0149\u014a\5P)\2\u014a"+
		"O\3\2\2\2\u014b\u014f\5T+\2\u014c\u014e\5R*\2\u014d\u014c\3\2\2\2\u014e"+
		"\u0151\3\2\2\2\u014f\u014d\3\2\2\2\u014f\u0150\3\2\2\2\u0150Q\3\2\2\2"+
		"\u0151\u014f\3\2\2\2\u0152\u0153\t\2\2\2\u0153\u0154\5T+\2\u0154S\3\2"+
		"\2\2\u0155\u0159\5X-\2\u0156\u0158\5V,\2\u0157\u0156\3\2\2\2\u0158\u015b"+
		"\3\2\2\2\u0159\u0157\3\2\2\2\u0159\u015a\3\2\2\2\u015aU\3\2\2\2\u015b"+
		"\u0159\3\2\2\2\u015c\u015d\t\3\2\2\u015d\u015e\5X-\2\u015eW\3\2\2\2\u015f"+
		"\u0163\5\\/\2\u0160\u0162\5Z.\2\u0161\u0160\3\2\2\2\u0162\u0165\3\2\2"+
		"\2\u0163\u0161\3\2\2\2\u0163\u0164\3\2\2\2\u0164Y\3\2\2\2\u0165\u0163"+
		"\3\2\2\2\u0166\u0167\t\4\2\2\u0167\u0168\5\\/\2\u0168[\3\2\2\2\u0169\u016d"+
		"\5`\61\2\u016a\u016c\5^\60\2\u016b\u016a\3\2\2\2\u016c\u016f\3\2\2\2\u016d"+
		"\u016b\3\2\2\2\u016d\u016e\3\2\2\2\u016e]\3\2\2\2\u016f\u016d\3\2\2\2"+
		"\u0170\u0171\t\5\2\2\u0171\u0172\5`\61\2\u0172_\3\2\2\2\u0173\u0177\5"+
		"b\62\2\u0174\u0177\5d\63\2\u0175\u0177\5f\64\2\u0176\u0173\3\2\2\2\u0176"+
		"\u0174\3\2\2\2\u0176\u0175\3\2\2\2\u0177a\3\2\2\2\u0178\u0179\7\6\2\2"+
		"\u0179\u017a\5f\64\2\u017ac\3\2\2\2\u017b\u017c\7\16\2\2\u017c\u017d\5"+
		"f\64\2\u017de\3\2\2\2\u017e\u0182\5j\66\2\u017f\u0181\5h\65\2\u0180\u017f"+
		"\3\2\2\2\u0181\u0184\3\2\2\2\u0182\u0180\3\2\2\2\u0182\u0183\3\2\2\2\u0183"+
		"g\3\2\2\2\u0184\u0182\3\2\2\2\u0185\u0186\7\34\2\2\u0186\u0187\7*\2\2"+
		"\u0187i\3\2\2\2\u0188\u018f\5l\67\2\u0189\u018f\5n8\2\u018a\u018f\7$\2"+
		"\2\u018b\u018f\5p9\2\u018c\u018f\7&\2\2\u018d\u018f\7\'\2\2\u018e\u0188"+
		"\3\2\2\2\u018e\u0189\3\2\2\2\u018e\u018a\3\2\2\2\u018e\u018b\3\2\2\2\u018e"+
		"\u018c\3\2\2\2\u018e\u018d\3\2\2\2\u018fk\3\2\2\2\u0190\u0191\7\21\2\2"+
		"\u0191\u0192\5H%\2\u0192\u0193\7\22\2\2\u0193m\3\2\2\2\u0194\u0196\7*"+
		"\2\2\u0195\u0197\5> \2\u0196\u0195\3\2\2\2\u0196\u0197\3\2\2\2\u0197o"+
		"\3\2\2\2\u0198\u0199\7\31\2\2\u0199\u019a\7*\2\2\u019aq\3\2\2\2\35z\u008a"+
		"\u0097\u009c\u00a8\u00b1\u00b8\u00c1\u00c9\u00d4\u00e0\u00f9\u0108\u0115"+
		"\u011f\u0127\u0131\u013b\u0145\u014f\u0159\u0163\u016d\u0176\u0182\u018e"+
		"\u0196";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}