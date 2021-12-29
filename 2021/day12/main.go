package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cave struct {
	isBig bool
	name  string
	paths []string
}

type path struct {
	start  string
	finish string
}

// Day 12 challenge 1
func main() {
	input, err := os.Open("./small-input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var caves []cave

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		aName := line[0]
		aIsBig := isBigCave(aName)
		bName := line[1]
		bIsBig := isBigCave(bName)

		start, startFound := getCave(aName, caves)
		finish, finishFound := getCave(bName, caves)
		aCave := cave{
			isBig: aIsBig,
			name:  aName,
			paths: []string{bName},
		}
		bCave := cave{
			isBig: bIsBig,
			name:  bName,
			paths: []string{aName},
		}
		if !startFound {
			caves = append(caves, aCave)
		} else {
			destFound := caveInPath(bName, caves[start].paths)
			if !destFound {
				caves[start].paths = append(caves[start].paths, bName)
			}
		}

		if !finishFound {
			caves = append(caves, bCave)
		} else {
			sourceFound := caveInPath(aName, caves[finish].paths)
			if !sourceFound {
				caves[finish].paths = append(caves[finish].paths, aName)
			}
		}
	}

	for _, cave := range caves {
		fmt.Printf("Cave %s can go to %v\n", cave.name, cave.paths)

	}
}

func isBigCave(cave string) bool {
	return strings.ToUpper(cave) == cave
}

func getCave(name string, caves []cave) (int, bool) {
	for i, v := range caves {
		if v.name == name {
			return i, true
		}
	}
	return -1, false
}

func caveInPath(name string, paths []string) bool {
	for _, v := range paths {
		if name == v {
			return true
		}
	}
	return false
}
