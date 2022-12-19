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
		score += 1
		switch i {
		case a:
			score += 3 // DRAW
		case b:
			score += 0 // LOSE
		case c:
			score += 6 // WIN
		}
	case y:
		score += 2
		switch i {
		case a:
			score += 6 // WIN
		case b:
			score += 3 // DRAW
		case c:
			score += 0 // LOSE
		}
	case z:
		score += 3
		switch i {
		case a:
			score += 0 // LOSE
		case b:
			score += 6 // WIN
		case c:
			score += 3 // DRAW
		}
	}
	return score
}
