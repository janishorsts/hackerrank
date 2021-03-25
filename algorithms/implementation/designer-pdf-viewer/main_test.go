package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_designerPdfViewer(t *testing.T) {
	for _, v := range []struct {
		h    []int32
		word string
		out  int32
	}{{
		h:    []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		word: "abc",
		out:  9,
	}, {
		h:    []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7},
		word: "zaba",
		out:  28,
	}} {
		t.Run(v.word, func(t *testing.T) {
			assert.Equal(t, v.out, designerPdfViewer(v.h, v.word))
		})
	}
}
