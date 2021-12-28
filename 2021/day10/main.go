package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Day 10 challenge 1 takes in a list of strings containing various brackets. The goal
// is to try to verify the opening and closing brackets line up properly. You then have
// to find the corrupted lines where the next closing bracket does not line up with the
// last open bracket. There are different point values based on what character is mismatched.
// For small-input.txt, the total error score was 26397.
//
// Day 10 challenge 2 takes the linnes that are not corrupted with syntax errors and calculates
// the characters required to complete the incomplete line of characters. With that list of characters,
// you calculate the score of the autocomplete calculation by starting a 0, and then for each character
// multiply the previous score by 5, then add the value of the character found. Finally, when you have
// all of the scores, the "final" score is the middle score of all the autocomplete lines found. For example
// in small-input.txt, there are 5 lines with syntax, leaving 5 lines that are incomplete. The middle score
// is 288957 on line 1.
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

	autocompleteValues := map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	errorTotal := 0
	autoCompleteScores := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		tracker := make([]byte, 0)
		unclosed := make([]byte, 0)
		for _, character := range line {
			if _, ok := legend[byte(character)]; ok {
				tracker = append(tracker, byte(character))
				unclosed = append(unclosed, legend[byte(character)])
			} else if legend[tracker[len(tracker)-1]] == byte(character) {
				tracker = tracker[:len(tracker)-1]
				unclosed = unclosed[:len(unclosed)-1]
			} else {
				errorTotal = errorTotal + errorValues[byte(character)]
				fmt.Println("ERROR...TRACKER", string(tracker), "CHARACTER", string(character), "POINTS", errorValues[byte(character)])
				unclosed = make([]byte, 0)
				break
			}
		}
		if len(unclosed) != 0 {
			unclosed = reverse(unclosed)
			lineTotal := 0
			for _, v := range unclosed {
				lineTotal = (lineTotal * 5) + autocompleteValues[v]
			}
			fmt.Println("AUTOCOMPLETE", string(unclosed), "POINTS", lineTotal)
			autoCompleteScores = append(autoCompleteScores, lineTotal)
		}
	}
	fmt.Println("ERROR TOTAL", errorTotal)

	sort.Ints(autoCompleteScores)
	fmt.Println("AUTOCOMPLETE TOTAL", autoCompleteScores)
	fmt.Println("MIDDLE SCORE", autoCompleteScores[len(autoCompleteScores)/2])
}

func reverse(in []byte) []byte {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}
