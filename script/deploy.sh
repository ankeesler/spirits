#!/usr/bin/env bash

set -exuo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

cd "$REPO_ROOT"

if [[ -n "$(git status --porcelain)" ]]; then
  echo "error: uncommited git changes"
  exit 1
fi

git push heroku main
./script/acceptance-test.sh