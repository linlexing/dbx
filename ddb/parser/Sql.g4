/*
  Author : Wdj
*/

grammar Sql;

AS : A S ;
SELECT : S E L E C T ;
FROM : F R O M ;
MAX : M A X ;
SUM : S U M ;
AVG : A V G ;
MIN : M I N ;
COUNT : C O U N T ;
DISTINCT : D I S T I N C T ;
WHERE : W H E R E ;
GROUP : G R O U P ;
BY : B Y ;
ORDER : O R D E R ;
NULLS : N U L L S ;
FIRST : F I R S T ;
LAST : L A S T ;
HAVING : H A V I N G ;
NOT : N O T ;
IS : I S ;
BETWEEN : B E T W E E N ;
AND : A N D ;
IN : I N ;
NULL : N U L L ;
OR : O R ;
ASC : A S C ;
DESC : D E S C ;
LIMIT : L I M I T ;
OFFSET : O F F S E T ;
LIKE : L I K E ;
EXISTS : E X I S T S ;
CAST : C A S T ;
INNER : I N N E R ;
LEFT : L E F T ;
RIGHT : R I G H T ;
JOIN : J O I N ;
ON : O N ;
UNION : U N I O N ;
ALL : A L L ;
CASE : C A S E ;
WHEN : W H E N ;
THEN : T H E N ;
ELSE : E L S E ;
END : E N D ;

fragment A : [aA] ;
fragment B : [bB] ;
fragment C : [cC] ;
fragment D : [dD] ;
fragment E : [eE] ;
fragment F : [fF] ;
fragment G : [gG] ;
fragment H : [hH] ;
fragment I : [iI] ;
fragment J : [jJ] ;
fragment K : [kK] ;
fragment L : [lL] ;
fragment M : [mM] ;
fragment N : [nN] ;
fragment O : [oO] ;
fragment P : [pP] ;
fragment Q : [qQ] ;
fragment R : [rR] ;
fragment S : [sS] ;
fragment T : [tT] ;
fragment U : [uU] ;
fragment V : [vV] ;
fragment W : [wW] ;
fragment X : [xX] ;
fragment Y : [yY] ;
fragment Z : [zZ] ;
fragment DEC_DIGIT : [0-9] ;  //10
fragment HEX_DIGIT : [0-9A-F] ;  //16
fragment LETTER : [a-zA-Z] ;

DECIMAL_LITERAL: DEC_DIGIT+ ;  //10

ID : ('a'..'z' | 'A'..'Z' | '\u4E00'..'\u9FA5')
 ('a'..'z' | 'A'..'Z' | '\u4E00'..'\u9FA5' | '\uFF08'..'\uFF09' | '0'..'9' | '.' | '_' | '*' | '$')* ;
//Chinese | bracket:('\u4E00'..'\u9FA5' | '\uFF08'..'\uFF09')

TEXT_STRING : ('\'' (('\\' '\\') | ('\'' '\'') | ('\\' '\'') | ~('\''))* '\'') ;

TEXT_ALIAS : ('"' ~[" \t\r\n]+ '"') ;  //Oracle, PostgreSQL

BIND_VARIABLE : (':' ~[: \t\r\n]+) ;  //Oracle, PostgreSQL

columnName : (star='*' | ID) ;
tableName : ID ;
typeName : ID ;
functionName : ID ;
alias : ID | TEXT_ALIAS ;

join : (INNER | LEFT | RIGHT) JOIN ;

union : UNION ALL? ;

decimalLiteral : DECIMAL_LITERAL ;
textLiteral : TEXT_STRING ;
bind_variables : BIND_VARIABLE ;

selectStatement :
 SELECT
 selectElements
 FROM tableSources joinClause
 (whereClause)?
 (groupByClause)?
 (havingClause)?
 (orderByClause)?
 (limitClause)?
 | selectStatement union selectStatement
 ;

selectElements : (selectElement)(',' selectElement)* ;

selectElement : expr (AS? alias)? ;

expr
 : columnName
 | functionCall
 | value
 | expr ('*'|'/') expr
 | expr ('+'|'-') expr
 | expr ('||') expr
 | '(' expr ')'
 | CASE WHEN logicExpression THEN expr (WHEN logicExpression THEN expr)* (ELSE expr)? END
 | CASE expr WHEN expr THEN expr (WHEN expr THEN expr)* (ELSE expr)? END
 | NULL
 | selectStatement
 ;

value
 : decimalLiteral
 | textLiteral
 | bind_variables
 ;

functionCall
 : aggregateFunction
 | commonFunction
 ;

aggregateFunction
 : (AVG | MAX | MIN | SUM) '(' functionArg ')'
 | COUNT '(' (starArg='*' | functionArg) ')'
 ;

commonFunction
 : functionName '(' functionArg ')'
 | DISTINCT functionArg
 | CAST '(' functionArg AS typeName ')'
 ;

functionArg : expr (',' expr)* ;

tableSources : tableSource alias? (',' tableSource alias?)* ;

tableSource
 : tableName
 | '(' selectStatement ')'
 ;

joinClause : (join tableSource alias ON logicExpression)* ;

whereClause : WHERE logicExpression ;

logicExpression
 : expr comparisonOperator expr
 | expr NOT? BETWEEN expr AND expr
 | expr NOT? IN '(' expr (',' expr)* ')'
 | expr NOT? IN '(' selectStatement ')'
 | expr NOT? LIKE expr
 | expr IS NOT? NULL
 | EXISTS leftBracket='(' selectStatement rightBracket=')'
 | COMMENT? leftBracket='(' logicExpression rightBracket=')'
 | NOT logicExpression
 | logicExpression logicalOperator=AND logicExpression
 | logicExpression logicalOperator=OR logicExpression
 ;

comparisonOperator
 : '=' | '>' | '<'
 | '>=' | '<=' | '<>'
 | '~' | '!~' | '?[]'
 | '!?[]'
 ;
//'~' => regexp

groupByClause : GROUP BY groupByItem (',' groupByItem)* ;

groupByItem : expr ;

havingClause : HAVING logicExpression ;

orderByClause : ORDER BY orderByExpression (',' orderByExpression)* ;

orderByExpression : expr order=(ASC | DESC)?
 ((NULLS FIRST) | (NULLS LAST))? ;

limitClause : LIMIT
 (
 (offset=decimalLiteral ',')? limit=decimalLiteral
 |
 limit=decimalLiteral OFFSET offset=decimalLiteral
 )
 ;
//MySQL, PostgreSQL
//Oracle supports rownum

COMMENT : '/*' .*? '*/' ;

WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines

//antlr4 -Dlanguage=Go -visitor Sql.g4