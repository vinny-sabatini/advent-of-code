package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "day5",
	Short: "Solve day 5 for Advent Of Code",
	Long: `day 5 involves stacking and moving crates
The input is in two parts, the first part is the starting positions of the crates.
The second part is how the crates are moved around. Moves consist of how many crates
are being moved, what stack those are being moved from, and what stack those are being
moved to.

For challenge one, each crate is moved one at a time.
For challenge two, all of the crates are moved at the same time.

The program will print the starting and finished stacks, and the answer
to the challenge is the top container for each stack after all of the moves.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
