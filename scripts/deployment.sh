#!/usr/bin/env bash

KIND=$1

function deploy_gateway() {
  GATEWAY_DIR="$(pwd)/api/gateway"

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

SERVICE=$2

function deploy_service() {
  K8S_DIR="$(pwd)/services/$SERVICE/deployment/k8s"

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

  # apply deployment to kubernetes.
  kubectl apply -f "$K8S_DIR/deployment.yaml"

  # apply env config map to kubernetes.
  kubectl apply -f "$K8S_DIR/env.yaml"
}

if [[ $KIND -eq "gateway" ]]; then
  deploy_gateway
elif [ $KIND -eq "service" ]; then
  deploy_service
fi
