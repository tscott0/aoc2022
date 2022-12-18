package readers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GroupedInts reads a file containing one integer per line, in groups separated
// by empty lines.
//
// Example:
// 123
// 456
//
// 7
// 89
func GroupedInts(filename string) [][]int {
	output := [][]int{}

	readFile, err := os.Open("input/01.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	current := []int{}
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		if i, err := strconv.Atoi(line); err == nil {
			current = append(current, i)
		} else {
			output = append(output, current)
			current = []int{}
		}
	}

	if len(current) > 0 {
		output = append(output, current)
	}

	readFile.Close()

	return output
}
