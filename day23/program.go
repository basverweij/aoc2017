package main

type program struct {
	instructions []*instruction
	regs         map[rune]int
}

func newProgram(instructions []*instruction) *program {
	return &program{
		instructions: instructions,
		regs:         make(map[rune]int, 8),
	}
}

// run returns the number of times the 'mul' instruction was run
func (p *program) run() int {
	mulCount := 0

	var i *instruction
	pc := 0
	for pc >= 0 && pc < len(p.instructions) {
		i = p.instructions[pc]

		switch i.opc {
		case set:
			p.regs[i.operand1.reg] = i.operand2.val(p)
		case sub:
			p.regs[i.operand1.reg] -= i.operand2.val(p)
		case mul:
			p.regs[i.operand1.reg] *= i.operand2.val(p)
			mulCount++
		}

		if i.opc == jnz && i.operand1.val(p) != 0 {
			pc += i.operand2.val(p)
		} else {
			pc++
		}
	}

	return mulCount
}
