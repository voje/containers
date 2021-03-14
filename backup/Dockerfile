FROM alpine:3.13

RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
        rsync \
        openssh-client \
        ca-certificates \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*
