#!/usr/bin/env bash

K8S_DIR="$(pwd)/deployment/k8s"

# check deployment folder.
if [[ ! -d "$K8S_DIR" ]]; then
  echo >&2 "error: deployment/k8s folder not found"
  exit 1
fi

# check deployment file.
if [[ ! -f "$K8S_DIR/deployment.yaml" ]]; then
  echo >&2 "error: deployment/k8s/deployment.yaml not found"
  exit 1
fi

# check env file.
if [[ ! -f "$K8S_DIR/env.yaml" ]]; then
  echo >&2 "error: deployment/k8s/env.yaml not found"
  exit 1
fi

# apply deployment to kubernetes.
kubectl apply -f "$K8S_DIR/deployment.yaml"

# apply env config map to kubernetes.
kubectl apply -f "$K8S_DIR/env.yaml"