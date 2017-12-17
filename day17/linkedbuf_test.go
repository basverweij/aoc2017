package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedBuf(t *testing.T) {
	b := newLinkedBuf()
	assert.NotNil(t, b)
	assert.NotNil(t, b.zero)
	assert.NotNil(t, b.current)
	assert.Equal(t, b.zero, b.current)
}

func TestLinkedBufMove(t *testing.T) {
	b := newLinkedBuf()

	b.move(3)
	b.insert(1)
	assert.Equal(t, []int{1, 0}, b.buf())

	b.move(3)
	b.insert(2)
	assert.Equal(t, []int{2, 1, 0}, b.buf())

	b.move(3)
	b.insert(3)
	assert.Equal(t, []int{3, 1, 0, 2}, b.buf())
}
