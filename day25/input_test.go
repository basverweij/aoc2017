package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.`

func TestInput(t *testing.T) {
	m := input(testInput)

	assert.NotNil(t, m)
	assert.Equal(t, 0, m.startState)
	assert.Equal(t, 6, m.diagnosticsAfter)

	assert.Equal(t, []*state{
		{[]*rules{
			{1, right, 1},
			{0, left, 1},
		}},
		{[]*rules{
			{1, left, 0},
			{1, right, 0},
		}},
	}, m.states)
}
