#!/bin/bash

REPO="my-repo"
SNAP="my-snap"


# Drop if exists
if aptly snapshot show $SNAP; then
	echo "[*] Dropping snapshot $SNAP"
	aptly snapshot drop $SNAP
fi


# Create snapshot
if ! aptly snapshot show $SNAP; then
	echo "[*] Creating snapshot $SNAP from repo $REPO"
	aptly snapshot create $SNAP from repo $REPO
fi


# Verify dependencies
aptly snapshot verify $SNAP
echo "[*] If there are MISSING DEPENDENCIES, add them to your packages list and redo steps 03 and 04"


# TODO publish: add gpg key!
