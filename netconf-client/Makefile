VERSION=0.1.0
IMG=voje/netconf-client:${VERSION}

build:
	docker build . -t ${IMG}

push: build
	docker push ${IMG}
