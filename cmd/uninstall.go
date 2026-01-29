package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bash-shortcuts/ah/pkg/manager"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Completely remove Alias Hub and all data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⚠️  DANGER: This will delete:")
		fmt.Println("  - All installed alias packages")
		fmt.Println("  - The registry cache")
		fmt.Println("  - The entire ~/.ah directory")
		fmt.Println("")
		fmt.Print("Are you sure? Type 'DELETE' to confirm: ")

		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if response != "DELETE" {
			fmt.Println("Uninstall cancelled.")
			return
		}

		root, _ := manager.GetRootDir()
		fmt.Printf("Removing %s...\n", root)
		if err := os.RemoveAll(root); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("✅ Uninstall complete.")
		fmt.Println("NOTE: You must manually remove the 'ah init' lines from your .zshrc/.bashrc")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
