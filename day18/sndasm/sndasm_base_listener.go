// Code generated from .\SndAsm.g4 by ANTLR 4.7.1. DO NOT EDIT.

package sndasm // SndAsm
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSndAsmListener is a complete listener for a parse tree produced by SndAsmParser.
type BaseSndAsmListener struct{}

var _ SndAsmListener = &BaseSndAsmListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSndAsmListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSndAsmListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSndAsmListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSndAsmListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCode is called when production code is entered.
func (s *BaseSndAsmListener) EnterCode(ctx *CodeContext) {}

// ExitCode is called when production code is exited.
func (s *BaseSndAsmListener) ExitCode(ctx *CodeContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseSndAsmListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseSndAsmListener) ExitInstruction(ctx *InstructionContext) {}

// EnterUnaryRegExpression is called when production unaryRegExpression is entered.
func (s *BaseSndAsmListener) EnterUnaryRegExpression(ctx *UnaryRegExpressionContext) {}

// ExitUnaryRegExpression is called when production unaryRegExpression is exited.
func (s *BaseSndAsmListener) ExitUnaryRegExpression(ctx *UnaryRegExpressionContext) {}

// EnterUnaryRegInstruction is called when production unaryRegInstruction is entered.
func (s *BaseSndAsmListener) EnterUnaryRegInstruction(ctx *UnaryRegInstructionContext) {}

// ExitUnaryRegInstruction is called when production unaryRegInstruction is exited.
func (s *BaseSndAsmListener) ExitUnaryRegInstruction(ctx *UnaryRegInstructionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseSndAsmListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseSndAsmListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryInstruction is called when production unaryInstruction is entered.
func (s *BaseSndAsmListener) EnterUnaryInstruction(ctx *UnaryInstructionContext) {}

// ExitUnaryInstruction is called when production unaryInstruction is exited.
func (s *BaseSndAsmListener) ExitUnaryInstruction(ctx *UnaryInstructionContext) {}

// EnterBinaryRegExpression is called when production binaryRegExpression is entered.
func (s *BaseSndAsmListener) EnterBinaryRegExpression(ctx *BinaryRegExpressionContext) {}

// ExitBinaryRegExpression is called when production binaryRegExpression is exited.
func (s *BaseSndAsmListener) ExitBinaryRegExpression(ctx *BinaryRegExpressionContext) {}

// EnterBinaryRegInstruction is called when production binaryRegInstruction is entered.
func (s *BaseSndAsmListener) EnterBinaryRegInstruction(ctx *BinaryRegInstructionContext) {}

// ExitBinaryRegInstruction is called when production binaryRegInstruction is exited.
func (s *BaseSndAsmListener) ExitBinaryRegInstruction(ctx *BinaryRegInstructionContext) {}

// EnterBinaryExpression is called when production binaryExpression is entered.
func (s *BaseSndAsmListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production binaryExpression is exited.
func (s *BaseSndAsmListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterBinaryInstruction is called when production binaryInstruction is entered.
func (s *BaseSndAsmListener) EnterBinaryInstruction(ctx *BinaryInstructionContext) {}

// ExitBinaryInstruction is called when production binaryInstruction is exited.
func (s *BaseSndAsmListener) ExitBinaryInstruction(ctx *BinaryInstructionContext) {}

// EnterReg is called when production reg is entered.
func (s *BaseSndAsmListener) EnterReg(ctx *RegContext) {}

// ExitReg is called when production reg is exited.
func (s *BaseSndAsmListener) ExitReg(ctx *RegContext) {}

// EnterRegOrValue is called when production regOrValue is entered.
func (s *BaseSndAsmListener) EnterRegOrValue(ctx *RegOrValueContext) {}

// ExitRegOrValue is called when production regOrValue is exited.
func (s *BaseSndAsmListener) ExitRegOrValue(ctx *RegOrValueContext) {}
