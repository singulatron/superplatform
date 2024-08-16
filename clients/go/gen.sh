#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

LOCALTRON_DIR="$SCRIPT_DIR/../../localtron"

cd "$LOCALTRON_DIR"
swag init

cd "$SCRIPT_DIR"
openapi-generator-cli generate -i "$LOCALTRON_DIR/docs/swagger.yaml" -g go -o .
rm -rf api docs go.mod
cp go.mod.template go.mod