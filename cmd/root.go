package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ah",
	Short: "Alias Hub - The ultimate shell alias manager",
	Long: `Alias Hub (ah) helps you manage, share, and sync shell aliases across your machines.
It features conflict detection, live updates, and a public registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
