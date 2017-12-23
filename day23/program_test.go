package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProgram(t *testing.T) {
	p := newProgram(input())
	assert.NotNil(t, p)
}
