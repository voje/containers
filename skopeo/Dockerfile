FROM golang:1.12-alpine AS builder

RUN apk add --no-cache \
    git \
    make \
    gcc \
    musl-dev \
    btrfs-progs-dev \
    lvm2-dev \
    gpgme-dev \
    glib-dev || apk update && apk upgrade

WORKDIR /go/src/github.com/containers/skopeo
RUN git clone https://github.com/containers/skopeo.git .
RUN git checkout v0.1.41
RUN make binary-local-static DISABLE_CGO=1


FROM alpine:3.11.3
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/containers/skopeo/skopeo /usr/local/bin/skopeo
COPY --from=builder /go/src/github.com/containers/skopeo/default-policy.json /etc/containers/policy.json
