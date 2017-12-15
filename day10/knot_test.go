package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnot(t *testing.T) {
	l := knot(5, []int{3, 4, 1, 5})

	assert.Equal(t, 3, l[0])
	assert.Equal(t, 4, l[1])
}
