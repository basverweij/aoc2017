package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	is := input()
	assert.NotNil(t, is)

	fmt.Println(is)
}
