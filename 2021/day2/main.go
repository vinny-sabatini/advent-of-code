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

// Day 2 challenge 2 changes the way the inputs are intepreted from part 1
// up and down simply change the "aim" by the quantity
// forward will change position by the quantity still, but will also change depth by the quantity times aim
// Using `small-input.txt`, horizontal = 15, depth = 60, total = 900
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	var direction string
	var quantity int
	aim := 0
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
			aim = aim - quantity
		case "down":
			aim = aim + quantity
		case "forward":
			position = position + quantity
			depth = depth + (aim * quantity)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan input:", err)
		os.Exit(1)
	}

	fmt.Printf("Our position is %d and our depth is %d which equals %d\n", position, depth, (depth * position))
}
