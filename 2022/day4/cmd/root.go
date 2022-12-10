package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "day4",
	Short: "Solve day 4 for Advent Of Code",
	Long: `day 4 involves finding overlap in integer ranges.
Each line represents two ranges of sections a pair of elves is assigned to clean.
For example, given the string "6-6,4-6" means
- The first elf has section 6
- The second elf has sections 4, 5, and 6.

In part 1, you have to determine if either of the elf ranges completely
overlaps the other range in its pair.

In part 2, you have to determine if either of the elf ranges have any
overlap in its pair.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
