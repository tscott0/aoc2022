package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	assert.Equal(t, 19, solve("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(t, 23, solve("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 23, solve("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 29, solve("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 26, solve("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}
