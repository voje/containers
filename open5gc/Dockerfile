FROM ubuntu:22.04 as builder

run apt-get update && apt-get upgrade -y
run apt-get install -y \
    git

run apt-get install -y python3-pip python3-setuptools python3-wheel ninja-build build-essential flex bison git cmake libsctp-dev libgnutls28-dev libgcrypt-dev libssl-dev libidn11-dev libmongoc-dev libbson-dev libyaml-dev libnghttp2-dev libmicrohttpd-dev libcurl4-gnutls-dev libnghttp2-dev libtins-dev libtalloc-dev meson

workdir /project

run git clone https://github.com/open5gs/open5gs.git --depth 1 -b v2.5.8
workdir /project/open5gs
run meson build --prefix=`pwd`/install
run ninja -C build

workdir /project/open5gs/build
run ninja install

# workdir /project/open5gs/build
# run meson test -v || true

