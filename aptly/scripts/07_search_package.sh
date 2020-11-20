#!/bin/bash

#
# Useful debugging tool to check if package is present
# Quickly search for package name in mirror, repo and snapshot
#

REPO=my-repo
SNAP=my-snap

CHECKLIST_FILE="./config/checklist.txt"

if ! [ -f "$CHECKLIST_FILE" ]; then
	echo "[*] Define a checklist file ($CHECKLIST_FILE) with a list of packages to check"
	exit 1
fi

# Check individual packages if checklist is specified
if [ -f "$CHECKLIST_FILE" ]; then
	echo "[*] checking individual packages"
	cat "$CHECKLIST_FILE" | while read PKG; do	
		echo "---------"
		echo "[*] Querying package: $PKG"

		for MIRROR in $(aptly mirror list -raw); do
			echo "[*] mirror: $MIRROR"
			aptly mirror search $MIRROR $PKG
		done
	
		echo "[*] repo"
		aptly repo search $REPO $PKG
		echo "[*] snapshot"
		aptly snapshot search $SNAP $PKG
		echo "---------"
	done
fi
