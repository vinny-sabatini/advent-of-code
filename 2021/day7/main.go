package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Day 7 challenge 1 takes in a list of integers representing the horizontal
// position of crabs. The goal is to determine the position all the crabs should
// be locationed to spend the least amount of gas assuming moving one crab one
// spot costs one unit of gas. For example, using small-input.txt, it is cheapest
// to move the crabs to position 2, and it would cost 37 units of gas.
//
// Day 7 part 2 does the same thing, however moving one position costs 1 fuel,
// moving 2 positions costs 3 (1 + 2), moving 3 positions costs 6 (1+2+3), etc.
// For example, using small-input.txt, it is cheapest to move the crabs to location
// 5, and it costs 168 fuel.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	tracker := make(map[string]int)
	cost := make(map[string]int)
	maxPosition := 0

	scanner.Scan()
	inputText := strings.Split(scanner.Text(), ",")

	for _, v := range inputText {
		if _, ok := tracker[v]; ok {
			tracker[v] = tracker[v] + 1
		} else {
			tracker[v] = 1
		}
		vInt, _ := strconv.Atoi(v)
		if vInt > maxPosition {
			maxPosition = vInt
		}
	}

	for positionStart := 0; positionStart < maxPosition; positionStart++ {
		positionStartString := strconv.Itoa(positionStart)
		cost[positionStartString] = 0
		for positionEnd, countEnd := range tracker {
			positionEndInt, _ := strconv.Atoi(positionEnd)
			distance := int(math.Abs(float64(positionStart - positionEndInt)))
			fuelUsed := calculateFuelUsed(distance)
			currentCost := fuelUsed * countEnd
			cost[positionStartString] = cost[positionStartString] + currentCost
		}
	}

	//for i, v := range cost {
	//	fmt.Println("Position", i, "Cost", v)
	//}

	cheapestCost := 0
	cheapestPosition := ""
	for i, v := range cost {
		if cheapestCost == 0 || cheapestCost > v {
			cheapestCost = v
			cheapestPosition = i
		}
	}
	fmt.Println("Position", cheapestPosition, "Cost", cheapestCost)
}

func calculateFuelUsed(distance int) int {
	total := 0
	for i := 0; i < distance+1; i++ {
		total += i
	}
	return total
}
