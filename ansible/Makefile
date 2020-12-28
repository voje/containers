IMG=voje/ansible:2.9

build:
	docker build . -t ${IMG}
	docker push ${IMG}

run: build
	docker run --rm -it ${IMG} /bin/bash