/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// merakeyCmd represents the merakey command
var merakeyCmd = &cobra.Command{
	Use:   "merakey",
	Short: "Run a command line application that lists some great qualities about Merakey consumers.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mainPeopleAtMerakey()
	},
}

func init() {
	rootCmd.AddCommand(merakeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// merakeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// merakeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
