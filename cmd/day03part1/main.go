package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/tscott0/aoc2022/internal/readers"
)

func main() {
	fmt.Println("Day 3 part 1")

	i := readers.ReadAll("input/03.txt")
	fmt.Printf("sum of priorities: %d\n", solve(i))
}

func solve(input string) (total int) {
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		if len(l) == 0 {
			break
		}

		mid := len(l) / 2

		c := commonRune(l[:mid], l[mid:])
		total += scoreRune(c)
	}

	return total
}

func commonRune(a, b string) rune {
	for _, i := range a {
		for _, j := range b {
			if i == j {
				return i
			}
		}
	}

	log.Fatal("failed to find common rune")
	return '0'
}

func scoreRune(r rune) int {
	answer := int(r) - 96   // a to z (97 to 122)
	if r >= 65 && r <= 90 { // A to Z (65 to 90)
		answer = int(r) - 38
	}

	// fmt.Printf("%v = %d (%d)\n", string(r), answer, r)
	return answer
}
