%{
package main

import (
    "fmt"
    "strconv"
)
%}

%union {
    String string
    Number float64
}

%left '+' '-'
%left '*' '/'

%token '(' ')'

%token <String> NUMBER IDENTIFIER
%type <Number> expr

%%

start:
     | expr             { fmt.Println($1) }
     | assignment
     ;

expr:
      NUMBER                { $$, _ = strconv.ParseFloat($1, 64) }
    | IDENTIFIER            { $$ = 0        }
    | expr '+' expr         { $$ = $1 + $3  }
    | expr '-' expr         { $$ = $1 - $3  }
    | expr '*' expr         { $$ = $1 * $3  }
    | expr '/' expr         { $$ = $1 / $3  }
    | '(' expr ')'          { $$ = $2       }
    | '-' expr   %prec '*'  { $$ = -$2      }
    ;

assignment:
          IDENTIFIER '=' expr
%%
