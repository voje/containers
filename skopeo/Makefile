IMG=voje/skopeo

build:
	docker build . -t ${IMG}

run: build
	docker run -it --rm ${IMG} /bin/sh
