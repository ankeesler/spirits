#!/usr/bin/env bash

set -euo pipefail

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Run from root of repo
cd "${MY_DIR}/.."

test -z "$(go fmt ./...)"
go vet -v ./...
go test -race -v ./...
