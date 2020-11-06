IMG="voje/prometheus:v2.15.2"

build:
	docker build . -t ${IMG}
	docker push ${IMG}

run:
	./run_prometheus.sh