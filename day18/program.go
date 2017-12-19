package main

import "fmt"

type executable interface {
	jump(int)
	get(string) int
	set(string, int)
	send(int)
	receive() int
}

type program struct {
	pc       int
	regs     map[string]int
	code     []instruction
	sndChan  chan<- int
	sndCount int
	rcvChan  <-chan int
	pgm      int
	rcving   bool
	other    *program
	deadlock bool
}

func newProgram(pgm int, code []instruction, sndChan chan<- int, rcvChan <-chan int) *program {
	p := &program{
		regs:    map[string]int{"p": pgm},
		code:    code,
		sndChan: sndChan,
		rcvChan: rcvChan,
		pgm:     pgm,
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

func (p *program) send(value int) {
	p.sndCount++

	p.sndChan <- value
}

func (p *program) receive() int {
	p.rcving = true
	defer func() {
		if !p.deadlock {
			p.rcving = false
		}
	}()

	for {
		select {
		case value := <-p.rcvChan:
			return value
		default:
		}

		if p.other != nil && p.other.rcving {
			p.deadlock = true
			return -1
		}
	}
}

func (p *program) String() string {
	return fmt.Sprintf("%d@%03d - %v (%d)", p.pgm, p.pc, p.regs, p.sndCount)
}

func (p *program) run(other *program) {
	p.other = other

	fmt.Println(p)

	for p.pc >= 0 && p.pc < len(p.code) {
		instr := p.code[p.pc]

		if instr != nil {
			instr(p)
		}

		if p.deadlock {
			fmt.Printf("DEADLOCK: %v\n", p)
			return
		}

		p.pc++

		fmt.Println(p)
	}
}
