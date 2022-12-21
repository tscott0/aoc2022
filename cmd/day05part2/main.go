package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/tscott0/aoc2022/internal/readers"
)

// move 1 from 2 to 1
var moveRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func main() {
	fmt.Println("Day 5 part 1")

	answer := readers.ReadAll("input/05.txt")
	fmt.Printf("answer: %v\n", solve(answer))
}

func solve(input string) (answer string) {
	lines := strings.Split(input, "\n")

	var stacks [][]rune
	for i, l := range lines {
		if i == 0 {
			numStacks := (len(l) + 1) / 4
			fmt.Printf("line 0 has length %d, assuming %d stacks\n", len(l), numStacks)
			stacks = make([][]rune, numStacks)
		}
		fmt.Println()
		fmt.Println(l)

		if len(l) > 0 && strings.Contains(l, "[") {
			for stackIdx, r := range stackRow(l) {
				stacks[stackIdx] = append(stacks[stackIdx], r)
			}
		} else if strings.HasPrefix(l, "move") {
			matches := moveRegex.FindStringSubmatch(l)
			if len(matches) != 4 {
				log.Fatal("expected 4 matches")
			}

			count, _ := strconv.Atoi(matches[1])
			from, _ := strconv.Atoi(matches[2])
			to, _ := strconv.Atoi(matches[3])

			fromIdx := from - 1
			toIdx := to - 1

			pretty(stacks)

			// create temporary slice to avoid problems with overlapping slices
			moving := make([]rune, count)
			for i := 0; i < count; i++ {
				moving[i] = stacks[fromIdx][i]
			}

			stacks[toIdx] = append(moving, stacks[toIdx]...)
			pretty(stacks)

			if len(stacks[fromIdx]) == count {
				stacks[fromIdx] = []rune{}
			} else {
				stacks[fromIdx] = stacks[fromIdx][count:]
			}

			pretty(stacks)

		}
	}

	for _, s := range stacks {
		answer += string(s[0])
	}

	return answer
}

// "    [D]    " returns map[int]rune{ 1: 'D' }
func stackRow(r string) map[int]rune {
	row := map[int]rune{}
	for i, c := range r {
		if (i-1)%4 == 0 {
			index := (i - 1) / 4
			if c != ' ' {
				row[index] = c
			}
		}
	}

	return row
}

func pretty(stacks [][]rune) {
	for _, s := range stacks {
		fmt.Print("[")
		for _, c := range s {
			fmt.Print(string(c))
		}
		fmt.Print("]")
	}

	fmt.Println()
}
