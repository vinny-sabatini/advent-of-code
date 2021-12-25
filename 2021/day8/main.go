package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	input  []string
	output []string
}

// Day 8 challenge 1 takes in a journal of inputs and outputs. Each letter corresponds
// to a segment of a digital number. Eg, 1 has the two segments going down the right side,
// 8 has all 7 segments. The inputs are on the left side of the "|" and the outputs are
// on the right side. The goal of part 1 is to figure out how many of the "unique" numbers
// are in the outputs. The unique numbers are 1, 4, 7, and 8 because there are no other numbers
// that use the same number of segments to create the number. For example, using small-input.txt
// the output is 26.
//
// Day 8 challenge 2 is to try to decipher the sum of all of the outputs. Based on the unique
// numbers, we should be able to correspond a letter to what segment it represents, and then
// figure out the 4 digit output numbers. For example, using small-input.txt, the total is
// 61229
func main() {
	input, err := os.Open("./small-input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var notebook []Entry
	tracker := make(map[int]string)

	for scanner.Scan() {
		entryString := strings.Split(scanner.Text(), "|")
		notebook = append(notebook, Entry{
			input:  strings.Fields(entryString[0]),
			output: strings.Fields(entryString[1]),
		})
	}

	for _, entry := range notebook {
		var lenFive []string = make([]string, 0)
		var lenSix []string = make([]string, 0)
		for _, input := range entry.input {
			switch len(input) {
			case 2:
				tracker[1] = input
			case 3:
				tracker[7] = input
			case 4:
				tracker[4] = input
			case 5:
				found := false
				for _, v := range lenFive {
					if v == input {
						found = true
					}
				}
				if !found {
					lenFive = append(lenFive, input)
				}
			case 6:
				found := false
				for _, v := range lenSix {
					if v == input {
						found = true
					}
				}
				if !found {
					lenSix = append(lenSix, input)
				}
			case 7:
				tracker[8] = input
			}
		}
		fmt.Println("Tracker", tracker, "Five", lenFive, "Six", lenSix)
	}
}
