package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repeatedString(t *testing.T) {
	assert.Equal(t, int64(4), repeatedString("abcac", 10))
	assert.Equal(t, int64(7), repeatedString("aba", 10))
	assert.Equal(t, int64(1000000000000), repeatedString("a", 1000000000000))
	assert.Equal(t, int64(1), repeatedString("abcac", 1))
}
