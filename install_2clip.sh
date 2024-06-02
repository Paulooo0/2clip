#!/bin/bash

# Ensure the script exits if any command fails
set -e

# Compile the Go tool
echo "Compiling 2clip..."
go build -o 2clip cmd/2clip/main.go

# Move the compiled binary to /usr/local/bin
echo "Moving 2clip to /usr/local/bin..."
sudo mv 2clip /usr/local/bin/

# Ensure /usr/local/bin is in PATH
if [[ ":$PATH:" != *":/usr/local/bin:"* ]]; then
    echo "Adding /usr/local/bin to PATH..."
    echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc
    source ~/.bashrc
fi

echo "2clip installation complete. You can now use '2clip' from anywhere in your terminal."