IMG=voje/aptly:0.0.2
CNT=voje_aptly
TEST_CNT=voje_aptly_test

build-docker:
	cd ./dockerfiles; docker build . -t ${IMG} -f Dockerfile

run-docker:
	-docker rm -f ${CNT}
	-mkdir persistent_data gpg_keys
	docker run \
		--rm \
		-it \
		--name ${CNT} \
		-v $(shell pwd)/scripts:/scripts \
		-v $(shell pwd)/persistent_data:/root \
		-v $(shell pwd)/gpg_keys:/gpg_keys \
		-w /scripts \
		${IMG} /bin/bash

# Install packages in a new ubuntu container using the created repo
run-tests:
	-docker rm -f ${TEST_CNT}
	docker run \
		--rm \
		-it \
		--name ${TEST_CNT} \
		-v $(shell pwd)/scripts:/scripts \
		-v $(shell pwd)/persistent_data:/root \
		-w /scripts/test_scripts \
		${IMG} /bin/bash
