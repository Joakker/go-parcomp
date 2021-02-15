grammar Expr;

file: stmt* ;   // A file consists of many statements
stmt: expr                      # ExprStmt
        // If a statement corresponds to an expression, print it's result
    | namelist '=' exprlist     # AssignStmt
        // Assign exprlist to namelist
    ;

// The order in which options are defined denotes the precedence they have
// compared to the others
expr: '(' expr ')'                  # NoOp
        // Parentheses have the highest precedence
    | <assoc=right> expr '^' expr   # Pow
        // Exponents go next
    | expr op=('*' | '/') expr      # MulDiv
        // Multiplication and division have the same precedence
    | expr op=('+' | '-') expr      # AddSub
        // Ditto with addition and subtraction
    | NUM                           # Number
    | ID                            # Identifier
        // Finally, numbers and identifiers
    ;

namelist: ID (',' ID)*      ;   // Define this so we can use GetAllID() in the code
exprlist: expr (',' expr)*  ;   // This one isn't necessary but for the sake of
                                // consistency

// Define these so we can easily refer to them by name in the code
POW : '^' ;
MUL : '*' ;
DIV : '/' ;
ADD : '+' ;
SUB : '-' ;
LPR : '(' ;
RPR : ')' ;
COM : ',' ;
EQU : '=' ;

ID  : ID_START ID_CONTINUE*         ;
NUM : [0-9]+ ('.' [0-9]+)?          ;
WS  : [\p{White-Space}]+    -> skip ;

// Use \p{Something} to select characters inside range Something or that have a
// the corresponding attribute
fragment ID_START       : [\p{Alpha}\p{General-Category=Other-Letter}]  ;
fragment ID_CONTINUE    : [\p{Alnum}\p{General-Category=Other-Letter}]  ;
