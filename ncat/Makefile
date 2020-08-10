IMG="voje/ncat"
CNT="voje_ncat"

build-docker:
	docker build . -t ${IMG}

push-docker:
	docker push ${IMG}

run-dev:
	-docker rm ${CNT}
	docker run -it --name ${CNT} ${IMG} /bin/sh

# This command runs a SCTP server listening on port 9876
run-docker:
	-docker rm ${CNT}
	docker run -it --name ${CNT} ${IMG} ncat --sctp -l 9876
