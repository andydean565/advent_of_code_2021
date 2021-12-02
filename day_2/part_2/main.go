package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	var horizontal int = 0
	var depth int = 0
	var aim int = 0

	for _, line := range lines {

		words := strings.Fields(line)
		move, err := strconv.Atoi(words[1])

		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		if words[0] == "forward" {
			fmt.Println("forward", move)
			horizontal += move
			depth += (aim * move)

		} else if words[0] == "down" {
			fmt.Println("down", move)
			aim += move
		} else if words[0] == "up" {
			fmt.Println("up", move)
			aim -= move
		}

	}

	fmt.Println("result:", (horizontal * depth))
}
