#!/usr/bin/env bash

set -euo pipefail

ME="$( basename "${BASH_SOURCE[0]}" )"
MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Run from root of repo
cd "${MY_DIR}/.."

usage() {
  echo "usage: $ME [-abefh]"
  echo "  -a    run all checks"
  echo "  -b    run backend checks"
  echo "  -e    run e2e checks"
  echo "  -f    run frontend checks"
  echo "  -h    print this message"
}

backend=0
e2e=0
frontend=0
while getopts "abefh" o; do
  case "$o" in
    a)
      backend=1
      e2e=1
      frontend=1
      ;;
    b)
      backend=1
      ;;
    e)
      e2e=1
      ;;
    f)
      frontend=1
      ;;
    h)
      usage
      exit 1
      ;;
    [?])
      usage
      exit 1
      ;;
  esac
done

if [[ "$backend" -ne 0 ]]; then
  echo "running backend tests..."
  pushd api
    go mod download
    test -z "$(go fmt ./...)" || (echo "'go fmt' failed" && exit 1)
    go vet ./...
    go test -race -v ./...
  popd
fi

if [[ "$frontend" -ne 0 ]]; then
  echo "running frontend tests..."
  pushd web
    npm install
    npm run check --all
  popd
fi

if [[ "$e2e" -ne 0 ]]; then
  echo "running e2e tests..."
  pushd test
    docker build -t spirits-test -f ../Dockerfile ..
    docker run --rm -d -p 12345:12345 --name spirits-test spirits-test
    cleanup() {
      docker stop spirits-test
    }
    trap cleanup EXIT
    npm install
    SPIRITS_TEST_URL=http://localhost:12345 npm run check
  popd
fi