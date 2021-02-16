package main

import "fmt"

type Parser struct {
	input    *Lexer
	lah, cur Token
}

func NewParser(input *Lexer) (p *Parser) {
	p = &Parser{input: input}
	p.Advance()
	p.Advance()
	return
}

func (p *Parser) Advance() {
	p.lah, p.cur = p.input.Next(), p.lah
}

// Stmt parses statements of the form:
//
//		stmt := IDN '=' expr
//		stmt := expr
func (p *Parser) Stmt() {
	switch p.cur.Type {
	case TokIdn:
		if p.lah.Type == TokEqu {
			fmt.Println("Found assignment")
			break
		}
		fallthrough
	default:
		fmt.Println("It's an expression")
	}
}

// Expr parses statements of the form:
//
//		expr := '(' expr ')'
//		expr := expr ('*'|'/'|'+'|'-') expr
//		expr := IDN
//		expr := NUM
func (p *Parser) Expr() {
	// TODO: Implement this
}
