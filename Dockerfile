# Build the manager binary
FROM golang:1.18 as builder

WORKDIR /workspace

# Copy the source
COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY internal/ internal/

# Build
RUN \
  --mount=type=cache,target=/cache/gocache \
  --mount=type=cache,target=/cache/gomodcache \
  GOCACHE=/cache/gocache \
  GOMODCACHE=/cache/gomodcache \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  go build -a -o manager ./cmd/manager

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
