#!/bin/bash

docker run \
    -it \
    -v $PWD/clouds.yaml:/etc/clouds.yaml \
    -p 9180:9180 \
    quay.io/niedbalski/openstack-exporter-linux-amd64:master \
    testcore \
    --os-client-config="/etc/clouds.yaml" \
    --web.listen-address=":9180"

