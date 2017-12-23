// Code generated from .\CoAsm.g4 by ANTLR 4.7.1. DO NOT EDIT.

package coasm // CoAsm
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CoAsmListener is a complete listener for a parse tree produced by CoAsmParser.
type CoAsmListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterRegOrConst is called when entering the regOrConst production.
	EnterRegOrConst(c *RegOrConstContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitRegOrConst is called when exiting the regOrConst production.
	ExitRegOrConst(c *RegOrConstContext)
}
