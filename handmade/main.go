package main

func main() {
	l := NewLexer("a = (1+2.2)")
	p := NewParser(l)
	p.Stmt()
}
