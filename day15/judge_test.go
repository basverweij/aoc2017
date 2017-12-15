package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudge(t *testing.T) {
	genA := newGen(factorGenA, testInputGenA)
	genB := newGen(factorGenB, testInputGenB)
	assert.Equal(t, 1, judge(genA, genB, 5))

	genA = newGen(factorGenA, testInputGenA)
	genB = newGen(factorGenB, testInputGenB)
	assert.Equal(t, 588, judge(genA, genB, 40000000))
}
