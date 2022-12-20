package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

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

	boilerplate := fmt.Sprintf(`package main

import "fmt"

func main() {
	fmt.Println("Day %d part %d")
}
`, *day, *part)

	if err := os.WriteFile(dir+"/main.go", []byte(boilerplate), 0644); err != nil {
		log.Fatal(err)
	}
}
