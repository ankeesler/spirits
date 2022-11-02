#!/usr/bin/env bash

set -euo pipefail

protoc api/*.proto -Iapi --go_out=paths=source_relative:internal/api --go-grpc_out=paths=source_relative:internal/api
