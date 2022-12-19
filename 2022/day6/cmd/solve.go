package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve day6",
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

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		tracker := ""
		line := scanner.Text()
		for i, v := range line {
			checkIndex := strings.Index(tracker, string(v))
			if checkIndex != -1 {
				tracker = tracker[checkIndex+1:] + string(v)
				continue
			} else {
				tracker = tracker + string(v)
			}
			if len(tracker) == 4 {
				fmt.Println("Found it!", i+1, string(v), tracker)
				break
			}
		}
	}
}

func partTwo(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open", inputFile, err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		tracker := ""
		line := scanner.Text()
		for i, v := range line {
			checkIndex := strings.Index(tracker, string(v))
			if checkIndex != -1 {
				tracker = tracker[checkIndex+1:] + string(v)
				continue
			} else {
				tracker = tracker + string(v)
			}
			if len(tracker) == 14 {
				fmt.Println("Found it!", i+1, string(v), tracker)
				break
			}
		}
	}
}
