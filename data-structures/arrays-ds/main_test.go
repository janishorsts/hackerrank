package main

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

func Test_reverseArray(t *testing.T) {
	assert.NoError(t, quick.Check(func(v []int32) bool {
		original := make([]int32, len(v))
		copy(original, v)

		reverseArray(v)

		for i := range original {
			if original[i] != v[len(v)-1-i] {
				return false
			}
		}

		return true
	}, nil))
}
