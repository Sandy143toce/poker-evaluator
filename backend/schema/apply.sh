#!/bin/bash
set -euo pipefail

# Variables for PostgreSQL connection with default values if not set
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-postgres}"
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_NAME="${DB_NAME:-poker_evaluator}"
SSL_MODE="${SSL_MODE:-disable}"

# Dynamically generate the PostgreSQL URL
POSTGRES_URL="postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$SSL_MODE"

# Function to install Atlas on Linux
install_atlas_linux() {
    # Check if curl is installed
    if ! command -v curl &> /dev/null; then
        echo "curl is not installed. Installing curl..."
        if [[ -f /etc/debian_version ]]; then
            sudo apt-get update && sudo apt-get install -y curl
        elif [[ -f /etc/redhat-release ]]; then
            sudo yum install -y curl
        else
            echo "Unsupported Linux distribution. Please install curl manually."
            exit 1
        fi
    fi

    # Check if sh is available (though most systems will have it)
    if ! command -v sh &> /dev/null; then
        echo "'sh' shell is not found. Please install 'sh' manually."
        exit 1
    fi

    # Install Atlas
    echo "Installing Atlas on Linux..."
    curl -sSf https://atlasgo.sh | sh
}

# Function to install Atlas based on OS
install_atlas() {
    if [[ "$OSTYPE" == "darwin"* ]]; then
        echo "Installing Atlas on macOS using Homebrew..."
        brew install ariga/tap/atlas
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        install_atlas_linux
    else
        echo "Unsupported OS: $OSTYPE"
        exit 1
    fi
}

# Check if 'atlas' command is available
if command -v atlas &> /dev/null; then
    echo "Atlas is already installed."
else
    echo "Atlas command not found, installing..."
    install_atlas
fi

# Run Atlas schema apply command
echo "Running Atlas schema apply..."

atlas schema apply --url "$POSTGRES_URL" --to "file://schema.hcl" --auto-approve

# Check if the last command was successful
if [ $? -eq 0 ]; then
    echo "Atlas schema applied successfully."
else
    echo "Failed to apply Atlas schema."
    exit 1
fi