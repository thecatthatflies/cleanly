#!/bin/bash

set -e

echo "removing cleanly..."
sudo rm -f /usr/local/bin/cleanly

read -p "remove history and config? (~/.cleanly) [y/N] " confirm
if [[ $confirm == "y" || $confirm == "Y" ]]; then
    rm -rf ~/.cleanly
    echo "config removed."
fi

echo "cleanly uninstalled."