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
WORKDIR /git/JohnTheRipper/run

# Crunch password generator.
RUN apt-get install -y crunch
