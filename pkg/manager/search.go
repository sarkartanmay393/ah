package manager

import (
	"os"
	"path/filepath"
	"strings"
)

type ValidPackage struct {
	Name        string
	Description string
}

func SearchPackages(query string) ([]ValidPackage, error) {
	// Auto-update registry ensures we search fresh data
	if err := UpdateRegistry(); err != nil {
		return nil, err
	}

	// ~/.ah/registry/registry
	baseDir, err := GetRegistryContentDir()
	if err != nil {
		return nil, err
	}

	var matches []ValidPackage

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []ValidPackage{}, nil
		}
		return nil, err
	}

	query = strings.ToLower(query)

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		// Load Metadata
		pkgPath := filepath.Join(baseDir, e.Name())
		meta, err := LoadMetadata(pkgPath)
		if err != nil {
			// Skip invalid packages (missing ah.yaml or too large)
			continue
		}

		nameMatch := strings.Contains(strings.ToLower(e.Name()), query)
		descMatch := strings.Contains(strings.ToLower(meta.Description), query)

		if nameMatch || descMatch {
			matches = append(matches, ValidPackage{Name: e.Name(), Description: meta.Description})
		}
	}
	return matches, nil
}
