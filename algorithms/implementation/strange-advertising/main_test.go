package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_viralAdvertising(t *testing.T) {
	for _, v := range []struct{ in, out int32 }{
		{1, 2},
		{2, 5},
		{3, 9},
		{4, 15},
		{5, 24},
	} {
		t.Run(strconv.Itoa(int(v.in)), func(t *testing.T) {
			assert.Equal(t, v.out, viralAdvertising(v.in))
		})
	}
}
