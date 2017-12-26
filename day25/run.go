package main

import (
	"math/big"
	"math/bits"
)

func runDiagnostics(m *machine) int {
	for n := 0; n < m.diagnosticsAfter; n++ {
		rules := m.states[m.currentState].rules[m.read()]

		m.write(rules.write)

		m.move(rules.move)

		m.currentState = rules.nextState
	}

	return bitCount(m.tapePos) + bitCount(m.tapeNeg)
}

func bitCount(b *big.Int) int {
	count := 0
	for _, w := range b.Bits() {
		count += bits.OnesCount64(uint64(w))
	}

	return count
}
