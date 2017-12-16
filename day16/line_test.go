package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLine(t *testing.T) {
	l := newLine(16)
	assert.NotNil(t, l)
}
