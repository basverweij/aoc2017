// Code generated from .\SndAsm.g4 by ANTLR 4.7.1. DO NOT EDIT.

package sndasm // SndAsm
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SndAsmListener is a complete listener for a parse tree produced by SndAsmParser.
type SndAsmListener interface {
	antlr.ParseTreeListener

	// EnterCode is called when entering the code production.
	EnterCode(c *CodeContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterUnaryInstruction is called when entering the unaryInstruction production.
	EnterUnaryInstruction(c *UnaryInstructionContext)

	// EnterBinaryRegExpression is called when entering the binaryRegExpression production.
	EnterBinaryRegExpression(c *BinaryRegExpressionContext)

	// EnterBinaryRegInstruction is called when entering the binaryRegInstruction production.
	EnterBinaryRegInstruction(c *BinaryRegInstructionContext)

	// EnterBinaryExpression is called when entering the binaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterBinaryInstruction is called when entering the binaryInstruction production.
	EnterBinaryInstruction(c *BinaryInstructionContext)

	// EnterReg is called when entering the reg production.
	EnterReg(c *RegContext)

	// EnterRegOrValue is called when entering the regOrValue production.
	EnterRegOrValue(c *RegOrValueContext)

	// ExitCode is called when exiting the code production.
	ExitCode(c *CodeContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitUnaryInstruction is called when exiting the unaryInstruction production.
	ExitUnaryInstruction(c *UnaryInstructionContext)

	// ExitBinaryRegExpression is called when exiting the binaryRegExpression production.
	ExitBinaryRegExpression(c *BinaryRegExpressionContext)

	// ExitBinaryRegInstruction is called when exiting the binaryRegInstruction production.
	ExitBinaryRegInstruction(c *BinaryRegInstructionContext)

	// ExitBinaryExpression is called when exiting the binaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitBinaryInstruction is called when exiting the binaryInstruction production.
	ExitBinaryInstruction(c *BinaryInstructionContext)

	// ExitReg is called when exiting the reg production.
	ExitReg(c *RegContext)

	// ExitRegOrValue is called when exiting the regOrValue production.
	ExitRegOrValue(c *RegOrValueContext)
}
