#!/bin/bash

set -e

# Get the directory of the current script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Define relevant paths
LOCALTRON_DIR="$SCRIPT_DIR/../localtron"
DOCS_SOURCE_DIR="$SCRIPT_DIR/../docs-source"
DOCS_DIR="$SCRIPT_DIR/../docs"
BUILD_DIR="$DOCS_SOURCE_DIR/build"
SWAGGER_FILE="$LOCALTRON_DIR/docs/swagger.yaml"
EXAMPLES_DIR="$DOCS_SOURCE_DIR/examples"
CNAME_FILE="$DOCS_SOURCE_DIR/CNAME"

# Error handler
trap 'echo "Error occurred in script at line $LINENO"; exit 1' ERR

# Initialize Swagger in localtron
echo "Initializing Swagger in $LOCALTRON_DIR"
cd "$LOCALTRON_DIR"
swag init

# Copy Swagger file to docs-source examples
echo "Copying Swagger file to $EXAMPLES_DIR"
cp "$SWAGGER_FILE" "$EXAMPLES_DIR/singulatron.yaml"

# Clean and generate API documentation
echo "Cleaning and generating API documentation"
cd "$DOCS_SOURCE_DIR"
yarn clean-api-docs singulatron
yarn gen-api-docs singulatron

# Build the project
echo "Building the project"
npm run build

# Clean and update docs directory
echo "Cleaning up old docs in $DOCS_DIR"
rm -rf "$DOCS_DIR"/*

echo "Copying CNAME file to $DOCS_DIR"
cp "$CNAME_FILE" "$DOCS_DIR/CNAME"

echo "Copying new build files to $DOCS_DIR"
cp -r "$BUILD_DIR"/* "$DOCS_DIR/"

echo "Documentation generation complete."
