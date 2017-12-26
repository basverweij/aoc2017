// Code generated from .\Turing.g4 by ANTLR 4.7.1. DO NOT EDIT.

package turing // Turing
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTuringListener is a complete listener for a parse tree produced by TuringParser.
type BaseTuringListener struct{}

var _ TuringListener = &BaseTuringListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTuringListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTuringListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTuringListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTuringListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBlueprint is called when production blueprint is entered.
func (s *BaseTuringListener) EnterBlueprint(ctx *BlueprintContext) {}

// ExitBlueprint is called when production blueprint is exited.
func (s *BaseTuringListener) ExitBlueprint(ctx *BlueprintContext) {}

// EnterStartState is called when production startState is entered.
func (s *BaseTuringListener) EnterStartState(ctx *StartStateContext) {}

// ExitStartState is called when production startState is exited.
func (s *BaseTuringListener) ExitStartState(ctx *StartStateContext) {}

// EnterDiagnostics is called when production diagnostics is entered.
func (s *BaseTuringListener) EnterDiagnostics(ctx *DiagnosticsContext) {}

// ExitDiagnostics is called when production diagnostics is exited.
func (s *BaseTuringListener) ExitDiagnostics(ctx *DiagnosticsContext) {}

// EnterState is called when production state is entered.
func (s *BaseTuringListener) EnterState(ctx *StateContext) {}

// ExitState is called when production state is exited.
func (s *BaseTuringListener) ExitState(ctx *StateContext) {}

// EnterStateRules is called when production stateRules is entered.
func (s *BaseTuringListener) EnterStateRules(ctx *StateRulesContext) {}

// ExitStateRules is called when production stateRules is exited.
func (s *BaseTuringListener) ExitStateRules(ctx *StateRulesContext) {}

// EnterRules is called when production rules is entered.
func (s *BaseTuringListener) EnterRules(ctx *RulesContext) {}

// ExitRules is called when production rules is exited.
func (s *BaseTuringListener) ExitRules(ctx *RulesContext) {}

// EnterWrite is called when production write is entered.
func (s *BaseTuringListener) EnterWrite(ctx *WriteContext) {}

// ExitWrite is called when production write is exited.
func (s *BaseTuringListener) ExitWrite(ctx *WriteContext) {}

// EnterMove is called when production move is entered.
func (s *BaseTuringListener) EnterMove(ctx *MoveContext) {}

// ExitMove is called when production move is exited.
func (s *BaseTuringListener) ExitMove(ctx *MoveContext) {}

// EnterNextState is called when production nextState is entered.
func (s *BaseTuringListener) EnterNextState(ctx *NextStateContext) {}

// ExitNextState is called when production nextState is exited.
func (s *BaseTuringListener) ExitNextState(ctx *NextStateContext) {}
