IMG=voje/aptly:0.0.1
CNT=voje_aptly

build-docker:
	cd ./dockerfiles; docker build . -t ${IMG} -f Dockerfile

run-docker:
	-docker rm -f ${CNT}
	-mkdir persistens_data gpg_keys
	docker run \
		--rm \
		-it \
		--name ${CNT} \
		-v $(shell pwd)/scripts:/scripts \
		-v $(shell pwd)/persistent_data:/root \
		-v $(shell pwd)/gpg_keys:/gpg_keys \
		-w /scripts \
		${IMG} /bin/bash
