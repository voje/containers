VER=0.0.1
IMG=voje/coreos-installer:${VER}
NAME=coreos_installer

.PHONY: build push run

build:
	docker build . -t ${IMG}

push: build
	docker push ${IMG}

run:
	-docker rm -f ${NAME}
	docker run \
		--hostname coreos-installer \
		-it \
		--name=${NAME} \
		-v $(shell pwd)/workdir:/workdir \
		${IMG} \
		/bin/bash
