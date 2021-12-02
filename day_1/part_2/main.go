package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var count int = 0
	var length int = len(lines)
	var prev int = 0

	for i, line := range lines {
		// TODO A
		if i >= (length - 2) {
			continue
		}
		if prev != 0 && (line+lines[i+1]+lines[i+2] > prev) {
			count++
		}
		prev = line + lines[i+1] + lines[i+2]
	}
	fmt.Println("result:", count)
}
