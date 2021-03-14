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
ssh-keygen -t ed25519 -C "your_email@example.com"
```
