IMG=voje/aptly:0.0.1
CNT=voje_aptly

build-docker:
	docker build . -t ${IMG} -f Dockerfile

run-docker:
	-docker rm -f ${CNT}
	docker run \
		--rm \
		-it \
		--name ${CNT} \
		-v $(shell pwd)/scripts:/scripts \
		-w /scripts \
		${IMG} /bin/bash
