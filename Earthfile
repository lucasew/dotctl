FROM golang:latest

build:
    WORKDIR /build
    COPY . .
    RUN go build -o dotctl
    SAVE ARTIFACT ./dotctl AS LOCAL ./dotctl
