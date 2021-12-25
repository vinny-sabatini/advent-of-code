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

const (
	LEN_ONE   int = 2
	LEN_SEVEN int = 3
	LEN_FOUR  int = 4
	LEN_EIGHT int = 7
)

// Day 8 challenge 1 takes in a journal of inputs and outputs. Each letter corresponds
// to a segment of a digital number. Eg, 1 has the two segments going down the right side,
// 8 has all 7 segments. The inputs are on the left side of the "|" and the outputs are
// on the right side. The goal of part 1 is to figure out how many of the "unique" numbers
// are in the outputs. The unique numbers are 1, 4, 7, and 8 because there are no other numbers
// that use the same number of segments to create the number. For example, using small-input.txt
// the output is 26.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var notebook []Entry

	for scanner.Scan() {
		entryString := strings.Split(scanner.Text(), "|")
		notebook = append(notebook, Entry{
			input:  strings.Fields(entryString[0]),
			output: strings.Fields(entryString[1]),
		})
	}

	countUnique := 0
	for _, entry := range notebook {
		for _, output := range entry.output {
			switch len(output) {
			case LEN_EIGHT,
				LEN_FOUR,
				LEN_ONE,
				LEN_SEVEN:
				countUnique = countUnique + 1
			}
		}
	}
	fmt.Println(countUnique)
}
