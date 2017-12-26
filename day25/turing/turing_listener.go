// Code generated from .\Turing.g4 by ANTLR 4.7.1. DO NOT EDIT.

package turing // Turing
import "github.com/antlr/antlr4/runtime/Go/antlr"

// TuringListener is a complete listener for a parse tree produced by TuringParser.
type TuringListener interface {
	antlr.ParseTreeListener

	// EnterBlueprint is called when entering the blueprint production.
	EnterBlueprint(c *BlueprintContext)

	// EnterStartState is called when entering the startState production.
	EnterStartState(c *StartStateContext)

	// EnterDiagnostics is called when entering the diagnostics production.
	EnterDiagnostics(c *DiagnosticsContext)

	// EnterState is called when entering the state production.
	EnterState(c *StateContext)

	// EnterStateRules is called when entering the stateRules production.
	EnterStateRules(c *StateRulesContext)

	// EnterRules is called when entering the rules production.
	EnterRules(c *RulesContext)

	// EnterWrite is called when entering the write production.
	EnterWrite(c *WriteContext)

	// EnterMove is called when entering the move production.
	EnterMove(c *MoveContext)

	// EnterNextState is called when entering the nextState production.
	EnterNextState(c *NextStateContext)

	// ExitBlueprint is called when exiting the blueprint production.
	ExitBlueprint(c *BlueprintContext)

	// ExitStartState is called when exiting the startState production.
	ExitStartState(c *StartStateContext)

	// ExitDiagnostics is called when exiting the diagnostics production.
	ExitDiagnostics(c *DiagnosticsContext)

	// ExitState is called when exiting the state production.
	ExitState(c *StateContext)

	// ExitStateRules is called when exiting the stateRules production.
	ExitStateRules(c *StateRulesContext)

	// ExitRules is called when exiting the rules production.
	ExitRules(c *RulesContext)

	// ExitWrite is called when exiting the write production.
	ExitWrite(c *WriteContext)

	// ExitMove is called when exiting the move production.
	ExitMove(c *MoveContext)

	// ExitNextState is called when exiting the nextState production.
	ExitNextState(c *NextStateContext)
}
