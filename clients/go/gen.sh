#!/bin/bash

set -e

cd ../../localtron;
swag init

cd ../clients/go;
openapi-generator-cli generate -i ../../localtron/docs/swagger.yaml -g go -o .
rm -rf api docs go.mod
cp go.mod.template go.mod
