package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Day 1 challenge 1 takes in an input file that has one int per line
// The goal is to return the number of times the next value increases
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var current int
	previous := -1
	increaseCount := 0

	for scanner.Scan() {
		current, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Failed to convert string to int", err)
		}
		if previous != -1 {
			if current > previous {
				increaseCount = increaseCount + 1
			}
		}

		previous = current
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan input:", err)
		os.Exit(1)
	}

	fmt.Printf("The value increased %d times\n", increaseCount)
}
