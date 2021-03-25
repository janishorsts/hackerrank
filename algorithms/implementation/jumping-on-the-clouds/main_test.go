package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jumpingOnClouds(t *testing.T) {
	assert.Equal(t, int32(4), jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
	assert.Equal(t, int32(3), jumpingOnClouds([]int32{0, 0, 0, 0, 1, 0}))
}
