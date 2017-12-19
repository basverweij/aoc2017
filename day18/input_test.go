package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	assert.Len(t, input(), 41)
}

var testInput = `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`

func TestParse(t *testing.T) {
	code := parse(testInput)
	assert.Len(t, code, 10)

	p := newProgram(code)
	p.run()
}
