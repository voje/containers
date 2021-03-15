#!/bin/bash

mkdir -p ./testdirs/some_backups
for i in {0..5}; do
    head -c 1M </dev/urandom >./testdirs/some_backups/some_file_$(date --date="-$i second" +%F-%s)
done

echo "Created testdata:"
ls -lah ./testdirs/some_backups
