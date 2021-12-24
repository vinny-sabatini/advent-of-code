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
//
// Day 6 challenge 2 does the same thing as part 1, however you have to track
// how many fish would exist after 256 days. The answer for small-input.txt would
// be 26984457539.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	tracker := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	scanner.Scan()
	inputText := strings.Split(scanner.Text(), ",")
	for _, v := range inputText {
		vInt, _ := strconv.Atoi(v)
		tracker[vInt] = tracker[vInt] + 1
	}

	for day := 0; day < 256; day++ {
		newFish := tracker[0]
		for i := 0; i < len(tracker)-1; i++ {
			tracker[i] = tracker[i+1]
		}
		tracker[8] = newFish
		tracker[6] = tracker[6] + newFish
	}
	fmt.Println("Tracker", tracker)
	count := 0
	for _, v := range tracker {
		count = count + v
	}
	fmt.Println(count)
}
