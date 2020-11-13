#!/bin/bash

set -e

# This script takes a while to download all of the mirror data.

aptly mirror update bionic-main-1
aptly mirror update bionic-updates-1
aptly mirror update bionic-docker-main-1
aptly mirror update bionic-ansible-28
