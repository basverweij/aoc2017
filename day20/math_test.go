package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQz(t *testing.T) {
	// no roots
	_, _, err := qroots(1, -3, 4)
	assert.Equal(t, errNoRoots, err)

	// one root
	t0, t1, err := qroots(-4, 12, -9)
	assert.Nil(t, err)
	assert.Equal(t, 1.5, t0)
	assert.Equal(t, 1.5, t1)

	// two roots
	t0, t1, err = qroots(2, -11, 5)
	assert.Nil(t, err)
	assert.Equal(t, .5, t0)
	assert.Equal(t, 5.0, t1)
}
