VERSION=0.0.1
IMG="voje/backup:${VERSION}"
CNT="voje_backup"

build-docker:
	docker build . -t ${IMG}

push-docker:
	docker push ${IMG}

run-dev:
	-docker rm ${CNT}
	-mkdir mounted
	docker run \
		-it \
		--name ${CNT} \
		-v ${PWD}/mounted:/mounted \
		${IMG} \
		/bin/sh
