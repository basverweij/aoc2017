package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosest(t *testing.T) {
	c := closest([]*particle{testPart0, testPart1})
	assert.Equal(t, 0, c)
}
