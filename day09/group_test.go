package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGroup(t *testing.T) {
	assert.NotNil(t, newGroup(nil))
}

func TestGroupString(t *testing.T) {
	g := newGroup(nil)
	assert.Equal(t, "{}", g.String())

	g.children = append(g.children, newGroup(g))
	g.children = append(g.children, newGroup(g))
	assert.Equal(t, "{{},{}}", g.String())
}
