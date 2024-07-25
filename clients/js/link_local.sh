#!/bin/bash

set -e

cd types;
npm install
npm run build;
npm link;
cd ..

cd node;
npm link @singulatron/types;
npm install
npm run build;
npm link;
cd ..

cd client;
npm link @singulatron/types;
npm install
npm run build;
npm link;
cd ..

cd client-example;
npm link @singulatron/types;
npm link @singulatron/client;
npm install
npm run build
