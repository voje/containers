from golang:alpine3.12 as builder

workdir src

# go.mod and go.sum change less frequently, we can cache the reults in a docker layer
copy go.mod go.sum ./
run go mod download

# Dynamic
copy hello.go ./
run CGO_ENABLED=0 go build -o /bin/main hello.go

from scratch
copy --from=builder /bin/main /main

entrypoint ["/main"]

