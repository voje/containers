# Backup
Simple backup container using Linux-native tools:
* rsync
* cronjob
* logrotate
* openssl-client

## Quick-start
We have three mounted folders: `ssh_keys` contains this service's ssh key for accessing remote locations (we need to generate that key).   
The folder `data` contains backed-up data. We're in control of the structure.   

```bash
cd ./ssh_keys
ssh-keygen -t ed25519 -C "backup-user@backup-server" -f ./id_ed25519

# Copy this part to remote machines' athorized_hosts` file
cat ./id_ed25519.pub
```

## Testing
Quickly generate some structures for testing:
```bash
mkdir -p ./testdirs/some_backups
cd ./testdirs/some_backups
for i in {0..5}; do head -c 1M </dev/urandom >some_file_$(date --date="-$i day" +%F-%s); done
```

## Use #!/bin/sh
When writing rsync scripts in `scripts` folder, make sure to use the `#!/bin/sh` instead of `#!/bin/bash`.   

## run-parts
We're using the `run-parts` command to run all of the scripts in a folder.   
This command has trouble processing `.sh` extensions! Avoid `.sh` extensions in `scripts`.   
