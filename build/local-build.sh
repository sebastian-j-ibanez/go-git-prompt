#!/bin/bash

# May 8 2025
# Sebastian J Ibanez
# This script builds and copies the go-git-prompt binary to ~/.local/bin

echo "[INFO] Building go-git-prompt binary..."
build_time=$( { time go build; } 2>&1 )
real_time=$( echo "$build_time" | grep real | awk '{print $2}' )

echo "[INFO] Copying go-git-prompt binary to ~/.local/bin..."
if ! [ -f ~/.local/bin ]; then
    mkdir -p ~/.local/bin
fi
cp go-git-prompt ~/.local/bin
rm ./go-git-prompt

echo "[INFO] Setup complete!"
echo "[INFO] Build time: $real_time"