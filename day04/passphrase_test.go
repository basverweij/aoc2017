package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassphrase(t *testing.T) {
	assert.Equal(t, passphrase{"aa", "bb", "cc", "dd", "ee"}, newPassphrase("aa bb cc dd ee"))
	assert.Equal(t, passphrase{"aa", "bb", "cc", "dd", "aa"}, newPassphrase("aa bb cc dd aa"))
	assert.Equal(t, passphrase{"aa", "bb", "cc", "dd", "aaa"}, newPassphrase("aa bb cc dd aaa"))
}
