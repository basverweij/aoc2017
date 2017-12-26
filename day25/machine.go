package main

import "math/big"

type machine struct {
	startState       int
	diagnosticsAfter int

	states []*state

	tapePos, tapeNeg *big.Int
	cursor           int
	currentState     int
}

func newMachine(startState, diagnosticsAfter int) *machine {
	return &machine{
		startState:       startState,
		diagnosticsAfter: diagnosticsAfter,
		tapePos:          new(big.Int),
		tapeNeg:          new(big.Int),
	}
}

type state struct {
	rules []*rules
}

type direction int

const (
	left direction = iota
	right
)

type rules struct {
	write     uint
	move      direction
	nextState int
}

func (m *machine) read() uint {
	if m.cursor < 0 {
		return m.tapeNeg.Bit(-m.cursor)
	}

	return m.tapePos.Bit(m.cursor)
}

func (m *machine) write(value uint) {
	if m.cursor < 0 {
		m.tapeNeg.SetBit(m.tapeNeg, -m.cursor, value)
	} else {
		m.tapePos.SetBit(m.tapePos, m.cursor, value)
	}
}

func (m *machine) move(dir direction) {
	if dir == right {
		m.cursor++
	} else {
		m.cursor--
	}
}
