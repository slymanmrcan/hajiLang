package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int // Satır numarası
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Tanımlayıcılar
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	FLOAT  = "FLOAT"

	// Operatörler
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	PERCENT  = "%"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="
	LT_EQ  = "<="
	GT_EQ  = ">="
	AND    = "&&"
	OR     = "||"
	COLON  = ":"

	// Ayıraçlar
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	LBRACKET = "[" //
	RBRACKET = "]" //

	// Keywordler
	FUNCTION = "FUNCTION"
	HAJI     = "HAJI"
	KATI     = "KATI"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FOR      = "FOR"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"haji":   HAJI,
	"kati":   KATI,
	"for":    FOR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
