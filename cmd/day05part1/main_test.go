package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	s := solve(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)

	assert.Equal(t, "CMZ", s)
}
