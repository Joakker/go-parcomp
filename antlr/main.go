package main

//go:generate java org.antlr.v4.Tool Expr.g4 -Dlanguage=Go -o parser

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Joakker/go-parcomp/antlr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	it := &Interpreter{Variable: map[string]float64{}}
	for {
		fmt.Print("> ")
		line, err := in.ReadString('\n')
		if err == io.EOF {
			break
		} else if len(line) == 1 {
			continue
		}
		is := antlr.NewInputStream(line)
		lx := parser.NewExprLexer(is)
		ts := antlr.NewCommonTokenStream(lx, antlr.TokenDefaultChannel)
		pr := parser.NewExprParser(ts)
		antlr.ParseTreeWalkerDefault.Walk(it, pr.Stmt())
	}
}
