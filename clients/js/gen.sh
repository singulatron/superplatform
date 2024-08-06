#!/bin/bash

set -e

cd ../../localtron;
swag init

cd ../clients/js;
openapi-generator-cli generate -i ../../localtron/docs/swagger.yaml -g typescript-fetch -o client/src/
