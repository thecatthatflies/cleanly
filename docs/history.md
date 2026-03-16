# history & undo

## how history works

every time you run `cleanly`, it saves a record of every file that was moved to:

```
~/.cleanly/data/history.json
```

each run is stored as an entry with a timestamp and a list of file moves:

```json
[
  {
    "time": "2025-01-01T12:00:00Z",
    "entries": [
      {
        "from": "/Users/you/Downloads/photo.jpg",
        "to": "/Users/you/Downloads/Images/photo.jpg"
      }
    ]
  }
]
```

## undoing a sort

to reverse the last sort:

```bash
cleanly --undo
```

this moves every file back to where it came from and removes any folders that are now empty.

only the most recent run is undone at a time. run `--undo` multiple times to go further back.

## remove history

`cleanly remove` also saves to history, with one difference — permanently deleted files (`-p`) are recorded as `PERMANENT` and cannot be undone.