CRI="docker"
VERSION=0.0.1
IMG=voje/pgbench:${VERSION}
CNT=pgbench

build-image:
	${CRI} build . -t ${IMG}

push-image: build-image
	${CRI} push ${IMG}

