FROM ubuntu as protoc

WORKDIR /workspace

RUN apt-get update && apt-get install unzip curl -y
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.9/protoc-21.9-linux-x86_64.zip > protoc.zip
RUN unzip protoc.zip

FROM golang:1.19 as protoc-plugin

WORKDIR /workspace

COPY go.mod go.mod
COPY tool.go tool.go
RUN go mod tidy

RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/envoyproxy/protoc-gen-validate

# gcr.io/distroless/static:nonroot
FROM debian:bullseye-slim
WORKDIR /workspace

RUN mkdir /workspace/bin
ENV PATH=${PATH}:/workspace/bin

COPY --from=protoc /workspace/bin/protoc /workspace/bin
COPY --from=protoc /workspace/include /workspace

COPY --from=protoc-plugin /go/bin/ /workspace/bin
COPY --from=protoc-plugin /go/pkg/mod/github.com/envoyproxy/protoc-gen-validate\@v0.9.0/validate /workspace/validate

USER 65532:65532

ENTRYPOINT ["protoc", "-I/workspace"]