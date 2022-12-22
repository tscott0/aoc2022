package main

import (
	"fmt"
	"strings"

	"github.com/tscott0/aoc2022/internal/readers"
)

func main() {
	fmt.Println("Day 6 part 1")

	answer := readers.ReadAll("input/06.txt")
	fmt.Printf("answer: %d\n", solve(answer))
}

func solve(input string) (total int) {
	lines := strings.Split(input, "\n")
	l := lines[0]

	for i, _ := range l {
		if allUnique(l[i : i+4]) {
			return i + 4
		}
	}

	return total
}

func allUnique(runes string) bool {
	for i, a := range runes {
		for j, b := range runes {
			if i == j {
				continue
			}

			if a == b {
				return false
			}
		}
	}
	return true
}
