package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewZeroBuf(t *testing.T) {
	b := newZeroBuf()
	assert.NotNil(t, b)
}

func TestZeroBufMoveAndInsert(t *testing.T) {
	b := newZeroBuf()

	b.move(3)
	b.insert(1)
	assert.Equal(t, 1, b.valueAfterZero)

	b.move(3)
	b.insert(2)
	assert.Equal(t, 2, b.valueAfterZero)

	b.move(3)
	b.insert(3)
	assert.Equal(t, 2, b.valueAfterZero)

	b.move(3)
	b.insert(4)
	assert.Equal(t, 2, b.valueAfterZero)

	b.move(3)
	b.insert(5)
	assert.Equal(t, 5, b.valueAfterZero)

	b.move(3)
	b.insert(6)
	assert.Equal(t, 5, b.valueAfterZero)

	b.move(3)
	b.insert(7)
	assert.Equal(t, 5, b.valueAfterZero)

	b.move(3)
	b.insert(8)
	assert.Equal(t, 5, b.valueAfterZero)

	b.move(3)
	b.insert(9)
	assert.Equal(t, 9, b.valueAfterZero)
}
