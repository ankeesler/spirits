#!/usr/bin/env bash

set -exuo pipefail

SPIRITS_TEST_BASE_URL=https://oh-great-spirits.herokuapp.com go test -count 1 -v ./internal/test
