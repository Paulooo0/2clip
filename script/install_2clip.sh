#!/bin/bash

VERSION="1.3.1"

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
	ARCH="amd64"
elif [ "$ARCH" = "arm64" ]; then
	ARCH="arm64"
else
	echo "Unsupported architecture: $ARCH"
	exit 1
fi

if [ "$OS" = "linux" ]; then
	BINARY="2clip-$VERSION-linux-$ARCH"
elif [ "$OS" = "darwin" ]; then
	BINARY="2clip-$VERSION-darwin-$ARCH"
else
	echo "Unsupported OS: $OS"
	exit 1
fi

URL="https://github.com/Paulooo0/2clip/releases/download/v$VERSION/$BINARY"
echo "Downloading $BINARY from $URL"
curl -L -o "$BINARY" "$URL"

chmod +x "$BINARY"

sudo mv "$BINARY" /usr/local/bin/2clip
echo "2clip v$VERSION installed successfully!"
