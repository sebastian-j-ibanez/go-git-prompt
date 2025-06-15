#!/bin/bash

output_dir="build"
mkdir -p "$output_dir"

# Windows build
echo "Building for Windows x64..."
GOOS=windows GOARCH=amd64 go build -o "$output_dir/go-git-prompt_windows_amd64.exe"
if [ $? -ne 0 ]; then
    echo "Error building for Windows."
    exit 1
fi
echo "Windows build complete."

# Linux build
echo "Building for Linux x64..."
GOOS=linux GOARCH=amd64 go build -o "$output_dir/go-git-prompt_linux_amd64"
if [ $? -ne 0 ]; then
    echo "Error building for Linux."
    exit 1
fi
echo "Linux build complete."

echo "All builds complete."
