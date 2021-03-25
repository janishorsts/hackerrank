package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_appendAndDelete(t *testing.T) {
	for _, v := range []struct {
		s, t, expected string
		k              int32
	}{
		{s: "hackerhappy", t: "hackerrank", k: 9, expected: "Yes"},
		{s: "hackerhappy", t: "hackerhappy", k: 0, expected: "Yes"},
		{s: "hackerhapps", t: "hackerrank", k: 8, expected: "No"},
		{s: "ashley", t: "ash", k: 2, expected: "No"},
		{s: "ash", t: "ashley", k: 3, expected: "Yes"},
		{s: "aba", t: "aba", k: 7, expected: "Yes"},
		{s: "aba", t: "abas", k: 8, expected: "Yes"},
		{s: "", t: "aba", k: 8, expected: "Yes"},
		{s: "", t: "happy", k: 5, expected: "Yes"},
		{s: "", t: "happy", k: 4, expected: "No"},
		{s: "happy", t: "", k: 5, expected: "Yes"},
		{s: "happy", t: "", k: 4, expected: "No"},
		{s: "happy", t: "happ", k: 1, expected: "Yes"},
		{s: "happ", t: "happy", k: 1, expected: "Yes"},
	} {
		t.Run(v.s, func(t *testing.T) {
			assert.Equal(t, v.expected, appendAndDelete(v.s, v.t, v.k))
		})
	}
}
