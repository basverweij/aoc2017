package main

import "fmt"

type executable interface {
	jump(int)
	get(string) int
	set(string, int)
	sound(int)
	recover()
	isRecovered() (bool, int)
}

type program struct {
	pc        int
	regs      map[string]int
	code      []instruction
	freq      int
	recovered bool
}

func newProgram(code []instruction) *program {
	p := &program{
		regs: make(map[string]int),
		code: code,
	}

	return p
}

func (p *program) jump(n int) {
	p.pc += n
}

func (p *program) get(reg string) int {
	return p.regs[reg]
}

func (p *program) set(reg string, value int) {
	p.regs[reg] = value
}

func (p *program) sound(freq int) {
	p.freq = freq
}

func (p *program) recover() {
	p.recovered = true
}

func (p *program) isRecovered() (bool, int) {
	if !p.recovered {
		return false, 0
	}

	return true, p.freq
}

func (p *program) String() string {
	return fmt.Sprintf("@%03d - %v (%d)", p.pc, p.regs, p.freq)
}

func (p *program) run() {
	for n := 0; n < 10000 && p.pc >= 0 && p.pc < len(p.code); n++ {
		instr := p.code[p.pc]

		if instr != nil {
			instr(p)
		}

		p.pc++
		// fmt.Printf("#%03d: %v\n", n, p)

		if p.recovered {
			return
		}
	}
}
