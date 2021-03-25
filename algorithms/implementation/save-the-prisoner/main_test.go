package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_saveThePrisoner(t *testing.T) {
	for _, v := range []struct{ n, m, s, out int32 }{
		{5, 2, 1, 2},
		{5, 2, 2, 3},
		{5, 2, 3, 4},
		{5, 2, 4, 5},
		{5, 2, 5, 1},
		{7, 19, 2, 6},
		{7, 19, 4, 1},
		{7, 19, 6, 3},
		{7, 19, 7, 4},
		{7, 20, 7, 5},
		{7, 21, 7, 6},
		{7, 22, 7, 7},
		{7, 23, 7, 1},
		{3, 7, 3, 3},
		{3, 7, 2, 2},
		{3, 7, 1, 1},
		{1, 1, 1, 1},
		{4, 3, 3, 1},
		{2, 2, 1, 2},
	} {
		t.Run(fmt.Sprintf("n%d-m%d-s%d", v.n, v.m, v.s), func(t *testing.T) {
			assert.Equal(t, v.out, saveThePrisoner(v.n, v.m, v.s))
		})
	}
}
