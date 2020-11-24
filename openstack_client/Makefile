IMG="voje/openstack-client"
CNT="voje_ops_cli"

build-docker:
	docker build . -t ${IMG}

push-docker:
	docker push ${IMG}

run-dev:
	-docker rm ${CNT}
	docker run -it --name ${CNT} ${IMG} /bin/sh

