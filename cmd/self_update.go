package cmd

import (
	"fmt"
	"os"

	"github.com/sarkartanmay393/ah/pkg/updater"
	"github.com/spf13/cobra"
)

var selfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Update ah to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		if err := updater.SelfUpdate(); err != nil {
			fmt.Fprintf(os.Stderr, "Error updating ah: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(selfUpdateCmd)
}
