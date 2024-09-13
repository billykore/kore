#!/usr/bin/env bash

KIND=$1
SERVICE=$2

function deploy_gateway() {
  GATEWAY_DIR="$(pwd)/infra/gateway"
  # apply deployment to kubernetes.
  kubectl apply -f "$GATEWAY_DIR/deployment.yaml"
}

function deploy_service() {
  K8S_DIR="$(pwd)/infra/services/$SERVICE"
  # apply to kubernetes.
  kubectl apply -f "$K8S_DIR"
}

function deploy_database() {
  K8S_DIR="$(pwd)/infra/database/$SERVICE"
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
