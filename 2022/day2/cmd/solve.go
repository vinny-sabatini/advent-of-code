package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve day2",
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
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	totalScore := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		round := scanner.Text()
		// A and X == Rock
		// B and Y == Paper
		// C and Z == Scissors
		win := []string{"A Y", "B Z", "C X"}
		draw := []string{"A X", "B Y", "C Z"}
		roundScore := 0
		if strings.Contains(round, "X") {
			roundScore += 1
		} else if strings.Contains(round, "Y") {
			roundScore += 2
		} else {
			roundScore += 3
		}

		if slices.Contains(win, round) {
			roundScore += 6
		} else if slices.Contains(draw, round) {
			roundScore += 3
		}
		fmt.Println(round, roundScore)
		totalScore += roundScore
	}
	fmt.Println("Final Score", totalScore)
}

func partTwo(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	totalScore := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		round := scanner.Text()
		roundScore := 0
		roundParts := strings.Split(round, " ")

		// A == Rock
		// B == Paper
		// C == Scissors
		// X == Lose
		// Y == Draw
		// Z == Win
		winMap := map[string]int{"A": 2, "B": 3, "C": 1}
		drawMap := map[string]int{"A": 1, "B": 2, "C": 3}
		loseMap := map[string]int{"A": 3, "B": 1, "C": 2}

		if roundParts[1] == "Y" {
			roundScore += 3
			roundScore += drawMap[roundParts[0]]
		} else if roundParts[1] == "Z" {
			roundScore += 6
			roundScore += winMap[roundParts[0]]
		} else {
			roundScore += 0
			roundScore += loseMap[roundParts[0]]
		}

		fmt.Println(round, roundScore)
		totalScore += roundScore
	}
	fmt.Println("Final Score", totalScore)
}
