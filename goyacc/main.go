package main

import (
	"io"
	"os"
)

//go:generate goyacc -o expr.go -p "Expr" grammar.go.y

func main() {
	parser := ExprNewParser()
	lexer := NewExprLexer(os.Stdin)
	for {
		parser.Parse(lexer)
		if lexer.GetMore() == io.EOF {
			break
		}
	}
}
