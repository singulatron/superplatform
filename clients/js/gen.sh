#!/bin/bash

set -e

cd ../../localtron;
swag init

cd ../clients/js;
rm -r client/src/*
openapi-generator-cli generate -i ../../localtron/docs/swagger.yaml -g typescript-fetch -o client/src/

rm -r node/src/*
openapi-generator-cli generate -i ../../localtron/docs/swagger.yaml -g typescript-node -o node/src/
