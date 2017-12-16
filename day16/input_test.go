package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	i := input()

	assert.Equal(t, 10000, len(i))
	assert.Equal(t, newPartner('o', 'k'), i[0])
	assert.Equal(t, newExchange(4, 0), i[1])
	assert.Equal(t, newSpin(12), i[2])
}
