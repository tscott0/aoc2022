package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	s := solve(`30373
25512
65332
33549
35390`)

	assert.Equal(t, 8, s)
}
