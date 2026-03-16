# AGENTS.md

Instructions for AI agents working on this codebase.

---

## Project Overview

cleanly is a CLI tool written in Go and bash that sorts folders by file type. it moves files into category subfolders, saves a history of every run, and supports undoing the last sort.

---

## Structure

```
main.go              # entry point, subcommand routing, flag handling
config/
└── config.go        # extension → category map
tools/
├── input.go         # CLI flag parsing, Input struct
├── scan.go          # reads directory, returns []os.DirEntry
├── identify.go      # maps files to categories using config
├── make.go          # creates category folders
├── sort.go          # moves files, saves history, handles --undo
├── clean.go         # verifies moves, retries failures, removes empty folders
└── remove.go        # removes files by extension or category, saves history
data/
└── history.json     # starter file, real history lives in ~/.cleanly/data/
docs/                # documentation pages
scripts/
├── install.sh       # builds and installs to /usr/local/bin
└── uninstall.sh     # removes binary and optionally ~/.cleanly
```

---

## Flow

```
main.go → input.go → make.go → scan.go → identify.go → sort.go → clean.go
```

subcommands break out of the flow early:
- `help` — prints usage and exits
- `version` / `-v` — prints version and exits
- `update` — runs go install and exits
- `remove` — calls tools.Remove() and exits

---

## Language & Style

- Go only. no CGo.
- all files in `tools/` use `package tools`
- `config/config.go` uses `package config`
- `main.go` uses `package main`
- run `gofmt` before committing
- no unused imports, no dead code
- exported functions need doc comments

---

## Key Rules

1. only `sort.go` and `remove.go` write to history. nothing else touches `~/.cleanly/data/history.json`
2. `remove.go` saves `"PERMANENT"` as the `To` field for permanently deleted files. `Undo()` in `sort.go` must skip these entries
3. `identify.go` defaults unrecognized extensions to `"Other"`
4. `clean.go` only removes folders that are empty after sorting — never remove folders with files in them
5. all user-facing output uses `fmt.Println` or `fmt.Printf`, never `log.Fatal` except for unrecoverable errors in `scan.go`

---

## History Format

```json
[
  {
    "time": "2025-01-01T00:00:00Z",
    "entries": [
      { "from": "/Users/user/Downloads/photo.jpg", "to": "/Users/user/Downloads/Images/photo.jpg" }
    ]
  }
]
```

---

## Adding a New Category

1. add extensions to `config/config.go`
2. that's it — `identify.go`, `sort.go`, and `make.go` all read from config dynamically

---

## What NOT to Do

- do not hardcode paths — always use `os.Getenv("HOME")` or the path passed in
- do not delete files permanently unless `-p` flag is explicitly passed
- do not write to history from `clean.go`, `make.go`, or `scan.go`
- do not add Windows or Linux specific code — macOS only for now
- do not change the `HistoryEntry` or `HistoryRun` struct fields without updating `Undo()` and `saveRemoveHistory()`

---

## Building

```bash
go build -o cleanly .
go run . .
```

## Installing

```bash
chmod +x scripts/install.sh
./scripts/install.sh
```