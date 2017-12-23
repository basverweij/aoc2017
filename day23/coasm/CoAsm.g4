grammar CoAsm ;

program
  : instruction*
  ;

instruction
  : Opcode regOrConst regOrConst
  ;

Opcode
  : 'set'
  | 'sub'
  | 'mul'
  | 'jnz'
  ;

Reg
  : 'a'
  | 'b'
  | 'c'
  | 'd'
  | 'e'
  | 'f'
  | 'g'
  | 'h'
  ;

regOrConst
  : Reg
  | Const
  ;

Const
  : '-'?[0-9]+
  ;

WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines
