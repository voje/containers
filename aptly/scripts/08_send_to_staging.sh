#!/bin/bash

STAGING="k-vm-repo-server.docker.iskratel.mak"

#
# Send snapshot in /root/.aptly/public to staging server - this will overwrite the data on staging!
# Prereq: requires a passwordless ssh to the staging server.
#

echo "[*] May require a passwordless ssh to the staging server:"
echo "[*] $ ssh-copy-id ubuntu@$STAGING"
rsync -rv --delete /root/.aptly/public ubuntu@$STAGING:/home/ubuntu/repo-server/static/ubuntu/.
