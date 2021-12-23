package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Day 4 challenge 1 takes in a list of bingo calls, and then a list of bingo boards.
// The goal is to find the board that will win first. When the winning board is found,
// you can calculate the "answer" by summing all of the uncalled numbers on the winning
// board, and then multiplying that by the last called number. For example, on small-input.txt
// the third board is the winning board. It sums up to 188, the last number called is 24, and
// the final total is 4512.
func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	bingoCalls := strings.Split(scanner.Text(), ",")
	var bingoBoards [][]string
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		bingoBoards = append(bingoBoards, row)
	}
	for _, call := range bingoCalls {
		bingoBoards = makeCall(call, bingoBoards)
		bingo, board := checkBingo(bingoBoards)
		if bingo {
			fmt.Println("Board", board, "Call", call)
			printWinningBoard(board, bingoBoards)
			break
		}
	}
}

func makeCall(call string, boards [][]string) [][]string {
	for i, row := range boards {
		for j, spot := range row {
			if spot == call {
				boards[i][j] = "x"
			}
		}
	}
	return boards
}

func checkBingo(boards [][]string) (bool, int) {
	board := 0
	columnCheck := []int{0, 0, 0, 0, 0}
	bingoRow := []string{"x", "x", "x", "x", "x"}

	for _, row := range boards {
		if len(row) == 0 {
			board = board + 1
			columnCheck = []int{0, 0, 0, 0, 0}
			continue
		}
		if reflect.DeepEqual(row, bingoRow) {
			return true, board
		}
		for j, v := range row {
			if v == "x" {
				columnCheck[j] = columnCheck[j] + 1
				if columnCheck[j] == 5 {
					return true, board
				}
			}
		}
	}
	return false, 0
}

func printWinningBoard(board int, boards [][]string) {
	sum := 0
	count := 0
	for _, row := range boards {
		if len(row) == 0 {
			count = count + 1
		}
		if count == board {
			fmt.Println(row)
			for _, v := range row {
				value, _ := strconv.Atoi(v)
				sum = sum + value
			}
		}
	}
	fmt.Println("The board sum", sum)
}
