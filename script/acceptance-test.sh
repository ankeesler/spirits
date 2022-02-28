#!/usr/bin/env bash

set -exuo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

cd "$REPO_ROOT"
SPIRITS_TEST_BASE_URL=https://oh-great-spirits.herokuapp.com go test -count 1 -v ./test
