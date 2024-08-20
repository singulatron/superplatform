#!/bin/bash

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Exit if any command fails
set -e

# Check if 'go' is installed
if ! command_exists go; then
    echo "Error: 'go' is not installed. Please install Go to proceed."
    exit 1
fi

# Check if 'npm' is installed
if ! command_exists npm; then
    echo "Error: 'npm' is not installed. Please install npm to proceed."
    exit 1
fi

# Install swag
echo "Installing swag..."
go install github.com/swaggo/swag/cmd/swag@latest

# Install openapi-generator-cli
echo "Installing openapi-generator-cli..."
npm install @openapitools/openapi-generator-cli -g

# Set the latest version of openapi-generator-cli
echo "Setting openapi-generator-cli to use the latest version..."
openapi-generator-cli version-manager set latest

# Run the generation scripts
echo "Running Go client generation script..."
bash clients/go/gen.sh

echo "Running JavaScript client generation script..."
bash clients/js/gen.sh

echo "Running documentation source generation script..."
bash docs-source/gen.sh

echo "All tasks completed successfully."
