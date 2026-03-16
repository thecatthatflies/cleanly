package tools

import (
	"os"
	"path/filepath"
)

func MakeFolders(basePath string, identified map[string]string) error {
	seen := map[string]bool{}

	for _, category := range identified {
		if seen[category] {
			continue
		}
		seen[category] = true

		folderPath := filepath.Join(basePath, category)
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
