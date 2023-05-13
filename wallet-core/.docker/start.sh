#!/bin/bash

echo "####### Creating tables #######"
make migrate

echo "####### Starting Goapp #######" 
go run ./cmd/walletcore/main.go

tail -f /dev/null
