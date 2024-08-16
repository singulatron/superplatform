#!/bin/bash

set -e

# Get the directory of the current script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Define relevant paths
LOCALTRON_DIR="$SCRIPT_DIR/../../localtron"
JS_CLIENT_DIR="$SCRIPT_DIR/../js"
TYPESCRIPT_CLIENT_DIR="$JS_CLIENT_DIR/client/src"
TYPESCRIPT_NODE_DIR="$JS_CLIENT_DIR/node/src"
LIBRARIES_DIR="$SCRIPT_DIR/../libraries"
SWAGGER_FILE="$LOCALTRON_DIR/docs/swagger.yaml"

# Error handler
trap 'echo "Error occurred in script at line $LINENO"; exit 1' ERR

# Initialize Swagger in localtron
echo "Initializing Swagger in $LOCALTRON_DIR"
cd "$LOCALTRON_DIR"
swag init

# Generate TypeScript Fetch client
echo "Generating TypeScript Fetch client in $TYPESCRIPT_CLIENT_DIR"
cd "$JS_CLIENT_DIR"
rm -r "$TYPESCRIPT_CLIENT_DIR"/* || true
openapi-generator-cli generate -i "$SWAGGER_FILE" -g typescript-fetch -o "$TYPESCRIPT_CLIENT_DIR"

# Generate TypeScript Node client
echo "Generating TypeScript Node client in $TYPESCRIPT_NODE_DIR"
rm -r "$TYPESCRIPT_NODE_DIR"/* || true
openapi-generator-cli generate -i "$SWAGGER_FILE" -g typescript-node -o "$TYPESCRIPT_NODE_DIR"

# Step into the node directory, install dependencies and build
echo "Installing dependencies and building in node directory"
cd "$JS_CLIENT_DIR/node"
npm install
npm run build

# Step into the client directory, install dependencies and build
echo "Installing dependencies and building in client directory"
cd "$JS_CLIENT_DIR/client"
npm install
npm run build

echo "All operations completed successfully."
