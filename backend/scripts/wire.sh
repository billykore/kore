#!/usr/bin/env bash

TARGET=$1

# run wire dependency injection
echo "Run wire dependency injection..."
wire ./services/$TARGET/cmd
