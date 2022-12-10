package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve day4",
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

	totalMatching := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		elfPairSections := strings.Split(scanner.Text(), ",")
		elfOne := generateElfString(elfPairSections[0])
		elfTwo := generateElfString(elfPairSections[1])
		if strings.Contains(elfOne, elfTwo) || strings.Contains(elfTwo, elfOne) {
			fmt.Printf("One: %s\nTwo: %s\n\n", elfOne, elfTwo)
			totalMatching += 1
			continue
		}
	}
	fmt.Println("Total matching:", totalMatching)
}

func generateElfString(elf string) string {
	elfString := ","
	rangeSplit := strings.Split(elf, "-")
	start, _ := strconv.Atoi(rangeSplit[0])
	end, _ := strconv.Atoi(rangeSplit[1])
	for i := start; i <= end; i++ {
		elfString += fmt.Sprintf("%v,", i)
	}
	return elfString
}

func partTwo(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open", inputFile, err)
		os.Exit(1)
	}
	defer input.Close()

	totalMatching := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		elfPairSections := strings.Split(scanner.Text(), ",")
		elfOne := generateElfString(elfPairSections[0])
		elfTwo := generateElfString(elfPairSections[1])
		for _, v := range strings.Split(elfOne, ",") {
			check := fmt.Sprintf(",%v,", v)
			if strings.Contains(elfTwo, check) {
				totalMatching += 1
				break
			}
		}
	}
	fmt.Println("Total matching:", totalMatching)
}
