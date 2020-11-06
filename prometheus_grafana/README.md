# Prometheus
Monitoring.   
Prometheus is the polling server that reads data from node-exporters.   
Node-exporters export data on a certain endpoint.   
Grafana displays nice graphs.   

## TODO

- [x] Docker compose with prometheus on backend network, exposing 9090.
- [x] Grafana on backend network, exposing 3000.
- [ ] Persist prometheus storage.
- [ ] Persist grafana storage (credentials).

## Quick start
```
$ make run
```

## Run the Prometheus server docker container
Run the wrapper script.   
```bash
./run_prometheus
```

You can enter the container for debugging with:
```bash
$ docker exec -it prometheus /bin/bash
```

The container is running the following process:
```
/bin/prometheus --config.file=/etc/prometheus/prometheus.yml --storage.tsdb.path=/prometheus --web.console.libraries=/
```

## Docker mount problem
This is a bit problematic... if we want to use our own `prometheus.yml` file, we will want to mount it from outside the container.   
Mounting files is bad because changes only get applied inside the container if we restart it.   
The better solution is mounting a folder, for example `/etc/prometheus`. The problem is that folder in the container already contains some other folders and we probably need them.   

### Workaround
`Dockerfile` with a custom `ENTRYPOINT`.   
`config` is mounted as `/home/config`.   

## Configuration
Edit the config file in `./config`.   
For example, let's add a docker-daemon exporter running on our host machine.
```yaml
  - job_name: 'docker'
    static_configs:
      - targets:
        - "127.0.0.1:9323"
```
After adding configuration, send a reload signal to the server.   
```bash
$ curl -X POST http://localhost:9090/-/reload
```
Note: when done with configuration, close the 9090 port in docker-compose and restart the stack.   


## Setup from scratch
```bash
$ docker-compose up -d
```

### Grafana setup
Enter Grafana dashboard on `localhost:3000`.   
Set admin password.   
Datasources -> Prometheus (Select) -> url: `http://prometheus:9090` -> Save & Test

### Adding Grafana dashboards
Grafana dashboards go hand-in-hand with exporters.   
There's usually instructions on how to install the exporter on the dashboard page.   

Some good dashboards:

* node_exporter: https://grafana.com/grafana/dashboards/11074


## Adding exporters
To add exporters:   

1. Install exporter.
2. Tell prometheus to listen to new exporter.
3. Add Grafana dashboard for new exporter.

Making your own dashboard is time-consuming.   
There are a lot of good ones out there, just make sure 
to get a version compatible with the exporter (exporter metrics change quite a lot 
between versions).   


### node_exporter
Basic exporter that monitors OS metrics.   

Make sure to use the right combination of dashboard-exporter.   

node_exporter releases:
`https://github.com/prometheus/node_exporter/releases`

install as daemon:
`https://www.digitalocean.com/community/tutorials/how-to-install-prometheus-on-ubuntu-16-04`


### docker daemon metrics
Monitors the docker daemon.   
https://ops.tips/gists/how-to-collect-docker-daemon-metrics/


### docker cAdvisor
Monitors docker containers.   
Google their github page - bring up 1 docker container per host and you're set.   
