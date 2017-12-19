package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProgram(t *testing.T) {
	p := createProgram()
	assert.NotNil(t, p)

	assert.Equal(t, "0@000 - map[p:0] (0)", p.String())
}

func TestProgramRegs(t *testing.T) {
	p := createProgram()

	assert.Equal(t, 0, p.get("a"))

	p.set("a", 1)
	assert.Equal(t, 1, p.get("a"))
}

func TestProgramSend(t *testing.T) {
	p := createProgram()

	assert.Equal(t, p.sndCount, 0)

	// TODO
	p.send(1)
	assert.Equal(t, p.sndCount, 1)
}

func TestProgramJump(t *testing.T) {
	p := createProgram()

	assert.Equal(t, 0, p.pc)

	p.jump(2)
	assert.Equal(t, 2, p.pc)

	p.jump(-1)
	assert.Equal(t, 1, p.pc)
}
