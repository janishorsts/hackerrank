package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_matchingStrings(t *testing.T) {
	for i, test := range []struct {
		strings, queries []string
		expected         []int32
	}{{
		strings:  []string{"ab", "ab", "abc"},
		queries:  []string{"ab", "abc", "bc"},
		expected: []int32{2, 1, 0},
	}, {
		strings:  []string{"def", "de", "fgh"},
		queries:  []string{"de", "lmn", "fgh"},
		expected: []int32{1, 0, 1},
	}, {
		strings:  []string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"},
		queries:  []string{"abcde", "sdaklfj", "asdjf", "na", "basdn"},
		expected: []int32{1, 3, 4, 3, 2},
	}} {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, test.expected, matchingStrings(test.strings, test.queries))
		})
	}
}
