package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	u0 = uint(0)
	u1 = uint(1)
)

func TestMachineReadWriteMove(t *testing.T) {
	m := newMachine(0, 6)

	assert.Equal(t, 0, m.cursor)

	assert.Equal(t, u0, m.read())
	m.write(u1)
	assert.Equal(t, u1, m.read())
	m.write(u0)
	assert.Equal(t, u0, m.read())

	m.move(right)
	assert.Equal(t, 1, m.cursor)

	m.write(u1)
	assert.Equal(t, u1, m.read())
	m.write(u0)
	assert.Equal(t, u0, m.read())

	m.move(left)
	m.move(left)
	assert.Equal(t, -1, m.cursor)

	m.write(u1)
	assert.Equal(t, u1, m.read())
	m.write(u0)
	assert.Equal(t, u0, m.read())
}
