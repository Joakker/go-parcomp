package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type MyExprLexer struct {
	buf   *bufio.Reader
	input string
	index int
}

func NewExprLexer(buf io.Reader) *MyExprLexer {
	l := &MyExprLexer{buf: bufio.NewReader(buf)}
	l.GetMore()
	return l
}

func (l *MyExprLexer) GetMore() error {
	fmt.Print("> ")
	line, err := l.buf.ReadString('\n')
	if err == io.EOF {
		return err
	}
	l.input = line[:len(line)-1] // We get rid of the '\n' at the end
	l.index = 0
	return nil
}

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
			buf.WriteByte(c)
		case '.':
			if !didDot {
				l.index++
				didDot = true
				buf.WriteByte(c)
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

func (l *MyExprLexer) Identifier(yylval *ExprSymType) int {
	buf := strings.Builder{}
	buf.WriteByte(l.input[l.index])
	l.index++
forloop:
	for {
		c := l.input[l.index]
		switch {
		case 'a' <= c && c <= 'z':
		case 'A' <= c && c <= 'Z':
		case '0' <= c && c <= '9':
			buf.WriteByte(c)
		default:
			break forloop
		}
	}
	yylval.String = buf.String()
	return IDENTIFIER
}

func (l *MyExprLexer) Error(s string) {
	fmt.Fprintln(os.Stderr, "Expr parsing error:", s)
}
