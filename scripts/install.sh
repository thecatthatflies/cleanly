#!/bin/bash

set -e

echo "installing cleanly..."
go build -o cleanly .

echo "installing to /usr/local/bin..."
sudo mv cleanly /usr/local/bin/cleanly

echo "creating config directory..."
mkdir -p ~/.cleanly/data
cp data/history.json ~/.cleanly/data/history.json 2>/dev/null || echo "[]" > ~/.cleanly/data/history.json

echo "cleanly installed. run 'cleanly .' to get started."