#!/bin/bash

REPO="my-repo"
SNAP="my-snap"
PUBL="bionic"


# Drop if exists
if aptly snapshot show $SNAP; then
	echo "[*] Dropping snapshot $SNAP"
	aptly snapshot drop $SNAP

	# Drop published snapshot
	aptly publish drop $PUBL
fi

sleep 3

# TODO
# For some reason, snapshot takes the wrong ansible version from repo
echo "[*] Removing undesired version of ansible"
aptly repo remove my-repo ansible_2.5.1+dfsg-1_all
aptly repo remove my-repo ansible_2.5.1+dfsg-1ubuntu0.1_all

# Create snapshot
if ! aptly snapshot show $SNAP; then
	echo "[*] Creating snapshot $SNAP from repo $REPO"
	aptly snapshot create $SNAP from repo $REPO
fi

sleep 2

# Verify dependencies
aptly snapshot verify $SNAP
echo "[*] If there are MISSING DEPENDENCIES, add them to your packages list and redo steps 03 and 04"

