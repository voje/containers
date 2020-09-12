#!/bin/bash

#
# Serve the snapshot
# add key
# try apt update
#

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

echo "[*] Stopping server"
# Stop serving
pkill aptly

# Restore sources.list
mv /etc/apt/sources.list.orig /etc/apt/sources.list

