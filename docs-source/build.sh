cd ../localtron
swag init
cd ../docs-source
cp ../localtron/docs/swagger.yaml examples/singulatron.yaml
npm run build