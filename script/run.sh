#!/usr/bin/env bash

set -exuo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

cd "$REPO_ROOT"

pushd web >/dev/null
 npm run build
popd >/dev/null

pushd api >/dev/null
  go build -o /tmp/spirits .
popd >/dev/null

/tmp/spirits -web-assets-dir web/build