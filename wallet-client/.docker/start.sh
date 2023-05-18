#!/bin/bash

echo "####### Creating tables #######"
make migrate

echo "####### Starting Goclient #######" 
# go run ./cmd/walletclient/main.go

tail -f /dev/null
