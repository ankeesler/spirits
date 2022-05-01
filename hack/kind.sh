#!/usr/bin/env bash

set -euo pipefail

CLUSTER_NAME="spirits-kind"
REG_NAME="kind-registry"

up() {
  # From https://kind.sigs.k8s.io/docs/user/local-registry...
  
  # create registry container unless it already exists
  reg_port='5001'
  if [ "$(docker inspect -f '{{.State.Running}}' "${REG_NAME}" 2>/dev/null || true)" != 'true' ]; then
    docker run \
      -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${REG_NAME}" \
      registry:2
  fi
  
  # create a cluster with the local registry enabled in containerd
  if ! grep -qw ${CLUSTER_NAME} <<<"$(kind get clusters 2>/dev/null)"; then
    cat <<EOF | kind create cluster --name "${CLUSTER_NAME}" --config=-
    kind: Cluster
    apiVersion: kind.x-k8s.io/v1alpha4
    containerdConfigPatches:
    - |-
      [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
        endpoint = ["http://${REG_NAME}:5000"]
  EOF
  fi
  
  # connect the registry to the cluster network if not already connected
  if [ "$(docker inspect -f='{{json .NetworkSettings.Networks.kind}}' "${REG_NAME}")" = 'null' ]; then
    docker network connect "kind" "${REG_NAME}"
  fi
  
  # Document the local registry
  # https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry
  cat <<EOF | kubectl apply -f -
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: local-registry-hosting
    namespace: kube-public
  data:
    localRegistryHosting.v1: |
      host: "localhost:${reg_port}"
      help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
  EOF
}

down() {
  kind delete cluster --name "$CLUSTER_NAME"
  docker stop "$REG_NAME"
  docker rm "$REG_NAME"
}

"$@"
