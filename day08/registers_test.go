package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	is, _ := parse(bytes.NewBufferString(input))

	regs := newRegisters()

	for _, i := range is {
		regs.apply(i)
	}

	assert.Equal(t, 1, regs.largestValue())
}
