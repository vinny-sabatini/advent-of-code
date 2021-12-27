package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type point struct {
	x     int
	y     int
	value int
}

// Day 9 challenge 1 takes in lines of numbers, and the goal is to find the low points.
// A low point is when all of the directly surrounding numbers (not including diagonals)
// are larger than a given number. If the number is on the edge, you do not have to worry
// about the "missing" edges. With those low points, the "risk score" is the number plus 1.
// For example, in small-input.txt, the low points are the 1 and 0 in the first row, the 5
// in the third row, and the 5 in the last row. With that, the risk score is 15.
//
// Day 9 challenge 2 takes all of the low points and finds the "basins" that go with them.
// A basin is all of the points that are larger than the low points that are smaller than 9.
// For example, in small-input.txt, the top-left basin has size 3, the top-right basin has size 9
// the middle basin has size 14, and the bottom-right basin has size 9. The final answer is
// the top 3 sized basins multiplied together, so 9 * 14 * 9 = 1134 for small-input.txt.
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
	var lowPoints []point
	for y, line := range lines {
		for x, character := range line {
			adjacentRows := []int{10, 10, 10, 10}
			current, _ := strconv.Atoi(string(character))
			if x != 0 {
				adjacentRows[0], _ = strconv.Atoi(string(lines[y][x-1]))
			}
			if x != rowLength-1 {
				adjacentRows[1], _ = strconv.Atoi(string(lines[y][x+1]))
			}
			if y != 0 {
				adjacentRows[2], _ = strconv.Atoi(string(lines[y-1][x]))
			}
			if y != rowCount-1 {
				adjacentRows[3], _ = strconv.Atoi(string(lines[y+1][x]))
			}
			isLowPoint := true
			for _, v := range adjacentRows {
				if math.Min(float64(current), float64(v)) == float64(v) {
					isLowPoint = false
					break
				}
			}
			if isLowPoint {
				lowPoints = append(lowPoints, point{x: x, y: y, value: current})
			}
		}
	}

	var basin []point
	var basinLengths []int
	for _, p := range lowPoints {
		basin = make([]point, 0)
		basin = append(basin, getBasin(p, lines, basin)...)
		fmt.Printf("LowPoint %d, %d w/ value %d has basin size %d\n", p.x, p.y, p.value, len(basin))
		basinLengths = append(basinLengths, len(basin))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinLengths)))
	fmt.Println(basinLengths)
	total := basinLengths[0] * basinLengths[1] * basinLengths[2]
	fmt.Println(total)

}

func getBasin(p point, lines []string, basin []point) []point {
	basin = append(basin, p)
	rowLength := len(lines[0])
	rowCount := len(lines)
	if p.x != 0 {
		leftValue, _ := strconv.Atoi(string(lines[p.y][p.x-1]))
		leftPoint := point{x: p.x - 1, y: p.y, value: leftValue}
		if leftValue > p.value && leftValue != 9 && !checkPointExists(leftPoint, basin) {
			basin = append(basin, getBasin(leftPoint, lines, basin)...)
		}
	}
	if p.x != rowLength-1 {
		rightValue, _ := strconv.Atoi(string(lines[p.y][p.x+1]))
		rightPoint := point{x: p.x + 1, y: p.y, value: rightValue}
		if rightValue > p.value && rightValue != 9 && !checkPointExists(rightPoint, basin) {
			basin = append(basin, getBasin(rightPoint, lines, basin)...)
		}
	}
	if p.y != 0 {
		upValue, _ := strconv.Atoi(string(lines[p.y-1][p.x]))
		upPoint := point{x: p.x, y: p.y - 1, value: upValue}
		if upValue > p.value && upValue != 9 && !checkPointExists(upPoint, basin) {
			basin = append(basin, getBasin(upPoint, lines, basin)...)
		}
	}
	if p.y != rowCount-1 {
		downValue, _ := strconv.Atoi(string(lines[p.y+1][p.x]))
		downPoint := point{x: p.x, y: p.y + 1, value: downValue}
		if downValue > p.value && downValue != 9 && !checkPointExists(downPoint, basin) {
			basin = append(basin, getBasin(downPoint, lines, basin)...)
		}
	}
	return uniqueBasin(basin)
}

func checkPointExists(p point, basin []point) bool {
	for _, v := range basin {
		if v.x == p.x && v.y == p.y {
			return true
		}
	}
	return false
}

func uniqueBasin(basin []point) []point {
	var b []point
	for _, p := range basin {
		if !checkPointExists(p, b) {
			b = append(b, p)
		}
	}
	return b
}
