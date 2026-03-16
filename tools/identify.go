package tools

import (
	"os"
	"path/filepath"
	"cleanly/config"
)

func identify(files []os.DirEntry) map[string]string {

	result := map[string]string{}
	
	for _, file := range files {
		ext := filepath.Ext(file.Name())
    		category := config.Categories[ext]
    		result[file.Name()] = category
	}
	return result
}