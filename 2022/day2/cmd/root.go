package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "day2",
	Short: "Solve day 2 for Advent Of Code",
	Long: `Day 2 involves rigging a rock, paper, scissors tournament.
Each line represents a single round of rock paper scissors.

Points are calculated based a total of what is thrown, and the result of the round
Rock = 1, Paper = 2, Scissors = 3
Lose = 0, Draw = 3, Win = 6

In part 1, the first column is the opponents play, and the second is your play
where A and X == rock, B and Y == paper, C and Z == scissors.

In part 2, the first column is the opponents play, and the second column is the result
where A == rock, B == paper, C == scissors, X == lose, Y == draw, Z == win

Both programs will print the total score of all rounds`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
