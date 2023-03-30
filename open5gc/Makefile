CRI="docker"
VERSION=5.2.8
IMG=voje/open5gc:${VERSION}
CNT=open5gc

build-image:
	${CRI} build . -t ${IMG}

push-image: build-image
	${CRI} push ${IMG}

dev-run:
	-docker rm -f ${CNT}
	docker run -d --name=${CNT} ${IMG} sleep inf
	docker exec -it ${CNT} /bin/bash

run:
	-docker rm -f ${CNT}
	docker run -it --name=${CNT} ${IMG} /project/open5gs/install/bin/open5gs-amfd

