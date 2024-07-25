#!/bin/bash

set -e

cd types;
npm install
npm run build;
npm link;
cd ..

cd node;
npm install
npm link @singulatron/types;
npm run build;
npm link;
cd ..

cd client;
npm install
npm link @singulatron/types;
npm run build;
npm link;
cd ..

cd client-example;
npm install
npm link @singulatron/types;
npm link @singulatron/client;
npm run build
