FROM ubuntu:22.04

WORKDIR /root

RUN apt-get update && apt-get install -y \
  curl

# Install Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs > rustup.sh && sh rustup.sh -y && rm rustup.sh
ENV PATH=${PATH}:/root/.cargo/bin

RUN apt-get install -y gcc
RUN cargo install \
  tree-sitter-cli \
  ripgrep

WORKDIR /root/git

RUN apt-get install -y git

##
## Compile and install neovim
##
RUN apt-get install -y ninja-build gettext libtool libtool-bin autoconf automake cmake g++ pkg-config unzip curl doxygen
RUN git clone https://github.com/neovim/neovim.git
RUN cd /root/git/neovim && \
  make CMAKE_BUILD_TYPE=RelWithDebInfo && \
  make install
RUN nvim --version
RUN rm -rf /root/git/neovim
RUN echo "alias vim='nvim' >> ~/.bash_aliases"

