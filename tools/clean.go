package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

func Clean(basePath string, identified map[string]string) error {
	var failed []string

	for filename, category := range identified {
		expected := filepath.Join(basePath, category, filename)
		if _, err := os.Stat(expected); os.IsNotExist(err) {
			failed = append(failed, filename)
		}
	}

	if len(failed) > 0 {
		fmt.Println("some files were not moved, retrying...")
		for _, filename := range failed {
			category := identified[filename]
			from := filepath.Join(basePath, filename)
			to := filepath.Join(basePath, category, filename)
			err := os.Rename(from, to)
			if err != nil {
				return fmt.Errorf("could not move %s on retry: %w", filename, err)
			}
			fmt.Printf("  retried %s → %s/\n", filename, category)
		}
	}

	return removeEmptyFolders(basePath)
}

func removeEmptyFolders(basePath string) error {
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		folderPath := filepath.Join(basePath, entry.Name())
		contents, err := os.ReadDir(folderPath)
		if err != nil {
			continue
		}

		if len(contents) == 0 {
			os.Remove(folderPath)
		}
	}

	return nil
}
