grammar ThreadComputation;

expressions
 : root_term*
;

root_term
 : ( fun
   | dblock
   | dmap
);

fun
 : (mod=MOD)? fname=FUNC_NAME ('[' (param+=fun_term)* ']')?
;

dblock_term
 : ( fun
   | dblock
   | dmap
);

fun_term
 : ( fun
   | dmap
   | pos_term
);

pos_term
 : '#' pname=NAME ;

dblock
 : (':' bname=NAME)?'(' (param+=dblock_term)* ')'
;

dmap
 : '{' (param+=key_term)* '}'
;

key_term: KEY=NAME ':' VALUE=fun_term ;

FUNC_NAME
  : NAME
  | OPS
  | INTEGER
  | FLOAT_NUMBER
  | STRING
  ;

OPS
  : (OP)+
  ;

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

MOD
  : '?'
  | '`'
  | '~'
  | '@'
  | '$'
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
 ;

fragment ID_CONTINUE
 : ID_START
 | '.'
 ;
