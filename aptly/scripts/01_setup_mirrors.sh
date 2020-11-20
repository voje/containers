#!/bin/bash

#wget -O -  https://download.docker.com/linux/ubuntu/gpg |  gpg --no-default-keyring --keyring /usr/share/keyrings/ubuntu-archive-keyring.gpg --import
#wget -O -  http://si.archive.ubuntu.com/ubuntu/project/ubuntu-archive-keyring.gpg.sig  |  gpg --no-default-keyring --keyring /usr/share/keyrings/ubuntu-archive-keyring.gpg --import 
#wget -O -  http://si.archive.ubuntu.com/ubuntu/project/ubuntu-archive-keyring.gpg  |  gpg --no-default-keyring --keyring /usr/share/keyrings/ubuntu-archive-keyring.gpg --import
#gpg  --keyring /usr/share/keyrings/ubuntu-archive-keyring.gpg -k
#gpg --no-default-keyring --keyring /usr/share/keyrings/ubuntu-archive-keyring.gpg --export | gpg --no-default-keyring --keyring trustedkeys.gpg --import

# Ubuntu
wget -O -  http://si.archive.ubuntu.com/ubuntu/project/ubuntu-archive-keyring.gpg.sig  | gpg --no-default-keyring --keyring trustedkeys.gpg --import
wget -O -  http://si.archive.ubuntu.com/ubuntu/project/ubuntu-archive-keyring.gpg  | gpg --no-default-keyring --keyring trustedkeys.gpg --import
# create DEB mirrors
aptly mirror create -architectures=amd64 -filter-with-deps bionic-main-1 http://si.archive.ubuntu.com/ bionic main universe multiverse
aptly mirror create -architectures=amd64 -filter-with-deps bionic-updates-1 http://si.archive.ubuntu.com/ bionic-updates main universe multiverse

# Docker
wget -O -  https://download.docker.com/linux/ubuntu/gpg | gpg --no-default-keyring --keyring trustedkeys.gpg --import
aptly mirror create -architectures=amd64 -filter-with-deps bionic-docker-main-1 https://download.docker.com/linux/ubuntu/ bionic stable

# Ansible 2.8
gpg --no-default-keyring --keyring trustedkeys.gpg --keyserver keys.gnupg.net --recv-keys 93C4A3FD7BB9C367
aptly mirror create -architectures=amd64 -filter-with-deps bionic-ansible-28 http://ppa.launchpad.net/ansible/ansible-2.8/ubuntu bionic main

gpg --keyring trustedkeys.gpg  -k

