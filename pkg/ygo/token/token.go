package token

// type Token int

const (
	EOF = iota
	COMMENT

	LPAREN // (
	RPAREN // )
	LBRACK // [
	RBRACK // ]
	LBRACE // {
	RBRACE // }

	COLON // :
	COMMA // ,

	QUOTE // "
)

var tokens = [...]string{
	EOF:     "EOF",
	COMMENT: "COMMENT",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	RPAREN: ")",
	RBRACK: "]",
	RBRACE: "}",

	COLON: ":",
	COMMA: ",",
	QUOTE: "\"",
}
