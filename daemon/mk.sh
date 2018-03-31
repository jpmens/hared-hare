#!/bin/sh

export GOPATH=`pwd`/gopath

if [ ! -d $GOPATH ]; then
	mkdir -p $GOPATH
	go get github.com/eclipse/paho.mqtt.golang
	go get gopkg.in/gcfg.v1
fi

GOOS=darwin go build -o hared-darwin hared.go

GOOS=linux GOARCH=amd64 go build -o hared-linux-amd64 hared.go
