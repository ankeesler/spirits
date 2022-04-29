#!/usr/bin/env bash

set -euo pipefail

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
SPIRITSPKG="github.com/ankeesler/spirits"
CODEGENPKG="$(grep k8s.io/code-generator go.mod | tr " " "@" | tr -d " \t")"

note() {
  echo "generate.sh > $@"
}

# Run from root of repo
cd "${MY_DIR}/.."

# Load go environment so we can access go mod cache
eval "$(go env)"
go mod download
go mod tidy

# Run codegen for external types
note "running codegen for external types..."
bash ${GOMODCACHE}/${CODEGENPKG}/generate-groups.sh \
  deepcopy,defaulter,conversion \
  ${SPIRITSPKG}/pkg/apis \
  ${SPIRITSPKG}/pkg/apis \
  spirits:v1alpha1 \
  --go-header-file hack/boilerplate.go.txt -v 1

# Run codegen for internal types
note "running codegen for internal types..."
bash ${GOMODCACHE}/${CODEGENPKG}/generate-internal-groups.sh \
  deepcopy,defaulter,conversion \
  ${SPIRITSPKG}/internal/apis \
  ${SPIRITSPKG}/internal/apis \
  ${SPIRITSPKG}/pkg/apis \
  spirits:v1alpha1 \
  --go-header-file hack/boilerplate.go.txt -v 1

# Generate CRDs
note "running codegen for CRDs..."
go run sigs.k8s.io/controller-tools/cmd/controller-gen \
  paths=./pkg/apis/spirits/v1alpha1 crd output:crd:artifacts:config=./config/crd

# Generate RBAC
note "running codegen for RBAC..."
go run sigs.k8s.io/controller-tools/cmd/controller-gen \
  paths=./pkg/controller +rbac:roleName=spirits-controller-manager output:rbac:dir=./config
mv config/role.yaml config/zz_generated.role.yaml