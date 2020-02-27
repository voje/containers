VERSION=0.1.41
IMG=voje/skopeo:${VERSION}

build:
	docker build . -t ${IMG}
	docker push ${IMG}

run: build
	docker run -it --rm ${IMG} skopeo --help
