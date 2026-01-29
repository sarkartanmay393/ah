#!/bin/bash
set -e

REPO="sarkartanmay393/alias-hub"
BINARY_NAME="ah"
INSTALL_DIR="/usr/local/bin"

# Detect OS and Arch
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

if [ "$ARCH" == "x86_64" ]; then
    ARCH="amd64"
elif [ "$ARCH" == "aarch64" ] || [ "$ARCH" == "arm64" ]; then
    ARCH="arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

echo "Detected: $OS $ARCH"
URL="https://github.com/$REPO/releases/latest/download/${BINARY_NAME}-${OS}-${ARCH}"

echo "Downloading $BINARY_NAME from $URL..."
curl -L -o "$BINARY_NAME" "$URL"
chmod +x "$BINARY_NAME"

echo "Installing to $INSTALL_DIR..."
sudo mv "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"

echo "Success! Run 'ah init' to get started."
