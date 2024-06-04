FROM golang:1.22.2-alpine
RUN apk add --update --no-cache ca-certificates git

WORKDIR /go/src/myapp

RUN go install github.com/cosmtrek/air@v1.51.0

ENTRYPOINT air -c .air.toml