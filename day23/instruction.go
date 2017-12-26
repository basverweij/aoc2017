package main

import (
	"fmt"
	"strconv"
)

type opcode int

const (
	set opcode = iota
	sub
	mul
	jnz
)

func (o opcode) String() string {
	switch o {
	case set:
		return "set"
	case sub:
		return "sub"
	case mul:
		return "mul"
	case jnz:
		return "jnz"
	}

	panic("invalid opcode: " + strconv.Itoa(int(o)))
}

type regOrConst struct {
	isReg bool
	reg   int
	cnst  int
}

func (roc *regOrConst) String() string {
	if roc.reg != 0 {
		return string(rune(roc.reg) + 'a')
	}

	return strconv.Itoa(roc.cnst)
}

func (roc *regOrConst) val(p *program) int {
	if roc.isReg {
		return p.regs[roc.reg]
	}

	return roc.cnst
}

type instruction struct {
	opc      opcode
	operand1 *regOrConst
	operand2 *regOrConst
}

func (i *instruction) String() string {
	return fmt.Sprintf("%v %v %v", i.opc, i.operand1, i.operand2)
}
