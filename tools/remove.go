package tools

import (
	"cleanly/config"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Remove(args []string) {
	if len(args) < 2 {
		fmt.Println("usage: cleanly remove -f png jpeg | cleanly remove -c images audio")
		return
	}

	flag := args[0]
	targets := args[1:]
	path := "."

	files := Scan(path)
	var toDelete []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		category := ""
		if cat, ok := config.Categories[ext]; ok {
			category = cat
		}

		switch flag {
		case "-f":
			for _, t := range targets {
				if ext == "."+strings.ToLower(t) {
					toDelete = append(toDelete, file.Name())
				}
			}
		case "-c":
			for _, t := range targets {
				   if category != "" && strings.EqualFold(strings.TrimSpace(category), strings.TrimSpace(t)) {
					toDelete = append(toDelete, file.Name())
				}
			}
		}
	}

	if len(toDelete) == 0 {
		fmt.Println("no files matched.")
		return
	}

	for _, filename := range toDelete {
		err := os.Remove(filepath.Join(path, filename))
		if err != nil {
			fmt.Println("error removing", filename, ":", err)
			continue
		}
		fmt.Println("  removed", filename)
	}
}
