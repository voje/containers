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