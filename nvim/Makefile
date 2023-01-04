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

