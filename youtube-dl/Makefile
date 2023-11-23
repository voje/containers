SHORT_COMMIT=aaed488
IMG=voje/youtube-dl:${SHORT_COMMIT}

CRI=podman

build:
	${CRI} build . -t ${IMG} --build-arg short_commit=${SHORT_COMMIT}
	${CRI} push ${IMG}

run: build
	${CRI} run -it --rm ${IMG} youtube-dl --help

install-shim:
	sudo install youtube-dl.sh /usr/local/bin/.
