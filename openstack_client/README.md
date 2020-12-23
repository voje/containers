# Openstack Client
It can be a pain to install a working openstack-client on a system transitioning from python2 to python3.   

Here's a working version locked in a docker container:
```bash
# pip3 list | grep openstack
openstackclient                4.0.0
openstacksdk                   0.52.0
python-openstackclient         5.4.0
```

## Quick usage
Prepare a folder with `openstackrc.sh`.  
This is the file with your user login information.  
You can get it from the openstack dashboard.   

Run the container and mount the folder containing `openstackrc.sh`.   
```bash
docker run -it -v $PWD:/openstackrc voje/openstack-client:0.1.0 /bin/sh
```
Inside the container
```bash
source openstackrc.sh

# Check if the connection works
openstack server list

# If you've configured your Openstack with a self-signed certificate, this might help:
openstack server list --insecure

alias o='openstack --insecure'
o server list
```


