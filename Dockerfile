FROM golang:1.13.7 as build-env


RUN mkdir /app

WORKDIR /app

COPY go.mod /app

#Download dependencies from go.mod
RUN GO111MODULE=on go mod download

COPY . /app

#Build app
RUN GO111MODULE=on CGO_ENABLED=0 go build -v -o /bin/imgsrv

#Stage-2 --> Copy app binary from previous stage
FROM alpine:3.8

#Add ca certificates
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

COPY --from=build-env /bin/imgsrv /imgsrv
RUN apk add --no-cache tini bash

ENTRYPOINT ["/sbin/tini", "--"]



