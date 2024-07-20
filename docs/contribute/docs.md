---
title: Contribute to the Docs
nav_order: 1
parent: Contribute
---

# Contribute to the Docs

## Documentation site

In repo root:

```sh
cd docs
bundle install
bundle exec jekyll serve --config _config_development.yml
```

## Api documentation

Go to the `localtron` folder and run:

```sh
swag init && swagger serve --host=localhost --port=8080 ./docs/swagger.yaml
```

Edit the Go Swagger annotations and restart the above command to see your changes reflected.
