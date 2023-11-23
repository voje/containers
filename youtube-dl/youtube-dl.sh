#!/bin/bash
# Shim for youtube-downloader

VERSION="aaed488"
IMG="voje/youtube-dl:$VERSION"
CRI="podman"

"$CRI" run \
    --rm \
    -it \
    -v "$PWD":/project \
    -u "$(id -u):$(id -g)" \
    "$IMG" /bin/sh
