IMG="voje/openstack-client:0.1.0"
CNT="voje_ops_cli"

build-docker:
	docker build . -t ${IMG}

push-docker:
	docker push ${IMG}

run-dev:
	-docker rm ${CNT}
	docker run \
	    -it \
	    --name ${CNT} \
	    -v $(shell pwd)/mount:/root/mount \
	    ${IMG} \
	    /bin/sh

