cd ../localtron
swag init
cd ../docs-source
cp ../localtron/docs/swagger.yaml examples/singulatron.yaml
npm run build
yarn gen-api-docs singulatron
rm -rf ../docs/*
cp CNAME ../docs/CNAME
cp -r ./build/* ../docs/