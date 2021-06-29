IMG="voje/pyrit"
CNT="voje_pyrit"

build-docker:
	docker build . -t ${IMG}

push-docker:
	docker push ${IMG}

run-dev:
	-docker rm ${CNT}
	-mkdir docker_mount
	docker run -it -v ${PWD}/docker_mount:/docker_mount --name ${CNT} ${IMG} /bin/bash

