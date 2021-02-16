package main

import "fmt"

type TokenType int

// Token types
const (
	TokEof TokenType = iota // EOF
	TokIdn                  // a, k2, €
	TokNum                  // 1, 2.2
	TokAdd                  // '+'
	TokSub                  // '-'
	TokMul                  // '*'
	TokDiv                  // '/'
	TokPow                  // '^'
	TokEqu                  // '='
	TokLPr                  // '('
	TokRPr                  // ')'
)

// Token represents a word in the input
type Token struct {
	Type TokenType
	Text string
}

var tokNames = [...]string{"EOF", "IDN", "NUM", "+", "-", "*", "/", "^", "=", "(", ")"}

// String returns the string representation of t
func (t Token) String() string {
	if len(t.Text) != 0 {
		return fmt.Sprintf("< %s, %s >", t.Type, t.Text)
	}
	return fmt.Sprintf("< %s >", t.Type)
}

// String returns the name of string
func (t TokenType) String() string { return tokNames[t] }
