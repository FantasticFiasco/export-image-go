#!/bin/bash

# Any subsequent(*) commands which fail will cause the shell script to exit immediately
set -e

echo "`go version`"

go vet ./...

go get golang.org/x/lint/golint
golint ./...

env GOOS=linux GOARCH=386 go build -o ./artifacts/exportimage ./cmd/exportimage
env GOOS=windows GOARCH=386 go build -o ./artifacts/exportimage.exe ./cmd/exportimage