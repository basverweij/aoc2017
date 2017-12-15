package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	l := newList(3)

	assert.NotNil(t, l)
	assert.Equal(t, 3, len(l))
}

func TestListString(t *testing.T) {
	l := newList(3)

	assert.Equal(t, "0 1 2", l.String())
}

func TestListReverse(t *testing.T) {
	l := newList(3)
	assert.Equal(t, "0 1 2", l.String())

	l.Reverse(0, 0)
	assert.Equal(t, "0 1 2", l.String())

	l.Reverse(0, 1)
	assert.Equal(t, "0 1 2", l.String())

	l.Reverse(0, 2)
	assert.Equal(t, "1 0 2", l.String())

	l.Reverse(1, 2)
	assert.Equal(t, "1 2 0", l.String())

	l.Reverse(2, 2)
	assert.Equal(t, "0 2 1", l.String())
}
