# Aptly

Aptly container with automatic mirror + exported repository.   

`make-build` to build the container
`make-run` to run the container

The folder `persistent_data` will contain mirrorred data (stored in container's root).   

Follow the scripts from 01 on to prepare mirror, download, create repo, create key, create signed snapshot.   

## Understanding docker usage
Aptly creates its database (including a huge mirror) in `./root/.aptly` - this folder is mapped to `./persistent_data` on docker host (we can set the mount folder in `Makefile`).   
The docker container brings the aptly server on a debian base image and guarantees a repeatable environment (no more conflicts with aptly and gpgv1/gpgv2, etc...).   
To use aptly, we bring up a container with `make-run` (assuming we've alredy built it), the run all the scripts from inside the container.   
There's no need to interact with data in `./persistent_data` from the docker host.   

## Using aptly
After running `make run`, you will be greeted with a shell insode the docker container.   
This container has aptly installed on a base debian image (aptly likes debian).   

Use the scripts in `./scripts` folder to run desired steps:   
* 01 create a mirror (run the scripts inside the container), to create a mirror.   
* 02 (!!!STORAGE!!! populate the mirror (check Makefile for where `/root` is mounted on host - that folder needs to have cca 200GB of space - docker will download packages into `/root/.aptly`, this will translate to the mounted folder on host).  
* 03 create a repo with packages, defined in `./config/packages.txt`.   
* 04 create a snapshot from the repo and check package dependencies dependencies (if deps are missing, fix `packages.txt` and redo steps 03 and 04.   
* 05 sign and publish the snapshot.  TODO  
* 06 test the published snapshot.   TODO
* 07 send the created `./public` dir to other server.   TODO   

