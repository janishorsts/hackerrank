package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hurdleRace(t *testing.T) {
	for _, v := range []struct {
		k, out  int32
		heights []int32
	}{
		{k: 4, heights: []int32{1, 6, 3, 5, 2}, out: 2},
		{k: 7, heights: []int32{2, 5, 4, 5, 2}},
	} {
		t.Run(strconv.Itoa(int(v.k)), func(t *testing.T) {
			assert.Equal(t, v.out, hurdleRace(v.k, v.heights))
		})
	}
}
