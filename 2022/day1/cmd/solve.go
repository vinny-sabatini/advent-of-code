package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve day1",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input-file")
		challenge, _ := cmd.Flags().GetInt("challenge")
		switch challenge {
		case 1:
			partOne(inputFile)
		case 2:
			topElfCount, _ := cmd.Flags().GetInt("top-elf-count")
			partTwo(inputFile, topElfCount)
		default:
			fmt.Println("Invalid Challenge Number", challenge)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)
	solveCmd.Flags().IntP("challenge", "c", 1, "The which part of the challenge to solve")
	solveCmd.MarkFlagRequired("challenge")
	solveCmd.Flags().StringP("input-file", "f", "", "The path to the file with the inputs")
	solveCmd.MarkFlagRequired("input-file")
	solveCmd.Flags().IntP("top-elf-count", "t", 3, "How many of the top elves should be counted (only for part 2)")
}

func partOne(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	maxCaloriesCarried := 0
	currentCaloriesCarried := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		calories := scanner.Text()
		if calories == "" {
			fmt.Printf("Elf has %d calories\n", currentCaloriesCarried)
			if currentCaloriesCarried > maxCaloriesCarried {
				maxCaloriesCarried = currentCaloriesCarried
			}
			currentCaloriesCarried = 0

		} else {
			c, _ := strconv.Atoi(calories)
			currentCaloriesCarried += c
		}
	}

	fmt.Printf("The top elf is carrying %d calories\n", maxCaloriesCarried)
}

func partTwo(inputFile string, topElfCount int) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	var topCalories []int
	currentCaloriesCarried := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		calories := scanner.Text()
		if calories == "" {
			fmt.Printf("Elf has %d calories\n", currentCaloriesCarried)
			if len(topCalories) < topElfCount {
				topCalories = append(topCalories, currentCaloriesCarried)
			} else {
				sort.IntSlice(topCalories).Sort()
				if currentCaloriesCarried > topCalories[0] {
					topCalories[0] = currentCaloriesCarried
				}
			}
			currentCaloriesCarried = 0

		} else {
			c, _ := strconv.Atoi(calories)
			currentCaloriesCarried += c
		}
	}

	totalCalories := 0
	sort.IntSlice(topCalories).Sort()
	for i, calories := range topCalories {
		fmt.Printf("Elf %d had %d calories\n", i, calories)
		totalCalories += calories
	}
	fmt.Println("Total Calories:", totalCalories)
}
