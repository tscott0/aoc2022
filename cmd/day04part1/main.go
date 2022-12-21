package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/tscott0/aoc2022/internal/readers"
)

// 2-4,6-8
var lineRegex = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func main() {
	fmt.Println("Day 4 part 1")

	answer := readers.ReadAll("input/04.txt")
	fmt.Printf("answer: %d\n", solve(answer))
}

func solve(input string) (total int) {
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		if len(l) == 0 {
			break
		}

		//2-4,6-8
		matches := lineRegex.FindStringSubmatch(l)
		if len(matches) != 5 {
			log.Fatal("expected 5 matches")
		}

		a1, _ := strconv.Atoi(matches[1])
		a2, _ := strconv.Atoi(matches[2])
		b1, _ := strconv.Atoi(matches[3])
		b2, _ := strconv.Atoi(matches[4])

		if isContained(a1, a2, b1, b2) || isContained(b1, b2, a1, a2) {
			total++
		}
	}

	return total
}

func isContained(x1, x2, y1, y2 int) bool {
	return x1 >= y1 && x2 <= y2
}
