#!/usr/bin/env bash

set -euo pipefail

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Run from root of repo
cd "${MY_DIR}/.."

docker run -v "${PWD}:/spirits" spirits-codegen:latest \
  --go_out=paths=source_relative:/spirits/pkg/api \
  --go-grpc_out=paths=source_relative:/spirits/pkg/api \
  --validate_out=lang=go,paths=source_relative:/spirits/pkg/api \
  --grpc-gateway_out=paths=source_relative:/spirits/pkg/api/gateway \
  -I/spirits/api \
  /spirits/api/spirits/v1/meta.proto \
  /spirits/api/spirits/v1/action.proto \
  /spirits/api/spirits/v1/spirit.proto \
  /spirits/api/spirits/v1/battle.proto
