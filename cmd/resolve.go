package cmd

import (
	"fmt"

	"github.com/bash-shortcuts/ah/pkg/server"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve [package]",
	Short: "Launch the Conflict Resolution Web UI",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		fmt.Printf("Starting resolution UI for %s...\n", pkgName)
		if err := server.Start(pkgName); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(resolveCmd)
}
