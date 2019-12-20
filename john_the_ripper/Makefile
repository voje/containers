IMG="voje/jon"

build:
	docker build . -t ${IMG}
	docker push ${IMG}

run:
	docker run \
		--rm \
		-it \
		-v $(shell realpath workdir):/workdir \
		${IMG} \
		/bin/bash
