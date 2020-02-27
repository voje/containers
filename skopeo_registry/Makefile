VERSION=0.0.1
IMG=voje/skopeo-registry:${VERSION}

build:
	docker build . -t ${IMG}
	docker push ${IMG}

run: build
	docker run -it --rm ${IMG} sh -c 'skopeo -v && registry -v'
