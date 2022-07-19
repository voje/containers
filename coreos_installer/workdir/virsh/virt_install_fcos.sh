#!/bin/bash

NAME=fcos

./virt_cleanup.sh $NAME

virt-install \
    --connect qemu:///session \
    --name fcos \
    --graphics none \
    --autoconsole none \
    --memory 2048 --vcpu 2 --disk size=5 \
    --os-variant fedora36 \
    --cdrom ../fedora-coreos-36.20220703.3.1-live.x86_64.iso
