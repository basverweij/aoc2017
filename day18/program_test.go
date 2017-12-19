package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProgram(t *testing.T) {
	p := newProgram(nil)
	assert.NotNil(t, p)
}

func TestProgramRegs(t *testing.T) {
	p := newProgram(nil)

	assert.Equal(t, 0, p.get("a"))

	p.set("a", 1)
	assert.Equal(t, 1, p.get("a"))
}

func TestProgramSound(t *testing.T) {
	p := newProgram(nil)

	assert.Equal(t, 0, p.freq)
	assert.False(t, p.recovered)

	p.sound(1)
	assert.Equal(t, 1, p.freq)
	assert.False(t, p.recovered)

	p.recover()
	assert.Equal(t, 1, p.freq)
	assert.True(t, p.recovered)
}

func TestProgramJump(t *testing.T) {
	p := newProgram(nil)

	assert.Equal(t, 0, p.pc)

	p.jump(2)
	assert.Equal(t, 2, p.pc)

	p.jump(-1)
	assert.Equal(t, 1, p.pc)
}
