package token

// TokenType トークンタイプ。INTにしたら高速になるかも
type TokenType string

// Token トークンのデータ構造
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL 未知の文字
	ILLEGAL = "ILLEGAL"
	// EOF 文末
	EOF = "EOF"

	// IDENT 識別子
	IDENT = "IDENT"
	// INT INT
	INT = "INT"

	// ASSIGN 代入
	ASSIGN = "="
	// PLUS 足し算
	PLUS = "+"

	// COMMA カンマ
	COMMA = ","
	// SEMICOLON 行末
	SEMICOLON = ";"

	// LPAREN 左括弧
	LPAREN = "("
	// RPAREN 右括弧
	RPAREN = ")"
	// LBRACE 左波括弧
	LBRACE = "{"
	// RBRACE 右波括弧
	RBRACE = "}"

	// FUNCTION 関数
	FUNCTION = "FUNCTION"
	// LET 変数
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent 適切な識別子を探す
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
