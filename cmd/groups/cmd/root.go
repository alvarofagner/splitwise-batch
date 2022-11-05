/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// TODO: Add support for:
// - Get information about a group
// - Create a group
// - Delete a group
// - Resotre a group
// - Add a user to a group
// - Remove user from a group
var rootCmd = &cobra.Command{
	Use: "categories",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
