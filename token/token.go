package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Tokens we don't recognize
	EOF     = "EOF"     // End of File

	// Identifiers + Literals...
	IDENT = "IDENT" // add, foobar, x, y,... (function names + variables)
	INT   = "INT"

	// Operators
	ASSIGN = ""
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
