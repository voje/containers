VERSION="0.1.1"
IMG="voje/benchmark:${VERSION}"
CNT="voje_xmr"

build-docker:
	docker build . -t ${IMG}

push-docker: build-docker
	docker push ${IMG}

run-dev:
	-docker rm ${CNT}
	docker run -it --name ${CNT} ${IMG} /bin/sh

# This command runs a SCTP server listening on port 9876
run-docker:
	sysctl -w vm.nr_hugepages=2000
	-docker rm ${CNT}
	docker run \
		--cpus="4.0" \
		-v /dev/hugepages:/dev/hugepages \
		--name ${CNT} \
		${IMG}

		# ./xmrig -c /config/config.json
		# -v $(shell realpath ./config):/config \


		

		# /bin/bash


#		./xmrig \
#			--algo="rx/0" \
#			--coin="monero" \
#			--url="xmr-eu1.nanopool.org:14433" \
#			--user="46E3XMJAMs77JkGmgCvd1KKddst29JmP1NqSRsEuMQpTaHibziUJkRJf2Fufdi3vD9dfDYNaK6A1NQbLqEVyKdCvGYuAd2s" \
#			--pass="x" \

