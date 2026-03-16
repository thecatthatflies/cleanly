# cleanly
a simple CLI tool that sorts any folder by file type, category, or custom rules

## install

**homebrew**
```bash
brew install thecatthatflies/tap/cleanly
```

**go install**
```bash
go install github.com/thecatthatflies/cleanly@latest
```

**manual**
```bash
git clone https://github.com/thecatthatflies/cleanly
cd cleanly
chmod +x scripts/install.sh scripts/uninstall.sh
./scripts/install.sh
```

## usage

```bash
cleanly .                          # sort current directory
cleanly ~/Downloads                # sort a specific folder
cleanly . -c "images audio"        # only sort specific categories
cleanly . --no-clean               # skip cleanup verification
cleanly --undo                     # reverse the last sort
cleanly remove -f png jpeg         # trash files by extension
cleanly remove -c images           # trash files by category
cleanly remove -f png -p           # permanently delete
```

## categories

| category  | extensions                                      |
|-----------|-------------------------------------------------|
| Images    | jpg, jpeg, png, gif, webp, svg, psd, fig...     |
| Audio     | mp3, wav, flac, aac, ogg, m4a...                |
| Video     | mp4, mkv, mov, avi, webm...                     |
| Documents | pdf, docx, txt, md, xlsx, pptx, epub...         |
| Archives  | zip, rar, 7z, tar, gz, dmg, iso...              |
| Apps      | exe, dmg, apk, ipa, deb, pkg...                 |
| Data      | json, xml, yaml, sql, sqlite...                 |
| Code      | go, js, ts, py, html, css, sh...                |
| Other     | anything unrecognized                           |

> full list of all file types can be found in [config.go](https://github.com/thecatthatflies/cleanly/config/config.go)

## flags

| flag          | description                          |
|---------------|--------------------------------------|
| `-c`          | only sort specific categories        |
| `--no-clean`  | skip cleanup verification            |
| `--undo`      | reverse the last sort                |
| `-v`          | print version                        |
| `help`        | show help                            |
| `update`      | update to latest version             |
| `remove -f`   | remove files by extension            |
| `remove -c`   | remove files by category             |
| `remove -p`   | permanently delete instead of trash  |

## license

MIT