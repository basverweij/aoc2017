package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupSize(t *testing.T) {
	programs := map[int][]int{
		0: []int{2},
		1: []int{1},
		2: []int{0, 3, 4},
		3: []int{2, 4},
		4: []int{2, 3, 6},
		5: []int{6},
		6: []int{4, 5},
	}

	assert.Equal(t, 6, groupSize(programs, 0))
}
