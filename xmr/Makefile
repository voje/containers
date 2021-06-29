IMG="voje/xmr"
CNT="voje_xmr"

build-docker:
	docker build . -t ${IMG}

push-docker:
	docker push ${IMG}

run-dev:
	-docker rm ${CNT}
	docker run -it --name ${CNT} ${IMG} /bin/sh

# This command runs a SCTP server listening on port 9876
run-docker:
	sysctl -w vm.nr_hugepages=2000
	-docker rm ${CNT}
	docker run -it \
		--cpus="4.0" \
		-v $(shell realpath ./config):/config \
		-v /dev/hugepages:/dev/hugepages \
		--name ${CNT} ${IMG} \
		./xmrig -c /config/config.json
		# /bin/bash

