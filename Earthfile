FROM golang:latest
WORKDIR /build
COPY . .

build-windows:
    ARG GOOS=windows
    RUN go build -o dotctl-windows.exe
    SAVE ARTIFACT dotctl-windows.exe AS LOCAL ./dist/dotctl-windows.exe

build-linux:
    RUN go build -o dotctl-linux
    SAVE ARTIFACT dotctl-linux AS LOCAL ./dist/dotctl-linux

build:
    RUN mkdir dist
    BUILD +build-windows
    BUILD +build-linux

test:
    RUN go test -v

ci:
    BUILD +build
    BUILD +test