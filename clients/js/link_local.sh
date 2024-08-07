#!/bin/bash

set -e

cd client;
npm install
npm run build;
npm link;
cd ..

cd client-example;
npm install
npm link @singulatron/client;
npm run build
cd ..

cd ../../desktop/workspaces/angular-app/
npm install
npm link @singulatron/client
cd ../../../clients/js