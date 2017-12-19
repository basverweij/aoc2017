package main

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/basverweij/aoc2017/day18/sndasm"
)

func input() []instruction {
	f, _ := os.Open("input.txt")
	s, _ := ioutil.ReadAll(f)
	return parse(string(s))
}

func parse(data string) []instruction {
	input := antlr.NewInputStream(data)
	lexer := sndasm.NewSndAsmLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := sndasm.NewSndAsmParser(stream)
	p.BuildParseTrees = true

	code := p.Code().(*sndasm.CodeContext)
	listener := &instructionListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, code)

	return listener.instrs
}

func constValue(value int) value {
	return func(executable) int {
		return value
	}
}

func regValue(reg string) value {
	return func(x executable) int {
		return x.get(reg)
	}
}

type instructionListener struct {
	*sndasm.BaseSndAsmListener
	instrs []instruction

	instr     instruction
	reg       string
	val, val2 value
}

func (l *instructionListener) ExitInstruction(ctx *sndasm.InstructionContext) {
	if l.instr == nil {
		panic("instruction must be non-nil")
	}

	l.instrs = append(l.instrs, l.instr)
	l.instr = nil

	l.reg = ""
	l.val, l.val2 = nil, nil
}

func (l *instructionListener) EnterReg(ctx *sndasm.RegContext) {
	l.reg = ctx.GetText()
}

func (l *instructionListener) EnterRegOrValue(ctx *sndasm.RegOrValueContext) {
	var v value
	if ctx.Reg() != nil {
		v = regValue(ctx.Reg().GetText())
	} else if ctx.Value() != nil {
		i, _ := strconv.Atoi(ctx.Value().GetText())
		v = constValue(i)
	}

	if l.val == nil {
		l.val = v
	} else {
		l.val2 = v
	}
}

func (l *instructionListener) ExitUnaryExpression(ctx *sndasm.UnaryExpressionContext) {
	if l.val == nil {
		panic("value must be non-nil for unary expression")
	}

	instrCtx := ctx.UnaryInstruction().(*sndasm.UnaryInstructionContext)

	if instrCtx.Snd() != nil {
		l.instr = newSnd(l.val)
	} else if instrCtx.Rcv() != nil {
		l.instr = newRcv(l.val)
	}
}

func (l *instructionListener) ExitBinaryRegExpression(ctx *sndasm.BinaryRegExpressionContext) {
	if l.reg == "" {
		panic("reg must be non-empty for binary reg expression")
	}

	if l.val == nil {
		panic("value must be non-nil for binary reg expression")
	}

	instrCtx := ctx.BinaryRegInstruction().(*sndasm.BinaryRegInstructionContext)

	if instrCtx.Set() != nil {
		l.instr = newSet(l.reg, l.val)
	} else if instrCtx.Add() != nil {
		l.instr = newAdd(l.reg, l.val)
	} else if instrCtx.Mul() != nil {
		l.instr = newMul(l.reg, l.val)
	} else if instrCtx.Mod() != nil {
		l.instr = newMod(l.reg, l.val)
	}
}

func (l *instructionListener) ExitBinaryExpression(ctx *sndasm.BinaryExpressionContext) {
	if l.val == nil {
		panic("value must be non-nil for binary expression")
	}

	if l.val2 == nil {
		panic("second value must be non-nil for binary expression")
	}

	instrCtx := ctx.BinaryInstruction().(*sndasm.BinaryInstructionContext)

	if instrCtx.Jgz() != nil {
		l.instr = newJgz(l.val, l.val2)
	}
}
