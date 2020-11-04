#!/bin/bash

#
# Create repo out of mirrorred repositories
#

# Debugging
set -x

# Prompt user before continuing
read -p "[*] This process might take long (depending on packages.txt). Continue? [y/N]" yn
if [[ "$yn" != y ]]; then
	echo "[*] Aborting"
	exit
fi


# Variables

REPO="my-repo"
PKG_LIST="./config/packages.txt"


# Check packages file
if ! [ -f $PKG_LIST ]; then
	echo "[*] Expecting a list of packages in file: $PKG_LIST"
	exit 1
else
	echo "[*] Using package list: $PKG_LIST"
fi


# Create a new repo - first drop the old one
if aptly repo show $REPO; then
	echo "[*] Dropping repo: $REPO"
	aptly repo drop $REPO
fi

# Create a new one
if ! aptly repo show $REPO; then
	echo "[*] Creating repo: $REPO"
	aptly repo create -component="main" -distribution="bionic" -comment="Bionic main repo" "$REPO"
fi


# Import packages from mirrors, created in step 01

# Import all packages from docker mirror
cat <<EOF
aptly repo import \
	--architectures="amd64" \
	--with-deps \
	bionic-docker-main-1 \
	"$REPO" \
	'$PackageType (%*)'
EOF

# Import all packages from ansible mirror
cat <<EOF
aptly repo import \
	--architectures="amd64" \
	--with-deps \
	bionic-ansible-28 \
	"$REPO" \
	'$PackageType (%*)'
EOF

# Import subset of packages from ubuntu main and ubuntu update and ubuntu security mirrors
cat $PKG_LIST | while read package; do
	# Skip empty lines
	if [[ "${#package}" == "0" ]]; then
		echo "[*] Skipping empty line"
		continue
	fi

	# Skip comments
	if echo $package | grep -q -e "^#"; then
		echo "[*] Skipping comment: $package"
		continue
	fi

	echo "[*] Importing package: $package"

	# Import from main
	aptly repo import \
		--architectures="amd64" \
		--with-deps \
		--dep-follow-recommends \
		bionic-main-1 \
		"$REPO" \
		"$package"

	# Import from updates
	aptly repo import \
		--architectures="amd64" \
		--with-deps \
		--dep-follow-recommends \
		bionic-updates-1 \
		"$REPO" \
		"$package"

	# Import from security
	aptly repo import \
		--architectures="amd64" \
		--with-deps \
		--dep-follow-recommends \
		bionic-security \
		"$REPO" \
		"$package"
done

echo "[*] Repo $REPO is ready! Time to create a snapshot."


