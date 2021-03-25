package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_formingMagicSquare(t *testing.T) {
	for _, v := range []struct {
		in  [][]int32
		out int32
	}{{
		in:  [][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}},
		out: 7,
	}, {
		in:  [][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 5}},
		out: 1,
	}, {
		in:  [][]int32{{4, 8, 2}, {4, 5, 7}, {6, 1, 6}},
		out: 4,
	}} {
		t.Run(strconv.Itoa(int(v.out)), func(t *testing.T) {
			assert.Equal(t, v.out, formingMagicSquare(v.in))
		})
	}
}
