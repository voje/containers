FROM alpine:3.13

RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
        rsync \
        openssh-client \
        ca-certificates \
	run-parts \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*

COPY project /project

CMD ["/bin/sh", "/project/entrypoint.sh"]
