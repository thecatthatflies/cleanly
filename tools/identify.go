package tools

import (
	"github.com/thecatthatflies/cleanly/config"
	"os"
	"path/filepath"
	"strings"
)

func Identify(files []os.DirEntry, filterCategories []string) map[string]string {
	result := map[string]string{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		category, ok := config.Categories[ext]
		if !ok {
			category = "Other"
		}

		if len(filterCategories) > 0 && !containsCategory(filterCategories, strings.ToLower(category)) {
			continue
		}

		result[file.Name()] = category
	}

	return result
}

func containsCategory(filter []string, category string) bool {
	for _, f := range filter {
		if f == category {
			return true
		}
	}
	return false
}
