package cmd

import (
	"fmt"
	"github.com/bash-shortcuts/ah/pkg/manager"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable [package]",
	Short: "Disable an alias package (without removing it)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		if err := manager.DisablePackage(packageName); err != nil {
			fmt.Printf("Error disabling package: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
