package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Day 14 challenge 1 takes in a polymer and a list of transformations. You have to perform transformations on the
// polymer 10 times, and then figure out the most and least frequently occurring element. The final answer is found
// bu subtracting the count of the least frequently occurring element from the count of the most frequently occurring
// element. A transformation takes two characters from the polymer, and injects in the corresponding character between
// those two characters. This has to be done for all characters in the polymer. For small-input.txt, the most frequent
// charcter is B with 1749, and the least frequent character is H with 161, so the final answer is 1588.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	polymer := scanner.Text()
	scanner.Scan()

	polymerLegend := make(map[string]string)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		polymerLegend[line[0]] = line[1]
	}

	for i := 0; i < 10; i++ {
		polymer = transformPolymer(polymer, polymerLegend)
		fmt.Println("Length", len(polymer), "after", i+1)
	}
	fmt.Println(diffMostAndLeastFrequentElements(polymer))
}

func transformPolymer(polymer string, legend map[string]string) string {
	var transformation string
	for i := 0; i < len(polymer)-1; i++ {
		key := fmt.Sprintf("%s%s", string(polymer[i]), string(polymer[i+1]))
		transformation = transformation + string(polymer[i]) + legend[key]
	}
	return transformation + string(polymer[len(polymer)-1])
}

func diffMostAndLeastFrequentElements(polymer string) int {
	tracker := make(map[string]int)
	for _, element := range polymer {
		if v, ok := tracker[string(element)]; !ok {
			tracker[string(element)] = 1
		} else {
			tracker[string(element)] = v + 1
		}
	}

	min := math.MaxInt
	max := math.MinInt
	for _, v := range tracker {
		min = int(math.Min(float64(v), float64(min)))
		max = int(math.Max(float64(v), float64(max)))
	}
	return max - min
}
