package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	s := solve(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)

	assert.Equal(t, 13, s)
}
