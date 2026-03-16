# getting started

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
chmod +x scripts/install.sh
./scripts/install.sh
```

## your first sort

navigate to the folder you want to sort:

```bash
cd ~/Downloads
cleanly .
```

cleanly will create subfolders like `Images/`, `Documents/`, `Video/` and move your files into them.

to reverse it:

```bash
cleanly --undo
```