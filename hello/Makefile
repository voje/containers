CRI="docker"
VERSION=0.0.2
IMG=voje/hello:${VERSION}
CNT=hello

build-image:
	${CRI} build . -t ${IMG}

push-image: build-image
	${CRI} push ${IMG}

run-docker:
	echo "curl localhost:8080/mypath"
	${CRI} run \
		-d \
		--rm \
		--name ${CNT} \
		-e HELLO='Hola !' \
		-e HELLO_PATH='/mypath' \
		-p 8080:8080 \
	${IMG}
	${CRI} logs -f ${CNT}

