IMG=voje/mkdocs:0.1.0
NAME=mkdocs

build:
	docker build . -t ${IMG}

run:
	-docker rm -f ${NAME}
	docker run \
		-d \
		--name=${NAME} \
		-p 8000:8000 \
		-v $(shell pwd)/testdir:/var/mkdocs/raw \
		${IMG} \
		sleep inf
	@echo "docker exec -it ${NAME} /bin/bash"
