# Node.js version
FROM node:18.5.0-alpine3.16
# apk update and package install
RUN apk add --update &&  apk add git tzdata
# create working directory
RUN mkdir -p /usr/src/web
# set working directory
WORKDIR /usr/src/web

