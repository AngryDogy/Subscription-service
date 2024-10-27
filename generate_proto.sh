#!/bin/sh

protoc proto/api/v1/*.proto --go_out=protogen --go_opt=paths=source_relative
protoc proto/api/v1/*.proto  --go-grpc_out=protogen --go-grpc_opt=paths=source_relative