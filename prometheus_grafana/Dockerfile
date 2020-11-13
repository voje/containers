FROM prom/prometheus:v2.15.2

# Moves config file to /etc/config for easier volume mount.   
# Enables lifecycle API for reloading config via REST.   

RUN mkdir /home/config
RUN cp /etc/prometheus/prometheus.yml /home/config/.

ENTRYPOINT [ \
    "/bin/prometheus", \
    "--config.file=/home/config/prometheus.yml", \
    "--storage.tsdb.path=/prometheus", \
    "--web.console.libraries=/", \
    "--web.enable-lifecycle" \
]