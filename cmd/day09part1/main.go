package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tscott0/aoc2022/internal/readers"
)

func main() {
	fmt.Println("Day 9 part 1")

	answer := readers.ReadAll("input/09.txt")
	fmt.Printf("answer: %d\n", solve(answer))
}

func solve(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	head, tail := coord{}, coord{}
	visited := map[coord]struct{}{}
	for _, l := range lines {
		fmt.Println(l)

		parts := strings.Split(l, " ")
		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])

		switch direction {
		case "U":
			up(&head, &tail, steps, visited)
		case "D":
			down(&head, &tail, steps, visited)
		case "L":
			left(&head, &tail, steps, visited)
		case "R":
			right(&head, &tail, steps, visited)
		}
	}

	return len(visited)
}

type coord struct {
	x, y int
}

func up(head, tail *coord, steps int, visited map[coord]struct{}) {
	for i := 0; i < steps; i++ {
		head.y++
		moveTail(head, tail, visited)
	}
}

func down(head, tail *coord, steps int, visited map[coord]struct{}) {
	for i := 0; i < steps; i++ {
		head.y--
		moveTail(head, tail, visited)
	}
}

func left(head, tail *coord, steps int, visited map[coord]struct{}) {
	for i := 0; i < steps; i++ {
		head.x--
		moveTail(head, tail, visited)
	}
}

func right(head, tail *coord, steps int, visited map[coord]struct{}) {
	for i := 0; i < steps; i++ {
		head.x++
		moveTail(head, tail, visited)
	}
}

func moveTail(head, tail *coord, visited map[coord]struct{}) {
	if head.y > tail.y+1 {
		tail.y++
		tail.x = head.x
	} else if head.y < tail.y-1 {
		tail.y--
		tail.x = head.x
	} else if head.x < tail.x-1 {
		tail.x--
		tail.y = head.y
	} else if head.x > tail.x+1 {
		tail.x++
		tail.y = head.y
	}
	visited[*tail] = struct{}{}
}
