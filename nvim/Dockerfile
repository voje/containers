FROM ubuntu:22.04

WORKDIR /root

RUN apt-get update && apt-get install -y \
  curl

# Install Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs > rustup.sh && sh rustup.sh -y && rm rustup.sh

