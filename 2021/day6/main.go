package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Day 6 challenge 1 takes in a list of integers representing the timers for
// a fish before it will spawn a new fish. When the timer hits 0, a new fish
// is spawned with a timer of 8, and the parent fish timer resets to 6. The
// goal is to determine how many fish will exist after 80 days. For example
// on small-input.txt, the final fish count is 5934.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var liveFish, newFish []int

	scanner.Scan()
	inputText := strings.Split(scanner.Text(), ",")
	for _, v := range inputText {
		vInt, _ := strconv.Atoi(v)
		liveFish = append(liveFish, vInt)
	}

	for day := 0; day < 80; day++ {
		newFish = make([]int, 0)
		for fish, val := range liveFish {
			liveFish[fish] = val - 1
			if val == 0 {
				liveFish[fish] = 6
				newFish = append(newFish, 8)
			}
		}
		liveFish = append(liveFish, newFish...)
	}
	fmt.Println("Fish Count", len(liveFish))
}
