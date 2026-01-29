package manager

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	RegistryDir = "registry"
)

// UpdateRegistry ensures the registry is cloned and up to date
func UpdateRegistry() error {
	repoURL := os.Getenv("AH_REGISTRY_URL")
	if repoURL == "" {
		repoURL = RegistryRepo // now defined in manager.go (wait, need to export it or move it)
	}
	// Default to built-in path if env var not set

	root, err := GetRootDir()
	if err != nil {
		return err
	}
	registryPath := filepath.Join(root, RegistryDir)

	if _, err := os.Stat(registryPath); os.IsNotExist(err) {
		// Clone
		fmt.Printf("Cloning registry from %s...\n", repoURL)
		cmd := exec.Command("git", "clone", repoURL, registryPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	// Pull
	fmt.Println("Updating registry...")
	cmd := exec.Command("git", "-C", registryPath, "pull")
	// Capture stderr to show detail if needed, but don't fail hard
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to update registry (using cached data): %v\n", err)
		return nil // Soft fail: proceed with existing data
	}
	return nil
}

// GetRegistryContentDir returns the path where the actual packages are located (~/.ah/registry/registry)
func GetRegistryContentDir() (string, error) {
	root, err := GetRootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, RegistryDir, "registry"), nil
}

// GetRegistryPackagePath returns the absolute path to a package in the local registry
func GetRegistryPackagePath(packageName string) (string, error) {
	contentDir, err := GetRegistryContentDir()
	if err != nil {
		return "", err
	}

	pkgPath := filepath.Join(contentDir, packageName)

	if _, err := os.Stat(pkgPath); os.IsNotExist(err) {
		return "", fmt.Errorf("package '%s' not found in registry", packageName)
	}
	return pkgPath, nil
}
