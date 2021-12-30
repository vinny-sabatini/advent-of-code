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
//
// Day 14 challenge 2 does the same thing, but tracks out to 40 days. For small-input.txt, the most frequent character
// is B with 2192039569602 and the least frequent is H with 3849876073 so the final answer is 2188189693529.
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
	pairs := make(map[string]int)
	elementCounter := make(map[string]int)

	for i := 0; i < len(polymer)-1; i++ {
		elementCounter[string(polymer[i])]++
		pairs[string(polymer[i])+string(polymer[i+1])]++
	}
	elementCounter[string(polymer[len(polymer)-1])]++
	fmt.Println(pairs, elementCounter)
	newPairs := pairs

	for i := 0; i < 40; i++ {
		current := newPairs
		newPairs = make(map[string]int)
		for pair, count := range current {
			pair1 := string(pair[0]) + string(polymerLegend[pair])
			pair2 := string(polymerLegend[pair]) + string(pair[1])
			newPairs[pair1] = newPairs[pair1] + count
			newPairs[pair2] = newPairs[pair2] + count
			elementCounter[polymerLegend[pair]] = elementCounter[polymerLegend[pair]] + count
		}
		fmt.Println(newPairs, elementCounter)
	}
	max := math.MinInt
	min := math.MaxInt

	for _, v := range elementCounter {
		max = int(math.Max(float64(v), float64(max)))
		min = int(math.Min(float64(v), float64(min)))
	}
	fmt.Println(max - min)
}
