#!/bin/bash

echo "####### Creating tables #######"
make migrate

echo "####### Starting Goclient #######" 
go run ./cmd/balance/main.go

tail -f /dev/null
