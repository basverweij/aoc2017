package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var parseInput = "ab\r\ncd\r\n"

func TestParse(t *testing.T) {
	assert.Equal(t,
		[][]byte{[]byte{'a', 'b'}, []byte{'c', 'd'}},
		parse(parseInput))
}

func TestInput(t *testing.T) {
	tubes := input()
	assert.Equal(t, 201, len(tubes))
	assert.Equal(t, 201, len(tubes[0]))
}
