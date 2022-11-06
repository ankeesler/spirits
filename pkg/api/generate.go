package api

//go:generate protoc ../../api/action.proto ../../api/battle.proto ../../api/meta.proto ../../api/spirit.proto -I../../api --go_out=paths=source_relative:../../pkg/api --go-grpc_out=paths=source_relative:../../pkg/api
