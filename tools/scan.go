package tools

import (
	"log"
	"os"
)

func Scan(path string) []os.DirEntry {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("error reading directory:", err)
	}
	return files
}
