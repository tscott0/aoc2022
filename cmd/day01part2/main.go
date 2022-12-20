package main

import (
	"fmt"
	"sort"

	"github.com/tscott0/aoc2022/internal/readers"
)

func main() {
	elves := readers.GroupedInts("input/01.txt")
	totals := make([]int, len(elves))

	for i, e := range elves {
		for _, c := range e {
			totals[i] += c
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	answer := totals[0] + totals[1] + totals[2]
	fmt.Printf("%d calories", answer)
}
