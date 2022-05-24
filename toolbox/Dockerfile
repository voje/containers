FROM rust as rust-builder

WORKDIR /git
RUN git clone https://github.com/svenstaro/dummyhttp.git
WORKDIR /git/dummyhttp
RUN cargo build 
RUN ls target/debug/dummyhttp

FROM ubuntu

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && apt-get install -y \
    vim \
    iputils-ping \
    nmap \
    traceroute \
    iproute2 \
    netcat \
    ncat \
    mtr \
    git \
    nghttp2 \
    dnsutils \
    curl \
    gnupg2 \
    wget \
    tcpdump 

RUN apt-get update -y && apt-get install -y \
    jq \
    tmux \
    socat \
    arping

RUN apt-get update -y && apt-get install -y \
  iperf3 \
  iptables \
  nftables

COPY --from=rust-builder /git/dummyhttp/target/debug/dummyhttp /usr/local/bin/.

