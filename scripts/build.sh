#!/usr/bin/env bash

KIND=$1
SERVICE=$2

build_gateway() {
  GATEWAY_DIR="$(pwd)/api/gateway"
  DOCKERFILE="$GATEWAY_DIR/Dockerfile"

  # check Dockerfile.
  if [[ ! -f "$DOCKERFILE" ]]; then
    echo >&2 "error: Dockerfile not found"
    exit 1
  fi

  IMAGE_NAME=billykore/monorepo-gateway:latest

  # build Docker image.
  echo "Build image..."
  docker build -t $IMAGE_NAME -f $DOCKERFILE .

  # push image to DockerHub.
  echo "Push image..."
  docker push $IMAGE_NAME

  # delete local Docker image.
  echo "Delete image $IMAGE_NAME..."
  docker image rm $IMAGE_NAME
}

build_service() {
  DOCKERFILE="$(pwd)/services/$SERVICE/Dockerfile"

  # check Dockerfile.
  if [[ ! -f "$DOCKERFILE" ]]; then
  echo >&2 "error: Dockerfile not found"
  exit 1
  fi

  IMAGE_NAME=billykore/monorepo-$SERVICE:latest

  # build Docker image.
  echo "Build image..."
  docker build -t $IMAGE_NAME -f $DOCKERFILE .

  # push image to DockerHub.
  echo "Push image..."
  docker push $IMAGE_NAME

  # delete local Docker image.
  echo "Delete image $IMAGE_NAME..."
  docker image rm $IMAGE_NAME
}

if [[ $KIND == "gateway" ]]; then
  build_gateway
fi
if [[ $KIND == "service" ]]; then
  build_service
fi
