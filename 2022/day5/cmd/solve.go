package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var filler = "[ ]"

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve day5",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input-file")
		challenge, _ := cmd.Flags().GetInt("challenge")
		stackCount, _ := cmd.Flags().GetInt("stack-count")

		switch challenge {
		case 1:
			partOne(inputFile, stackCount)
		case 2:
			partTwo(inputFile, stackCount)
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
	solveCmd.Flags().IntP("stack-count", "s", 3, "The path to the file with the inputs")
	solveCmd.MarkFlagRequired("stack-count")
}

func partOne(inputFile string, stackCount int) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open", inputFile, err)
		os.Exit(1)
	}
	defer input.Close()
	stacks := make([][]string, stackCount)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "move") {
			stringSplit := strings.Split(line, " ")
			count, _ := strconv.Atoi(stringSplit[1])
			from, _ := strconv.Atoi(stringSplit[3])
			to, _ := strconv.Atoi(stringSplit[5])
			for i := 0; i < count; i++ {
				popValue := stacks[from-1][len(stacks[from-1])-1]
				stacks[to-1] = append(stacks[to-1], popValue)
				stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
			}
		} else if strings.Contains(line, " 1") {
			for i := range stacks {
				for i2, j := 0, len(stacks[i])-1; i2 < j; i2, j = i2+1, j-1 {
					stacks[i][i2], stacks[i][j] = stacks[i][j], stacks[i][i2]
				}
			}
			fmt.Println("Start", stacks)
			scanner.Scan()
		} else {
			for index, character := range line {
				if !strings.ContainsRune(filler, character) {
					stack := index / 4
					stacks[stack] = append(stacks[stack], string(character))
				}
			}
		}
	}
	fmt.Println("Finish", stacks)
	for _, v := range stacks {
		fmt.Print(v[len(v)-1])
	}
	fmt.Println()
}

func partTwo(inputFile string, stackCount int) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open", inputFile, err)
		os.Exit(1)
	}
	defer input.Close()
	stacks := make([][]string, stackCount)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "move") {
			stringSplit := strings.Split(line, " ")
			count, _ := strconv.Atoi(stringSplit[1])
			from, _ := strconv.Atoi(stringSplit[3])
			to, _ := strconv.Atoi(stringSplit[5])
			popValues := stacks[from-1][len(stacks[from-1])-count:]
			stacks[to-1] = append(stacks[to-1], popValues...)
			stacks[from-1] = stacks[from-1][:len(stacks[from-1])-count]
		} else if strings.Contains(line, " 1") {
			for i := range stacks {
				for i2, j := 0, len(stacks[i])-1; i2 < j; i2, j = i2+1, j-1 {
					stacks[i][i2], stacks[i][j] = stacks[i][j], stacks[i][i2]
				}
			}
			fmt.Println("Start", stacks)
			scanner.Scan()
		} else {
			for index, character := range line {
				if !strings.ContainsRune(filler, character) {
					stack := index / 4
					stacks[stack] = append(stacks[stack], string(character))
				}
			}
		}
	}
	fmt.Println("Finish", stacks)
	for _, v := range stacks {
		fmt.Print(v[len(v)-1])
	}
	fmt.Println()
}
