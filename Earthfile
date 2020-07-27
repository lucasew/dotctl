FROM golang:latest
WORKDIR /build
COPY . .

build:
    RUN go build -o dotctl
    SAVE ARTIFACT ./dotctl AS LOCAL ./dotctl

test:
    RUN go test -v