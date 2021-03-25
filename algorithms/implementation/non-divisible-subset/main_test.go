package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nonDivisibleSubset(t *testing.T) {
	for _, test := range []struct {
		k, result int32
		s         []int32
	}{{
		k:      3,
		s:      []int32{1, 7, 2, 4},
		result: 3,
	}, {
		k:      4,
		s:      []int32{19, 10, 12, 10, 24, 25, 22},
		result: 3,
	}, {
		k:      7,
		s:      []int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436},
		result: 11,
	}, {
		k:      9,
		s:      []int32{422346306, 940894801, 696810740, 862741861, 85835055, 313720373},
		result: 5,
	}} {
		t.Run(fmt.Sprintf("k: %d, r: %d", test.k, test.result), func(t *testing.T) {
			assert.Equal(t, test.result, nonDivisibleSubset(test.k, test.s))
		})
	}
}
