// Code generated by goyacc -o expr.go -p Expr grammar.go.y. DO NOT EDIT.

//line grammar.go.y:2
package main

import __yyfmt__ "fmt"

//line grammar.go.y:2

import (
	"fmt"
	"strconv"
)

//line grammar.go.y:10
type ExprSymType struct {
	yys    int
	String string
	Number float64
}

const NUMBER = 57346
const IDENTIFIER = 57347

var ExprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'('",
	"')'",
	"NUMBER",
	"IDENTIFIER",
	"'='",
}

var ExprStatenames = [...]string{}

const ExprEofCode = 1
const ExprErrCode = 2
const ExprInitialStackSize = 16

//line grammar.go.y:43

//line yacctab:1
var ExprExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const ExprPrivate = 57344

const ExprLast = 31

var ExprAct = [...]int{
	2, 8, 9, 10, 11, 12, 21, 13, 15, 16,
	17, 18, 19, 20, 7, 3, 1, 6, 7, 4,
	14, 6, 0, 4, 5, 8, 9, 10, 11, 10,
	11,
}

var ExprPact = [...]int{
	13, -1000, 21, -1000, -1000, -7, 9, 9, 9, 9,
	9, 9, 9, -3, -1000, -1000, 23, 23, -1000, -1000,
	21, -1000,
}

var ExprPgo = [...]int{
	0, 0, 16, 15,
}

var ExprR1 = [...]int{
	0, 2, 2, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 3,
}

var ExprR2 = [...]int{
	0, 0, 1, 1, 1, 1, 3, 3, 3, 3,
	3, 2, 3,
}

var ExprChk = [...]int{
	-1000, -2, -1, -3, 10, 11, 8, 5, 4, 5,
	6, 7, 12, -1, 11, -1, -1, -1, -1, -1,
	-1, 9,
}

var ExprDef = [...]int{
	1, -2, 2, 3, 4, 5, 0, 0, 0, 0,
	0, 0, 0, 0, 5, 11, 6, 7, 8, 9,
	12, 10,
}

var ExprTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	8, 9, 6, 4, 3, 5, 3, 7, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 12,
}

var ExprTok2 = [...]int{
	2, 3, 10, 11,
}

var ExprTok3 = [...]int{
	0,
}

var ExprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	ExprDebug        = 0
	ExprErrorVerbose = false
)

type ExprLexer interface {
	Lex(lval *ExprSymType) int
	Error(s string)
}

type ExprParser interface {
	Parse(ExprLexer) int
	Lookahead() int
}

type ExprParserImpl struct {
	lval  ExprSymType
	stack [ExprInitialStackSize]ExprSymType
	char  int
}

func (p *ExprParserImpl) Lookahead() int {
	return p.char
}

func ExprNewParser() ExprParser {
	return &ExprParserImpl{}
}

const ExprFlag = -1000

func ExprTokname(c int) string {
	if c >= 1 && c-1 < len(ExprToknames) {
		if ExprToknames[c-1] != "" {
			return ExprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func ExprStatname(s int) string {
	if s >= 0 && s < len(ExprStatenames) {
		if ExprStatenames[s] != "" {
			return ExprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func ExprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !ExprErrorVerbose {
		return "syntax error"
	}

	for _, e := range ExprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + ExprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := ExprPact[state]
	for tok := TOKSTART; tok-1 < len(ExprToknames); tok++ {
		if n := base + tok; n >= 0 && n < ExprLast && ExprChk[ExprAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if ExprDef[state] == -2 {
		i := 0
		for ExprExca[i] != -1 || ExprExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; ExprExca[i] >= 0; i += 2 {
			tok := ExprExca[i]
			if tok < TOKSTART || ExprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if ExprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += ExprTokname(tok)
	}
	return res
}

func Exprlex1(lex ExprLexer, lval *ExprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = ExprTok1[0]
		goto out
	}
	if char < len(ExprTok1) {
		token = ExprTok1[char]
		goto out
	}
	if char >= ExprPrivate {
		if char < ExprPrivate+len(ExprTok2) {
			token = ExprTok2[char-ExprPrivate]
			goto out
		}
	}
	for i := 0; i < len(ExprTok3); i += 2 {
		token = ExprTok3[i+0]
		if token == char {
			token = ExprTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = ExprTok2[1] /* unknown char */
	}
	if ExprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", ExprTokname(token), uint(char))
	}
	return char, token
}

func ExprParse(Exprlex ExprLexer) int {
	return ExprNewParser().Parse(Exprlex)
}

func (Exprrcvr *ExprParserImpl) Parse(Exprlex ExprLexer) int {
	var Exprn int
	var ExprVAL ExprSymType
	var ExprDollar []ExprSymType
	_ = ExprDollar // silence set and not used
	ExprS := Exprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Exprstate := 0
	Exprrcvr.char = -1
	Exprtoken := -1 // Exprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Exprstate = -1
		Exprrcvr.char = -1
		Exprtoken = -1
	}()
	Exprp := -1
	goto Exprstack

ret0:
	return 0

ret1:
	return 1

Exprstack:
	/* put a state and value onto the stack */
	if ExprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", ExprTokname(Exprtoken), ExprStatname(Exprstate))
	}

	Exprp++
	if Exprp >= len(ExprS) {
		nyys := make([]ExprSymType, len(ExprS)*2)
		copy(nyys, ExprS)
		ExprS = nyys
	}
	ExprS[Exprp] = ExprVAL
	ExprS[Exprp].yys = Exprstate

Exprnewstate:
	Exprn = ExprPact[Exprstate]
	if Exprn <= ExprFlag {
		goto Exprdefault /* simple state */
	}
	if Exprrcvr.char < 0 {
		Exprrcvr.char, Exprtoken = Exprlex1(Exprlex, &Exprrcvr.lval)
	}
	Exprn += Exprtoken
	if Exprn < 0 || Exprn >= ExprLast {
		goto Exprdefault
	}
	Exprn = ExprAct[Exprn]
	if ExprChk[Exprn] == Exprtoken { /* valid shift */
		Exprrcvr.char = -1
		Exprtoken = -1
		ExprVAL = Exprrcvr.lval
		Exprstate = Exprn
		if Errflag > 0 {
			Errflag--
		}
		goto Exprstack
	}

Exprdefault:
	/* default state action */
	Exprn = ExprDef[Exprstate]
	if Exprn == -2 {
		if Exprrcvr.char < 0 {
			Exprrcvr.char, Exprtoken = Exprlex1(Exprlex, &Exprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if ExprExca[xi+0] == -1 && ExprExca[xi+1] == Exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Exprn = ExprExca[xi+0]
			if Exprn < 0 || Exprn == Exprtoken {
				break
			}
		}
		Exprn = ExprExca[xi+1]
		if Exprn < 0 {
			goto ret0
		}
	}
	if Exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Exprlex.Error(ExprErrorMessage(Exprstate, Exprtoken))
			Nerrs++
			if ExprDebug >= 1 {
				__yyfmt__.Printf("%s", ExprStatname(Exprstate))
				__yyfmt__.Printf(" saw %s\n", ExprTokname(Exprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Exprp >= 0 {
				Exprn = ExprPact[ExprS[Exprp].yys] + ExprErrCode
				if Exprn >= 0 && Exprn < ExprLast {
					Exprstate = ExprAct[Exprn] /* simulate a shift of "error" */
					if ExprChk[Exprstate] == ExprErrCode {
						goto Exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if ExprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", ExprS[Exprp].yys)
				}
				Exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if ExprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", ExprTokname(Exprtoken))
			}
			if Exprtoken == ExprEofCode {
				goto ret1
			}
			Exprrcvr.char = -1
			Exprtoken = -1
			goto Exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production Exprn */
	if ExprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Exprn, ExprStatname(Exprstate))
	}

	Exprnt := Exprn
	Exprpt := Exprp
	_ = Exprpt // guard against "declared and not used"

	Exprp -= ExprR2[Exprn]
	// Exprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Exprp+1 >= len(ExprS) {
		nyys := make([]ExprSymType, len(ExprS)*2)
		copy(nyys, ExprS)
		ExprS = nyys
	}
	ExprVAL = ExprS[Exprp+1]

	/* consult goto table to find next state */
	Exprn = ExprR1[Exprn]
	Exprg := ExprPgo[Exprn]
	Exprj := Exprg + ExprS[Exprp].yys + 1

	if Exprj >= ExprLast {
		Exprstate = ExprAct[Exprg]
	} else {
		Exprstate = ExprAct[Exprj]
		if ExprChk[Exprstate] != -Exprn {
			Exprstate = ExprAct[Exprg]
		}
	}
	// dummy call; replaced with literal code
	switch Exprnt {

	case 2:
		ExprDollar = ExprS[Exprpt-1 : Exprpt+1]
//line grammar.go.y:26
		{
			fmt.Println(ExprDollar[1].Number)
		}
	case 4:
		ExprDollar = ExprS[Exprpt-1 : Exprpt+1]
//line grammar.go.y:31
		{
			ExprVAL.Number, _ = strconv.ParseFloat(ExprDollar[1].String, 64)
		}
	case 5:
		ExprDollar = ExprS[Exprpt-1 : Exprpt+1]
//line grammar.go.y:32
		{
			ExprVAL.Number = 0
		}
	case 6:
		ExprDollar = ExprS[Exprpt-3 : Exprpt+1]
//line grammar.go.y:33
		{
			ExprVAL.Number = ExprDollar[1].Number + ExprDollar[3].Number
		}
	case 7:
		ExprDollar = ExprS[Exprpt-3 : Exprpt+1]
//line grammar.go.y:34
		{
			ExprVAL.Number = ExprDollar[1].Number - ExprDollar[3].Number
		}
	case 8:
		ExprDollar = ExprS[Exprpt-3 : Exprpt+1]
//line grammar.go.y:35
		{
			ExprVAL.Number = ExprDollar[1].Number * ExprDollar[3].Number
		}
	case 9:
		ExprDollar = ExprS[Exprpt-3 : Exprpt+1]
//line grammar.go.y:36
		{
			ExprVAL.Number = ExprDollar[1].Number / ExprDollar[3].Number
		}
	case 10:
		ExprDollar = ExprS[Exprpt-3 : Exprpt+1]
//line grammar.go.y:37
		{
			ExprVAL.Number = ExprDollar[2].Number
		}
	case 11:
		ExprDollar = ExprS[Exprpt-2 : Exprpt+1]
//line grammar.go.y:38
		{
			ExprVAL.Number = -ExprDollar[2].Number
		}
	}
	goto Exprstack /* stack new state and value */
}
