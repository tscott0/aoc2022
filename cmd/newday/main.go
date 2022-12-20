package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var mainFile = `package main

import (
	"fmt"
	"strings"

	"github.com/tscott0/aoc2022/internal/readers"
)

func main() {
	fmt.Println("Day %d part %d")

	answer := readers.ReadAll("input/%02d.txt")
	fmt.Printf("answer: %%d\n", solve(answer))
}

func solve(input string) (total int) {
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		fmt.Println(l)
	}

	return total
}`

var mainTestFile = `package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	s := solve("")

	assert.Equal(t, 157, s)
}`

func main() {
	day := flag.Int("day", 0, "day number (1-25)")
	part := flag.Int("part", 0, "part number (1 or 2)")

	flag.Parse()

	if *day == 0 || *part == 0 {
		log.Fatal("day and part flags are required")
	}

	dir := fmt.Sprintf("cmd/day%02dpart%d", *day, *part)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	m := fmt.Sprintf(mainFile, *day, *part, *day)
	if err := os.WriteFile(dir+"/main.go", []byte(m), 0644); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(dir+"/main_test.go", []byte(mainTestFile), 0644); err != nil {
		log.Fatal(err)
	}
}
