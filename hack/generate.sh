#!/usr/bin/env bash

set -euo pipefail

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
SPIRITSPKG="github.com/ankeesler/spirits"
CODEGENPKG="$(grep k8s.io/code-generator go.mod | tr " " "@" | tr -d " \t")"

# Run from root of repo
cd "${MY_DIR}/.."

# Load go environment so we can access go mod cache
eval "$(go env)"
go mod download

# Run codegen for external types
echo "running codegen for external types..."
bash ${GOMODCACHE}/${CODEGENPKG}/generate-groups.sh \
  deepcopy,defaulter,conversion \
  ${SPIRITSPKG}/pkg/api \
  ${SPIRITSPKG}/pkg/api \
  v1alpha1 \
  --go-header-file hack/boilerplate.go.txt -v 1

# Run codegen for internal types
echo "running codegen for internal types..."
bash ${GOMODCACHE}/${CODEGENPKG}/generate-internal-groups.sh \
  deepcopy,defaulter,conversion \
  ${SPIRITSPKG}/pkg/api \
  ${SPIRITSPKG}/pkg/api \
  ${SPIRITSPKG}/pkg/api \
  v1alpha1 \
  --go-header-file hack/boilerplate.go.txt -v 1

# Generate CRDs
echo "running codegen for CRDs..."
go run sigs.k8s.io/controller-tools/cmd/controller-gen@v0.7.0 \
  paths=./pkg/api crd output:crd:artifacts:config=./config
