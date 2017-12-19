grammar SndAsm;

code : instruction* ;

instruction : unaryRegExpression | unaryExpression | binaryRegExpression | binaryExpression;

unaryRegExpression : unaryRegInstruction reg ;

unaryRegInstruction : Rcv ;

unaryExpression : unaryInstruction regOrValue ;

unaryInstruction : Snd ;

binaryRegExpression : binaryRegInstruction reg regOrValue ;

binaryRegInstruction : Set | Add | Mul | Mod ;

binaryExpression : binaryInstruction regOrValue regOrValue ;

binaryInstruction : Jgz ;

reg : Reg ;

regOrValue : Reg | Value ;

Snd : 'snd' ;

Rcv : 'rcv' ;

Set : 'set' ;

Add : 'add' ;

Mul : 'mul' ;

Mod : 'mod' ;

Jgz : 'jgz' ;

Reg : [a-z] ;

Value : '-'?[0-9]+ ;

WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines
