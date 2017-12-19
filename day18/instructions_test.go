package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testValue(value int) value {
	return func(executable) int {
		return value
	}
}

func TestSnd(t *testing.T) {
	s := newSnd(testValue(1))

	p := newProgram(nil)
	s(p)
	p.recover()

	isRecovered, freq := p.isRecovered()
	assert.True(t, isRecovered)
	assert.Equal(t, 1, freq)
}

func TestRcv(t *testing.T) {
	p := newProgram(nil)
	p.sound(1)

	r := newRcv(testValue(0))
	r(p)

	isRecovered, freq := p.isRecovered()
	assert.False(t, isRecovered)
	assert.Equal(t, 0, freq)

	r = newRcv(testValue(1))
	r(p)

	isRecovered, freq = p.isRecovered()
	assert.True(t, isRecovered)
	assert.Equal(t, 1, freq)
}

func TestSet(t *testing.T) {
	p := newProgram(nil)
	p.set("a", 1)
	assert.Equal(t, 1, p.get("a"))

	s := newSet("a", testValue(2))
	s(p)

	assert.Equal(t, 2, p.get("a"))
}

func TestAdd(t *testing.T) {
	p := newProgram(nil)
	p.set("a", 1)
	assert.Equal(t, 1, p.get("a"))

	a := newAdd("a", testValue(2))
	a(p)

	assert.Equal(t, 3, p.get("a"))
}

func TestMul(t *testing.T) {
	p := newProgram(nil)
	p.set("a", 2)
	assert.Equal(t, 2, p.get("a"))

	m := newMul("a", testValue(3))
	m(p)

	assert.Equal(t, 6, p.get("a"))
}

func TestMod(t *testing.T) {
	p := newProgram(nil)
	p.set("a", 3)
	assert.Equal(t, 3, p.get("a"))

	m := newMod("a", testValue(2))
	m(p)

	assert.Equal(t, 1, p.get("a"))
}

func TestJgz(t *testing.T) {
	p := newProgram(nil)
	assert.Equal(t, 0, p.pc)

	j := newJgz(testValue(-1), testValue(2))
	j(p)
	assert.Equal(t, 0, p.pc)

	j = newJgz(testValue(0), testValue(2))
	j(p)
	assert.Equal(t, 0, p.pc)

	j = newJgz(testValue(1), testValue(2))
	j(p)
	assert.Equal(t, 1, p.pc)
}
