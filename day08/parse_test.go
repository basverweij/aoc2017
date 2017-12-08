package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	input = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`
)

func TestParse(t *testing.T) {
	is, err := parse(bytes.NewBufferString(input))
	assert.Nil(t, err)

	assert.Equal(t, &instruction{"b", true, 5, "a", ">", 1}, is[0])
	assert.Equal(t, &instruction{"a", true, 1, "b", "<", 5}, is[1])
	assert.Equal(t, &instruction{"c", false, -10, "a", ">=", 1}, is[2])
	assert.Equal(t, &instruction{"c", true, -20, "c", "==", 10}, is[3])
}
