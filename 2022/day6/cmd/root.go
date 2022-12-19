package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "day6",
	Short: "Solve day 6 for Advent Of Code",
	Long: `day 6 involves finding substrings of unique characters

For challenge one, the substring of unique characters is 4 characters long
For challenge two, the substring of unique characters is 14 characters long`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
