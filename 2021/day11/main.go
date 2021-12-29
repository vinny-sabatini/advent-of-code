package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x, y int
}

// Day 11 challenge 1 takes in a 10x10 grid of numbers between 0 and 9. Every turn, all of the
// points increase by 1, then if any of the points are greater than 9 it "flashes", which results
// in all of the neighboring points increase by 1. After all of the "flashes" are done, one round
// is complete. The goal is to count how many flashes occur after 100 rounds. In small-input.txt,
// there are 1656 flashes.
//
// Day 11 challenge 2 has to find in which step do all points flash. In small-input.txt, this happens
// on step 195.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	grid := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		gridRow := make([]int, 0)
		for _, v := range line {
			num, _ := strconv.Atoi(string(v))
			gridRow = append(gridRow, num)
		}
		grid = append(grid, gridRow)
	}
	fmt.Println("START GRID")
	for _, v := range grid {
		fmt.Println(v)
	}

	totalFlashes := 0

	for i := 0; i < 1000; i++ {
		newFlash := false
		flashedPoints := make([]point, 0)

		// Increase everything by 1
		for i, row := range grid {
			for j, num := range row {
				grid[i][j] = num + 1
				if grid[i][j] > 9 {
					newFlash = true
				}
			}
		}

		// "Flash" all points > 9 and bump neighboring points until there are no new points to "flash"
		for newFlash {
			newFlash = false
			for i, row := range grid {
				for j, num := range row {
					if num > 9 && !alreadyFlashed(point{x: j, y: i}, flashedPoints) {
						flashedPoints = append(flashedPoints, point{x: j, y: i})
						newFlash = true
						if i != 9 {
							grid[i+1][j] = grid[i+1][j] + 1
							if j != 9 {
								grid[i+1][j+1] = grid[i+1][j+1] + 1
							}
							if j != 0 {
								grid[i+1][j-1] = grid[i+1][j-1] + 1
							}
						}
						if j != 9 {
							grid[i][j+1] = grid[i][j+1] + 1
							if i != 0 {
								grid[i-1][j+1] = grid[i-1][j+1] + 1
							}
						}
						if i != 0 {
							grid[i-1][j] = grid[i-1][j] + 1
							if j != 0 {
								grid[i-1][j-1] = grid[i-1][j-1] + 1
							}
						}
						if j != 0 {
							grid[i][j-1] = grid[i][j-1] + 1
						}
					}
				}
			}
		}

		// Reset all points that flashed
		for i, row := range grid {
			for j, num := range row {
				if num > 9 {
					grid[i][j] = 0
				}
			}
		}

		fmt.Println("END ROUND", i)
		for _, v := range grid {
			fmt.Println(v)
		}
		totalFlashes = totalFlashes + len(flashedPoints)
		fmt.Println("FLASHED POINTS", flashedPoints, "LENGTH", len(flashedPoints))
		if len(flashedPoints) == 100 {
			fmt.Println("WE ARE SYNCED", i+1)
			break
		}
	}

	fmt.Println("Total Flashes", totalFlashes)
}

func alreadyFlashed(p point, flashedPoints []point) bool {
	for _, v := range flashedPoints {
		if p.x == v.x && p.y == v.y {
			return true
		}
	}
	return false
}
