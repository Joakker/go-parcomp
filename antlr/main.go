package main

//go:generate java org.antlr.v4.Tool Expr.g4 -Dlanguage=Go -o parser

import (
	"github.com/Joakker/go-parcomp/antlr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	is := antlr.NewInputStream("(1 + 2)")
	lx := parser.NewExprLexer(is)
	ts := antlr.NewCommonTokenStream(lx, antlr.TokenDefaultChannel)
	pr := parser.NewExprParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(&Interpreter{
		Stack:    []float64{},
		Variable: map[string]float64{},
	}, pr.File())
}
