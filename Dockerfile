FROM golang:1.16-buster

ENV TZ Asia/Tokyo

WORKDIR /go/src/app
COPY ./ /go/src/app

RUN go build main.go
