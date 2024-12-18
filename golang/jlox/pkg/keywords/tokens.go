package keywords

type Token struct {
	TokenType string // should probably be an enum
	Lexeme    string
	Literal   string
	Line      int
}

const LEFT_PAREN = ""
const RIGHT_PAREN = ""
const LEFT_BRACE = ""
const RIGHT_BRACE = ""
const COMMA = ""
const DOT = ""
const MINUS = ""
const PLUS = ""
const SEMICOLON = ""
const SLASH = ""
const STAR = ""

// One or two character tokens.
const BANG = ""
const BANG_EQUAL = ""
const EQUAL = ""
const EQUAL_EQUAL = ""
const GREATER = ""
const GREATER_EQUAL = ""
const LESS = ""
const LESS_EQUAL = ""

// Literals.
const IDENTIFIER = ""
const STRING = ""
const NUMBER = ""

// Keywords.
const AND = ""
const CLASS = ""
const ELSE = ""
const FALSE = ""
const FUN = ""
const FOR = ""
const IF = ""
const NIL = ""
const OR = ""
const PRINT = ""
const RETURN = ""
const SUPER = ""
const THIS = ""
const TRUE = ""
const VAR = ""
const WHILE = ""
const EOF = ""
