# usage

## basic

```bash
cleanly .                    # sort current directory
cleanly ~/Downloads          # sort a specific folder
```

## flags

### `-c` — filter by category

only sort specific categories:

```bash
cleanly . -c "images audio"
cleanly . -c "documents"
```

category names are case insensitive. `Images`, `images`, and `IMAGES` all work.

### `--no-clean`

skip the cleanup verification step after sorting. useful if you want speed over safety:

```bash
cleanly . --no-clean
```

### `--undo`

reverse the last sort run. restores all files to their original locations and removes the folders that were created:

```bash
cleanly --undo
```

### `-v` / `version`

print the current version:

```bash
cleanly -v
cleanly version
```

### `update`

update cleanly to the latest version:

```bash
cleanly update
```

### `help`

print usage information:

```bash
cleanly help
```