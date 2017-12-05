package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAll(t *testing.T) {
	assert.Equal(t, 2, validateAll([]passphrase{
		passphrase{"aa", "bb", "cc", "dd", "ee"},
		passphrase{"aa", "bb", "cc", "dd", "aa"},
		passphrase{"aa", "bb", "cc", "dd", "aaa"},
	}))
}

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, isValid(passphrase{"aa", "bb", "cc", "dd", "ee"}))
	assert.Equal(t, false, isValid(passphrase{"aa", "bb", "cc", "dd", "aa"}))
	assert.Equal(t, true, isValid(passphrase{"aa", "bb", "cc", "dd", "aaa"}))
}
