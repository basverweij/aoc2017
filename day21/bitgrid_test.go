package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRule(t *testing.T) {
	b := newBitgrid(3)
	assert.NotNil(t, b)
}

func TestSetGetKeyString(t *testing.T) {
	b := newBitgrid(2)
	assert.False(t, b.get(0, 0))
	assert.False(t, b.get(0, 1))
	assert.False(t, b.get(1, 0))
	assert.False(t, b.get(1, 1))
	assert.EqualValues(t, 0, b.key(0, 0, 2))
	assert.Equal(t, "../..", b.String())

	b.set(0, 0, true)
	assert.True(t, b.get(0, 0))
	assert.False(t, b.get(0, 1))
	assert.False(t, b.get(1, 0))
	assert.False(t, b.get(1, 1))
	assert.EqualValues(t, 1, b.key(0, 0, 2))
	assert.Equal(t, "#./..", b.String())

	b.set(0, 0, false)
	b.set(0, 1, true)
	assert.False(t, b.get(0, 0))
	assert.True(t, b.get(0, 1))
	assert.False(t, b.get(1, 0))
	assert.False(t, b.get(1, 1))
	assert.EqualValues(t, 4, b.key(0, 0, 2))
	assert.Equal(t, "../#.", b.String())

	b.set(0, 1, false)
	b.set(1, 0, true)
	assert.False(t, b.get(0, 0))
	assert.False(t, b.get(0, 1))
	assert.True(t, b.get(1, 0))
	assert.False(t, b.get(1, 1))
	assert.EqualValues(t, 2, b.key(0, 0, 2))
	assert.Equal(t, ".#/..", b.String())

	b.set(1, 0, false)
	b.set(1, 1, true)
	assert.False(t, b.get(0, 0))
	assert.False(t, b.get(0, 1))
	assert.False(t, b.get(1, 0))
	assert.True(t, b.get(1, 1))
	assert.EqualValues(t, 8, b.key(0, 0, 2))
	assert.Equal(t, "../.#", b.String())

	b.set(1, 1, false)
	assert.False(t, b.get(0, 0))
	assert.False(t, b.get(0, 1))
	assert.False(t, b.get(1, 0))
	assert.False(t, b.get(1, 1))
	assert.EqualValues(t, 0, b.key(0, 0, 2))
	assert.Equal(t, "../..", b.String())
}

func TestCountOn(t *testing.T) {
	start := parse(rawStart)

	assert.Equal(t, 5, start.countOn())
}

func TestCopyFrom(t *testing.T) {
	start := parse(rawStart)

	b := newBitgrid(4)
	b.copyFrom(start, 1, 1)

	assert.Equal(t, "..../..#./...#/.###", b.String())
}
