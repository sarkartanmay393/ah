package cmd

import (
	"fmt"

	"github.com/sarkartanmay393/ah/pkg/manager"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search packages in the registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		fmt.Printf("Searching for '%s'...\n", query)

		results, err := manager.SearchPackages(query)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if len(results) == 0 {
			fmt.Println("No matches found.")
			return
		}

		fmt.Println("\nFound Packages:")
		for _, pkg := range results {
			if pkg.Description != "" {
				fmt.Printf(" - %-15s : %s\n", pkg.Name, pkg.Description)
			} else {
				fmt.Printf(" - %s\n", pkg.Name)
			}
		}
		fmt.Printf("\nUse 'ah install <name>' to install.\n")
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
