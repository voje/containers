VERSION=0.0.1
IMG=voje/ping-listener:${VERSION}
CNT=ping_listener

# We can build directly in minikube (no need for registry or tagging)
# eval $(minikube docker-env)
build-docker:
	docker build . -t ${IMG}

push-docker: build-docker
	docker push ${IMG}

run-docker:
	docker run --rm --name ${CNT} -p 8080:8080 ${IMG}
	docker logs -f ${CNT}

