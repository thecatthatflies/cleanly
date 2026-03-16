package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type HistoryEntry struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type HistoryRun struct {
	Time    string         `json:"time"`
	Entries []HistoryEntry `json:"entries"`
}

var historyPath = filepath.Join(os.Getenv("HOME"), ".cleanly", "data", "history.json")

func Sort(basePath string, identified map[string]string) error {
	var entries []HistoryEntry

	for filename, category := range identified {
		from := filepath.Join(basePath, filename)
		to := filepath.Join(basePath, category, filename)

		err := os.Rename(from, to)
		if err != nil {
			return fmt.Errorf("could not move %s: %w", filename, err)
		}

		entries = append(entries, HistoryEntry{From: from, To: to})
		fmt.Printf("  moved %s → %s/\n", filename, category)
	}

	return saveHistory(entries)
}

func saveHistory(entries []HistoryEntry) error {
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

	out, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(historyPath, out, 0644)
}

func Undo() error {
	data, err := os.ReadFile(historyPath)
	if err != nil {
		return fmt.Errorf("no history found")
	}

	var history []HistoryRun
	err = json.Unmarshal(data, &history)
	if err != nil || len(history) == 0 {
		return fmt.Errorf("history is empty")
	}

	lastRun := history[len(history)-1]

	for _, entry := range lastRun.Entries {
		err := os.Rename(entry.To, entry.From)
		if err != nil {
			return fmt.Errorf("could not restore %s: %w", entry.From, err)
		}
		fmt.Printf("  restored %s\n", filepath.Base(entry.From))
	}

	seen := map[string]bool{}
	for _, entry := range lastRun.Entries {
		folder := filepath.Dir(entry.To)
		if seen[folder] {
			continue
		}
		seen[folder] = true

		contents, err := os.ReadDir(folder)
		if err == nil && len(contents) == 0 {
			os.Remove(folder)
		}
	}

	history = history[:len(history)-1]
	out, _ := json.MarshalIndent(history, "", "  ")
	os.WriteFile(historyPath, out, 0644)

	return nil
}
