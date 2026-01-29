package cmd

import (
	"fmt"

	"github.com/sarkartanmay393/ah/pkg/manager"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [package]",
	Short: "Install a package from the registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		if err := manager.InstallPackage(url); err != nil {
			fmt.Printf("Error installing package: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
