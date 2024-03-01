#!/bin/bash

# Detect the package manager and call the respective install script
if command -v apt &> /dev/null; then
    ./apt/install.sh
elif command -v pacman &> /dev/null; then
    ./pacman/install.sh
else
    echo "Unsupported package manager. Exiting."
    exit 1
fi
