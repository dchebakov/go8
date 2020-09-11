#!/usr/bin/env bash

cd ../
#protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=$GOPATH/src internal/proto-files/domain/repository.proto
#protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=plugins=grpc:. internal/proto-files/service/repository-service.proto
protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=plugins=grpc:. internal/proto-files/service/user.proto
