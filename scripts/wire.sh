#!/usr/bin/env bash

WIRE_FILE="$(pwd)/cmd/wire.go"

# check wire injector file.
if [[ ! -f "$WIRE_FILE" ]]; then
  echo >&2 "error: wire.go not found"
  exit 1
fi

# run wire dependency injection
echo "Run dependency injection..."
wire ./...
