/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runWebCmd represents the runWeb command
var runWebCmd = &cobra.Command{
	Use:   "runWeb",
	Short: "Run a Web Server to display hacker names and their specialties.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runWeb called")
		mainRunWebServer()
	},
}

func init() {
	rootCmd.AddCommand(runWebCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runWebCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runWebCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
