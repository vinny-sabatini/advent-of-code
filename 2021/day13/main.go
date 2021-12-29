package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type fold struct {
	direction string
	location  int
}

// Day 13 challenge 1 takes in a list of points and a list of folds. The goal is to figure out how many
// unique points remain after the first fold. A fold has an axis (x or y), and a position. If the fold is
// on the y axis, all points below the line fold up above the line, and if the fold is on the x axis, all
// points to the right are folded to the left. For small-input.txt, there are 17 unique points after the
// first fold.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var points []point
	var folds []fold

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if line[0] != "" {
			if len(line) == 2 {
				x, _ := strconv.Atoi(line[0])
				y, _ := strconv.Atoi(line[1])
				points = append(points, point{x: x, y: y})
			} else {
				foldInfo := strings.Split(strings.Fields(line[0])[2], "=")
				l, _ := strconv.Atoi(foldInfo[1])
				folds = append(folds, fold{location: l, direction: foldInfo[0]})
			}
		}
	}

	fmt.Println("POINTS", points)
	fmt.Println("FOLDS", folds)
	for _, fold := range folds {
		for i, p := range points {
			if fold.direction == "x" && p.x > fold.location {
				points[i] = point{
					x: p.x - (2 * (p.x - fold.location)),
					y: p.y,
				}
			}
			if fold.direction == "y" && p.y > fold.location {
				points[i] = point{
					x: p.x,
					y: p.y - (2 * (p.y - fold.location)),
				}
			}
		}
		uniquePoints := getUniquePoints(points)
		fmt.Println("POINTS", uniquePoints, "COUNT", len(uniquePoints))
		break
	}
}

func getUniquePoints(points []point) []point {
	var uniquePoints []point
	for _, p := range points {
		pointExists := false
		for _, check := range uniquePoints {
			if check.x == p.x && check.y == p.y {
				pointExists = true
				break
			}
		}
		if !pointExists {
			uniquePoints = append(uniquePoints, p)
		}
	}
	return uniquePoints
}
