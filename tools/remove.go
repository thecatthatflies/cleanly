package tools

import (
	"github.com/thecatthatflies/cleanly/config"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Remove(args []string) {
	if len(args) < 2 {
		fmt.Println("usage: cleanly remove -f png jpeg | cleanly remove -c images audio")
		return
	}

	permanent := false
	filteredArgs := []string{}
	for _, a := range args {
		if a == "--permanent" || a == "-p" {
			permanent = true
		} else {
			filteredArgs = append(filteredArgs, a)
		}
	}
	args = filteredArgs

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
		category, _ := config.Categories[ext]

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

	fmt.Printf("about to delete %d files. are you sure? [y/N] ", len(toDelete))
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToLower(confirm) != "y" {
		fmt.Println("cancelled.")
		return
	}

	trashDir := filepath.Join(os.Getenv("HOME"), ".Trash")
	var entries []HistoryEntry

	for _, filename := range toDelete {
		from := filepath.Join(path, filename)

		if permanent {
			err := os.Remove(from)
			if err != nil {
				fmt.Println("error removing", filename, ":", err)
				continue
			}
			fmt.Println("  permanently deleted", filename)
			entries = append(entries, HistoryEntry{From: from, To: "PERMANENT"})
		} else {
			to := filepath.Join(trashDir, filename)
			err := os.Rename(from, to)
			if err != nil {
				fmt.Println("error trashing", filename, ":", err)
				continue
			}
			fmt.Println("  trashed", filename)
			entries = append(entries, HistoryEntry{From: from, To: to})
		}
	}

	saveRemoveHistory(entries)
}

func saveRemoveHistory(entries []HistoryEntry) error {
	err := os.MkdirAll(filepath.Dir(historyPath), 0755)
	if err != nil {
		return err
	}

	var history []HistoryRun
	data, err := os.ReadFile(historyPath)
	if err == nil {
		json.Unmarshal(data, &history)
	}

	history = append(history, HistoryRun{
		Time:    time.Now().Format(time.RFC3339),
		Entries: entries,
	})

	out, _ := json.MarshalIndent(history, "", "  ")
	return os.WriteFile(historyPath, out, 0644)
}