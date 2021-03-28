package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_arrayManipulation(t *testing.T) {
	assert.Equal(t,
		int64(200),
		arrayManipulation(5, [][]int32{
			{1, 2, 100},
			{2, 5, 100},
			{3, 4, 100}}))

	assert.Equal(t,
		int64(10),
		arrayManipulation(10, [][]int32{
			{1, 5, 3},
			{4, 8, 7},
			{6, 9, 1}}))

	assert.Equal(t,
		int64(31),
		arrayManipulation(10, [][]int32{
			{2, 6, 8},
			{3, 5, 7},
			{1, 8, 1},
			{5, 9, 15},
		}))
}
