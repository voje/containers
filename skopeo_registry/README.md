# Skopeo-registry
This container adds the 'skopeo' image transferring tool to the 
official 'registry:2' image.   
In `go`, we also build some tools for a cleaner interface:   
* list images in registry
* pull images to registry (input is a list of images)
* push all images from registry to desired location

In order to preserve (and transport) docker images, we need to mount 
`/var/lib/registry`. On destination, simply run the registry with this mounted 
folder and you should be able to list and push the images.   

## Developing tools
When developing docker tools, make sure you have a skopeo bin on your local environment.   
Best option to get the right version (version with 'sync') is to run this docker container and copy 
the `/usr/local/bin/skopeo` to your system.   
Run the conainer with `--net host` so you can access the registry on `:5000`.   

## Run these inside the docker container:
```bash
# Official docker registry entrypoint:
$ registry serve /etc/docker/registry/config.yml &

$ skopeo sync --src-tls-verify=false --src-creds=deploy:deploypassword --src=yaml --dest-tls-verify=false --dest=docker skopeo_repo_server.yml 127.0.0.1:5000

# This works:
$ skopeo sync --src-creds=voje:hiddenPASSWORD --src=docker --dest=docker --dest-tls-verify=false --dest-no-creds=true vmgitent.
iskratel.si:4567/ai6212ax/bacula-infra-backup/it_bacula-db:1.1.0 127.0.0.1:5000

```

## Working examples
To run, we copied skopeo to our host machine and ran a `--net host` docker container.   

Push all images from registry to destination.
```bash
go run main.go push --dst-addr=k-vm-repo-server.docker.iskratel.mak:4567/deploy/portainer_templates --creds deploy:deploypassword
```