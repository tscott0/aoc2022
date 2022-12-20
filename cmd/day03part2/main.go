package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/tscott0/aoc2022/internal/readers"
)

func main() {
	fmt.Println("Day 3 part 2")

	i := readers.ReadAll("input/03.txt")
	fmt.Printf("sum of priorities: %d\n", solve(i))
}

func solve(input string) (total int) {
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i += 3 {
		if len(lines[i]) == 0 {
			break
		}
		total += scoreRune(commonRune(lines[i], lines[i+1], lines[i+2]))
	}

	return total
}

func commonRune(a, b, c string) rune {
	for _, i := range a {
		for _, j := range b {
			if i != j {
				continue
			}
			for _, k := range c {
				if j == k {
					return i
				}
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
