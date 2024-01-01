#!/bin/bash

GINIT_BINARY="https://github.com/rudrodip/ginit/releases/download/pre-release/ginit"
INSTALL_DIR="$HOME/.local/bin"
BIN_NAME="ginit"

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Check if curl is installed
if ! command -v curl &> /dev/null; then
    echo -e "${RED}Error: curl is not installed. Please install curl before running this script.${NC}"
    exit 1
fi

# Check if running on Windows
if [ -n "$WINDIR" ]; then
    echo -e "${RED}Error: This script is intended for Linux. Please use WSL (Windows Subsystem for Linux) to run it.${NC}"
    exit 1
fi

# Create the installation directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Download the binary with progress bar
curl -# -fSL "$GINIT_BINARY" -o "$BIN_NAME"

# Check if the download was successful
if [ $? -ne 0 ]; then
    echo -e "${RED}Error downloading the ginit binary.${NC}"
    rm -f "$BIN_NAME"  # Remove the binary if an error occurred
    exit 1
fi

# Make it executable
chmod +x "$BIN_NAME"

# Move it to the installation directory
mv "$BIN_NAME" "$INSTALL_DIR/"
if [ $? -eq 0 ]; then
    echo -e "${GREEN}$BIN_NAME installed successfully in $INSTALL_DIR${NC}"
    echo -e "${GREEN}Make sure to add $INSTALL_DIR to your PATH.${NC}"
    echo -e "${BLUE}\nRun 'ginit' in terminal\n${NC}"
else
    echo -e "${RED}Error moving $BIN_NAME to $INSTALL_DIR. Please check permissions.${NC}"
    rm -f "$INSTALL_DIR/$BIN_NAME"  # Remove the binary if the move operation failed
    exit 1
fi
