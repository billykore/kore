#!/usr/bin/env bash

# generate swagger documentation
echo "Generate swagger documentation..."
swag fmt
swag init -g ./api/spec/main.go -o ./api/spec/docs
