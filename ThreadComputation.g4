grammar ThreadComputation;

expressions
 : root_term*
;

root_term
 : ( fun
   | ufun
   | term
   | vars
   | dblock
   | dmap
);

fun
 : fname=(NAME|OPS) ('[' (param+=fun_term)* ']')?
;

ufun
 : '@' fname=(NAME|OPS) ('[' (param+=ufun_term)* ']')?
;

vars
 : '$' vname=NAME
;

dblock_term
 : ( fun
   | term
);

fun_term
 : ( fun
   | term
);

ufun_term
 : ( fun
   | term
   | vars
   | dblock
   | dmap
);

dblock
 : (':' bname=NAME)?'(' (param+=dblock_term)* ')'
;

dmap
 : '{' (param+=key_term)* '}'
;

integer_term: VALUE=INTEGER ;
float_term: VALUE=FLOAT_NUMBER ;
string_term: VALUE=STRING ;
boolean_term: VALUE=(TRUE|FALSE) ;
key_term: KEY=NAME ':' VALUE=(INTEGER|FLOAT_NUMBER|STRING|TRUE|FALSE) ;

term
 : ( integer_term
    | float_term
    | string_term
    | boolean_term
);

OPS
  : (OPS_START)? (OP)+
  ;
OPS_START
  : (
    '`'
  | '?'
  );
NAME
  : ID_START ID_CONTINUE*
  ;

INTEGER
  :  (SIGN)? DECIMAL_INTEGER
  ;

DECIMAL_INTEGER
  : NON_ZERO_DIGIT DIGIT*
  | '0'+
  ;

FLOAT_NUMBER
  : EXPONENT_OR_POINT_FLOAT
  ;

STRING
  : SHORT_STRING
  ;

OP
  : '+'
  | '-'
  | '*'
  | '/'
  | '|'
  | '^'
  | '&'
  | '!'
  | '~'
  | '↑'
  | '↓'
  | '×'
  | '÷'
  | '∆'
  | '∇'
  | ','
  | ';'
  | '_'
  | '<'
  | '>'
  | '='
  | '¬'
  | '∨'
  | '∧'
  ;

TRUE
  : '#true'
  | '#True'
  | '#TRUE'
  ;

FALSE
  : '#false'
  | '#False'
  | '#FALSE'
  ;

BLOCK_COMMENT
  :   '/*' .*? '*/' -> skip
  ;

WhiteSpace : [ \t]+ -> skip;
NewLine : ('\r'?'\n'|'\r') -> skip;

fragment NON_ZERO_DIGIT
  : [1-9]
  ;

fragment DIGIT
  : [0-9]
  ;

fragment SIGN
  : ('+' | '-')
  ;

fragment EXPONENT_OR_POINT_FLOAT
  : (SIGN)? ([0-9]+ | POINT_FLOAT) [eE] [+-]? [0-9]+
  | (SIGN)? POINT_FLOAT
  ;

fragment POINT_FLOAT
  : [0-9]* '.' [0-9]+
  | [0-9]+ '.'
  ;

fragment RN
  : '\r'? '\n'
  ;
fragment SHORT_STRING
  : '\'' ('\\' (RN | .) | ~[\\\r\n'])* '\''
  | '"'  ('\\' (RN | .) | ~[\\\r\n"])* '"'
  ;

fragment ID_START
 : ([A-Z]|[a-z])
 | [a-z]
 | '`'
 | '?'
 ;

fragment ID_CONTINUE
 : ID_START
 | [0-9]
 | '.'
 ;
