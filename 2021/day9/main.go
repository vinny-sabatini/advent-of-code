package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Day 9 challenge 1 takes in lines of numbers, and the goal is to find the low points.
// A low point is when all of the directly surrounding numbers (not including diagonals)
// are larger than a given number. If the number is on the edge, you do not have to worry
// about the "missing" edges. With those low points, the "risk score" is the number plus 1.
// For example, in small-input.txt, the low points are the 1 and 0 in the first row, the 5
// in the third row, and the 5 in the last row. With that, the risk score is 15.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	rowLength := len(lines[0])
	rowCount := len(lines)
	var lowPoints []int
	for i, line := range lines {
		for j, character := range line {
			adjacentRows := []int{10, 10, 10, 10}
			current, _ := strconv.Atoi(string(character))
			if j != 0 {
				adjacentRows[0], _ = strconv.Atoi(string(line[j-1]))
			}
			if j != rowLength-1 {
				adjacentRows[1], _ = strconv.Atoi(string(line[j+1]))
			}
			if i != 0 {
				adjacentRows[2], _ = strconv.Atoi(string(lines[i-1][j]))
			}
			if i != rowCount-1 {
				adjacentRows[3], _ = strconv.Atoi(string(lines[i+1][j]))
			}
			lowPoint := true
			for _, v := range adjacentRows {
				if math.Min(float64(current), float64(v)) == float64(v) {
					lowPoint = false
					break
				}
			}
			if lowPoint {
				lowPoints = append(lowPoints, current)
			}
		}
	}
	total := 0
	for _, v := range lowPoints {
		total += v + 1
	}
	fmt.Println(total)
}
