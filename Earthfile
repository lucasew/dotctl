FROM golang:latest
WORKDIR /build
COPY . .

build-windows:
    ARG GOOS=windows
    RUN go build -o dist/dotctl-windows.exe

build-linux:
    RUN go build -o dist/dotctl-linux

build:
    RUN mkdir dist
    BUILD +build-windows
    BUILD +build-linux
    SAVE ARTIFACT ./dist AS LOCAL ./dist

test:
    RUN go test -v