grammar Expr;

file: stmt* ;
stmt: expr                      # ExprStmt
    | namelist '=' exprlist     # AssignStmt
    ;
expr: '(' expr ')'                  # NoOp
    | <assoc=right> expr '^' expr   # Pow
    | expr op=('*' | '/') expr      # MulDiv
    | expr op=('+' | '-') expr      # AddSub
    | NUM                           # Number
    | ID                            # Identifier
    ;

namelist: ID (',' ID)*      ;
exprlist: expr (',' expr)*  ;

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

fragment ID_START       : [\p{Alpha}\p{General-Category=Other-Letter}]  ;
fragment ID_CONTINUE    : [\p{Alnum}\p{General-Category=Other-Letter}]  ;
