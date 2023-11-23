package token

type TokenType string

type Token struct {
	Type       TokenType
	Literal    string
	LineNumber int
	FileName   string
}

const (
	ILLEGAL      = "ILLEGAL"
	EOF          = "EOF"
	ID           = "ID"
	INT          = "INT"
	ASSIGN       = "="
	PLUS         = "+"
	COMMA        = ","
	SEMICOLON    = ";"
	LPARENTHESES = "("
	RPARENTHESES = ")"
	LBRACE       = "{"
	RBRACE       = "}"
	FUNCTION     = "FUNCTION"
	DEF          = "DEF"
	RETURN       = "RETURN"
)

var keywords map[string]TokenType = map[string]TokenType{
	"fun":    FUNCTION,
	"def":    DEF,
	"return": RETURN,
}

func LookupIdentifier(id string) TokenType {
	if token, ok := keywords[id]; ok {
		return token
	}
	return ID
}
