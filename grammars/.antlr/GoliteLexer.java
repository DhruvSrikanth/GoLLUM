// Generated from /Users/dhruvsrikanth/Work/Projects/MPCS/Quarter_5/Compilers/Project/gollum/grammars/GoliteLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GoliteLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		ASSIGN=1, AND=2, OR=3, NOT=4, LT=5, GT=6, LE=7, GE=8, EQ=9, NE=10, PLUS=11, 
		MINUS=12, MULT=13, DIV=14, LPAREN=15, RPAREN=16, LBRACE=17, RBRACE=18, 
		SEMICOLON=19, COMMA=20, VAR=21, TYPE=22, NEW=23, DELETE=24, STRUCT=25, 
		DOT=26, FUNC=27, RET=28, IF=29, ELSE=30, FOR=31, SCAN=32, PRINTF=33, INT_LIT=34, 
		BOOL_LIT=35, NIL_LIT=36, INT=37, STRING=38, BOOL=39, PTR=40, IDENT=41, 
		WS=42, COMMENT=43;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"ASSIGN", "AND", "OR", "NOT", "LT", "GT", "LE", "GE", "EQ", "NE", "PLUS", 
			"MINUS", "MULT", "DIV", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "SEMICOLON", 
			"COMMA", "VAR", "TYPE", "NEW", "DELETE", "STRUCT", "DOT", "FUNC", "RET", 
			"IF", "ELSE", "FOR", "SCAN", "PRINTF", "INT_LIT", "BOOL_LIT", "NIL_LIT", 
			"INT", "STRING", "BOOL", "PTR", "IDENT", "WS", "COMMENT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'&&'", "'||'", "'!'", "'<'", "'>'", "'<='", "'>='", "'=='", 
			"'!='", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'{'", "'}'", "';'", 
			"','", "'var'", "'type'", "'new'", "'delete'", "'struct'", "'.'", "'func'", 
			"'return'", "'if'", "'else'", "'for'", "'scan'", "'printf'", null, null, 
			"'nil'", "'int'", "'string'", "'bool'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ASSIGN", "AND", "OR", "NOT", "LT", "GT", "LE", "GE", "EQ", "NE", 
			"PLUS", "MINUS", "MULT", "DIV", "LPAREN", "RPAREN", "LBRACE", "RBRACE", 
			"SEMICOLON", "COMMA", "VAR", "TYPE", "NEW", "DELETE", "STRUCT", "DOT", 
			"FUNC", "RET", "IF", "ELSE", "FOR", "SCAN", "PRINTF", "INT_LIT", "BOOL_LIT", 
			"NIL_LIT", "INT", "STRING", "BOOL", "PTR", "IDENT", "WS", "COMMENT"
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


	public GoliteLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "GoliteLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2-\u0121\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3"+
		"\b\3\t\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3"+
		"\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3"+
		"\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\31\3"+
		"\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3"+
		"\33\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3"+
		"\36\3\36\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3"+
		"\"\3\"\3\"\3\"\3\"\3#\5#\u00ca\n#\3#\3#\7#\u00ce\n#\f#\16#\u00d1\13#\3"+
		"#\5#\u00d4\n#\3#\5#\u00d7\n#\3$\3$\3$\3$\3$\3$\3$\3$\3$\5$\u00e2\n$\3"+
		"%\3%\3%\3%\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3)\3"+
		")\3)\3*\3*\7*\u00fd\n*\f*\16*\u0100\13*\3+\6+\u0103\n+\r+\16+\u0104\3"+
		"+\3+\3,\3,\3,\3,\7,\u010d\n,\f,\16,\u0110\13,\3,\7,\u0113\n,\f,\16,\u0116"+
		"\13,\3,\7,\u0119\n,\f,\16,\u011c\13,\3,\3,\3,\3,\3\u010e\2-\3\3\5\4\7"+
		"\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22"+
		"#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C"+
		"#E$G%I&K\'M(O)Q*S+U,W-\3\2\t\3\2//\3\2\63;\3\2\62;\3\2\62\62\4\2C\\c|"+
		"\5\2\62;C\\c|\5\2\13\f\17\17\"\"\2\u012a\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3"+
		"\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2"+
		"\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35"+
		"\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)"+
		"\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2"+
		"\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2"+
		"A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3"+
		"\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\3Y\3\2\2"+
		"\2\5[\3\2\2\2\7^\3\2\2\2\ta\3\2\2\2\13c\3\2\2\2\re\3\2\2\2\17g\3\2\2\2"+
		"\21j\3\2\2\2\23m\3\2\2\2\25p\3\2\2\2\27s\3\2\2\2\31u\3\2\2\2\33w\3\2\2"+
		"\2\35y\3\2\2\2\37{\3\2\2\2!}\3\2\2\2#\177\3\2\2\2%\u0081\3\2\2\2\'\u0083"+
		"\3\2\2\2)\u0085\3\2\2\2+\u0087\3\2\2\2-\u008b\3\2\2\2/\u0090\3\2\2\2\61"+
		"\u0094\3\2\2\2\63\u009b\3\2\2\2\65\u00a2\3\2\2\2\67\u00a4\3\2\2\29\u00a9"+
		"\3\2\2\2;\u00b0\3\2\2\2=\u00b3\3\2\2\2?\u00b8\3\2\2\2A\u00bc\3\2\2\2C"+
		"\u00c1\3\2\2\2E\u00d6\3\2\2\2G\u00e1\3\2\2\2I\u00e3\3\2\2\2K\u00e7\3\2"+
		"\2\2M\u00eb\3\2\2\2O\u00f2\3\2\2\2Q\u00f7\3\2\2\2S\u00fa\3\2\2\2U\u0102"+
		"\3\2\2\2W\u0108\3\2\2\2YZ\7?\2\2Z\4\3\2\2\2[\\\7(\2\2\\]\7(\2\2]\6\3\2"+
		"\2\2^_\7~\2\2_`\7~\2\2`\b\3\2\2\2ab\7#\2\2b\n\3\2\2\2cd\7>\2\2d\f\3\2"+
		"\2\2ef\7@\2\2f\16\3\2\2\2gh\7>\2\2hi\7?\2\2i\20\3\2\2\2jk\7@\2\2kl\7?"+
		"\2\2l\22\3\2\2\2mn\7?\2\2no\7?\2\2o\24\3\2\2\2pq\7#\2\2qr\7?\2\2r\26\3"+
		"\2\2\2st\7-\2\2t\30\3\2\2\2uv\7/\2\2v\32\3\2\2\2wx\7,\2\2x\34\3\2\2\2"+
		"yz\7\61\2\2z\36\3\2\2\2{|\7*\2\2| \3\2\2\2}~\7+\2\2~\"\3\2\2\2\177\u0080"+
		"\7}\2\2\u0080$\3\2\2\2\u0081\u0082\7\177\2\2\u0082&\3\2\2\2\u0083\u0084"+
		"\7=\2\2\u0084(\3\2\2\2\u0085\u0086\7.\2\2\u0086*\3\2\2\2\u0087\u0088\7"+
		"x\2\2\u0088\u0089\7c\2\2\u0089\u008a\7t\2\2\u008a,\3\2\2\2\u008b\u008c"+
		"\7v\2\2\u008c\u008d\7{\2\2\u008d\u008e\7r\2\2\u008e\u008f\7g\2\2\u008f"+
		".\3\2\2\2\u0090\u0091\7p\2\2\u0091\u0092\7g\2\2\u0092\u0093\7y\2\2\u0093"+
		"\60\3\2\2\2\u0094\u0095\7f\2\2\u0095\u0096\7g\2\2\u0096\u0097\7n\2\2\u0097"+
		"\u0098\7g\2\2\u0098\u0099\7v\2\2\u0099\u009a\7g\2\2\u009a\62\3\2\2\2\u009b"+
		"\u009c\7u\2\2\u009c\u009d\7v\2\2\u009d\u009e\7t\2\2\u009e\u009f\7w\2\2"+
		"\u009f\u00a0\7e\2\2\u00a0\u00a1\7v\2\2\u00a1\64\3\2\2\2\u00a2\u00a3\7"+
		"\60\2\2\u00a3\66\3\2\2\2\u00a4\u00a5\7h\2\2\u00a5\u00a6\7w\2\2\u00a6\u00a7"+
		"\7p\2\2\u00a7\u00a8\7e\2\2\u00a88\3\2\2\2\u00a9\u00aa\7t\2\2\u00aa\u00ab"+
		"\7g\2\2\u00ab\u00ac\7v\2\2\u00ac\u00ad\7w\2\2\u00ad\u00ae\7t\2\2\u00ae"+
		"\u00af\7p\2\2\u00af:\3\2\2\2\u00b0\u00b1\7k\2\2\u00b1\u00b2\7h\2\2\u00b2"+
		"<\3\2\2\2\u00b3\u00b4\7g\2\2\u00b4\u00b5\7n\2\2\u00b5\u00b6\7u\2\2\u00b6"+
		"\u00b7\7g\2\2\u00b7>\3\2\2\2\u00b8\u00b9\7h\2\2\u00b9\u00ba\7q\2\2\u00ba"+
		"\u00bb\7t\2\2\u00bb@\3\2\2\2\u00bc\u00bd\7u\2\2\u00bd\u00be\7e\2\2\u00be"+
		"\u00bf\7c\2\2\u00bf\u00c0\7p\2\2\u00c0B\3\2\2\2\u00c1\u00c2\7r\2\2\u00c2"+
		"\u00c3\7t\2\2\u00c3\u00c4\7k\2\2\u00c4\u00c5\7p\2\2\u00c5\u00c6\7v\2\2"+
		"\u00c6\u00c7\7h\2\2\u00c7D\3\2\2\2\u00c8\u00ca\t\2\2\2\u00c9\u00c8\3\2"+
		"\2\2\u00c9\u00ca\3\2\2\2\u00ca\u00cb\3\2\2\2\u00cb\u00cf\t\3\2\2\u00cc"+
		"\u00ce\t\4\2\2\u00cd\u00cc\3\2\2\2\u00ce\u00d1\3\2\2\2\u00cf\u00cd\3\2"+
		"\2\2\u00cf\u00d0\3\2\2\2\u00d0\u00d7\3\2\2\2\u00d1\u00cf\3\2\2\2\u00d2"+
		"\u00d4\t\2\2\2\u00d3\u00d2\3\2\2\2\u00d3\u00d4\3\2\2\2\u00d4\u00d5\3\2"+
		"\2\2\u00d5\u00d7\t\5\2\2\u00d6\u00c9\3\2\2\2\u00d6\u00d3\3\2\2\2\u00d7"+
		"F\3\2\2\2\u00d8\u00d9\7v\2\2\u00d9\u00da\7t\2\2\u00da\u00db\7w\2\2\u00db"+
		"\u00e2\7g\2\2\u00dc\u00dd\7h\2\2\u00dd\u00de\7c\2\2\u00de\u00df\7n\2\2"+
		"\u00df\u00e0\7u\2\2\u00e0\u00e2\7g\2\2\u00e1\u00d8\3\2\2\2\u00e1\u00dc"+
		"\3\2\2\2\u00e2H\3\2\2\2\u00e3\u00e4\7p\2\2\u00e4\u00e5\7k\2\2\u00e5\u00e6"+
		"\7n\2\2\u00e6J\3\2\2\2\u00e7\u00e8\7k\2\2\u00e8\u00e9\7p\2\2\u00e9\u00ea"+
		"\7v\2\2\u00eaL\3\2\2\2\u00eb\u00ec\7u\2\2\u00ec\u00ed\7v\2\2\u00ed\u00ee"+
		"\7t\2\2\u00ee\u00ef\7k\2\2\u00ef\u00f0\7p\2\2\u00f0\u00f1\7i\2\2\u00f1"+
		"N\3\2\2\2\u00f2\u00f3\7d\2\2\u00f3\u00f4\7q\2\2\u00f4\u00f5\7q\2\2\u00f5"+
		"\u00f6\7n\2\2\u00f6P\3\2\2\2\u00f7\u00f8\7,\2\2\u00f8\u00f9\5S*\2\u00f9"+
		"R\3\2\2\2\u00fa\u00fe\t\6\2\2\u00fb\u00fd\t\7\2\2\u00fc\u00fb\3\2\2\2"+
		"\u00fd\u0100\3\2\2\2\u00fe\u00fc\3\2\2\2\u00fe\u00ff\3\2\2\2\u00ffT\3"+
		"\2\2\2\u0100\u00fe\3\2\2\2\u0101\u0103\t\b\2\2\u0102\u0101\3\2\2\2\u0103"+
		"\u0104\3\2\2\2\u0104\u0102\3\2\2\2\u0104\u0105\3\2\2\2\u0105\u0106\3\2"+
		"\2\2\u0106\u0107\b+\2\2\u0107V\3\2\2\2\u0108\u0109\7\61\2\2\u0109\u010a"+
		"\7\61\2\2\u010a\u010e\3\2\2\2\u010b\u010d\13\2\2\2\u010c\u010b\3\2\2\2"+
		"\u010d\u0110\3\2\2\2\u010e\u010f\3\2\2\2\u010e\u010c\3\2\2\2\u010f\u0114"+
		"\3\2\2\2\u0110\u010e\3\2\2\2\u0111\u0113\7\13\2\2\u0112\u0111\3\2\2\2"+
		"\u0113\u0116\3\2\2\2\u0114\u0112\3\2\2\2\u0114\u0115\3\2\2\2\u0115\u011a"+
		"\3\2\2\2\u0116\u0114\3\2\2\2\u0117\u0119\7\17\2\2\u0118\u0117\3\2\2\2"+
		"\u0119\u011c\3\2\2\2\u011a\u0118\3\2\2\2\u011a\u011b\3\2\2\2\u011b\u011d"+
		"\3\2\2\2\u011c\u011a\3\2\2\2\u011d\u011e\7\f\2\2\u011e\u011f\3\2\2\2\u011f"+
		"\u0120\b,\2\2\u0120X\3\2\2\2\r\2\u00c9\u00cf\u00d3\u00d6\u00e1\u00fe\u0104"+
		"\u010e\u0114\u011a\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}