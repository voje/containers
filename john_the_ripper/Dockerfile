FROM ubuntu

# Install dependencies.
RUN apt-get update && apt-get install -y \
curl \
vim \
git \
build-essential \
libssl-dev

RUN mkdir /git
WORKDIR /git

# Clone the srouces code.
RUN git clone https://github.com/magnumripper/JohnTheRipper.git

# Compile.
WORKDIR /git/JohnTheRipper/src
RUN ./configure && make
RUN ln -s -T /git/JohnTheRipper/run/john /usr/local/bin/john

RUN echo "john --help"
RUN mkdir /workdir
WORKDIR /workdir

