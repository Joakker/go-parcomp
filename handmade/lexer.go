package main

import (
	"strings"
	"unicode"
)

// Lexer scans the input string given to NewLexer or SetInput for Tokens
type Lexer struct {
	input *strings.Reader
	cur   rune
	idx   uint64
	buf   strings.Builder
}

// NewLexer creates and initializes a new instance of Lexer reading from input
func NewLexer(input string) (l *Lexer) {
	l = &Lexer{
		input: strings.NewReader(input),
		buf:   strings.Builder{},
	}
	l.Advance()
	return
}

// SetInput resets l to start reading from input
func (l *Lexer) SetInput(input string) {
	l.input.Reset(input)
	l.Advance()
	l.buf.Reset()
}

// Next returns the next Token it can detect from the input string. If it has
// reached the end of the string, it will emit a Token with type TokEof
func (l *Lexer) Next() Token {
	for ; l.cur != 0; l.Advance() {
		if unicode.IsSpace(l.cur) {
			continue
		}
		// If the index didn't advance, i.e., we stumbled upon a 1 character
		// long token, then we have to consume it
		defer func(idx uint64) {
			if l.idx == idx {
				l.Advance()
			}
		}(l.idx)
		switch l.cur {
		case '+':
			return Token{Type: TokAdd}
		case '-':
			return Token{Type: TokSub}
		case '*':
			return Token{Type: TokMul}
		case '/':
			return Token{Type: TokDiv}
		case '=':
			return Token{Type: TokEqu}
		case '(':
			return Token{Type: TokLPr}
		case ')':
			return Token{Type: TokRPr}
		default:
			if unicode.IsLetter(l.cur) {
				return l.Name()
			} else if unicode.IsDigit(l.cur) {
				return l.Num()
			} else {
				panic("No viable alternative!")
			}
		}
	}
	return Token{Type: TokEof}
}

// Name parses an identifier
func (l *Lexer) Name() Token {
	defer l.buf.Reset()
	for {
		l.buf.WriteRune(l.cur)
		if cur := l.Advance(); !(unicode.IsDigit(cur) || unicode.IsLetter(cur)) {
			break
		}
	}
	return Token{Type: TokIdn, Text: l.buf.String()}
}

// Num parses a number
func (l *Lexer) Num() (t Token) {
	defer l.buf.Reset()
	for {
		l.buf.WriteRune(l.cur)
		if cur := l.Advance(); !unicode.IsDigit(cur) {
			break
		}
	}
	if l.cur == '.' {
		for {
			l.buf.WriteRune(l.cur)
			if cur := l.Advance(); !unicode.IsDigit(cur) {
				break
			}
		}
	}
	return Token{Type: TokNum, Text: l.buf.String()}
}

func (l *Lexer) Advance() rune {
	cur, _, err := l.input.ReadRune()
	if err != nil {
		cur = 0
	}
	l.cur = cur
	l.idx++
	return l.cur
}
