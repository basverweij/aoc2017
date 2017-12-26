grammar Turing;

blueprint
  : startState diagnostics state*
  ;

startState
  : 'Begin in state ' State '.'
  ;

diagnostics
  : 'Perform a diagnostic checksum after ' Num ' steps.'
  ;

state
  : 'In state ' State ':' stateRules*
  ;

stateRules
  : 'If the current value is ' Num ':' rules
  ;

rules
  : write move nextState
  ;

write
  : '- Write the value ' Num '.'
  ;

move
  : '- Move one slot to the ' Dir '.'
  ;

nextState
  : '- Continue with state ' State '.'
  ;

State : [A-F] ;

Num : [0-9]+ ;

Dir : 'left' | 'right' ;

WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines
