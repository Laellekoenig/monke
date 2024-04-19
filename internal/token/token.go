package token

type TokenType string

type Token struct {
	Type TokenType
	// for better performance use int or byte
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers and literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// operators
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	BANG    = "!"
	ASTERIX = "*"
	SLASH   = "/"
	LT      = "<"
	GT      = ">"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

	// conditional
	IF   = "if"
	ELSE = "else"

	//boolean
	TRUE  = "true"
	FALSE = "false"
	EQ    = "=="
	NEQ   = "!="

	RETURN = "return"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupKeyword(s *string) TokenType {
	if token, ok := keywords[*s]; ok {
		return token
	}

	return IDENTIFIER
}
