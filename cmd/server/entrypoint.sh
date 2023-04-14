#!/bin/bash
if [ $APP_STAGE = "local" ]; \
    then 
        go install github.com/cosmtrek/air@latest && \
        export PATH=$PATH:$(go env GOPATH)/bin && \
        "air" "-c" ".air.toml"
    else 
        go build -o ./bin/${APP_NAME} ./cmd/server/main.go && \
        "./bin/${APP_NAME}"
fi