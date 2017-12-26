package main

import "fmt"

type program struct {
	instructions []*instruction
	regs         []int
}

func newProgram(instructions []*instruction) *program {
	return &program{
		instructions: instructions,
		regs:         make([]int, 8),
	}
}

// run returns the number of times the `mul` instruction was run,
// and the value of the `h` register
func (p *program) run(start map[rune]int) (int, int) {
	if start != nil {
		for reg, val := range start {
			p.regs[int(reg-'a')] = val
		}
	}

	mulCount := 0

	var i *instruction
	pc := 0
	n := 0
	for ; n < 1000 && pc >= 0 && pc < len(p.instructions); n++ {
		i = p.instructions[pc]

		// if n%50000000 == 0 {
		fmt.Printf("@%02d;%d;%d;%d;%d;%d;%d;%d;%d;%v\r\n",
			pc,
			p.regs[0], p.regs[1], p.regs[2], p.regs[3],
			p.regs[4], p.regs[5], p.regs[6], p.regs[7],
			i)
		// }

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

	fmt.Println(n)

	return mulCount, p.regs[7]
}
