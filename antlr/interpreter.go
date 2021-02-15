package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/Joakker/go-parcomp/antlr/parser"
)

// Interpreter listens to events in the parse tree and carries the logic for
// defining variables and calculating expressions
type Interpreter struct {
	parser.BaseExprListener
	Variable map[string]float64 // Contains all the named variables and their values
	Stack    []float64          // Contains values while an expression is being calculated
}

// EnterNumber calculates the float64 value of the number and pushes
// it onto the Stack
func (i *Interpreter) EnterNumber(ctx *parser.NumberContext) {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		panic(err)
	}
	i.Push(val)
}

// EnterIdentifier retrieves the value associated with the identifier, or
// math.NaN() if it doesn't exist, and pushes it onto the Stack
func (i *Interpreter) EnterIdentifier(ctx *parser.IdentifierContext) {
	val, ok := i.Variable[ctx.GetText()]
	if !ok {
		panic("No such variable")
	}
	i.Push(val)
}

// ExitMulDiv groups the operands '*' and '/' together, applying the
// corresponding operation to the 2 topmost values on the Stack. Since we have
// to have visited both of them before doing this action, it uses `Exit' instead
// of `Enter'.
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

// ExitAddSub groups the operands '+' and '-' together, applying the
// corresponding operation to the 2 topmost values on the Stack. Since we have
// to have visited both of them before doing this action, it uses `Exit' instead
// of `Enter'.
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

// ExitAssignStmt defines all the names in the left hand side and assigns the
// expressions in the right to them in order. The length of both lists must be
// the same, otherwise no variable will be assigned
func (i *Interpreter) ExitAssignStmt(ctx *parser.AssignStmtContext) {
	name := ctx.Namelist().(*parser.NamelistContext).AllID()
	if len(name) != len(i.Stack) {
		fmt.Fprintf(os.Stderr, "Lengths mismatch")
		return
	}
	for j := len(name) - 1; j >= 0; j-- {
		i.Variable[name[j].GetText()] = i.Pop()
	}
}

// ExitExprStmt prints the resulting value of a statement when it consists of an
// expression
func (i *Interpreter) ExitExprStmt(ctx *parser.ExprStmtContext) {
	fmt.Println(i.Pop())
}

// Push appends a value val to the Stack
func (i *Interpreter) Push(val float64) {
	i.Stack = append(i.Stack, val)
}

// Pop deletes the topmost value off the Stack and returns it
func (i *Interpreter) Pop() float64 {
	val := i.Stack[len(i.Stack)-1]
	i.Stack = i.Stack[:len(i.Stack)-1]
	return val
}
