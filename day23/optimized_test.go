package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptimized(t *testing.T) {
	h := optimized()

	fmt.Println(h)

	assert.True(t, h > 887)
	assert.True(t, h < 999)
}

func TestGenPrimes(t *testing.T) {
	primes := genPrimes(11)

	assert.EqualValues(t, map[int]struct{}{
		2:  struct{}{},
		3:  struct{}{},
		5:  struct{}{},
		7:  struct{}{},
		11: struct{}{},
	}, primes)
}
