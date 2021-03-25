package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jumpingOnClouds(t *testing.T) {
	for _, v := range []struct {
		c      []int32
		k, out int32
	}{
		{c: []int32{0, 0, 1, 0, 0, 1, 1, 0}, k: 2, out: 92},
		{c: []int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, k: 3, out: 80},
	} {
		t.Run(strconv.Itoa(int(v.k)), func(t *testing.T) {
			assert.Equal(t, v.out, jumpingOnClouds(v.c, v.k))
		})
	}
}
