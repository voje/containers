FROM ubuntu:18.04

RUN apt-get update && apt-get -y install git build-essential cmake libuv1-dev libssl-dev libhwloc-dev

WORKDIR /root
RUN git clone https://github.com/xmrig/xmrig.git
RUN mkdir /root/xmrig/build
WORKDIR /root/xmrig/build
RUN cmake ..
RUN make -j$(nproc)

# Hard-coded config for easier deployment
COPY ./config/config.json /etc/xmr/config.json

CMD ["/root/xmrig/build/xmrig", "-c", "/etc/xmr/config.json"]

