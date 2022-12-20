package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	s := solve(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

	assert.Equal(t, 157, s)
}

func TestScoreRune(t *testing.T) {
	assert.Equal(t, 16, scoreRune('p'))
	assert.Equal(t, 38, scoreRune('L'))
	assert.Equal(t, 42, scoreRune('P'))
	assert.Equal(t, 22, scoreRune('v'))
	assert.Equal(t, 20, scoreRune('t'))
	assert.Equal(t, 19, scoreRune('s'))
}
