#!/bin/bash

set -e

bash clients/go/gen.sh
bash clients/js/gen.sh
bash docs-source/gen.sh