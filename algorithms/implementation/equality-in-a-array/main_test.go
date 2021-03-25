package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_equalizeArray(t *testing.T) {
	assert.Equal(t, int32(2), equalizeArray([]int32{1, 2, 2, 3}))
	assert.Equal(t, int32(2), equalizeArray([]int32{3, 3, 2, 1, 3}))
}
