FROM golang:1.10-alpine3.9

RUN mkdir /build

ADD ./src/go.mod ./src/go.sum /build/ 

WORKDIR /build

RUN go build

