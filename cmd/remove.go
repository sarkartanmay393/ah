package cmd

import (
	"fmt"

	"github.com/sarkartanmay393/ah/pkg/manager"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [package]",
	Short: "Remove an installed alias package",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		if err := manager.RemovePackage(packageName); err != nil {
			fmt.Printf("Error removing package: %v\n", err)
			return
		}
		fmt.Printf("Package '%s' removed.\n", packageName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
