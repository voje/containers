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
# Copy prometheus config to mounted system folder
```bash
sudo mkdir /etc/prometheus
sudo install 755 ./config/prometheus.yml /etc/prometheus/prometheus.yml
vim /etc/prometheus/prometheus.yml
```
```bash
# Start docker-compose
docker-compose up -d
```

## Prometheus configuration
Edit config: `/etc/prometheus/prometheus.yml`.   
For example, let's add a docker-daemon exporter running on our host machine.
```yaml
  - job_name: 'docker'
    static_configs:
      - targets:
        - "127.0.0.1:9323"
```

```bash
# Enter prometheus container, prepare config and send restart
docker-compose exec prometheus vi /etc/prometheus/prometheus.yml

# Send reload command to prometheus
curl -X POST http://localhost:9090/-/reload
```

## Safety
When done with configuration, close the 9090 port in docker-compose and restart the stack.   
Prometheus doesn't handle authentication to the web console so it's better to close the port.   
We will access prometheus data via Grafana.   


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
