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
//
// Day 4 challenge two takes in the same inputs as challenge 1, but the goal is to figure out
// the last board that would win. For small-input.txt, the second board would be the last to
// win, it wins on 13, the remaining tiles sums up to 148, and the final total is 1924.
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
	findWinner(bingoBoards, bingoCalls)
	findLoser(bingoBoards, bingoCalls)
}

func findWinner(boards [][]string, calls []string) {
	for _, call := range calls {
		boards = makeCall(call, boards)
		bingo, board := checkBingo(boards)
		if bingo {
			fmt.Println("Board", board, "Call", call)
			printBoard(board, boards)
			break
		}
	}
}

func findLoser(boards [][]string, calls []string) {
	boardCount := countBoards(boards)
	var lastBoard int
	for _, call := range calls {
		boards = makeCall(call, boards)
		bingoCount := countBingos(boards)
		if boardCount == bingoCount {
			lastBoard = getLastBoard(boards)
			fmt.Println("Last board", lastBoard)
		}
		if boardCount+1 == bingoCount {
			fmt.Println("Call", call)
			printBoard(lastBoard, boards)
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

func countBoards(boards [][]string) int {
	count := 0
	for _, v := range boards {
		if len(v) == 0 {
			count = count + 1
		}
	}
	return count
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

func countBingos(boards [][]string) int {
	bingoCount := 0
	bingoCheck := false
	columnCheck := []int{0, 0, 0, 0, 0}
	bingoRow := []string{"x", "x", "x", "x", "x"}

	for _, row := range boards {
		if len(row) == 0 {
			bingoCheck = false
			columnCheck = []int{0, 0, 0, 0, 0}
			continue
		}
		if !bingoCheck {
			if reflect.DeepEqual(row, bingoRow) {
				bingoCheck = true
				bingoCount = bingoCount + 1
				continue
			}
			for j, v := range row {
				if v == "x" {
					columnCheck[j] = columnCheck[j] + 1
					if columnCheck[j] == 5 {
						bingoCheck = true
						bingoCount = bingoCount + 1
					}
				}
			}
		}
	}
	return bingoCount
}

func getLastBoard(boards [][]string) int {
	board := 0
	columnCheck := []int{0, 0, 0, 0, 0}
	bingoRow := []string{"x", "x", "x", "x", "x"}
	isBingo := true

	for _, row := range boards {
		if len(row) == 0 {
			if isBingo {
				board = board + 1
				columnCheck = []int{0, 0, 0, 0, 0}
				isBingo = false
				continue
			} else {
				return board
			}
		}
		if reflect.DeepEqual(row, bingoRow) {
			isBingo = true
		}
		for j, v := range row {
			if v == "x" {
				columnCheck[j] = columnCheck[j] + 1
				if columnCheck[j] == 5 {
					isBingo = true
				}
			}
		}
	}
	return 0
}

func printBoard(board int, boards [][]string) {
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
