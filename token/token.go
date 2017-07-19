package token

type TokenType string
type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 1343456
	// Operators
	ASSIGN = "ASSIGN"
	PLUS = "+"
	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

var symbols = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	":=": ASSIGN,
}
func LookupIdent(ident string) TokenType {
	if tok, ok := symbols[ident]; ok {
		return tok
	}
	return IDENT
}