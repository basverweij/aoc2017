package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDance(t *testing.T) {
	l := dance(5, []move{
		newSpin(1),
		newExchange(3, 4),
		newPartner('e', 'b'),
	})

	assert.Equal(t, []byte{'b', 'a', 'e', 'd', 'c'}, l.d)
}
