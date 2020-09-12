#!/bin/bash

#
# Prereq: a gpg key for signing the package (bring your own!)
# Import a gpg key before this step
# $ gpg --import secret.key
# $ gpg --list-keys
#

GPG_KEY_ID=
SNAP="my-snap"

# Publish a (signed) snapshot
# /root/.aptly/public will be created
aptly publish snapshot --gpg-key="$GPG_KEY_ID" $SNAP

# Add public key to the snapshot
gpg --output /root/.aptly/public/ubuntu.gpg.key --armor --export $GPG_KEY_ID
