#!/usr/bin/env bash

PROTO_NAME=$1
PROTO_DIR="$(pwd)/libs/proto/v1"

# check protobuf directory.
if [[ ! -d "$PROTO_DIR" ]]; then
  echo >&2 "error: $PROTO_DIR not found"
  exit 1
fi

generate() {
  # generate gRPC code.
  echo "Generate gRPC code..."
  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
    libs/proto/v1/*.proto
}

generate_proto_name() {
  if [[ ! -f "$PROTO_DIR/$PROTO_NAME" ]]; then
    echo >&2 "error: $PROTO_DIR/$PROTO_NAME not found"
    exit 1
  fi

  # generate gRPC code.
  echo "Generate gRPC code..."
  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
    libs/proto/v1/$PROTO_NAME
}

if [[ $PROTO_NAME == "" ]]; then
  generate
else
  generate_proto_name
fi
