package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRuntime(t *testing.T) {
	p := make(map[string]*program)
	p["a"] = &program{"a", 1, nil}
	p["b"] = &program{"b", 2, []string{"c", "d", "e"}}

	assert.Equal(t, &runtime{
		programs:     p,
		totalWeights: make(map[string]int),
	}, newRuntime([]*program{
		&program{"a", 1, nil},
		&program{"b", 2, []string{"c", "d", "e"}},
	}))
}

func TestTotalWeight(t *testing.T) {
	r := newRuntime([]*program{
		&program{"a", 1, nil},
		&program{"b", 2, []string{"c", "d"}},
		&program{"c", 3, nil},
		&program{"d", 4, nil},
	})

	assert.Equal(t, 1, r.totalWeight("a"))
	assert.Equal(t, 9, r.totalWeight("b"))
}
