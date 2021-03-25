package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_beautifulDays(t *testing.T) {
	for _, v := range []struct{ i, j, k, out int32 }{
		{20, 23, 6, 2},
	} {
		t.Run(strconv.Itoa(int(v.k)), func(t *testing.T) {
			assert.Equal(t, v.out, beautifulDays(v.i, v.j, v.k))
		})
	}
}

func Benchmark_beautifulDays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		beautifulDays(20, 23, 6)
	}
}
