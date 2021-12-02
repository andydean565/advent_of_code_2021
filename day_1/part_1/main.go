package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var count int = 0

	for i, line := range lines {
		// fmt.Println(i, line)

		if i == 0 || line > lines[(i-1)] {
			count++
			fmt.Println(i, line, " (increased)")
		} else {
			fmt.Println(i, line, " (decreased)")

		}
	}

	fmt.Println("result:", count)
}
