package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagrams(t *testing.T) {
	assert.EqualValues(t,
		[]string{"abc", "bac", "cab", "acb", "bca", "cba"},
		anagrams("abc"))
}
