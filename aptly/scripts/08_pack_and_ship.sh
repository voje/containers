#!/bin/bash

PUBLIC="/root/.aptly/public"
TAG=$(date +%F-%H%M)
PAYLOAD_DIR="/root/payloads"
PAYLOAD_FILE="ubuntu-$TAG.tar"

STAGING_SERVER="ubuntu@k-vm-repo-server.docker.iskratel.mak"
STORAGE_SERVER="test@centos70mano.iskratel.si"

#
# Add a tag file inside the folder
#
echo $TAG > $PUBLIC/verison.txt

#
# Prepare payload dir and create a tar file in there
#
mkdir -p $PAYLOAD_DIR
tar cvf $PAYLOAD_DIR/$PAYLOAD_FILE -C $PUBLIC/.. ./public
 
#
# Send snapshot in /root/.aptly/public to staging server - this will overwrite the data on staging!
# Prereq: requires a passwordless ssh to the staging server.
#
echo
echo "[*] WARNING the following steps need to be executed manually"
echo

echo "[*] Send to staging server:"
echo "[*] $ rsync -rv --delete $PUBLIC $STAGING_SERVER:/home/ubuntu/repo-server/static/ubuntu/."
echo
echo "[*] Send to storage server:"
echo "[*] $ rsync -rv $PAYLOAD_DIR/$PAYLOAD_FILE $STORAGE_SERVER:/opt/IT_MANO/IT_OPENSTACK_KOLLA/repo_data/ubuntu_repo/."
