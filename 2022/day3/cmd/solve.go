package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve day3",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input-file")
		challenge, _ := cmd.Flags().GetInt("challenge")
		switch challenge {
		case 1:
			partOne(inputFile)
		case 2:
			partTwo(inputFile)
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
}

func partOne(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open", inputFile, err)
		os.Exit(1)
	}
	defer input.Close()

	prioritySum := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rucksack := scanner.Text()
		compartmentOne := rucksack[0 : len(rucksack)/2]
		compartmentTwo := rucksack[len(rucksack)/2:]
		for _, v := range compartmentOne {
			if strings.ContainsRune(compartmentTwo, v) {
				prioritySum += strings.Index(alphabet, string(v)) + 1
				break
			}
		}
	}
	fmt.Println("Score", prioritySum)
}

func partTwo(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open", inputFile, err)
		os.Exit(1)
	}
	defer input.Close()

	rucksackCounter := 0
	prioritySum := 0
	tracker := make([]string, 3)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rucksack := scanner.Text()
		trackerIndex := rucksackCounter % 3
		tracker[trackerIndex] = rucksack
		// Found end of group
		if trackerIndex == len(tracker)-1 {
			for _, v := range tracker[0] {
				if strings.ContainsRune(tracker[1], v) && strings.ContainsRune(tracker[2], v) {
					prioritySum += strings.Index(alphabet, string(v)) + 1
					break
				}
			}
		}
		rucksackCounter += 1
	}
	fmt.Println("Score", prioritySum)
}
