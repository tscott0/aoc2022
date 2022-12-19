package readers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// TwoRunes reads each line expecting a rune, followed by a space, followed by another rune
func TwoRunes(filename string) [][]rune {
	output := [][]rune{}

	readFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatal("expected lines to have two parts")
		}

		if len(parts[0]) != 1 || len(parts[1]) != 1 {
			log.Fatal("expected parts to have a single rune")
		}

		output = append(output, []rune{rune(parts[0][0]), rune(parts[1][0])})
	}

	readFile.Close()

	return output
}
