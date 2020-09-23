#!/bin/bash

#
# Serve the snapshot
# add key
# try apt update
#

CHECKLIST_FILE="./config/checklist.txt"

# Serve snapshot
echo "[*] Serving snapshot (/root/.aptly/public)"
aptly serve &
sleep 5

# Add gpg key
echo "[*] Fetching gpg key"
curl -s localhost:8080/ubuntu.gpg.key | apt-key add -

# Change sources.list
mv /etc/apt/sources.list /etc/apt/sources.list.orig
echo "deb http://localhost:8080 bionic main" > /etc/apt/sources.list

# Moment of truth...
echo "[*] apt update"
if apt update; then
	echo "[*] SUCCESS ! apt update works, key seems fine"
else
	echo "[*] apt update FAILED!!!"
fi

# Check individual packages if checklist is specified
if [ -f "$CHECKLIST_FILE" ]; then
	echo "[*] checking individual packages"
	cat "$CHECKLIST_FILE" | while read package; do	
		apt-cache policy $package
	done
fi

echo "[*] Stopping server"
# Stop serving
pkill aptly

# Restore sources.list
mv /etc/apt/sources.list.orig /etc/apt/sources.list

