package cmd

import (
	"fmt"

	"github.com/sarkartanmay393/ah/pkg/manager"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable [package]",
	Short: "Enable an installed alias package",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		if err := manager.EnablePackageFromRepo(packageName); err != nil {
			fmt.Printf("Error enabling package: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
