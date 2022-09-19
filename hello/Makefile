CRI="podman"
VERSION=0.0.1
IMG=voje/hello:${VERSION}
CNT=hello

build-image:
	${CRI} build . -t ${IMG}

push-image: build-image
	${CRI} push ${IMG}

run-docker:
	${CRI} run --rm --name ${CNT} -p 8080:8080 ${IMG}
	${CRI} logs -f ${CNT}

