package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBuf(t *testing.T) {
	b := newBuf(2018)
	assert.NotNil(t, b)

	// new buffer should have length 1
	assert.Equal(t, 1, b.len)

	// new buffer should contain zero as the first element
	assert.Equal(t, 0, b.d[0])

	// new buffer should start at zero
	assert.Equal(t, 0, b.pos)
}

func TestBufMove(t *testing.T) {
	b := newBuf(2018)

	b.move(3)
	b.insert(1)
	assert.Equal(t, []int{0, 1}, b.buf())

	b.move(3)
	b.insert(2)
	assert.Equal(t, []int{0, 2, 1}, b.buf())

	b.move(3)
	b.insert(3)
	assert.Equal(t, []int{0, 2, 3, 1}, b.buf())
}

func TestBufInsert(t *testing.T) {
	b := newBuf(2018)

	b.insert(1)
	assert.Equal(t, 2, b.len)
	assert.Equal(t, []int{0, 1}, b.buf())
}
