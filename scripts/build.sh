#!/usr/bin/env bash

DOCKERFILE="$(pwd)/Dockerfile"

# check Dockerfile.
if [[ ! -f "$DOCKERFILE" ]]; then
  echo >&2 "error: Dockerfile not found"
  exit 1
fi

# build Docker image.
echo "Build image..."
docker build -t billykore/todo-list-app:latest .

# push image to DockerHub.
echo "Push image..."
docker push billykore/todo-list-app
