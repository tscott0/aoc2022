package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	assert.Equal(t, 5, solve("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 6, solve("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 10, solve("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 11, solve("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}
