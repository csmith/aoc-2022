#!/bin/bash

set -e
cd /code

if [ -d "$1" ]; then
    if [ ! -f "${GOPATH}/bin/$1" ] || [ "${GOPATH}/bin/$1" -nt "$1/main.go" ]; then
        export HOME="/tmp"
        go get -d -v ./... >/dev/null 2>&1
        go install ./... >/dev/null
    fi
    export GOGC=off
    time "${GOPATH}/bin/$1"
fi
