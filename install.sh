#!/usr/bin/env bash

# Get the latest release version
LATEST_RELEASE=$(curl -s https://api.github.com/repos/brettearle/termadoro/releases/latest | jq -r .tag_name)

# Define the GitHub release URL
GITHUB_URL="https://github.com/brettearle/termadoro/releases/download/$LATEST_RELEASE"

# Detect OS and architecture
OS=$(uname -s)
ARCH=$(uname -m)

# Set binary name based on OS and architecture
if [[ "$OS" == "Linux" && "$ARCH" == "x86_64" ]]; then
    BIN_NAME="termadoro-linux-amd64"
elif [[ "$OS" == "Darwin" && "$ARCH" == "x86_64" ]]; then
    BIN_NAME="termadoro-darwin-amd64"
elif [[ "$OS" == "Darwin" && "$ARCH" == "arm64" ]]; then
    BIN_NAME="termadoro-darwin-arm64"
else
    echo "Unsupported OS or architecture. Exiting."
    exit 1
fi

# Download the appropriate binary
echo "Downloading $BIN_NAME..."
curl -LO "$GITHUB_URL/$BIN_NAME"

# Make it executable
chmod +x "$BIN_NAME"

# Move to /usr/local/bin as 'tdoro'
echo "Installing as /usr/local/bin/tdoro"
sudo mv "$BIN_NAME" /usr/local/bin/tdoro

echo "✅ Installation complete! You can now run 'tdoro'"
