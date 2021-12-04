package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Day 2 challenge 1 takes in an input file that has a direction followed by an int
// The goal is to find the final horizontal and depth location, as well as return the
// value of those two multiplied together.
// Using `small-input.txt`, horizontal = 15, depth = 10, total = 150.

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	var direction string
	var quantity int
	depth := 0
	position := 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		direction = line[0]
		quantity, err = strconv.Atoi(line[1])
		if err != nil {
			fmt.Println("Failed to convert quantity to int", err)
			os.Exit(1)
		}
		switch direction {
		case "up":
			depth = depth - quantity
		case "down":
			depth = depth + quantity
		case "forward":
			position = position + quantity
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan input:", err)
		os.Exit(1)
	}

	fmt.Printf("Our position is %d and our depth is %d which equals %d\n", position, depth, (depth * position))
}
