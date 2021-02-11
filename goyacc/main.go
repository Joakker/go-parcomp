package main

import (
	"io"
	"os"
)

//go:generate goyacc -o grammar.go -p "Expr" grammar.go.y

func main() {
	ExprErrorVerbose = true
	parser := ExprNewParser()
	lexer := NewExprLexer(os.Stdin)
	for {
		parser.Parse(lexer)
		if lexer.GetMore() == io.EOF {
			break
		}
	}
}
