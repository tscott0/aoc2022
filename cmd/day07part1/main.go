package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/tscott0/aoc2022/internal/readers"
)

var lsRegex = regexp.MustCompile(`^(\d+) .*`)

func main() {
	fmt.Println("Day 7 part 1")

	answer := readers.ReadAll("input/07.txt")
	fmt.Printf("answer: %d\n", solve(answer))
}

func solve(input string) (total int) {
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
				// pop
				directories = directories[:len(directories)-1]
			} else {
				// push
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

	for _, s := range dirSizeMap {
		if s <= 100000 {
			total += s
		}
	}

	return total
}
