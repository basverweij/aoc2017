package main

import (
	"bytes"
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

var input = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func TestUnbalanced(t *testing.T) {
	programs, err := parse(bytes.NewBufferString(input))
	assert.Nil(t, err)

	unb, w, err := unbalanced(programs)
	assert.Nil(t, err)

	assert.Equal(t, "ugml", unb.Name)
	assert.Equal(t, 60, w)
}
