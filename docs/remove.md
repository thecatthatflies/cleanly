# remove

the `remove` subcommand deletes files from the current directory by extension or category.

## by extension

```bash
cleanly remove -f png jpeg gif
```

removes all `.png`, `.jpeg`, and `.gif` files in the current directory.

## by category

```bash
cleanly remove -c images audio
```

removes all files that belong to the `images` or `audio` categories.

## trash vs permanent

by default, files are moved to `~/.Trash` so you can recover them if needed.

to permanently delete instead:

```bash
cleanly remove -f png -p
cleanly remove -c images --permanent
```

permanently deleted files cannot be undone.

## confirmation

cleanly will always ask for confirmation before removing files:

```
about to delete 12 files. are you sure? [y/N]
```

type `y` to confirm.