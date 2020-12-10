# ping-listener

Useful for debugging in kubernetes: the container listens on `:8080/ping`.   
Sends back `pong` and logs information about the caller.   

```bash
make build-docker
make run-docker
curl localhost:8080/ping
make push-docker
```

