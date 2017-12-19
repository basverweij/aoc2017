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

func createProgram(instrs ...instruction) *program {
	chan0 := make(chan int, 100)
	return newProgram(0, instrs, chan0, chan0)
}

func TestSnd(t *testing.T) {
	p := createProgram()

	s := newSnd(testValue(1))
	s(p)
}

func TestRcv(t *testing.T) {
	p := createProgram()
	p.send(1)

	r := newRcv("a")
	r(p)
}

func TestSet(t *testing.T) {
	p := createProgram()
	p.set("a", 1)
	assert.Equal(t, 1, p.get("a"))

	s := newSet("a", testValue(2))
	s(p)

	assert.Equal(t, 2, p.get("a"))
}

func TestAdd(t *testing.T) {
	p := createProgram()
	p.set("a", 1)
	assert.Equal(t, 1, p.get("a"))

	a := newAdd("a", testValue(2))
	a(p)

	assert.Equal(t, 3, p.get("a"))
}

func TestMul(t *testing.T) {
	p := createProgram()
	p.set("a", 2)
	assert.Equal(t, 2, p.get("a"))

	m := newMul("a", testValue(3))
	m(p)

	assert.Equal(t, 6, p.get("a"))
}

func TestMod(t *testing.T) {
	p := createProgram()
	p.set("a", 3)
	assert.Equal(t, 3, p.get("a"))

	m := newMod("a", testValue(2))
	m(p)

	assert.Equal(t, 1, p.get("a"))
}

func TestJgz(t *testing.T) {
	p := createProgram()
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
