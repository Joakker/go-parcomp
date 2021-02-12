package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Joakker/go-parcomp/antlr/parser"
)

type Interpreter struct {
	parser.BaseExprListener
	Variable map[string]float64
	Stack    []float64
}

func (i *Interpreter) EnterNumber(ctx *parser.NumberContext) {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		panic(err)
	}
	i.Push(val)
}

func (i *Interpreter) EnterIdentifier(ctx *parser.IdentifierContext) {
	val, ok := i.Variable[ctx.GetText()]
	if !ok {
		panic("No such variable")
	}
	i.Push(val)
	fmt.Println(i.Stack)
}

func (i *Interpreter) ExitMulDiv(ctx *parser.MulDivContext) {
	b := i.Pop()
	a := i.Pop()
	switch ctx.GetOp().GetTokenType() {
	case parser.ExprLexerMUL:
		i.Push(a * b)
	case parser.ExprLexerDIV:
		if b == 0.0 {
			i.Push(math.NaN())
		} else {
			i.Push(a / b)
		}
	}
}

func (i *Interpreter) ExitAddSub(ctx *parser.AddSubContext) {
	b := i.Pop()
	a := i.Pop()

	switch ctx.GetOp().GetTokenType() {
	case parser.ExprLexerADD:
		i.Push(a + b)
	case parser.ExprLexerSUB:
		i.Push(a - b)
	}
}

func (i *Interpreter) ExitExprStmt(ctx *parser.ExprStmtContext) {
	fmt.Println(i.Pop())
}

func (i *Interpreter) Push(val float64) {
	i.Stack = append(i.Stack, val)
}

func (i *Interpreter) Pop() float64 {
	val := i.Stack[len(i.Stack)-1]
	i.Stack = i.Stack[:len(i.Stack)-1]
	return val
}
