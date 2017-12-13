package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = map[int]int{
	0: 3,
	1: 2,
	4: 4,
	6: 4,
}

func TestNewFirewall(t *testing.T) {
	f := newFirewall(testInput)

	assert.Equal(t, 6, f.depth)
	assert.Equal(t, 4, len(f.layers))
}

func TestFirewallAdvance(t *testing.T) {
	f := newFirewall(testInput)

	f.advance()
	assert.Equal(t, 1, f.scanner(0))
	assert.Equal(t, 1, f.scanner(1))
	assert.Equal(t, -1, f.scanner(2))
	assert.Equal(t, -1, f.scanner(3))
	assert.Equal(t, 1, f.scanner(4))
	assert.Equal(t, -1, f.scanner(5))
	assert.Equal(t, 1, f.scanner(6))

	f.advance()
	assert.Equal(t, 2, f.scanner(0))
	assert.Equal(t, 0, f.scanner(1))
	assert.Equal(t, -1, f.scanner(2))
	assert.Equal(t, -1, f.scanner(3))
	assert.Equal(t, 2, f.scanner(4))
	assert.Equal(t, -1, f.scanner(5))
	assert.Equal(t, 2, f.scanner(6))
}
