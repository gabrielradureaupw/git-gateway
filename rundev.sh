#!/bin/bash

export $(cat example.env | xargs)

export GO111MODULE=on

if [ -n "$USE_SSL" ]; then
    openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem
fi
visit chrome://flags/#allow-insecure-localhost

go mod init
go mod tidy

go run main.go