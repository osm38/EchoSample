# go version
FROM golang:1.18.3-alpine
# apk update and package install
RUN apk add --update &&  apk add git alpine-sdk build-base
# create log directory
RUN mkdir -p /var/log/es
# set working directory
WORKDIR /go/src

