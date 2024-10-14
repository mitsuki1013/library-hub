FROM alpine:latest AS builder

RUN set -ex; \
        apk update; \
        apk add --no-cache --virtual .build-deps \
            curl \
            gnupg; \
        curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz; \
        apk del --no-network .build-deps; \
        mv migrate /usr/local/bin/migrate

ENTRYPOINT ["/bin/sh"]