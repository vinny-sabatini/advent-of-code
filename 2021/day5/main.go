package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type path struct {
	xStart, yStart, xFinish, yFinish string
}

// Day 5 challenge 1 takes in a list of starting and finishing coordinates
// The goal is to find how many points have at least 2 lines that cross that
// given point. For example, on small-input.txt, points 0,9 through 2,9 and
// points 4,3 and 4,7 have two intersections, so the final answer is 5.
//
// Day 5 challenge two does the same thing as challenge one, however you have
// to count diagonal lines (exactly 45*). On small-input.txt, the final answer
// is 12.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	var paths []path
	hits := make(map[string]int)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		start := strings.Split(row[0], ",")
		finish := strings.Split(row[2], ",")
		paths = append(paths, path{
			xStart:  start[0],
			yStart:  start[1],
			xFinish: finish[0],
			yFinish: finish[1],
		})
	}

	for _, p := range paths {
		yStart, _ := strconv.Atoi(p.yStart)
		yFinish, _ := strconv.Atoi(p.yFinish)
		xStart, _ := strconv.Atoi(p.xStart)
		xFinish, _ := strconv.Atoi(p.xFinish)
		if p.xStart == p.xFinish && p.yStart == p.yFinish {
			point := fmt.Sprintf("%s,%s", p.xStart, p.xFinish)
			addToHits(point, hits)
		} else if p.xStart == p.xFinish {
			if yStart > yFinish {
				for i := yFinish; i != yStart+1; i++ {
					point := fmt.Sprintf("%s,%d", p.xStart, i)
					addToHits(point, hits)
				}
			} else {
				for i := yStart; i != yFinish+1; i++ {
					point := fmt.Sprintf("%s,%d", p.xStart, i)
					addToHits(point, hits)
				}
			}
		} else if p.yStart == p.yFinish {
			if xStart > xFinish {
				for i := xFinish; i != xStart+1; i++ {
					point := fmt.Sprintf("%d,%s", i, p.yStart)
					addToHits(point, hits)
				}
			} else {
				for i := xStart; i != xFinish+1; i++ {
					point := fmt.Sprintf("%d,%s", i, p.yStart)
					addToHits(point, hits)
				}
			}
		} else {
			var xHits, yHits []string
			if xStart > xFinish {
				for i := xStart; i > xFinish-1; i-- {
					xHits = append(xHits, strconv.Itoa(i))
				}
			} else {
				for i := xStart; i < xFinish+1; i++ {
					xHits = append(xHits, strconv.Itoa(i))
				}
			}
			if yStart > yFinish {
				for i := yStart; i > yFinish-1; i-- {
					yHits = append(yHits, strconv.Itoa(i))
				}
			} else {
				for i := yStart; i < yFinish+1; i++ {
					yHits = append(yHits, strconv.Itoa(i))
				}
			}
			for i := 0; i < len(xHits); i++ {
				point := fmt.Sprintf("%s,%s", xHits[i], yHits[i])
				addToHits(point, hits)
			}
		}
	}
	counter := 0
	for _, v := range hits {
		if v > 1 {
			//fmt.Println("Point", i, "Hits", v)
			counter = counter + 1
		}
	}
	fmt.Println("Total hits", counter)
}

func addToHits(coordinate string, hits map[string]int) {
	if val, ok := hits[coordinate]; ok {
		hits[coordinate] = val + 1
	} else {
		hits[coordinate] = 1
	}
}
