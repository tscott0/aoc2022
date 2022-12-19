package main

import (
	"fmt"

	"github.com/tscott0/aoc2022/internal/readers"
)

const (
	a rune = 'A'
	b rune = 'B'
	c rune = 'C'
	x rune = 'X'
	y rune = 'Y'
	z rune = 'Z'
)

func main() {
	rounds := readers.TwoRunes("input/02.txt")

	total := 0
	for _, r := range rounds {
		total += scoreRound(r[0], r[1])
	}
	fmt.Printf("total: %d", total)
}

func scoreRound(i, j rune) (score int) {
	switch j {
	case x:
		score += 0 // LOSE
		switch i {
		case a:
			score += 3 // rock vs scissors
		case b:
			score += 1 // paper vs rock
		case c:
			score += 2 // scissors vs paper
		}
	case y:
		score += 3 // DRAW
		switch i {
		case a:
			score += 1 // rock vs rock
		case b:
			score += 2 // paper vs paper
		case c:
			score += 3 // scissors vs scissors
		}
	case z:
		score += 6 // WIN
		switch i {
		case a:
			score += 2 // rock vs paper
		case b:
			score += 3 // paper vs scissors
		case c:
			score += 1 // scissors vs rock
		}
	}
	return score
}
