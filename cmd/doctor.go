package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bash-shortcuts/ah/pkg/manager"
	"github.com/spf13/cobra"
)

var doctorFix bool

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system health and dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running doctor...")

		// Auto-fix/Ensure environment is consistent
		if err := manager.EnsureDirs(); err != nil {
			fmt.Printf("[ERROR] Failed to ensure directories: %v\n", err)
			return
		}

		// Check 1: Directory Structure
		root, err := manager.GetRootDir()
		if err != nil {
			fmt.Printf("[FAIL] Could not determine home directory: %v\n", err)
			return
		}

		if _, err := os.Stat(root); os.IsNotExist(err) {
			fmt.Printf("[FAIL] Root directory %s does not exist.\n", root)
			if doctorFix {
				fmt.Println("  -> Creating directories...")
				if err := manager.EnsureDirs(); err != nil {
					fmt.Printf("  [FAIL] Failed to create directories: %v\n", err)
				} else {
					fmt.Println("  [OK] Fixed.")
				}
			}
		} else {
			fmt.Printf("[OK] Root directory exists at %s\n", root)
		}

		// Check 2: Dependencies
		if _, err := exec.LookPath("git"); err != nil {
			fmt.Println("[FAIL] 'git' is not installed or not in PATH.")
		} else {
			fmt.Println("[OK] 'git' is installed.")
		}
	},
}

func init() {
	doctorCmd.Flags().BoolVar(&doctorFix, "fix", false, "Attempt to fix found issues automatically")
	rootCmd.AddCommand(doctorCmd)
}
