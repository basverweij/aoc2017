// Code generated from .\CoAsm.g4 by ANTLR 4.7.1. DO NOT EDIT.

package coasm // CoAsm
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCoAsmListener is a complete listener for a parse tree produced by CoAsmParser.
type BaseCoAsmListener struct{}

var _ CoAsmListener = &BaseCoAsmListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCoAsmListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCoAsmListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCoAsmListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCoAsmListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseCoAsmListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseCoAsmListener) ExitProgram(ctx *ProgramContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseCoAsmListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseCoAsmListener) ExitInstruction(ctx *InstructionContext) {}

// EnterRegOrConst is called when production regOrConst is entered.
func (s *BaseCoAsmListener) EnterRegOrConst(ctx *RegOrConstContext) {}

// ExitRegOrConst is called when production regOrConst is exited.
func (s *BaseCoAsmListener) ExitRegOrConst(ctx *RegOrConstContext) {}
