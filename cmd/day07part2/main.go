package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/tscott0/aoc2022/internal/readers"
)

var lsRegex = regexp.MustCompile(`^(\d+) .*`)

func main() {
	fmt.Println("Day 7 part 2")

	answer := readers.ReadAll("input/07.txt")
	fmt.Printf("answer: %d\n", solve(answer))
}

func solve(input string) int {
	lines := strings.Split(input, "\n")

	dirSizeMap := map[string]int{}
	directories := []string{}
	for _, l := range lines {
		if len(l) == 0 {
			break
		}
		fmt.Println(l)

		if strings.HasPrefix(l, "$ cd ") {
			dir := strings.TrimPrefix(l, "$ cd ")

			if dir == ".." {
				directories = directories[:len(directories)-1]
			} else {
				directories = append(directories, dir)
			}
			fmt.Println(directories)
		} else if lsRegex.MatchString(l) {
			matches := lsRegex.FindStringSubmatch(l)
			if len(matches) != 2 {
				log.Fatal("expected 2 matches")
			}

			size, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(err)
			}

			for i := range directories {
				current := strings.Join(directories[:i+1], "/")
				dirSizeMap[current] += size
			}

		}
	}

	diskSize := 70000000
	required := 30000000
	used := dirSizeMap["/"]
	currentFree := diskSize - used
	toFree := required - currentFree

	// find smallest dir that will free enough space if deleted
	smallest := math.MaxInt
	for _, s := range dirSizeMap {
		if s >= toFree && s < smallest {
			smallest = s
		}
	}

	return smallest
}
