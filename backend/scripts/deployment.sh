#!/usr/bin/env bash

KIND=$1
SERVICE=$2

function deploy_gateway() {
  GATEWAY_DIR="$(pwd)/infra/gateway"

  # check deployment folder.
  if [[ ! -d "$GATEWAY_DIR" ]]; then
    echo >&2 "error: $GATEWAY_DIR folder not found"
    exit 1
  fi

  # check deployment file.
  if [[ ! -f "$GATEWAY_DIR/deployment.yaml" ]]; then
    echo >&2 "error: $K8S_DIR/deployment.yaml not found"
    exit 1
  fi

  # apply deployment to kubernetes.
  kubectl apply -f "$GATEWAY_DIR/deployment.yaml"
}

function deploy_service() {
  K8S_DIR="$(pwd)/infra/services/$SERVICE"

  # check deployment folder.
  if [[ ! -d "$K8S_DIR" ]]; then
    echo >&2 "error: $K8S_DIR folder not found"
    exit 1
  fi

  # check deployment file.
  if [[ ! -f "$K8S_DIR/deployment.yaml" ]]; then
    echo >&2 "error: $K8S_DIR/deployment.yaml not found"
    exit 1
  fi

  # check env file.
  if [[ ! -f "$K8S_DIR/env.yaml" ]]; then
    echo >&2 "error: $K8S_DIR/env.yaml not found"
    exit 1
  fi

  # apply to kubernetes.
  kubectl apply -f "$K8S_DIR"
}

function deploy_database() {
  K8S_DIR="$(pwd)/infra/database/$SERVICE"

  # check deployment folder.
  if [[ ! -d "$K8S_DIR" ]]; then
    echo >&2 "error: $K8S_DIR folder not found"
    exit 1
  fi

  # check deployment file.
  if [[ ! -f "$K8S_DIR/deployment.yaml" ]]; then
    echo >&2 "error: $K8S_DIR/deployment.yaml not found"
    exit 1
  fi

  # check env file.
  if [[ ! -f "$K8S_DIR/env.yaml" ]]; then
    echo >&2 "error: $K8S_DIR/env.yaml not found"
    exit 1
  fi

  # check volume file.
  if [[ ! -f "$K8S_DIR/volume.yaml" ]]; then
    echo >&2 "error: $K8S_DIR/volume.yaml not found"
    exit 1
  fi

  # apply to kubernetes.
  kubectl apply -f "$K8S_DIR"
}

if [[ $KIND == "gateway" ]]; then
  deploy_gateway
elif [ $KIND == "service" ]; then
  deploy_service
elif [ $KIND == "database" ]; then
  deploy_database
fi
