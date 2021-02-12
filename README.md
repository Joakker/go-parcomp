# go-parcomp

`go-parcomp` is a project aimed at comparing different parsers and parser
generators in the Go language by implementing the same application, an
expression calculator.

Currently the list includes:

- [Goyacc](https://pkg.go.dev/golang.org/x/tools/cmd/goyacc) using a
  hand-written lexer
- [Antlr 4](https://www.antlr.org) with the `-Dlanguage=Go` option

Feel free to send a PR or Issue to add a particular parser generator :smiley:
