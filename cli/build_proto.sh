#!/bin/bash

echo "starting build proto"

PATH=$PATH:$GOPATH/bin/ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative delivery/proto/3party/*.proto

echo "successfully generate"