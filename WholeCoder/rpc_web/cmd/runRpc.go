/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runRpcCmd represents the runRpc command
var runRpcCmd = &cobra.Command{
	Use:   "runRpc",
	Short: "Run an Remote Procedure Call server for hackers and their specialties.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runRpc called")
		mainRunRpcServer()
	},
}

func init() {
	rootCmd.AddCommand(runRpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runRpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runRpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
