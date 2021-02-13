%{
package main

import (
    "fmt"
    "strconv"
    "math"
)
%}

%union {
    String string       // Number or name
    Number float64      // Value that an expr evaluates to
}

%left '+' '-'           /* Lowest operator precedence */
%left '*' '/'           /* Mext highest operator precedence*/

%token '(' ')'

%token  <String> NUMBER IDENTIFIER  /* These tokens use the field String */
%type   <Number> expr               /* While this one evaluates to a number*/

%%

start:
        /* A blank line is valid */
     | expr             { fmt.Println($1) }
        /*
        * Print the result of an expression when it
        * is the only element of the statement
        */

     | assignment
        /*
        * Do nothing on assignment, as it would be
        * taken care of in the rule itself
        */
     ;

expr:
      NUMBER                { $$, _ = strconv.ParseFloat($1, 64) }
        /*
        * When an expression consists of a number (string), it's value
        * corresponds to the float64 that string represents.
        */

    | IDENTIFIER            { $$ = Variables.Get($1)        }
        /*
         * IDENTIFIER corresponds to a string, so we can use it to index
         * the Variables map
         */

    | expr '+' expr         { $$ = $1 + $3  }
    | expr '-' expr         { $$ = $1 - $3  }
    | expr '*' expr         { $$ = $1 * $3  }
    | expr '/' expr         { $$ = $1 / $3  }
        /* Perform the corresponding binary operation */

    | '(' expr ')'          { $$ = $2       }
        /* Group expressions between parens to give them higher precedence */

    | '-' expr   %prec '*'  { $$ = -$2      }
        /* The unary minus has the same precedence as multiplication */
    ;

assignment:
          IDENTIFIER '=' expr   { if !math.IsNaN($3) { Variables[$1] = $3 } }
            /* Store the value of expr in the Variables map, but
             * only if it is defined (distinct from math.NaN())
             */
%%
