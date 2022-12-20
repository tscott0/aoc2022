package main

import (
	"fmt"

	"github.com/tscott0/aoc2022/internal/readers"
)

func main() {
	elves := readers.GroupedInts("input/01.txt")

	largest := 0
	for _, e := range elves {
		sum := 0
		for _, c := range e {
			sum += c
		}

		if sum > largest {
			largest = sum
		}
	}

	fmt.Printf("%d calories", largest)
}
