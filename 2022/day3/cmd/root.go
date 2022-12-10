package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "day3",
	Short: "Solve day 3 for Advent Of Code",
	Long: `day 3 involves finding common characters across various strings.
Each line represents a rucksack carried by an elf.
Each character in a line represents a single item with a priority.

Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52.

In part 1, the rucksack is split into two compartments (split string in half)
and then you have to find the common character, and add that priority to the total.

In part 2, the rucksacks are packaged into groups of three, you have to find the common
character between each trio of rucksacks, and add that priority to the total.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
