package cmd

import (
	"fmt"

	"github.com/sarkartanmay393/ah/pkg/manager"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the package registry and re-compile aliases",
	Long:  `Downloads the latest package definitions from the registry and re-generates your alias configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating registry...")
		if err := manager.UpdateRegistry(); err != nil {
			// UpdateRegistry handles its own warnings/errors mostly, but if hard fail:
			fmt.Printf("Error updating registry: %v\n", err)
			return
		}

		fmt.Println("Compiling aliases...")
		if err := manager.CompileAliases(); err != nil {
			fmt.Printf("Error compiling aliases: %v\n", err)
			return
		}

		fmt.Println("All set! Registry and aliases updated.")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
