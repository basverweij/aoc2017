package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootShouldReturnErrorOnNoRoot(t *testing.T) {
	_, err := root([]*program{})
	assert.Equal(t, errNoRoot, err)

	_, err = root([]*program{
		&program{"a", 1, []string{"b"}},
		&program{"b", 2, []string{"a"}},
	})
	assert.Equal(t, errNoRoot, err)
}

func TestRootShouldReturnErrorOnMultipleRoots(t *testing.T) {
	_, err := root([]*program{
		&program{"a", 1, []string{"b"}},
		&program{"c", 3, []string{"d"}},
	})
	assert.Equal(t, errMultipleRoots, err)
}

func TestRootShouldReturnRoot(t *testing.T) {
	root, err := root([]*program{
		&program{"a", 1, []string{"b", "c"}},
		&program{"b", 2, nil},
		&program{"c", 3, nil},
	})
	assert.Nil(t, err)

	assert.Equal(t, &program{"a", 1, []string{"b", "c"}}, root)
}
