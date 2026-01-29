package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/bash-shortcuts/ah/pkg/manager"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed alias packages",
	Run: func(cmd *cobra.Command, args []string) {
		packages, err := manager.ListPackages()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		if len(packages) == 0 {
			fmt.Println("No alias packages installed.")
			return
		}

		fmt.Println("Installed Packages:")
		root, _ := manager.GetRootDir()

		for _, pkg := range packages {
			path := filepath.Join(root, manager.ActiveDir, pkg)
			meta, err := manager.LoadMetadata(path)
			if err == nil && meta.Description != "" {
				fmt.Printf(" - %-15s : %s\n", pkg, meta.Description)
			} else {
				fmt.Printf(" - %s\n", pkg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
