#!/usr/bin/env bash

set -euo pipefail

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Run from root of repo
cd "${MY_DIR}/.."

docker build -t spirits .
docker run --rm -p 12345:12345 --name spirits spirits