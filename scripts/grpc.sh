#!/usr/bin/env bash

PROTO_FILE="$(pwd)/internal/proto/v1/todo.proto"

# check protobuf file.
if [[ ! -f "$PROTO_FILE" ]]; then
  echo >&2 "error: todo.proto not found"
  exit 1
fi

# generate gRPC code.
echo "Generate gRPC code..."
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
    internal/proto/v1/*.proto