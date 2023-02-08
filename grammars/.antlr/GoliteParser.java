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
		RULE_decl = 4, RULE_type = 5, RULE_declarations = 6, RULE_declaration = 7, 
		RULE_ids = 8, RULE_functions = 9, RULE_function = 10, RULE_parameters = 11, 
		RULE_returnType = 12, RULE_statements = 13, RULE_statement = 14, RULE_read = 15, 
		RULE_block = 16, RULE_delete = 17, RULE_assignment = 18, RULE_print = 19, 
		RULE_conditional = 20, RULE_loop = 21, RULE_returnRule = 22, RULE_invocation = 23, 
		RULE_arguments = 24, RULE_lValue = 25, RULE_expression = 26, RULE_boolTerm = 27, 
		RULE_equalTerm = 28, RULE_relationTerm = 29, RULE_simpleTerm = 30, RULE_term = 31, 
		RULE_unaryTerm = 32, RULE_selectorTerm = 33, RULE_factor = 34;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "types", "typeDeclaration", "fields", "decl", "type", "declarations", 
			"declaration", "ids", "functions", "function", "parameters", "returnType", 
			"statements", "statement", "read", "block", "delete", "assignment", "print", 
			"conditional", "loop", "returnRule", "invocation", "arguments", "lValue", 
			"expression", "boolTerm", "equalTerm", "relationTerm", "simpleTerm", 
			"term", "unaryTerm", "selectorTerm", "factor"
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
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public DeclarationsContext declarations() {
			return getRuleContext(DeclarationsContext.class,0);
		}
		public FunctionsContext functions() {
			return getRuleContext(FunctionsContext.class,0);
		}
		public TerminalNode EOF() { return getToken(GoliteParser.EOF, 0); }
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
			setState(70);
			types();
			setState(71);
			declarations();
			setState(72);
			functions();
			setState(73);
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
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TYPE) {
				{
				{
				setState(75);
				typeDeclaration();
				}
				}
				setState(80);
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
		public TerminalNode TYPE() { return getToken(GoliteParser.TYPE, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public TerminalNode STRUCT() { return getToken(GoliteParser.STRUCT, 0); }
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public FieldsContext fields() {
			return getRuleContext(FieldsContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
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
			setState(81);
			match(TYPE);
			setState(82);
			match(IDENT);
			setState(83);
			match(STRUCT);
			setState(84);
			match(LBRACE);
			setState(85);
			fields();
			setState(86);
			match(RBRACE);
			setState(87);
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
		public List<DeclContext> decl() {
			return getRuleContexts(DeclContext.class);
		}
		public DeclContext decl(int i) {
			return getRuleContext(DeclContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(GoliteParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(GoliteParser.SEMICOLON, i);
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
			setState(89);
			decl();
			setState(90);
			match(SEMICOLON);
			setState(96);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENT) {
				{
				{
				setState(91);
				decl();
				setState(92);
				match(SEMICOLON);
				}
				}
				setState(98);
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

	public static class DeclContext extends ParserRuleContext {
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
		enterRule(_localctx, 8, RULE_decl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(99);
			match(IDENT);
			setState(100);
			type();
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
		enterRule(_localctx, 10, RULE_type);
		try {
			setState(106);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				match(INT);
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				match(BOOL);
				}
				break;
			case MULT:
				enterOuterAlt(_localctx, 3);
				{
				{
				setState(104);
				match(MULT);
				setState(105);
				match(IDENT);
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
		enterRule(_localctx, 12, RULE_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(108);
				declaration();
				}
				}
				setState(113);
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
		public TerminalNode VAR() { return getToken(GoliteParser.VAR, 0); }
		public IdsContext ids() {
			return getRuleContext(IdsContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			match(VAR);
			setState(115);
			ids();
			setState(116);
			type();
			setState(117);
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
		public List<TerminalNode> IDENT() { return getTokens(GoliteParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(GoliteParser.IDENT, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public IdsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ids; }
	}

	public final IdsContext ids() throws RecognitionException {
		IdsContext _localctx = new IdsContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_ids);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			match(IDENT);
			setState(124);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(120);
				match(COMMA);
				setState(121);
				match(IDENT);
				}
				}
				setState(126);
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
		enterRule(_localctx, 18, RULE_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNC) {
				{
				{
				setState(127);
				function();
				}
				}
				setState(132);
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
		public TerminalNode FUNC() { return getToken(GoliteParser.FUNC, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public DeclarationsContext declarations() {
			return getRuleContext(DeclarationsContext.class,0);
		}
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
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
		enterRule(_localctx, 20, RULE_function);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			match(FUNC);
			setState(134);
			match(IDENT);
			setState(135);
			parameters();
			setState(137);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MULT) | (1L << INT) | (1L << BOOL))) != 0)) {
				{
				setState(136);
				returnType();
				}
			}

			setState(139);
			match(LBRACE);
			setState(140);
			declarations();
			setState(141);
			statements();
			setState(142);
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
		public List<DeclContext> decl() {
			return getRuleContexts(DeclContext.class);
		}
		public DeclContext decl(int i) {
			return getRuleContext(DeclContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public ParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters; }
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			match(LPAREN);
			setState(153);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENT) {
				{
				setState(145);
				decl();
				setState(150);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(146);
					match(COMMA);
					setState(147);
					decl();
					}
					}
					setState(152);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(155);
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

	public static class ReturnTypeContext extends ParserRuleContext {
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
		enterRule(_localctx, 24, RULE_returnType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(157);
			type();
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
		enterRule(_localctx, 26, RULE_statements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LBRACE) | (1L << DELETE) | (1L << RET) | (1L << IF) | (1L << FOR) | (1L << SCAN) | (1L << PRINTF) | (1L << IDENT))) != 0)) {
				{
				{
				setState(159);
				statement();
				}
				}
				setState(164);
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
		enterRule(_localctx, 28, RULE_statement);
		try {
			setState(174);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(165);
				block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(166);
				assignment();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(167);
				print();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(168);
				delete();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(169);
				read();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(170);
				conditional();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(171);
				loop();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(172);
				returnRule();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(173);
				invocation();
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
		enterRule(_localctx, 30, RULE_read);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(176);
			match(SCAN);
			setState(177);
			lValue();
			setState(178);
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
		enterRule(_localctx, 32, RULE_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(180);
			match(LBRACE);
			setState(181);
			statements();
			setState(182);
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
		enterRule(_localctx, 34, RULE_delete);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			match(DELETE);
			setState(185);
			expression();
			setState(186);
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
		enterRule(_localctx, 36, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			lValue();
			setState(189);
			match(ASSIGN);
			setState(190);
			expression();
			setState(191);
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
		enterRule(_localctx, 38, RULE_print);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(193);
			match(PRINTF);
			setState(194);
			match(LPAREN);
			setState(195);
			match(STRING_LIT);
			setState(200);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(196);
				match(COMMA);
				setState(197);
				expression();
				}
				}
				setState(202);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(203);
			match(RPAREN);
			setState(204);
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
		enterRule(_localctx, 40, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(206);
			match(IF);
			setState(207);
			match(LPAREN);
			setState(208);
			expression();
			setState(209);
			match(RPAREN);
			setState(210);
			block();
			setState(213);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(211);
				match(ELSE);
				setState(212);
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
		enterRule(_localctx, 42, RULE_loop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(215);
			match(FOR);
			setState(216);
			match(LPAREN);
			setState(217);
			expression();
			setState(218);
			match(RPAREN);
			setState(219);
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
		enterRule(_localctx, 44, RULE_returnRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(221);
			match(RET);
			setState(223);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NOT) | (1L << MINUS) | (1L << LPAREN) | (1L << NEW) | (1L << INT_LIT) | (1L << BOOL_LIT) | (1L << NIL_LIT) | (1L << IDENT))) != 0)) {
				{
				setState(222);
				expression();
				}
			}

			setState(225);
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
		enterRule(_localctx, 46, RULE_invocation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(227);
			match(IDENT);
			setState(228);
			arguments();
			setState(229);
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
		enterRule(_localctx, 48, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231);
			match(LPAREN);
			setState(240);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NOT) | (1L << MINUS) | (1L << LPAREN) | (1L << NEW) | (1L << INT_LIT) | (1L << BOOL_LIT) | (1L << NIL_LIT) | (1L << IDENT))) != 0)) {
				{
				setState(232);
				expression();
				setState(237);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(233);
					match(COMMA);
					setState(234);
					expression();
					}
					}
					setState(239);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(242);
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
		enterRule(_localctx, 50, RULE_lValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(244);
			match(IDENT);
			setState(249);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(245);
				match(DOT);
				setState(246);
				match(IDENT);
				}
				}
				setState(251);
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
		enterRule(_localctx, 52, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(252);
			boolTerm();
			setState(257);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(253);
				match(OR);
				setState(254);
				boolTerm();
				}
				}
				setState(259);
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
		enterRule(_localctx, 54, RULE_boolTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(260);
			equalTerm();
			setState(265);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(261);
				match(AND);
				setState(262);
				equalTerm();
				}
				}
				setState(267);
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
		enterRule(_localctx, 56, RULE_equalTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(268);
			relationTerm();
			setState(273);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EQ || _la==NE) {
				{
				{
				setState(269);
				_la = _input.LA(1);
				if ( !(_la==EQ || _la==NE) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(270);
				relationTerm();
				}
				}
				setState(275);
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
		enterRule(_localctx, 58, RULE_relationTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			simpleTerm();
			setState(281);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << LE) | (1L << GE))) != 0)) {
				{
				{
				setState(277);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << LE) | (1L << GE))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(278);
				simpleTerm();
				}
				}
				setState(283);
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
		enterRule(_localctx, 60, RULE_simpleTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(284);
			term();
			setState(289);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PLUS || _la==MINUS) {
				{
				{
				setState(285);
				_la = _input.LA(1);
				if ( !(_la==PLUS || _la==MINUS) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(286);
				term();
				}
				}
				setState(291);
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
		enterRule(_localctx, 62, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(292);
			unaryTerm();
			setState(297);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MULT || _la==DIV) {
				{
				{
				setState(293);
				_la = _input.LA(1);
				if ( !(_la==MULT || _la==DIV) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(294);
				unaryTerm();
				}
				}
				setState(299);
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
		enterRule(_localctx, 64, RULE_unaryTerm);
		try {
			setState(305);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				enterOuterAlt(_localctx, 1);
				{
				setState(300);
				match(NOT);
				setState(301);
				selectorTerm();
				}
				break;
			case MINUS:
				enterOuterAlt(_localctx, 2);
				{
				setState(302);
				match(MINUS);
				setState(303);
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
				setState(304);
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
		public List<TerminalNode> DOT() { return getTokens(GoliteParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(GoliteParser.DOT, i);
		}
		public List<TerminalNode> IDENT() { return getTokens(GoliteParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(GoliteParser.IDENT, i);
		}
		public SelectorTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectorTerm; }
	}

	public final SelectorTermContext selectorTerm() throws RecognitionException {
		SelectorTermContext _localctx = new SelectorTermContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_selectorTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(307);
			factor();
			setState(312);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(308);
				match(DOT);
				setState(309);
				match(IDENT);
				}
				}
				setState(314);
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

	public static class FactorContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public TerminalNode IDENT() { return getToken(GoliteParser.IDENT, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public TerminalNode INT_LIT() { return getToken(GoliteParser.INT_LIT, 0); }
		public TerminalNode NEW() { return getToken(GoliteParser.NEW, 0); }
		public TerminalNode BOOL_LIT() { return getToken(GoliteParser.BOOL_LIT, 0); }
		public TerminalNode NIL_LIT() { return getToken(GoliteParser.NIL_LIT, 0); }
		public FactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor; }
	}

	public final FactorContext factor() throws RecognitionException {
		FactorContext _localctx = new FactorContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_factor);
		int _la;
		try {
			setState(328);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				enterOuterAlt(_localctx, 1);
				{
				setState(315);
				match(LPAREN);
				setState(316);
				expression();
				setState(317);
				match(RPAREN);
				}
				break;
			case IDENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(319);
				match(IDENT);
				setState(321);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LPAREN) {
					{
					setState(320);
					arguments();
					}
				}

				}
				break;
			case INT_LIT:
				enterOuterAlt(_localctx, 3);
				{
				setState(323);
				match(INT_LIT);
				}
				break;
			case NEW:
				enterOuterAlt(_localctx, 4);
				{
				setState(324);
				match(NEW);
				setState(325);
				match(IDENT);
				}
				break;
			case BOOL_LIT:
				enterOuterAlt(_localctx, 5);
				{
				setState(326);
				match(BOOL_LIT);
				}
				break;
			case NIL_LIT:
				enterOuterAlt(_localctx, 6);
				{
				setState(327);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3,\u014d\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\3\2\3\2\3\2\3\2\3\2\3\3\7\3O\n\3\f\3\16\3R\13"+
		"\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\7\5a\n\5\f\5\16"+
		"\5d\13\5\3\6\3\6\3\6\3\7\3\7\3\7\3\7\5\7m\n\7\3\b\7\bp\n\b\f\b\16\bs\13"+
		"\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\7\n}\n\n\f\n\16\n\u0080\13\n\3\13\7"+
		"\13\u0083\n\13\f\13\16\13\u0086\13\13\3\f\3\f\3\f\3\f\5\f\u008c\n\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\7\r\u0097\n\r\f\r\16\r\u009a\13\r\5"+
		"\r\u009c\n\r\3\r\3\r\3\16\3\16\3\17\7\17\u00a3\n\17\f\17\16\17\u00a6\13"+
		"\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u00b1\n\20\3\21"+
		"\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\24\3\24\3\24"+
		"\3\24\3\24\3\25\3\25\3\25\3\25\3\25\7\25\u00c9\n\25\f\25\16\25\u00cc\13"+
		"\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u00d8\n\26"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\5\30\u00e2\n\30\3\30\3\30\3\31"+
		"\3\31\3\31\3\31\3\32\3\32\3\32\3\32\7\32\u00ee\n\32\f\32\16\32\u00f1\13"+
		"\32\5\32\u00f3\n\32\3\32\3\32\3\33\3\33\3\33\7\33\u00fa\n\33\f\33\16\33"+
		"\u00fd\13\33\3\34\3\34\3\34\7\34\u0102\n\34\f\34\16\34\u0105\13\34\3\35"+
		"\3\35\3\35\7\35\u010a\n\35\f\35\16\35\u010d\13\35\3\36\3\36\3\36\7\36"+
		"\u0112\n\36\f\36\16\36\u0115\13\36\3\37\3\37\3\37\7\37\u011a\n\37\f\37"+
		"\16\37\u011d\13\37\3 \3 \3 \7 \u0122\n \f \16 \u0125\13 \3!\3!\3!\7!\u012a"+
		"\n!\f!\16!\u012d\13!\3\"\3\"\3\"\3\"\3\"\5\"\u0134\n\"\3#\3#\3#\7#\u0139"+
		"\n#\f#\16#\u013c\13#\3$\3$\3$\3$\3$\3$\5$\u0144\n$\3$\3$\3$\3$\3$\5$\u014b"+
		"\n$\3$\2\2%\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\66"+
		"8:<>@BDF\2\6\3\2\13\f\3\2\7\n\3\2\r\16\3\2\17\20\2\u0151\2H\3\2\2\2\4"+
		"P\3\2\2\2\6S\3\2\2\2\b[\3\2\2\2\ne\3\2\2\2\fl\3\2\2\2\16q\3\2\2\2\20t"+
		"\3\2\2\2\22y\3\2\2\2\24\u0084\3\2\2\2\26\u0087\3\2\2\2\30\u0092\3\2\2"+
		"\2\32\u009f\3\2\2\2\34\u00a4\3\2\2\2\36\u00b0\3\2\2\2 \u00b2\3\2\2\2\""+
		"\u00b6\3\2\2\2$\u00ba\3\2\2\2&\u00be\3\2\2\2(\u00c3\3\2\2\2*\u00d0\3\2"+
		"\2\2,\u00d9\3\2\2\2.\u00df\3\2\2\2\60\u00e5\3\2\2\2\62\u00e9\3\2\2\2\64"+
		"\u00f6\3\2\2\2\66\u00fe\3\2\2\28\u0106\3\2\2\2:\u010e\3\2\2\2<\u0116\3"+
		"\2\2\2>\u011e\3\2\2\2@\u0126\3\2\2\2B\u0133\3\2\2\2D\u0135\3\2\2\2F\u014a"+
		"\3\2\2\2HI\5\4\3\2IJ\5\16\b\2JK\5\24\13\2KL\7\2\2\3L\3\3\2\2\2MO\5\6\4"+
		"\2NM\3\2\2\2OR\3\2\2\2PN\3\2\2\2PQ\3\2\2\2Q\5\3\2\2\2RP\3\2\2\2ST\7\30"+
		"\2\2TU\7*\2\2UV\7\33\2\2VW\7\23\2\2WX\5\b\5\2XY\7\24\2\2YZ\7\25\2\2Z\7"+
		"\3\2\2\2[\\\5\n\6\2\\b\7\25\2\2]^\5\n\6\2^_\7\25\2\2_a\3\2\2\2`]\3\2\2"+
		"\2ad\3\2\2\2b`\3\2\2\2bc\3\2\2\2c\t\3\2\2\2db\3\2\2\2ef\7*\2\2fg\5\f\7"+
		"\2g\13\3\2\2\2hm\7(\2\2im\7)\2\2jk\7\17\2\2km\7*\2\2lh\3\2\2\2li\3\2\2"+
		"\2lj\3\2\2\2m\r\3\2\2\2np\5\20\t\2on\3\2\2\2ps\3\2\2\2qo\3\2\2\2qr\3\2"+
		"\2\2r\17\3\2\2\2sq\3\2\2\2tu\7\27\2\2uv\5\22\n\2vw\5\f\7\2wx\7\25\2\2"+
		"x\21\3\2\2\2y~\7*\2\2z{\7\26\2\2{}\7*\2\2|z\3\2\2\2}\u0080\3\2\2\2~|\3"+
		"\2\2\2~\177\3\2\2\2\177\23\3\2\2\2\u0080~\3\2\2\2\u0081\u0083\5\26\f\2"+
		"\u0082\u0081\3\2\2\2\u0083\u0086\3\2\2\2\u0084\u0082\3\2\2\2\u0084\u0085"+
		"\3\2\2\2\u0085\25\3\2\2\2\u0086\u0084\3\2\2\2\u0087\u0088\7\35\2\2\u0088"+
		"\u0089\7*\2\2\u0089\u008b\5\30\r\2\u008a\u008c\5\32\16\2\u008b\u008a\3"+
		"\2\2\2\u008b\u008c\3\2\2\2\u008c\u008d\3\2\2\2\u008d\u008e\7\23\2\2\u008e"+
		"\u008f\5\16\b\2\u008f\u0090\5\34\17\2\u0090\u0091\7\24\2\2\u0091\27\3"+
		"\2\2\2\u0092\u009b\7\21\2\2\u0093\u0098\5\n\6\2\u0094\u0095\7\26\2\2\u0095"+
		"\u0097\5\n\6\2\u0096\u0094\3\2\2\2\u0097\u009a\3\2\2\2\u0098\u0096\3\2"+
		"\2\2\u0098\u0099\3\2\2\2\u0099\u009c\3\2\2\2\u009a\u0098\3\2\2\2\u009b"+
		"\u0093\3\2\2\2\u009b\u009c\3\2\2\2\u009c\u009d\3\2\2\2\u009d\u009e\7\22"+
		"\2\2\u009e\31\3\2\2\2\u009f\u00a0\5\f\7\2\u00a0\33\3\2\2\2\u00a1\u00a3"+
		"\5\36\20\2\u00a2\u00a1\3\2\2\2\u00a3\u00a6\3\2\2\2\u00a4\u00a2\3\2\2\2"+
		"\u00a4\u00a5\3\2\2\2\u00a5\35\3\2\2\2\u00a6\u00a4\3\2\2\2\u00a7\u00b1"+
		"\5\"\22\2\u00a8\u00b1\5&\24\2\u00a9\u00b1\5(\25\2\u00aa\u00b1\5$\23\2"+
		"\u00ab\u00b1\5 \21\2\u00ac\u00b1\5*\26\2\u00ad\u00b1\5,\27\2\u00ae\u00b1"+
		"\5.\30\2\u00af\u00b1\5\60\31\2\u00b0\u00a7\3\2\2\2\u00b0\u00a8\3\2\2\2"+
		"\u00b0\u00a9\3\2\2\2\u00b0\u00aa\3\2\2\2\u00b0\u00ab\3\2\2\2\u00b0\u00ac"+
		"\3\2\2\2\u00b0\u00ad\3\2\2\2\u00b0\u00ae\3\2\2\2\u00b0\u00af\3\2\2\2\u00b1"+
		"\37\3\2\2\2\u00b2\u00b3\7\"\2\2\u00b3\u00b4\5\64\33\2\u00b4\u00b5\7\25"+
		"\2\2\u00b5!\3\2\2\2\u00b6\u00b7\7\23\2\2\u00b7\u00b8\5\34\17\2\u00b8\u00b9"+
		"\7\24\2\2\u00b9#\3\2\2\2\u00ba\u00bb\7\32\2\2\u00bb\u00bc\5\66\34\2\u00bc"+
		"\u00bd\7\25\2\2\u00bd%\3\2\2\2\u00be\u00bf\5\64\33\2\u00bf\u00c0\7\3\2"+
		"\2\u00c0\u00c1\5\66\34\2\u00c1\u00c2\7\25\2\2\u00c2\'\3\2\2\2\u00c3\u00c4"+
		"\7#\2\2\u00c4\u00c5\7\21\2\2\u00c5\u00ca\7%\2\2\u00c6\u00c7\7\26\2\2\u00c7"+
		"\u00c9\5\66\34\2\u00c8\u00c6\3\2\2\2\u00c9\u00cc\3\2\2\2\u00ca\u00c8\3"+
		"\2\2\2\u00ca\u00cb\3\2\2\2\u00cb\u00cd\3\2\2\2\u00cc\u00ca\3\2\2\2\u00cd"+
		"\u00ce\7\22\2\2\u00ce\u00cf\7\25\2\2\u00cf)\3\2\2\2\u00d0\u00d1\7\37\2"+
		"\2\u00d1\u00d2\7\21\2\2\u00d2\u00d3\5\66\34\2\u00d3\u00d4\7\22\2\2\u00d4"+
		"\u00d7\5\"\22\2\u00d5\u00d6\7 \2\2\u00d6\u00d8\5\"\22\2\u00d7\u00d5\3"+
		"\2\2\2\u00d7\u00d8\3\2\2\2\u00d8+\3\2\2\2\u00d9\u00da\7!\2\2\u00da\u00db"+
		"\7\21\2\2\u00db\u00dc\5\66\34\2\u00dc\u00dd\7\22\2\2\u00dd\u00de\5\"\22"+
		"\2\u00de-\3\2\2\2\u00df\u00e1\7\36\2\2\u00e0\u00e2\5\66\34\2\u00e1\u00e0"+
		"\3\2\2\2\u00e1\u00e2\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\u00e4\7\25\2\2"+
		"\u00e4/\3\2\2\2\u00e5\u00e6\7*\2\2\u00e6\u00e7\5\62\32\2\u00e7\u00e8\7"+
		"\25\2\2\u00e8\61\3\2\2\2\u00e9\u00f2\7\21\2\2\u00ea\u00ef\5\66\34\2\u00eb"+
		"\u00ec\7\26\2\2\u00ec\u00ee\5\66\34\2\u00ed\u00eb\3\2\2\2\u00ee\u00f1"+
		"\3\2\2\2\u00ef\u00ed\3\2\2\2\u00ef\u00f0\3\2\2\2\u00f0\u00f3\3\2\2\2\u00f1"+
		"\u00ef\3\2\2\2\u00f2\u00ea\3\2\2\2\u00f2\u00f3\3\2\2\2\u00f3\u00f4\3\2"+
		"\2\2\u00f4\u00f5\7\22\2\2\u00f5\63\3\2\2\2\u00f6\u00fb\7*\2\2\u00f7\u00f8"+
		"\7\34\2\2\u00f8\u00fa\7*\2\2\u00f9\u00f7\3\2\2\2\u00fa\u00fd\3\2\2\2\u00fb"+
		"\u00f9\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fc\65\3\2\2\2\u00fd\u00fb\3\2\2"+
		"\2\u00fe\u0103\58\35\2\u00ff\u0100\7\5\2\2\u0100\u0102\58\35\2\u0101\u00ff"+
		"\3\2\2\2\u0102\u0105\3\2\2\2\u0103\u0101\3\2\2\2\u0103\u0104\3\2\2\2\u0104"+
		"\67\3\2\2\2\u0105\u0103\3\2\2\2\u0106\u010b\5:\36\2\u0107\u0108\7\4\2"+
		"\2\u0108\u010a\5:\36\2\u0109\u0107\3\2\2\2\u010a\u010d\3\2\2\2\u010b\u0109"+
		"\3\2\2\2\u010b\u010c\3\2\2\2\u010c9\3\2\2\2\u010d\u010b\3\2\2\2\u010e"+
		"\u0113\5<\37\2\u010f\u0110\t\2\2\2\u0110\u0112\5<\37\2\u0111\u010f\3\2"+
		"\2\2\u0112\u0115\3\2\2\2\u0113\u0111\3\2\2\2\u0113\u0114\3\2\2\2\u0114"+
		";\3\2\2\2\u0115\u0113\3\2\2\2\u0116\u011b\5> \2\u0117\u0118\t\3\2\2\u0118"+
		"\u011a\5> \2\u0119\u0117\3\2\2\2\u011a\u011d\3\2\2\2\u011b\u0119\3\2\2"+
		"\2\u011b\u011c\3\2\2\2\u011c=\3\2\2\2\u011d\u011b\3\2\2\2\u011e\u0123"+
		"\5@!\2\u011f\u0120\t\4\2\2\u0120\u0122\5@!\2\u0121\u011f\3\2\2\2\u0122"+
		"\u0125\3\2\2\2\u0123\u0121\3\2\2\2\u0123\u0124\3\2\2\2\u0124?\3\2\2\2"+
		"\u0125\u0123\3\2\2\2\u0126\u012b\5B\"\2\u0127\u0128\t\5\2\2\u0128\u012a"+
		"\5B\"\2\u0129\u0127\3\2\2\2\u012a\u012d\3\2\2\2\u012b\u0129\3\2\2\2\u012b"+
		"\u012c\3\2\2\2\u012cA\3\2\2\2\u012d\u012b\3\2\2\2\u012e\u012f\7\6\2\2"+
		"\u012f\u0134\5D#\2\u0130\u0131\7\16\2\2\u0131\u0134\5D#\2\u0132\u0134"+
		"\5D#\2\u0133\u012e\3\2\2\2\u0133\u0130\3\2\2\2\u0133\u0132\3\2\2\2\u0134"+
		"C\3\2\2\2\u0135\u013a\5F$\2\u0136\u0137\7\34\2\2\u0137\u0139\7*\2\2\u0138"+
		"\u0136\3\2\2\2\u0139\u013c\3\2\2\2\u013a\u0138\3\2\2\2\u013a\u013b\3\2"+
		"\2\2\u013bE\3\2\2\2\u013c\u013a\3\2\2\2\u013d\u013e\7\21\2\2\u013e\u013f"+
		"\5\66\34\2\u013f\u0140\7\22\2\2\u0140\u014b\3\2\2\2\u0141\u0143\7*\2\2"+
		"\u0142\u0144\5\62\32\2\u0143\u0142\3\2\2\2\u0143\u0144\3\2\2\2\u0144\u014b"+
		"\3\2\2\2\u0145\u014b\7$\2\2\u0146\u0147\7\31\2\2\u0147\u014b\7*\2\2\u0148"+
		"\u014b\7&\2\2\u0149\u014b\7\'\2\2\u014a\u013d\3\2\2\2\u014a\u0141\3\2"+
		"\2\2\u014a\u0145\3\2\2\2\u014a\u0146\3\2\2\2\u014a\u0148\3\2\2\2\u014a"+
		"\u0149\3\2\2\2\u014bG\3\2\2\2\35Pblq~\u0084\u008b\u0098\u009b\u00a4\u00b0"+
		"\u00ca\u00d7\u00e1\u00ef\u00f2\u00fb\u0103\u010b\u0113\u011b\u0123\u012b"+
		"\u0133\u013a\u0143\u014a";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}