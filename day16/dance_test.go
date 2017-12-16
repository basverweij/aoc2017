package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDance(t *testing.T) {
	l := newLine(5)

	moves := []move{
		newSpin(1),
		newExchange(3, 4),
		newPartner('e', 'b'),
	}

	dance(l, moves)
	assert.Equal(t, "baedc", l.String())

	dance(l, moves)
	assert.Equal(t, "ceadb", l.String())
}
