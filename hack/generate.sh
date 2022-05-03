#!/usr/bin/env bash

set -euo pipefail

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
SPIRITSPKG="$(grep '^module' go.mod | awk '{print $2}')"
CODEGENPKG="$(grep k8s.io/code-generator go.mod | tr " " "@" | tr -d " \t")"

_note() {
  echo "generate.sh > $@"
}

# Run codegen for external types
generate_groups() {
  _note "running codegen for external types..."
  bash ${GOMODCACHE}/${CODEGENPKG}/generate-groups.sh \
    deepcopy,defaulter,conversion,client \
    ${SPIRITSPKG}/pkg/apis \
    ${SPIRITSPKG}/pkg/apis \
    spirits:v1alpha1 \
    --go-header-file hack/boilerplate.go.txt -v 1
}

# Run codegen for internal types
generate_internal_groups() {
  _note "running codegen for internal types..."
  bash ${GOMODCACHE}/${CODEGENPKG}/generate-internal-groups.sh \
    deepcopy,defaulter,conversion \
    ${SPIRITSPKG}/internal/apis \
    ${SPIRITSPKG}/internal/apis \
    ${SPIRITSPKG}/pkg/apis \
    spirits:v1alpha1 \
    --go-header-file hack/boilerplate.go.txt -v 1
}

# Generate CRDs
generate_crds() {
  _note "running codegen for CRDs..."
  go run sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0 \
    paths=./pkg/apis/spirits/v1alpha1 crd output:crd:artifacts:config=./config/zz_generated_crds
}

# Generate RBAC
generate_rbac() {
  _note "running codegen for RBAC..."
  go run sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0 \
    paths=./pkg/controller +rbac:roleName=spirits-server output:rbac:dir=./config/zz_generated_rbac
}

# Run from root of repo
cd "${MY_DIR}/.."

# Load go environment so we can access go mod cache
eval "$(go env)"
go mod download
go mod tidy

if [[ "$#" == "0" ]]; then
  generate_groups
  generate_internal_groups
  generate_crds
  generate_rbac
else
  "$@"
fi
