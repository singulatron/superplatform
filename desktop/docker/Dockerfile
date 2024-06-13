FROM node:22 as build

WORKDIR /app
COPY . .

RUN npm i && npm run package:angular-app

FROM nginx:alpine

COPY --from=build /app/workspaces/angular-app/.dist/angular-app /usr/share/nginx/html

COPY ./docker/entrypoint.sh /entrypoint.sh
COPY ./docker/nginx.conf /etc/nginx/nginx.conf

RUN if [ ! -f /entrypoint.sh ]; then \
    echo "ERROR: entrypoint.sh not found"; \
    exit 1; \
    fi

RUN chmod +x /entrypoint.sh

EXPOSE 80

ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]
