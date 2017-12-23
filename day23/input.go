package main

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/basverweij/aoc2017/day23/coasm"
)

var rawInput = `set b 99
set c b
jnz a 2
jnz 1 5
mul b 100
sub b -100000
set c b
sub c -17000
set f 1
set d 2
set e 2
set g d
mul g e
sub g b
jnz g 2
set f 0
sub e -1
set g e
sub g b
jnz g -8
sub d -1
set g d
sub g b
jnz g -13
jnz f 2
sub h -1
set g b
sub g c
jnz g 2
jnz 1 3
sub b -17
jnz 1 -23`

func input() []*instruction {
	input := antlr.NewInputStream(rawInput)
	lexer := coasm.NewCoAsmLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := coasm.NewCoAsmParser(stream)
	p.BuildParseTrees = true

	program := p.Program()
	listener := &instructionListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, program)

	return listener.instructions
}

func parseOpcode(rawOpcode string) opcode {
	switch rawOpcode {
	case "set":
		return set
	case "sub":
		return sub
	case "mul":
		return mul
	case "jnz":
		return jnz
	}

	panic("invalid opcode: " + rawOpcode)
}

type instructionListener struct {
	*coasm.BaseCoAsmListener
	instructions []*instruction

	// temp values
	instr *instruction
}

func (l *instructionListener) EnterInstruction(ctx *coasm.InstructionContext) {
	l.instr = &instruction{
		opc: parseOpcode(ctx.Opcode().GetText()),
	}
}

func (l *instructionListener) ExitInstruction(ctx *coasm.InstructionContext) {
	l.instructions = append(l.instructions, l.instr)
}

func (l *instructionListener) EnterRegOrConst(ctx *coasm.RegOrConstContext) {
	roc := &regOrConst{}
	if ctx.Reg() != nil {
		roc.reg = rune(ctx.Reg().GetText()[0])
	} else if ctx.Const() != nil {
		roc.cnst, _ = strconv.Atoi(ctx.Const().GetText())
	} else {
		panic("invalid reg or const: " + ctx.GetText())
	}

	if l.instr.operand1 == nil {
		l.instr.operand1 = roc
	} else {
		l.instr.operand2 = roc
	}
}
