# Build the server binary
FROM golang:1.19 as builder

WORKDIR /workspace

# Copy the source
COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/
COPY api/builtin builtin/

# Build
RUN \
  --mount=type=cache,target=/cache/gocache \
  --mount=type=cache,target=/cache/gomodcache \
  GOCACHE=/cache/gocache \
  GOMODCACHE=/cache/gomodcache \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  go build -a -o server ./cmd/server

# Use distroless as minimal base image to package the server binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/server .
COPY --from=builder /workspace/builtin/ /etc/spirits/builtin/
USER 65532:65532

ENTRYPOINT ["/server"]
CMD ["-spirit-builtin-dir", "/etc/spirits/builtin/spirit", "-action-builtin-dir", "/etc/spirits/builtin/action"]
