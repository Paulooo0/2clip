#!/bin/bash

# Define the version you want to install
VERSION="v1.0.1"

# Determine the OS and architecture
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64"
fi

# Determine the download URL
if [ "$OS" = "linux" ]; then
  BINARY="2clip-linux-$ARCH"
elif [ "$OS" = "darwin" ]; then
  BINARY="2clip-darwin-$ARCH"
else
  echo "Unsupported OS: $OS"
  exit 1
fi

# Download the binary
URL="https://github.com/Paulooo0/2clip/releases/download/$VERSION/$BINARY"
echo "Downloading $BINARY from $URL"
curl -L -o "$BINARY" "$URL"

# Make the binary executable
chmod +x "$BINARY"

# Move the binary to /usr/local/bin (Linux/macOS) or another directory in PATH
sudo mv "$BINARY" /usr/local/bin/2clip
echo "2clip installed to /usr/local/bin/2clip"
