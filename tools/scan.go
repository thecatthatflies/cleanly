package tools

import (
	"log"
	"os"
)

func scan(path string) []os.DirEntry {

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}
