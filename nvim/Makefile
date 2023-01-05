CRI="docker"
VERSION=0.0.1
IMG=voje/nvim:${VERSION}
CNT=nvim

build-image:
	${CRI} build . -t ${IMG}

push-image: build-image
	${CRI} push ${IMG}

run: build-image
	docker run --rm -it --name ${CNT} ${IMG} /bin/bash

# After 'run', run nvim twice to install the packages
# Then run commit-push
commit-push:
	${CRI} commit ${CNT} ${IMG} 
	${CRI} push ${IMG}

