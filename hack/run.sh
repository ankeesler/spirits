#!/usr/bin/env bash

set -euo pipefail

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Run from root of repo
cd "${MY_DIR}/.."

# Ensure kind cluster to test with
./hack/kind.sh up

# Run tilt
tilt up spirits-server-compile spirits-server --stream
