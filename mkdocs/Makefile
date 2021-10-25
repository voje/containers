IMG=voje/mkdocs:0.1.0
NAME=mkdocs

# Where your .md files live (abs path)
# INPUT=$(shell pwd)/testdir
INPUT=/home/voje/git/foam/docs

build:
	docker build . -t ${IMG}

run:
	-docker rm -f ${NAME}
	docker run \
		-d \
		--name=${NAME} \
		-p 8000:8000 \
		-v ${INPUT}:/var/mkdocs/raw:ro \
		${IMG}
	@echo "docker exec -it ${NAME} /bin/bash"
