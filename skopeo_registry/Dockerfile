FROM golang:1.12-alpine AS go_builder

# Build skopeo
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

# Build tools
RUN mkdir /tools
WORKDIR /tools
ADD ./tools /tools/
RUN go build -o tools /tools/cmd/tools/main.go

# Take docker registry:2 as base, add go binaries.   
FROM registry:2

# Add skopeo
RUN apk add --no-cache ca-certificates
COPY --from=go_builder /go/src/github.com/containers/skopeo/skopeo /usr/local/bin/skopeo
COPY --from=go_builder /go/src/github.com/containers/skopeo/default-policy.json /etc/containers/policy.json

# Add tools
COPY --from=go_builder /tools/tools /usr/bin/tools

VOLUME ["/var/lib/registry"]
EXPOSE 5000
ENTRYPOINT ["registry", "serve", "/etc/docker/registry/config.yml"]

CMD ["skopeo", "-h", "tools", "-h"]
