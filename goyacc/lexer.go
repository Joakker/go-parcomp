package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// MyExprLexer reads from an io.Reader and tokenizes the input from it
type MyExprLexer struct {
	buf   *bufio.Reader
	input []rune
	index int
}

// NewExprLexer creates and initializes a MyExprLexer from buf
func NewExprLexer(buf io.Reader) *MyExprLexer {
	l := &MyExprLexer{buf: bufio.NewReader(buf)}
	l.GetMore()
	return l
}

// GetMore prints a prompt and reads from the input to get a line of text
func (l *MyExprLexer) GetMore() error {
	fmt.Print("> ")
	line, err := l.buf.ReadString('\n')
	if err == io.EOF {
		return err
	}
	l.input = []rune(line[:len(line)-1]) // We get rid of the '\n' at the end
	l.index = 0
	return nil
}

// Lex is called by ExprParserImpl to get the next token, initializing yylval
// accordingly
func (l *MyExprLexer) Lex(yylval *ExprSymType) int {
	if l.index >= len(l.input) {
		return -1
	}
	l.SkipWhitespace()
	switch c := l.input[l.index]; c {
	case '+', '-', '*', '/', '(', ')', '=':
		l.index++
		return int(c)
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return l.Number(yylval)
	default:
		return l.Identifier(yylval)
	}
}

// SkipWhitespace skips characters from the buffered line while they are
// whitespace.
func (l *MyExprLexer) SkipWhitespace() {
	for {
		if l.index >= len(l.input) {
			return
		}
		switch l.input[l.index] {
		case ' ', '\t', '\r', '\n':
			l.index++
		default:
			return
		}
	}
}

// Number reads a number in decimal representation from the buffered line,
// storing it's value into yylval.String.
func (l *MyExprLexer) Number(yylval *ExprSymType) int {
	buf := strings.Builder{}
	didDot := false
forloop:
	for {
		if l.index >= len(l.input) {
			break forloop
		}
		switch c := l.input[l.index]; c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.index++
			buf.WriteRune(c)
		case '.':
			if !didDot {
				l.index++
				didDot = true
				buf.WriteRune(c)
				continue
			}
			fallthrough
		default:
			break forloop
		}
	}
	yylval.String = buf.String()
	return NUMBER
}

// Identifier reads a name of the form [a-zA-Z][a-zA-Z0-9]*, storing it's value
// into yylval.String.
func (l *MyExprLexer) Identifier(yylval *ExprSymType) int {
	buf := strings.Builder{}
	buf.WriteRune(l.input[l.index])
	l.index++
forloop:
	for {
		c := l.input[l.index]
		switch {
		case 'a' <= c && c <= 'z':
		case 'A' <= c && c <= 'Z':
		case '0' <= c && c <= '9':
			buf.WriteRune(c)
		default:
			break forloop
		}
	}
	yylval.String = buf.String()
	return IDENTIFIER
}

// Error logs the error s to os.Stderr.
func (l *MyExprLexer) Error(s string) {
	fmt.Fprintln(os.Stderr, "Expr parsing error:", s)
}
