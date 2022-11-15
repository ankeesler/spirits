FROM node:18-alpine3.16 as react-builder

WORKDIR /workspace

COPY package.json package.json
COPY package-lock.json package-lock.json
COPY tsconfig.json tsconfig.json
COPY public/ public/
COPY src/ src/

RUN npm install
RUN npm run build --production

FROM golang:1.19-alpine3.16 as server-builder

WORKDIR /workspace

# Copy the source
COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/

# Build
RUN \
  --mount=type=cache,target=/cache/gocache \
  --mount=type=cache,target=/cache/gomodcache \
  GOCACHE=/cache/gocache \
  GOMODCACHE=/cache/gomodcache \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  go build -a -o server ./cmd/server/web

FROM gcr.io/distroless/static:nonroot

WORKDIR /
ENV NODE_ENV production

COPY --from=react-builder /workspace/build .
COPY --from=server-builder /workspace/server .

USER 65532:65532

ENTRYPOINT ["/server"]
CMD ["-web-assets-dir", "/build", "-upstream-api-server", "127.0.0.1:12345"]
