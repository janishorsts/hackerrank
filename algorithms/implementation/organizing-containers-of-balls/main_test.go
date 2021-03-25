package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_organizingContainers(t *testing.T) {
	assert.Equal(t, "Impossible", organizingContainers([][]int32{
		{1, 4},
		{2, 3}}))
	assert.Equal(t, "Possible", organizingContainers([][]int32{
		{1, 1},
		{1, 1}}))
	assert.Equal(t, "Impossible", organizingContainers([][]int32{
		{0, 2},
		{1, 1}}))
	assert.Equal(t, "Impossible", organizingContainers([][]int32{
		{1, 3, 1},
		{2, 1, 2},
		{3, 3, 3}}))
	assert.Equal(t, "Possible", organizingContainers([][]int32{
		{0, 2, 1},
		{1, 1, 1},
		{2, 0, 0}}))
	assert.Equal(t, "Possible", organizingContainers([][]int32{
		{1, 0},
		{0, 1}}))
	assert.Equal(t, "Possible", organizingContainers([][]int32{
		{1, 2},
		{2, 1}}))
	assert.Equal(t, "Possible", organizingContainers([][]int32{
		{0, 4, 0},
		{2, 0, 1},
		{1, 0, 2}}))
	assert.Equal(t, "Possible", organizingContainers([][]int32{
		{1, 2, 3, 4},
		{2, 1, 4, 3},
		{3, 4, 2, 1},
		{4, 3, 1, 2}}))
	assert.Equal(t, "Possible", organizingContainers([][]int32{
		{0, 0, 5, 0},
		{4, 0, 0, 0},
		{0, 2, 0, 1},
		{0, 1, 0, 2}}))
	assert.Equal(t, "Impossible", organizingContainers([][]int32{
		{2, 1},
		{0, 0}}))
	assert.Equal(t, "Impossible", organizingContainers([][]int32{
		{2, 1},
		{0, 1}}))
	assert.Equal(t, "Impossible", organizingContainers([][]int32{
		{1, 2, 3},
		{3, 2, 1},
		{2, 3, 1}}))
	assert.Equal(t, "Impossible", organizingContainers([][]int32{
		{1, 0, 0},
		{0, 2, 0},
		{0, 2, 0}}))
	assert.Equal(t, "Impossible", organizingContainers([][]int32{
		{1, 2, 1, 3},
		{2, 1, 3, 1},
		{1, 3, 2, 1},
		{3, 2, 1, 1}}))
}
