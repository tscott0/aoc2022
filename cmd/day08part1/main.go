package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tscott0/aoc2022/internal/readers"
)

func main() {
	fmt.Println("Day 8 part 1")

	answer := readers.ReadAll("input/08.txt")
	fmt.Printf("answer: %d\n", solve(answer))
}

func solve(input string) (total int) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	grid := [][]int{}
	for _, l := range lines {
		row := []int{}
		for _, c := range l {
			i, _ := strconv.Atoi(string(c))
			row = append(row, i)
		}

		grid = append(grid, row)
	}

	// fmt.Println(grid)
	for i := range grid {
		for j := range grid[i] {

			tree := grid[i][j]

			vUp, vDown, vLeft, vRight := true, true, true, true
			// up
			for i2 := i - 1; i2 >= 0; i2-- {
				if grid[i2][j] >= tree {
					vUp = false
					break
				}
			}

			// down
			for i2 := i + 1; i2 < len(grid); i2++ {
				if grid[i2][j] >= tree {
					vDown = false
					break
				}
			}

			// left
			for j2 := j - 1; j2 >= 0; j2-- {
				if grid[i][j2] >= tree {
					vLeft = false
					break
				}
			}

			// right
			for j2 := j + 1; j2 < len(grid[i]); j2++ {
				if grid[i][j2] >= tree {
					vRight = false
					break
				}
			}

			if vUp || vDown || vLeft || vRight {
				total++
			}
		}
	}

	return total
}
