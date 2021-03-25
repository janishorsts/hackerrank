package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permutationEquation(t *testing.T) {
	for _, v := range []struct{ p, out []int32 }{
		{[]int32{2, 3, 1}, []int32{2, 3, 1}},
		{[]int32{4, 3, 5, 1, 2}, []int32{1, 3, 5, 4, 2}},
	} {
		t.Run(fmt.Sprintf("%v", v.p), func(t *testing.T) {
			assert.Equal(t, v.out, permutationEquation(v.p))
		})
	}
}
