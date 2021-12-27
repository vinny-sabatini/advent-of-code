package main

import (
	"bufio"
	"fmt"
	"os"
)

// Day 10 challenge 1 takes in a list of strings containing various brackets. The goal
// is to try to verify the opening and closing brackets line up properly. You then have
// to find the corrupted lines where the next closing bracket does not line up with the
// last open bracket. There are different point values based on what character is mismatched.
// For small-input.txt, the total error score was 26397.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	legend := map[byte]byte{
		'{': '}',
		'[': ']',
		'<': '>',
		'(': ')',
	}

	errorValues := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	errorTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		tracker := make([]byte, 0)
		for _, character := range line {
			if _, ok := legend[byte(character)]; ok {
				tracker = append(tracker, byte(character))
			} else if legend[tracker[len(tracker)-1]] == byte(character) {
				tracker = tracker[:len(tracker)-1]
			} else {
				errorTotal = errorTotal + errorValues[byte(character)]
				fmt.Println("ERROR...TRACKER", string(tracker), "CHARACTER", string(character))
				break
			}
		}
	}
	fmt.Println("ERROR TOTAL", errorTotal)
}
