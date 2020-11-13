#!/bin/bash

IMG="voje/prometheus:v2.15.2"

docker rm -f prometheus

docker pull $IMG

docker run \
    -d \
    --name prometheus \
    -v $PWD/config:/home/config \
    -p 9090:9090 \
    $IMG

docker ps

cat <<EOF
*******************
Prometheus running.   

Web-console:
http://localhost:9090
EOF
