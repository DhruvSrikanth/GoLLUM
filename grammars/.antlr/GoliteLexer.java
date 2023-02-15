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
		STRING_LIT=35, BOOL_LIT=36, NIL_LIT=37, INT=38, BOOL=39, IDENT=40, WS=41, 
		COMMENT=42;
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
			"IF", "ELSE", "FOR", "SCAN", "PRINTF", "INT_LIT", "STRING_LIT", "BOOL_LIT", 
			"NIL_LIT", "INT", "BOOL", "IDENT", "WS", "COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2,\u010c\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\3"+
		"\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\t"+
		"\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17"+
		"\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26"+
		"\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\34"+
		"\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36"+
		"\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3"+
		"\"\3\"\3\"\3#\3#\7#\u00c9\n#\f#\16#\u00cc\13#\3#\5#\u00cf\n#\3$\3$\7$"+
		"\u00d3\n$\f$\16$\u00d6\13$\3$\3$\3%\3%\3%\3%\3%\3%\3%\3%\3%\5%\u00e3\n"+
		"%\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3)\3)\7)\u00f4\n)\f)\16)"+
		"\u00f7\13)\3*\6*\u00fa\n*\r*\16*\u00fb\3*\3*\3+\3+\3+\3+\7+\u0104\n+\f"+
		"+\16+\u0107\13+\3+\3+\3+\3+\4\u00d4\u0105\2,\3\3\5\4\7\5\t\6\13\7\r\b"+
		"\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26"+
		"+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S"+
		"+U,\3\2\b\3\2\63;\3\2\62;\3\2\62\62\4\2C\\c|\5\2\62;C\\c|\5\2\13\f\17"+
		"\17\"\"\2\u0112\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3"+
		"\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2"+
		"\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3"+
		"\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2"+
		"\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\2"+
		"9\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3"+
		"\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2"+
		"\2\2S\3\2\2\2\2U\3\2\2\2\3W\3\2\2\2\5Y\3\2\2\2\7\\\3\2\2\2\t_\3\2\2\2"+
		"\13a\3\2\2\2\rc\3\2\2\2\17e\3\2\2\2\21h\3\2\2\2\23k\3\2\2\2\25n\3\2\2"+
		"\2\27q\3\2\2\2\31s\3\2\2\2\33u\3\2\2\2\35w\3\2\2\2\37y\3\2\2\2!{\3\2\2"+
		"\2#}\3\2\2\2%\177\3\2\2\2\'\u0081\3\2\2\2)\u0083\3\2\2\2+\u0085\3\2\2"+
		"\2-\u0089\3\2\2\2/\u008e\3\2\2\2\61\u0092\3\2\2\2\63\u0099\3\2\2\2\65"+
		"\u00a0\3\2\2\2\67\u00a2\3\2\2\29\u00a7\3\2\2\2;\u00ae\3\2\2\2=\u00b1\3"+
		"\2\2\2?\u00b6\3\2\2\2A\u00ba\3\2\2\2C\u00bf\3\2\2\2E\u00ce\3\2\2\2G\u00d0"+
		"\3\2\2\2I\u00e2\3\2\2\2K\u00e4\3\2\2\2M\u00e8\3\2\2\2O\u00ec\3\2\2\2Q"+
		"\u00f1\3\2\2\2S\u00f9\3\2\2\2U\u00ff\3\2\2\2WX\7?\2\2X\4\3\2\2\2YZ\7("+
		"\2\2Z[\7(\2\2[\6\3\2\2\2\\]\7~\2\2]^\7~\2\2^\b\3\2\2\2_`\7#\2\2`\n\3\2"+
		"\2\2ab\7>\2\2b\f\3\2\2\2cd\7@\2\2d\16\3\2\2\2ef\7>\2\2fg\7?\2\2g\20\3"+
		"\2\2\2hi\7@\2\2ij\7?\2\2j\22\3\2\2\2kl\7?\2\2lm\7?\2\2m\24\3\2\2\2no\7"+
		"#\2\2op\7?\2\2p\26\3\2\2\2qr\7-\2\2r\30\3\2\2\2st\7/\2\2t\32\3\2\2\2u"+
		"v\7,\2\2v\34\3\2\2\2wx\7\61\2\2x\36\3\2\2\2yz\7*\2\2z \3\2\2\2{|\7+\2"+
		"\2|\"\3\2\2\2}~\7}\2\2~$\3\2\2\2\177\u0080\7\177\2\2\u0080&\3\2\2\2\u0081"+
		"\u0082\7=\2\2\u0082(\3\2\2\2\u0083\u0084\7.\2\2\u0084*\3\2\2\2\u0085\u0086"+
		"\7x\2\2\u0086\u0087\7c\2\2\u0087\u0088\7t\2\2\u0088,\3\2\2\2\u0089\u008a"+
		"\7v\2\2\u008a\u008b\7{\2\2\u008b\u008c\7r\2\2\u008c\u008d\7g\2\2\u008d"+
		".\3\2\2\2\u008e\u008f\7p\2\2\u008f\u0090\7g\2\2\u0090\u0091\7y\2\2\u0091"+
		"\60\3\2\2\2\u0092\u0093\7f\2\2\u0093\u0094\7g\2\2\u0094\u0095\7n\2\2\u0095"+
		"\u0096\7g\2\2\u0096\u0097\7v\2\2\u0097\u0098\7g\2\2\u0098\62\3\2\2\2\u0099"+
		"\u009a\7u\2\2\u009a\u009b\7v\2\2\u009b\u009c\7t\2\2\u009c\u009d\7w\2\2"+
		"\u009d\u009e\7e\2\2\u009e\u009f\7v\2\2\u009f\64\3\2\2\2\u00a0\u00a1\7"+
		"\60\2\2\u00a1\66\3\2\2\2\u00a2\u00a3\7h\2\2\u00a3\u00a4\7w\2\2\u00a4\u00a5"+
		"\7p\2\2\u00a5\u00a6\7e\2\2\u00a68\3\2\2\2\u00a7\u00a8\7t\2\2\u00a8\u00a9"+
		"\7g\2\2\u00a9\u00aa\7v\2\2\u00aa\u00ab\7w\2\2\u00ab\u00ac\7t\2\2\u00ac"+
		"\u00ad\7p\2\2\u00ad:\3\2\2\2\u00ae\u00af\7k\2\2\u00af\u00b0\7h\2\2\u00b0"+
		"<\3\2\2\2\u00b1\u00b2\7g\2\2\u00b2\u00b3\7n\2\2\u00b3\u00b4\7u\2\2\u00b4"+
		"\u00b5\7g\2\2\u00b5>\3\2\2\2\u00b6\u00b7\7h\2\2\u00b7\u00b8\7q\2\2\u00b8"+
		"\u00b9\7t\2\2\u00b9@\3\2\2\2\u00ba\u00bb\7u\2\2\u00bb\u00bc\7e\2\2\u00bc"+
		"\u00bd\7c\2\2\u00bd\u00be\7p\2\2\u00beB\3\2\2\2\u00bf\u00c0\7r\2\2\u00c0"+
		"\u00c1\7t\2\2\u00c1\u00c2\7k\2\2\u00c2\u00c3\7p\2\2\u00c3\u00c4\7v\2\2"+
		"\u00c4\u00c5\7h\2\2\u00c5D\3\2\2\2\u00c6\u00ca\t\2\2\2\u00c7\u00c9\t\3"+
		"\2\2\u00c8\u00c7\3\2\2\2\u00c9\u00cc\3\2\2\2\u00ca\u00c8\3\2\2\2\u00ca"+
		"\u00cb\3\2\2\2\u00cb\u00cf\3\2\2\2\u00cc\u00ca\3\2\2\2\u00cd\u00cf\t\4"+
		"\2\2\u00ce\u00c6\3\2\2\2\u00ce\u00cd\3\2\2\2\u00cfF\3\2\2\2\u00d0\u00d4"+
		"\7$\2\2\u00d1\u00d3\13\2\2\2\u00d2\u00d1\3\2\2\2\u00d3\u00d6\3\2\2\2\u00d4"+
		"\u00d5\3\2\2\2\u00d4\u00d2\3\2\2\2\u00d5\u00d7\3\2\2\2\u00d6\u00d4\3\2"+
		"\2\2\u00d7\u00d8\7$\2\2\u00d8H\3\2\2\2\u00d9\u00da\7v\2\2\u00da\u00db"+
		"\7t\2\2\u00db\u00dc\7w\2\2\u00dc\u00e3\7g\2\2\u00dd\u00de\7h\2\2\u00de"+
		"\u00df\7c\2\2\u00df\u00e0\7n\2\2\u00e0\u00e1\7u\2\2\u00e1\u00e3\7g\2\2"+
		"\u00e2\u00d9\3\2\2\2\u00e2\u00dd\3\2\2\2\u00e3J\3\2\2\2\u00e4\u00e5\7"+
		"p\2\2\u00e5\u00e6\7k\2\2\u00e6\u00e7\7n\2\2\u00e7L\3\2\2\2\u00e8\u00e9"+
		"\7k\2\2\u00e9\u00ea\7p\2\2\u00ea\u00eb\7v\2\2\u00ebN\3\2\2\2\u00ec\u00ed"+
		"\7d\2\2\u00ed\u00ee\7q\2\2\u00ee\u00ef\7q\2\2\u00ef\u00f0\7n\2\2\u00f0"+
		"P\3\2\2\2\u00f1\u00f5\t\5\2\2\u00f2\u00f4\t\6\2\2\u00f3\u00f2\3\2\2\2"+
		"\u00f4\u00f7\3\2\2\2\u00f5\u00f3\3\2\2\2\u00f5\u00f6\3\2\2\2\u00f6R\3"+
		"\2\2\2\u00f7\u00f5\3\2\2\2\u00f8\u00fa\t\7\2\2\u00f9\u00f8\3\2\2\2\u00fa"+
		"\u00fb\3\2\2\2\u00fb\u00f9\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fc\u00fd\3\2"+
		"\2\2\u00fd\u00fe\b*\2\2\u00feT\3\2\2\2\u00ff\u0100\7\61\2\2\u0100\u0101"+
		"\7\61\2\2\u0101\u0105\3\2\2\2\u0102\u0104\13\2\2\2\u0103\u0102\3\2\2\2"+
		"\u0104\u0107\3\2\2\2\u0105\u0106\3\2\2\2\u0105\u0103\3\2\2\2\u0106\u0108"+
		"\3\2\2\2\u0107\u0105\3\2\2\2\u0108\u0109\7\f\2\2\u0109\u010a\3\2\2\2\u010a"+
		"\u010b\b+\2\2\u010bV\3\2\2\2\n\2\u00ca\u00ce\u00d4\u00e2\u00f5\u00fb\u0105"+
		"\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}