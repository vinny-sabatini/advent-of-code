/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "day1",
	Short: "Solve day 1 for Advent Of Code",
	Long: `Day 1 involves tracking how many calories each elf is carrying.
Each line with numbers in the input file indicates how many calories per piece of food.
If a line is empty, that indicates the end of tracking for a single elf.

Part 1 will provide the total calories for the top elf
Part 2 will provide the total calories for the top 3 elves`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
