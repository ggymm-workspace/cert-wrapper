#!/bin/bash

export GOOS=windows
export GOARCH=amd64
go build -ldflags="-s -w" -o cert-wrapper.exe

export GOOS=linux
export GOARCH=arm
export GOARM=7
go build -ldflags="-s -w" -o cert-wrapper

zip -q cert-wrapper.zip cert-wrapper

