package manager

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type PackageMetadata struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Author      string `yaml:"author"`
	Website     string `yaml:"website"`
}

func LoadMetadata(packageDir string) (*PackageMetadata, error) {
	path := filepath.Join(packageDir, "ah.yaml")

	// Security: Prevent reading huge files
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
	if info.Size() > 10*1024 { // 10KB Limit
		return nil, fmt.Errorf("ah.yaml is too large (max 10KB)")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var meta PackageMetadata
	if err := yaml.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	if meta.Name == "" || meta.Version == "" {
		return nil, fmt.Errorf("ah.yaml must contain 'name' and 'version'")
	}

	return &meta, nil
}
