FROM ubuntu

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && apt-get install -y \
    vim \
    iputils-ping \
    nmap \
    traceroute \
    iproute2 \
    netcat \
    mtr \
    git \
    nghttp2

RUN apt-get update -y && apt-get install -y curl

RUN apt-get update -y && apt-get install -y gnupg2 wget
