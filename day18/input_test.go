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

	// chan0 := make(chan int, 100)
	// chan1 := make(chan int, 100)
	//
	// p0 := newProgram(0, code, chan1, chan0)
	// p1 := newProgram(1, code, chan0, chan1)
	//
	// p0.run(p1)
}
