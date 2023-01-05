FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
  wget ripgrep git

WORKDIR /root/downloads

# Install tree-sitter
RUN wget https://github.com/tree-sitter/tree-sitter/releases/download/v0.20.7/tree-sitter-linux-x64.gz && \
  gzip -d tree-sitter-linux-x64.gz && \
  install tree-sitter-linux-x64 /usr/local/bin/tree-sitter -m 755 && \
  rm tree-sitter-linux-x64

RUN wget https://github.com/neovim/neovim/releases/latest/download/nvim.appimage && \
  chmod u+x nvim.appimage && \
  ./nvim.appimage --appimage-extract && \
  rm nvim.appimage && \
  mv squashfs-root / && \
  ln -s /squashfs-root/AppRun /usr/bin/nvim && \
  nvim --version

# Nvim package deps
RUN apt-get install -y \
  build-essential \
  make \
  gcc

RUN echo "alias vim='nvim'" >> ~/.bashrc
COPY dotfiles/init.lua /root/.config/nvim/init.lua

WORKDIR /root/

