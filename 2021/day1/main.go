package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Day 1 challenge 1 takes in an input file that has one int per line
// The goal is to return the number of times the next value increases
// Using `small-input.txt`, the value is 7.

// Day 1 challenge 2 uses that same input file as challenge 1
// The goal is to return the number of times a group of 3 lines increases
// It also must stop when there aren't enough measurements left to create
// a new three-measurement sum.
// Using `small-input.txt`, the value is 5.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	value1 := 0
	value2 := 0
	value3 := 0
	currentSum := 0
	previousSum := -1
	increaseCount := 0

	for scanner.Scan() {
		value1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Failed to convert string to int", err)
		}
		if value1 > 0 && value2 > 0 && value3 > 0 {
			currentSum = value1 + value2 + value3
			if previousSum != -1 {
				if currentSum > previousSum {
					increaseCount = increaseCount + 1
				}
			}
			previousSum = currentSum
		}
		value3 = value2
		value2 = value1
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan input:", err)
		os.Exit(1)
	}

	fmt.Printf("The value increased %d times\n", increaseCount)
}
